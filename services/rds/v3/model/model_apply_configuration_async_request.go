package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type ApplyConfigurationAsyncRequest struct {

	// 语言
	XLanguage *string `json:"X-Language,omitempty" xml:"X-Language"`

	// 参数模板ID。
	ConfigId string `json:"config_id" xml:"config_id"`

	Body *ApplyConfigurationRequest `json:"body,omitempty" xml:"body"`
}

func (o ApplyConfigurationAsyncRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ApplyConfigurationAsyncRequest struct{}"
	}

	return strings.Join([]string{"ApplyConfigurationAsyncRequest", string(data)}, " ")
}
