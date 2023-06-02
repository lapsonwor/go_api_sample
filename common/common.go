package common

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"net/url"
	"os"
	"reflect"
	"strconv"
	"strings"
	"time"
)

var env string = os.Getenv("API_ENV")
var DEBUG bool = false

func IsEmpty(str string) bool {
	return len(str) == 0
}

func IsError(err error) bool {
	return err != nil
}

func PanicWhenError(err error) {
	if err != nil {
		panic(err)
	}
}

func SecToUnixTime(sec int) time.Time {
	return time.Unix(int64(sec), 0)
}

func MsecToUnixTime(msec int) time.Time {
	return time.Unix(int64(float64(msec)*0.001), 0)
}

func PrintJSON(v []byte) {
	if strings.Contains(strings.ToUpper(env), "DEBUG") || DEBUG {
		m := map[string]interface{}{}
		if err := json.Unmarshal(v, &m); err == nil {
			if b, err := json.MarshalIndent(m, "", "    "); err == nil {
				fmt.Printf("%s\n", b)
			}
		} else {
			fmt.Printf("%s\n", v)
		}
	}
}

func IsJSONArray(v []byte) bool {
	x := bytes.TrimLeft(v, " \t\r\n")
	return len(x) > 0 && x[0] == '['
}

const TARGET_TIME_LAYOUT_1 string = "2006-01-02 15:04:05"
const TARGET_TIME_LAYOUT_2 string = "2006-01-02"
const LEGACY_TIME_LAYOUT_1 string = time.RFC3339
const LEGACY_TIME_LAYOUT_2 string = "2006-01-02 15:04:05.999999-07"

func FromLegacyTimeLayout(value string) (time.Time, error) {
	if t, err := time.Parse(LEGACY_TIME_LAYOUT_1, value); err == nil {
		return t, err
	}
	if t, err := time.Parse(LEGACY_TIME_LAYOUT_2, value); err == nil {
		return t, err
	}
	if value, err := url.QueryUnescape(value); err == nil {
		if t, err := time.Parse(LEGACY_TIME_LAYOUT_1, value); err == nil {
			return t, err
		}
		if t, err := time.Parse(LEGACY_TIME_LAYOUT_2, value); err == nil {
			return t, err
		}
	}
	return time.Time{}, errors.New("unknown time format: " + value)
}

func FromTargetTimeLayout(value string) (time.Time, error) {
	if t, err := time.Parse(TARGET_TIME_LAYOUT_1, value); err == nil {
		return t, err
	}
	return time.Parse(TARGET_TIME_LAYOUT_2, value)
}

func ToTargetTime(t time.Time) string {
	return t.Format(TARGET_TIME_LAYOUT_1)
}

func ToLegacyTime(t time.Time) string {
	return t.Format(LEGACY_TIME_LAYOUT_1)
}

func ParseFloat64(s string) float64 {
	f, _ := strconv.ParseFloat(s, 64)
	return f
}

func ParseInt(s string) int {
	i, _ := strconv.Atoi(s)
	return i
}

func IfThenElse(_if bool, _then, _else interface{}) interface{} {
	if _if {
		return _then
	} else {
		return _else
	}
}

func ArrayMap(a interface{}, f func(int, interface{}) interface{}) (v []interface{}) {
	for i := 0; i < reflect.ValueOf(a).Len(); i++ {
		j := reflect.ValueOf(a).Index(i).Interface()
		if o := f(i, j); o != nil {
			v = append(v, o)
		}
	}
	return v
}

func ConvertStringArray(a []interface{}) (s []string) {
	for _, v := range a {
		if vs, ok := v.(string); ok {
			s = append(s, vs)
		}
	}
	return s
}

func IsExist(a interface{}, b interface{}, c func(interface{}, interface{}) bool) bool {
	for i := 0; i < reflect.ValueOf(a).Len(); i++ {
		j := reflect.ValueOf(a).Index(i).Interface()
		if c(j, b) {
			return true
		}
	}
	return false
}

func SwitchCase(value string, options []string, output []interface{}, _default interface{}) interface{} {
	for i, option := range options {
		if value == option {
			return output[i]
		}
	}
	return _default
}

type Float64 float64

func (f *Float64) UnmarshalJSON(bytes []byte) error {
	str := string(bytes)
	if bytes[0] == '"' && bytes[len(bytes)-1] == '"' {
		str = string(bytes[1 : len(bytes)-1])
	}
	o, _ := strconv.ParseFloat(str, 64)
	*f = Float64(o)
	return nil
}

func IsByteArrayChannelClosed(ch chan []byte) bool {
	select {
	case _, ok := <-ch:
		return !ok
	default:
	}
	return false
}

type KeyMarker map[string]bool

func (km *KeyMarker) IsExisted(k string) bool {
	_, ok := (*km)[k]
	return ok
}

func (km *KeyMarker) IsUsed(k string) bool {
	b, ok := (*km)[k]
	return ok && b
}

func (km *KeyMarker) Delete(k string) {
	if _, ok := (*km)[k]; ok {
		delete(*km, k)
	}
}

func (km *KeyMarker) Use(k string) {
	(*km)[k] = true
}

func (km *KeyMarker) New(k string) {
	(*km)[k] = false
}
