package tools

import (
	"github.com/iancurtis/wechat.v2/mch/core"
)

// 授权码查询OPENID接口.
func AuthCodeToOpenId(clt *core.Client, req map[string]string) (resp map[string]string, err error) {
	return clt.PostXML("https://api.mch.weixin.qq.com/tools/authcodetoopenid", req)
}
