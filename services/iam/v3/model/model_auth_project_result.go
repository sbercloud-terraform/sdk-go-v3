package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

//
type AuthProjectResult struct {

	// false.
	IsDomain bool `json:"is_domain" xml:"is_domain"`

	// 项目描述信息。
	Description string `json:"description" xml:"description"`

	Links *LinksSelf `json:"links" xml:"links"`

	// 项目是否可用。
	Enabled bool `json:"enabled" xml:"enabled"`

	// 项目ID。
	Id string `json:"id" xml:"id"`

	// 如果查询自己创建的项目，则此处返回所属区域的项目ID。  如果查询的是系统内置项目，如cn-north-4，则此处返回账号ID。
	ParentId string `json:"parent_id" xml:"parent_id"`

	// 项目所属账号ID。
	DomainId string `json:"domain_id" xml:"domain_id"`

	// 项目名称。
	Name string `json:"name" xml:"name"`
}

func (o AuthProjectResult) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AuthProjectResult struct{}"
	}

	return strings.Join([]string{"AuthProjectResult", string(data)}, " ")
}
