package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type BatchAddDeleteTagsResponse struct {
	Body           *string `json:"body,omitempty" xml:"body"`
	HttpStatusCode int     `json:"-"`
}

func (o BatchAddDeleteTagsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchAddDeleteTagsResponse struct{}"
	}

	return strings.Join([]string{"BatchAddDeleteTagsResponse", string(data)}, " ")
}
