package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

//
type Domains struct {

	// 是否启用账号，true为启用，false为停用，默认为true。
	Enabled bool `json:"enabled" xml:"enabled"`

	// 账号ID。
	Id string `json:"id" xml:"id"`

	// 账号名。
	Name string `json:"name" xml:"name"`

	Links *LinksSelf `json:"links" xml:"links"`

	// 账号的描述信息。
	Description string `json:"description" xml:"description"`
}

func (o Domains) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Domains struct{}"
	}

	return strings.Join([]string{"Domains", string(data)}, " ")
}
