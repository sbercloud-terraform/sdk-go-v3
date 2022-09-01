package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type HttpInfoResponseBody struct {

	// HTTPS证书是否启用。0：不启用，此时无需填写证书及私钥参数；1：启用HTTPS加速并协议跟随回源；2：启用HTTPS加速并HTTP回源；3：启用HTTPS加速并HTTPS回源，开启时需要传递证书及私钥
	HttpsStatus *int32 `json:"https_status,omitempty" xml:"https_status"`

	// 证书名称。（长度限制为3-32字符）。
	CertName *string `json:"cert_name,omitempty" xml:"cert_name"`

	// 证书内容。
	Certificate *string `json:"certificate,omitempty" xml:"certificate"`

	// 功能说明： HTTPS协议使用的私钥，不启用证书则无需输入。（为了客户信息安全，接口返回私钥为空）
	PrivateKey *string `json:"private_key,omitempty" xml:"private_key"`

	// 证书类型。1：代表华为云托管证书；0：表示自有证书。
	CertificateType *int32 `json:"certificate_type,omitempty" xml:"certificate_type"`

	// 客户端请求是否强制重定向。1是，0否。（如果为2，表示强制跳转HTTP）
	ForceRedirectHttps *int32 `json:"force_redirect_https,omitempty" xml:"force_redirect_https"`

	ForceRedirectConfig *ForceRedirect `json:"force_redirect_config,omitempty" xml:"force_redirect_config"`

	// 是否使用HTTP2.0。（1是，0否。）
	Http2 *int32 `json:"http2,omitempty" xml:"http2"`

	// 证书过期时间
	ExpirationTime *int64 `json:"expiration_time,omitempty" xml:"expiration_time"`
}

func (o HttpInfoResponseBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "HttpInfoResponseBody struct{}"
	}

	return strings.Join([]string{"HttpInfoResponseBody", string(data)}, " ")
}
