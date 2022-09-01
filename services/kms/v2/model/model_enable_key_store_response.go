package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type EnableKeyStoreResponse struct {
	Keystore       *KeyStoreStateInfo `json:"keystore,omitempty" xml:"keystore"`
	HttpStatusCode int                `json:"-"`
}

func (o EnableKeyStoreResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "EnableKeyStoreResponse struct{}"
	}

	return strings.Join([]string{"EnableKeyStoreResponse", string(data)}, " ")
}
