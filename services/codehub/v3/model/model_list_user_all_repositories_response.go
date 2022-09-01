package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type ListUserAllRepositoriesResponse struct {
	Error *Error `json:"error,omitempty" xml:"error"`

	Result *RepoListInfoV2 `json:"result,omitempty" xml:"result"`

	// 响应状态
	Status         *string `json:"status,omitempty" xml:"status"`
	HttpStatusCode int     `json:"-"`
}

func (o ListUserAllRepositoriesResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListUserAllRepositoriesResponse struct{}"
	}

	return strings.Join([]string{"ListUserAllRepositoriesResponse", string(data)}, " ")
}
