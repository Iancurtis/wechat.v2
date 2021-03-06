// +build !wechatdebug

package api

import (
	"io"

	"github.com/iancurtis/util"
)

func DebugPrintGetRequest(url string) {}

func DebugPrintPostXMLRequest(url string, body []byte) {}

func DecodeXMLHttpResponse(r io.Reader) (map[string]string, error) {
	return util.DecodeXMLToMap(r)
}
