package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type UpdateIssueV4Response struct {

	// 实际工时
	ActualWorkHours *float64 `json:"actual_work_hours,omitempty" xml:"actual_work_hours"`

	// 抄送人
	AssignedCcUser *[]IssueUser `json:"assigned_cc_user,omitempty" xml:"assigned_cc_user"`

	AssignedUser *IssueUser `json:"assigned_user,omitempty" xml:"assigned_user"`

	// 预计开始时间，年-月-日
	BeginTime *string `json:"begin_time,omitempty" xml:"begin_time"`

	// 创建时间 年-月-日 时:分:秒
	CreatedTime *string `json:"created_time,omitempty" xml:"created_time"`

	Creator *IssueUser `json:"creator,omitempty" xml:"creator"`

	// 自定义属性值,不建议使用，建议参考new_custom_fields字段
	CustomFields *[]CustomField `json:"custom_fields,omitempty" xml:"custom_fields"`

	// 自定义属性值
	NewCustomFields *[]NewCustomField `json:"new_custom_fields,omitempty" xml:"new_custom_fields"`

	Developer *IssueUser `json:"developer,omitempty" xml:"developer"`

	Domain *IssueItemSfV4Domain `json:"domain,omitempty" xml:"domain"`

	// 工作项进度值
	DoneRatio *int32 `json:"done_ratio,omitempty" xml:"done_ratio"`

	// 预计结束时间，年-月-日
	EndTime *string `json:"end_time,omitempty" xml:"end_time"`

	// 预计工时
	ExpectedWorkHours *float64 `json:"expected_work_hours,omitempty" xml:"expected_work_hours"`

	// 工作项项id
	Id *int32 `json:"id,omitempty" xml:"id"`

	Project *IssueProjectResponseV4 `json:"project,omitempty" xml:"project"`

	Iteration *IssueItemSfV4Iteration `json:"iteration,omitempty" xml:"iteration"`

	Module *IssueItemSfV4Module `json:"module,omitempty" xml:"module"`

	// 标题
	Name *string `json:"name,omitempty" xml:"name"`

	ParentIssue *CreateIssueResponseV4ParentIssue `json:"parent_issue,omitempty" xml:"parent_issue"`

	Priority *IssueItemSfV4Priority `json:"priority,omitempty" xml:"priority"`

	Severity *IssueItemSfV4Severity `json:"severity,omitempty" xml:"severity"`

	Status *IssueItemSfV4Status `json:"status,omitempty" xml:"status"`

	Tracker *CreateIssueResponseV4Tracker `json:"tracker,omitempty" xml:"tracker"`

	// 更新时间 年-月-日 时:分:秒
	UpdatedTime *string `json:"updated_time,omitempty" xml:"updated_time"`

	// 关闭时间 年-月-日 时:分:秒
	ClosedTime     *string `json:"closed_time,omitempty" xml:"closed_time"`
	HttpStatusCode int     `json:"-"`
}

func (o UpdateIssueV4Response) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateIssueV4Response struct{}"
	}

	return strings.Join([]string{"UpdateIssueV4Response", string(data)}, " ")
}
