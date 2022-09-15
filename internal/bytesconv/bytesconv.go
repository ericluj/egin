package bytesconv

import "unsafe"

// TODO:
func StringToBytes(s string) []byte {
	return *(*[]byte)(unsafe.Pointer(
		&struct {
			string
			Cap int
		}{s, len(s)},
	))
}

// TODO:
func BytesToString(b []byte) string {
	return *(*string)(unsafe.Pointer(&b))
}
