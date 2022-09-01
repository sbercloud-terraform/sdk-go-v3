package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type ListEnvTagsResponse struct {

	// 环境标签数据模型
	EnvTags *[]CmdbTagEntity `json:"env_tags,omitempty" xml:"env_tags"`

	// 总条数
	TotalCount     *int32 `json:"total_count,omitempty" xml:"total_count"`
	HttpStatusCode int    `json:"-"`
}

func (o ListEnvTagsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListEnvTagsResponse struct{}"
	}

	return strings.Join([]string{"ListEnvTagsResponse", string(data)}, " ")
}
