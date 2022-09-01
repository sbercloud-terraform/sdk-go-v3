package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type CreateManualBackupResponse struct {
	Backup *BackupInfo `json:"backup,omitempty" xml:"backup"`

	// 任务id。
	JobId          *string `json:"job_id,omitempty" xml:"job_id"`
	HttpStatusCode int     `json:"-"`
}

func (o CreateManualBackupResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateManualBackupResponse struct{}"
	}

	return strings.Join([]string{"CreateManualBackupResponse", string(data)}, " ")
}
