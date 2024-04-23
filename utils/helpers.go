package utils

import (
	"crypto/sha1"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"hash/fnv"
	"strconv"
	"time"
)

// IntStrFormtZero
// Convert int convert to string formt 0000000
func IntStrFormtZero(number int) string {
	return fmt.Sprintf("%06d", number)
}

// Convert Float64 convert to string
func FloatStr(fl float64) string {
	return fmt.Sprintf("%f", fl)
}

// Convert Float64 convert to string
func Int64ToStr(nt int64) string {
	return fmt.Sprintf("%v", nt)
}

// Current time
func CurTimeNow() time.Time {
	return time.Now()
}

// Current tim in Unix
func CurUnixTime() int64 {
	return time.Now().Unix()
}

// String number for document
func NumDoc() string {
	return fmt.Sprintf("%v", CurUnixTime())
}

// Curren time
func CurTm() string {
	return time.Now().Format("02-01-2006 15:04:05")
}

// Current time
// 2018-03-01T12:26:40+00:00
func CurTime() string {
	return time.Now().Format("2006-01-02T15:04:05+00:00")
}

// Code based to timedate
func CodeByTime() string {
	return time.Now().Format("02012006150405")
}

// Code based to timedate
func CodeByTimeShort() string {
	return time.Now().Format("02150405")
}

// Log
func Logg(txt string) {
	fmt.Printf("%s : %s \n", time.Now().Format("02-01-2006 15:04:05"), txt)
}

// Hash
func Hash(s string) uint32 {
	h := fnv.New32a()
	h.Write([]byte(s))
	return h.Sum32()
}

// Unixtime convert to string
func UnixString(unxtime int64) string {
	// var unix_timestamp int64 = 1234567890
	tm := time.Unix(unxtime, 0)
	return fmt.Sprintf("UNIX Timestamp %d указывает на %s",
		unxtime, tm.Format("2 January 2006 15:04"))
}

// Sha
func Sha(txt string) string {
	h := sha1.New()
	h.Write([]byte(txt))
	sha1_hash := hex.EncodeToString(h.Sum(nil))
	return sha1_hash
}

// Convertor string to int
func StrInt(s string) int {
	i, _ := strconv.Atoi(s)
	return i
}

// Конвертация
func Cnvrt(fl float64) string {
	result := fmt.Sprintf("%0.*f", 2, fl)
	return result
}

// Convertor string to int64
func Atoi(s string) int64 {
	i, _ := strconv.ParseInt(s, 10, 64)
	return i
}

// Convert str to float64
// https://yourbasic.org/golang/convert-string-to-float/
func StrToFloat(f string) float64 {
	s, err := strconv.ParseFloat(f, 64)
	if err != nil {
		return 0.0
	}
	return s
}

// Binding data
func Binding(bt []byte, interf *interface{}) interface{} {
	err := json.Unmarshal(bt, &interf)
	if err != nil {
		return nil
	}
	return interf
}

// Time stamp
type Timestamp time.Time

func (t *Timestamp) Param(src string) error {
	ts, err := time.Parse(time.RFC3339, src)
	*t = Timestamp(ts)
	return err
}
