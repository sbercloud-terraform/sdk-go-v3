package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type AddingBackendInstancesV2Response struct {

	// 本次返回的列表长度
	Size int32 `json:"size" xml:"size"`

	// 满足条件的记录数
	Total int64 `json:"total" xml:"total"`

	// 本次查询到的云服务器列表
	Members        *[]VpcMemberInfo `json:"members,omitempty" xml:"members"`
	HttpStatusCode int              `json:"-"`
}

func (o AddingBackendInstancesV2Response) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AddingBackendInstancesV2Response struct{}"
	}

	return strings.Join([]string{"AddingBackendInstancesV2Response", string(data)}, " ")
}
