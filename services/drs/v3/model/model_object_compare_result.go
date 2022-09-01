package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

//
type ObjectCompareResult struct {

	// 对象级对比任务的id。
	CompareTaskId string `json:"compare_task_id" xml:"compare_task_id"`

	// 对象对比结果概览。
	ObjectCompareOverview *[]ObjectCompareResultOverview `json:"object_compare_overview,omitempty" xml:"object_compare_overview"`

	// 对象对比结果详情。KEY值为对象对比结果概览中的对象类型。
	ObjectCompareDetails map[string][]ObjectCompareResultDetails `json:"object_compare_details,omitempty" xml:"object_compare_details"`

	// 错误码。
	ErrorCode *string `json:"error_code,omitempty" xml:"error_code"`

	// 错误信息。
	ErrorMsg *string `json:"error_msg,omitempty" xml:"error_msg"`
}

func (o ObjectCompareResult) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ObjectCompareResult struct{}"
	}

	return strings.Join([]string{"ObjectCompareResult", string(data)}, " ")
}
