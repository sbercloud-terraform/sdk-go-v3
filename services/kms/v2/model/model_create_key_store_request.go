package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type CreateKeyStoreRequest struct {
	Body *CreateKeyStoreRequestBody `json:"body,omitempty" xml:"body"`
}

func (o CreateKeyStoreRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateKeyStoreRequest struct{}"
	}

	return strings.Join([]string{"CreateKeyStoreRequest", string(data)}, " ")
}
