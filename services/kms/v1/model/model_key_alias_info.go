package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type KeyAliasInfo struct {

	// 密钥ID。
	KeyId *string `json:"key_id,omitempty" xml:"key_id"`

	// 密钥别名。
	KeyAlias *string `json:"key_alias,omitempty" xml:"key_alias"`
}

func (o KeyAliasInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "KeyAliasInfo struct{}"
	}

	return strings.Join([]string{"KeyAliasInfo", string(data)}, " ")
}
