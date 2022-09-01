package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type ListAssociatedTestCasesRequest struct {

	// devcloud项目的32位id
	ProjectId string `json:"project_id" xml:"project_id"`

	// 工作项ID
	IssueId int32 `json:"issue_id" xml:"issue_id"`

	// 每页数量
	Limit *int32 `json:"limit,omitempty" xml:"limit"`

	// 偏移量
	Offset *int32 `json:"offset,omitempty" xml:"offset"`
}

func (o ListAssociatedTestCasesRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListAssociatedTestCasesRequest struct{}"
	}

	return strings.Join([]string{"ListAssociatedTestCasesRequest", string(data)}, " ")
}
