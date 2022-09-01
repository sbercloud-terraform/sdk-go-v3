package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type UpdateStructConfigResponse struct {
	Body           *string `json:"body,omitempty" xml:"body"`
	HttpStatusCode int     `json:"-"`
}

func (o UpdateStructConfigResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateStructConfigResponse struct{}"
	}

	return strings.Join([]string{"UpdateStructConfigResponse", string(data)}, " ")
}
