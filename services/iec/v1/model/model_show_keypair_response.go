package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type ShowKeypairResponse struct {

	// 密钥名称。
	Name *string `json:"name,omitempty" xml:"name"`

	//   密钥对应publicKey信息。
	PublicKey *string `json:"public_key,omitempty" xml:"public_key"`

	// 用户ID。
	UserId *string `json:"user_id,omitempty" xml:"user_id"`

	//   密钥对应指纹信息。
	Fingerprint    *string `json:"fingerprint,omitempty" xml:"fingerprint"`
	HttpStatusCode int     `json:"-"`
}

func (o ShowKeypairResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowKeypairResponse struct{}"
	}

	return strings.Join([]string{"ShowKeypairResponse", string(data)}, " ")
}
