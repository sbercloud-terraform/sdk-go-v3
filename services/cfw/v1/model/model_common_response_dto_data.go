package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// object
type CommonResponseDtoData struct {

	// 防护对象id
	Id *string `json:"id,omitempty"`
}

func (o CommonResponseDtoData) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CommonResponseDtoData struct{}"
	}

	return strings.Join([]string{"CommonResponseDtoData", string(data)}, " ")
}
