package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type CreateUserResponse struct {

	// 用户id
	Id             *string `json:"id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o CreateUserResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateUserResponse struct{}"
	}

	return strings.Join([]string{"CreateUserResponse", string(data)}, " ")
}
