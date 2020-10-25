package util

import "encoding/json"

// JSONEncode returns the JSON formatted string for the msg
// If msg cannot be encoded to JSON, this retunrs an empty string
func JSONEncode(msg interface{}) string {
	b, err := json.Marshal(msg)
	if err != nil {
		return ""
	}
	return string(b)
}
