package hack

import (
	"encoding/json"
	"reflect"
	"strings"
	"unsafe"
)

// String converts slice to string without copy.
// Use at your own risk.
func String(b []byte) (s string) {
	if len(b) == 0 {
		return ""
	}
	pbytes := (*reflect.SliceHeader)(unsafe.Pointer(&b))
	pstring := (*reflect.StringHeader)(unsafe.Pointer(&s))
	pstring.Data = pbytes.Data
	pstring.Len = pbytes.Len
	return
}

// Slice converts string to slice without copy.
// Use at your own risk.
func Slice(s string) (b []byte) {
	pbytes := (*reflect.SliceHeader)(unsafe.Pointer(&b))
	pstring := (*reflect.StringHeader)(unsafe.Pointer(&s))
	pbytes.Data = pstring.Data
	pbytes.Len = pstring.Len
	pbytes.Cap = pstring.Len
	return
}

//格式例子："username=abc&password=abc"，如不符合key=value的参数会被抛弃, 重复的key以最后一个出现为准
func FormToJson(form []byte) ([]byte, error) {
	paramMap := map[string]interface{}{}
	params := strings.Split(string(form), "&")
	for _, param := range params {
		paramKeyVal := strings.Split(param, "=")
		if len(paramKeyVal) != 2 {
			continue
		}
		paramMap[paramKeyVal[0]] = paramKeyVal[1]
	}
	paramJson, err := json.Marshal(paramMap)
	if err != nil {
		return nil, err
	}
	paramByte := []byte(string(paramJson))
	return paramByte, nil
}