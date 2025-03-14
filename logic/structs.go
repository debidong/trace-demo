package logic

import "strings"

func FormatRequestURL(srvName, uri string) string {
	prefix := "http:/"
	return strings.Join([]string{prefix, C.Server[srvName].Addr, uri}, "/")
}
