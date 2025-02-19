package logic

import "strings"

func FormatRequestURL(srvName, uri string) string {
	config := MustLoadConfig("../../addr.yaml")
	prefix := "http:/"
	return strings.Join([]string{prefix, config.Server[srvName].Addr, uri}, "/")
}
