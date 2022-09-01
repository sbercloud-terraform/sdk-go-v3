package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type GetRepositoryByProjectIdResponse struct {
	Error *Error `json:"error,omitempty" xml:"error"`

	Result *RepoInfo `json:"result,omitempty" xml:"result"`

	// 响应状态
	Status         *string `json:"status,omitempty" xml:"status"`
	HttpStatusCode int     `json:"-"`
}

func (o GetRepositoryByProjectIdResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "GetRepositoryByProjectIdResponse struct{}"
	}

	return strings.Join([]string{"GetRepositoryByProjectIdResponse", string(data)}, " ")
}
