package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type UpdateMixJobRequest struct {

	// 使用AK/SK方式认证时必选，携带的鉴权信息。
	Authorization *string `json:"Authorization,omitempty" xml:"Authorization"`

	// 使用AK/SK方式认证时必选，请求的发生时间。
	XSdkDate *string `json:"X-Sdk-Date,omitempty" xml:"X-Sdk-Date"`

	// 使用AK/SK方式认证时必选，携带项目ID信息。
	XProjectId *string `json:"X-Project-Id,omitempty" xml:"X-Project-Id"`

	// 应用id
	AppId string `json:"app_id" xml:"app_id"`

	// 任务编号
	JobId string `json:"job_id" xml:"job_id"`

	Body *UpdateMixJobReq `json:"body,omitempty" xml:"body"`
}

func (o UpdateMixJobRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateMixJobRequest struct{}"
	}

	return strings.Join([]string{"UpdateMixJobRequest", string(data)}, " ")
}
