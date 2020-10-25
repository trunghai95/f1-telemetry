package main

import (
	"flag"
	"log"
	"net"
	"os"

	"github.com/trunghai95/f1-telemetry/config"
)

var (
	configPath = flag.String("configPath", "../.conf/config.yaml", "YAML config path")
	logPath    = flag.String("logPath", "./log/app.log", "Path to write log")
)

func main() {
	flag.Parse()

	err := initLog()
	if err != nil {
		log.Fatalf("Init log err: %v", err)
	}

	err = config.InitConfigYaml(*configPath)
	if err != nil {
		log.Fatalf("Init config err: %v", err)
	}

	log.Println("Starting...")
	startUDPServer()
}

func initLog() (err error) {
	log.SetFlags(log.Ldate | log.Ltime | log.Lmicroseconds | log.Lshortfile)
	if len(*logPath) == 0 {
		return
	}
	logFile, err := os.OpenFile(*logPath, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		return
	}
	log.SetOutput(logFile)
	return
}

func startUDPServer() {
	s, err := net.ResolveUDPAddr("udp4", config.GetAppConfig().UDPListen)
	if err != nil {
		log.Fatalf("ResolveUDPAddr err: %v", err)
	}

	conn, err := net.ListenUDP("udp4", s)
	if err != nil {
		log.Fatalf("ListenUDP err: %v", err)
	}

	buf := make([]byte, 2048) // Should be at least the size of the smallest packet possible
	for {
		// Listen to conn
		nBytes, addr, err := conn.ReadFromUDP(buf)
		if err != nil {
			log.Printf("Read UDP msg err: %v", err)
			continue
		}
		log.Printf("DEBUG | Received %v bytes from %v", nBytes, addr.String())

		// TODO: Parse the message
	}
}
