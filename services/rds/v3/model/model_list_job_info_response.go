package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type ListJobInfoResponse struct {
	Job            *GetJobInfoResponseBodyJob `json:"job,omitempty" xml:"job"`
	HttpStatusCode int                        `json:"-"`
}

func (o ListJobInfoResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListJobInfoResponse struct{}"
	}

	return strings.Join([]string{"ListJobInfoResponse", string(data)}, " ")
}
