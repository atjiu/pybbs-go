package utils

import (
	"crypto/md5"
	"encoding/hex"
	"strings"
)

func NoHtml(str string) string {
	return strings.Replace(strings.Replace(str, "<script", "&lt;script", -1), "script>", "script&gt;", -1)
}

func MD5(str string) string {
	h := md5.New()
	h.Write([]byte(str)) // 需要加密的字符串为 123456
	cipherStr := h.Sum(nil)
	return hex.EncodeToString(cipherStr)
}
