package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type EncryptDataResponse struct {

	// 密钥ID。
	KeyId *string `json:"key_id,omitempty" xml:"key_id"`

	// DEK密文16进制，两位表示1byte。
	CipherText     *string `json:"cipher_text,omitempty" xml:"cipher_text"`
	HttpStatusCode int     `json:"-"`
}

func (o EncryptDataResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "EncryptDataResponse struct{}"
	}

	return strings.Join([]string{"EncryptDataResponse", string(data)}, " ")
}
