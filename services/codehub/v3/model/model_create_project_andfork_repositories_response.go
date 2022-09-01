package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type CreateProjectAndforkRepositoriesResponse struct {
	Error *Error `json:"error,omitempty" xml:"error"`

	Result *ProjectRepository `json:"result,omitempty" xml:"result"`

	// 响应状态
	Status         *string `json:"status,omitempty" xml:"status"`
	HttpStatusCode int     `json:"-"`
}

func (o CreateProjectAndforkRepositoriesResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateProjectAndforkRepositoriesResponse struct{}"
	}

	return strings.Join([]string{"CreateProjectAndforkRepositoriesResponse", string(data)}, " ")
}
