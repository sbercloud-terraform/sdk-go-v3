package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// flavor数据结构说明
type FlavorInfos struct {

	// 裸金属服务器规格ID
	Id *string `json:"id,omitempty" xml:"id"`

	// 裸金属服务器规格名称
	Name *string `json:"name,omitempty" xml:"name"`

	// 该裸金属服务器规格对应要求系统盘大小，0为不限制。
	Disk *string `json:"disk,omitempty" xml:"disk"`

	// 该裸金属服务器规格对应的CPU核数
	Vcpus *string `json:"vcpus,omitempty" xml:"vcpus"`

	// 该裸金属服务器规格对应的内存大小，单位为MB
	Ram *string `json:"ram,omitempty" xml:"ram"`
}

func (o FlavorInfos) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "FlavorInfos struct{}"
	}

	return strings.Join([]string{"FlavorInfos", string(data)}, " ")
}
