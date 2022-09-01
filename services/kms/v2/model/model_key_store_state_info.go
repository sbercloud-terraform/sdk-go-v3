package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 密钥库状态详情
type KeyStoreStateInfo struct {

	// 密钥库ID
	KeystoreId *string `json:"keystore_id,omitempty" xml:"keystore_id"`

	// 密钥库状态
	KeystoreState *string `json:"keystore_state,omitempty" xml:"keystore_state"`
}

func (o KeyStoreStateInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "KeyStoreStateInfo struct{}"
	}

	return strings.Join([]string{"KeyStoreStateInfo", string(data)}, " ")
}
