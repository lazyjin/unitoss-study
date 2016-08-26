package common

import (
	"encoding/json"
)

const (
	NORMAL int = iota
	TIME_ERR
	EUI_ERR
	FMT_ERR
)

type UdrReqMsg struct {
	ErrorType int `json:"errortype"`
	Count     int `json:"count"`
}

func UdrReqMsgParse(fromMsg []byte) (*UdrReqMsg, error) {
	var toMsg = &UdrReqMsg{}

	err := json.Unmarshal(fromMsg, toMsg)
	if err != nil {
		return nil, err
	}
	return toMsg, nil
}
