package util

import "fmt"

func FormatStr(text string, data ...any) string {
	return fmt.Sprintf(text, data...)
}

func ByteToString(data []byte) string {
	return string(data)
}
