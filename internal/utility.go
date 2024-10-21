package internal

import (
	"encoding/json"
)

func UnmarshalPayload(payloadSource map[string]interface{}, result interface{}) error {
	payloadBytes, err := json.Marshal(payloadSource)
	if err != nil {
		return err
	}
	err = json.Unmarshal(payloadBytes, &result)
	return err
}
