package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/sdktime"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type PublishResp struct {

	// 发布记录的ID
	PublishId *string `json:"publish_id,omitempty" xml:"publish_id"`

	// API编号
	ApiId *string `json:"api_id,omitempty" xml:"api_id"`

	// API名称
	ApiName *string `json:"api_name,omitempty" xml:"api_name"`

	// 发布的环境编号
	EnvId *string `json:"env_id,omitempty" xml:"env_id"`

	// 发布描述
	Remark *string `json:"remark,omitempty" xml:"remark"`

	// 发布时间
	PublishTime *sdktime.SdkTime `json:"publish_time,omitempty" xml:"publish_time"`

	// 在线的版本号
	VersionId *string `json:"version_id,omitempty" xml:"version_id"`
}

func (o PublishResp) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PublishResp struct{}"
	}

	return strings.Join([]string{"PublishResp", string(data)}, " ")
}
