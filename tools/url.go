package tools

import (
	"github.com/pkg/errors"
	"strings"
)

type UrlOptions struct {
	IsJoinSubPrefix bool
	SubDomainPrefix string
	Protocol        string
}

// BuildHttpUrl 区分是域名还是正确的 url ，并将正确的 url 返回
func BuildHttpUrl(url string, options UrlOptions) (string, error) {
	if !IsCarryHost(url) {
		return url, errors.New("It is illegal to url")
	}
	if IsHttp(url) {
		return url, nil
	}
	if options.IsJoinSubPrefix {
		url = options.Protocol + "://" + options.SubDomainPrefix + "." + url
	}

	return "", nil
}

func IsHost() bool {
	return false
}

func IsHostAndPort() bool {
	return false
}

// IsHttp 是否是http协议
func IsHttp(url string) bool {
	return strings.HasPrefix(url, "http") || strings.HasPrefix(url, "https")
}

// IsCarryHost url 是否携带域名
func IsCarryHost(url string) bool {
	return strings.LastIndex(url, ".") != -1
}
