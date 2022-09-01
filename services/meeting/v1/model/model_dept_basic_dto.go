package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 部门基本信息, 查询企业级别的管理员时需要显示部门信息
type DeptBasicDto struct {

	// 部门编码
	DeptCode *string `json:"deptCode,omitempty" xml:"deptCode"`

	// 企业id
	CorpId *string `json:"corpId,omitempty" xml:"corpId"`

	// 部门名称
	DeptName *string `json:"deptName,omitempty" xml:"deptName"`

	// 部门名称路径
	DeptNamePath *string `json:"deptNamePath,omitempty" xml:"deptNamePath"`

	// 父部门编码
	ParentDeptCode *string `json:"parentDeptCode,omitempty" xml:"parentDeptCode"`
}

func (o DeptBasicDto) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeptBasicDto struct{}"
	}

	return strings.Join([]string{"DeptBasicDto", string(data)}, " ")
}
