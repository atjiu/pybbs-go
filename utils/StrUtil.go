package utils

import "strings"

func NoHtml(str string) string {
	return strings.Replace(strings.Replace(str, "<script", "&lt;script", -1), "script>", "script&gt;", -1)
}
