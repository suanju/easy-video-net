package conversion

import (
	"reflect"
	"strings"
	"unsafe"
)

// StringConversionMap 字符串转数组
func StringConversionMap(s string) []string {
	list := strings.Split(s, ",")
	return list
}

// StringImgConversionMap 字符串图片转数组
func StringImgConversionMap(s string) []string {
	list := strings.Split(s, ",")
	for k, v := range list {
		list[k] = FormattingSrc(v)
	}
	return list
}

//BoolTurnInt8 布尔类型转int8
func BoolTurnInt8(is bool) int8 {
	if is {
		return 1
	} else {
		return 0
	}
}

func IntTurnBool(i int) bool {
	if i > 0 {
		return true
	} else {
		return false
	}
}

// String2Bytes 强转化
func String2Bytes(s string) []byte {
	sh := (*reflect.StringHeader)(unsafe.Pointer(&s))
	bh := reflect.SliceHeader{
		Data: sh.Data,
		Len:  sh.Len,
		Cap:  sh.Len,
	}
	return *(*[]byte)(unsafe.Pointer(&bh))
}

// Bytes2String 强转化
func Bytes2String(b []byte) string {
	return *(*string)(unsafe.Pointer(&b))
}
