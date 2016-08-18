package udr

import (
	"common"
	"encoding/json"
)

type UdrRaw struct {
	Eui       uint32 `json:"eui"`       // device identifier 1000000~2000000
	StartTime string `json:"startTime"` // packet trans start time
	EndTime   string `json:"endTime"`   // packet trans end time
	ByteCount uint32 `json:"byteCount"` // packet size
	Gateway   string `json:"gateway"`   // gateway ID
}

type UdrRated struct {
	Userid    uint32  `json:"userid"` // user ID 10000000~99999999
	Eui       uint32  `json:"eui"`
	StartTime string  `json:"startTime"`
	EndTime   string  `json:"endTime"`
	ByteCount uint32  `json:"byteCount"`
	Gateway   string  `json:"gateway"`
	ChargeAmt int32   `json:"chargeAmt"` // charge amount
	Latitude  float64 `json:"latitude"`  // gateway latitude
	Longitude float64 `json:"longitude"` // gateway longitude
}

const (
	EUI_BASE    = 1000000
	USERID_BASE = 10000000
	USERID_MAX  = 90000000
)

func GetEmptyUdrRaw() UdrRaw {
	var ur UdrRaw

	ur.SetUdrRaw(0, "", "", 0, "")

	return ur
}

func (ur *UdrRaw) SetUdrRaw(eui uint32, start string, end string, byteCnt uint32, gateway string) {
	ur.Eui = eui
	ur.StartTime = start
	ur.EndTime = end
	ur.ByteCount = byteCnt
	ur.Gateway = gateway
}

func (ur *UdrRaw) ConvToJsonStr() string {
	jsonUdr, err := json.Marshal(ur)
	common.CheckErrPanic(err)

	return string(jsonUdr)
}
