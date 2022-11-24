package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type SubTaskRuntimeDto struct {

	// 作业子任务的并发实例名称
	SubTaskName *string `json:"sub_task_name,omitempty"`

	// 作业子任务的并发实例运行创建时间
	CreateTime *string `json:"create_time,omitempty"`

	// 作业子任务的并发实例运行结束时间
	FinishTime *string `json:"finish_time,omitempty"`

	// 作业子任务的并发实例实际运行时间
	ActualRunningTime *int32 `json:"actual_running_time,omitempty"`

	// 作业子任务的并发实例运行状态
	Status *string `json:"status,omitempty"`

	// 作业日志存储链接
	LogStorageLink *string `json:"log_storage_link,omitempty"`
}

func (o SubTaskRuntimeDto) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SubTaskRuntimeDto struct{}"
	}

	return strings.Join([]string{"SubTaskRuntimeDto", string(data)}, " ")
}
