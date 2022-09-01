package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type SchemaPath1 struct {

	// 导入OBS文件对应的jobId。
	JobId string `json:"jobId" xml:"jobId"`

	// OBS存储路径，不包含OBS Endpoint。
	Path string `json:"path" xml:"path"`

	// OBS文件导入状态。 - success：完全导入成功 - partiallyFailed：部分失败 - failed：完全导入失败
	Status string `json:"status" xml:"status"`
}

func (o SchemaPath1) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SchemaPath1 struct{}"
	}

	return strings.Join([]string{"SchemaPath1", string(data)}, " ")
}
