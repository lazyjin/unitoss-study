package udr

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"time"
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
	TIME_FMT    = "%d%02d%02d%02d%02d%02d%1d"
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

func (ur *UdrRaw) ConvToJsonStr() (string, error) {
	jsonUdr, err := json.Marshal(ur)
	if err != nil {
		return "", err
	}

	return string(jsonUdr), nil
}

func MakeRandomUdr() (UdrRaw, error) {
	tmpUdr := GetEmptyUdrRaw()

	r := rand.New(rand.NewSource(time.Now().UnixNano()))

	// make random EUI && byte count
	randEui := r.Uint32()%EUI_BASE + EUI_BASE
	randByte := r.Uint32() % 10 * 100

	// make time fields
	start, end, _ := getUdrTime()

	tmpUdr.SetUdrRaw(randEui, start, end, randByte, "")

	return tmpUdr, nil
}

func MakeTimeErrUdr() (UdrRaw, error) {
	tmpUdr := GetEmptyUdrRaw()

	r := rand.New(rand.NewSource(time.Now().UnixNano()))

	// make random EUI && byte count
	randEui := r.Uint32()%EUI_BASE + EUI_BASE
	randByte := r.Uint32() % 10 * 100
	randType := r.Uint32() % 2

	start, end, _ := getWrongTime(randType)

	tmpUdr.SetUdrRaw(randEui, start, end, randByte, "")

	return tmpUdr, nil
}

func MakeFmtErrUdr() (UdrRaw, error) {
	tmpUdr := GetEmptyUdrRaw()

	r := rand.New(rand.NewSource(time.Now().UnixNano()))

	randType := r.Uint32() % 2

	var randEui uint32
	var gwid = ""

	if randType == 0 { // EUI length error
		randEui = r.Uint32()%EUI_BASE + EUI_BASE + USERID_BASE
	} else {
		randEui = r.Uint32()%EUI_BASE + EUI_BASE
	}

	start, end, _ := getUdrTime()

	if randType == 1 {
		gwid = "ERRGWLONGLENGTH"
	}

	randByte := r.Uint32() % 10 * 100

	tmpUdr.SetUdrRaw(randEui, start, end, randByte, gwid)

	return tmpUdr, nil
}

func MakeEuiErrUdr() (UdrRaw, error) {
	tmpUdr := GetEmptyUdrRaw()

	r := rand.New(rand.NewSource(time.Now().UnixNano()))

	// make random out of range EUI && byte count
	randEui := r.Uint32()%EUI_BASE + EUI_BASE*2
	randByte := r.Uint32() % 10 * 100

	start, end, _ := getUdrTime()

	tmpUdr.SetUdrRaw(randEui, start, end, randByte, "")

	return tmpUdr, nil
}

func getUdrTime() (string, string, error) {
	now := time.Now()
	start := fmt.Sprintf(
		TIME_FMT,
		now.Year(),
		now.Month(),
		now.Day(),
		now.Hour(),
		now.Minute(),
		now.Second(),
		now.Nanosecond()/1000000000)

	d, err := time.ParseDuration("5s")
	if err != nil {
		return "", "", err
	}

	then := now.Add(d)
	end := fmt.Sprintf(
		TIME_FMT,
		then.Year(),
		then.Month(),
		then.Day(),
		then.Hour(),
		then.Minute(),
		then.Second(),
		then.Nanosecond()/1000000000)

	return start, end, nil
}

func getWrongTime(typ uint32) (string, string, error) {
	var (
		start string
		end   string
	)

	if typ == 0 {
		end, start, _ = getUdrTime()
	} else {
		now := time.Now()
		start = fmt.Sprintf(
			TIME_FMT,
			"2020",
			now.Month(),
			now.Day(),
			now.Hour(),
			now.Minute(),
			now.Second(),
			now.Nanosecond()/1000000000)

		d, err := time.ParseDuration("5s")
		if err != nil {
			return "", "", err
		}

		then := now.Add(d)
		end = fmt.Sprintf(
			TIME_FMT,
			"2020",
			then.Month(),
			then.Day(),
			then.Hour(),
			then.Minute(),
			then.Second(),
			then.Nanosecond()/1000000000)
	}

	return start, end, nil
}
