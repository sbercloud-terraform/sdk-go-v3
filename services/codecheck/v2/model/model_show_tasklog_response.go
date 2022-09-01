package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type ShowTasklogResponse struct {
	ParamInfo *ParamInfo `json:"param_info,omitempty" xml:"param_info"`

	// 日志信息
	LogInfo        *[]LogInfo `json:"log_info,omitempty" xml:"log_info"`
	HttpStatusCode int        `json:"-"`
}

func (o ShowTasklogResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowTasklogResponse struct{}"
	}

	return strings.Join([]string{"ShowTasklogResponse", string(data)}, " ")
}
