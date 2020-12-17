package awsutils

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"strings"
)

// VerifyUserInGroup Check if user is in a group
func VerifyUserInGroup(token string, groupName string) bool {
	var payload string
	payload = strings.Split(token, ".")[1]
	sDec, err := base64.RawStdEncoding.DecodeString(payload)
	if err != nil {
		fmt.Println("payload:", payload)
		fmt.Println("error:", err)
		return false
	}
	var result map[string]interface{}
	json.Unmarshal([]byte(sDec), &result)

	if groups, ok := result["cognito:groups"]; ok {
		for _, v := range groups.([]interface{}) {
			if groupName == v.(string) {
				return true
			}
		}
	}
	return false
}
