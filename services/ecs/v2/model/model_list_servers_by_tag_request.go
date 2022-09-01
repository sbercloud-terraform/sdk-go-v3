package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type ListServersByTagRequest struct {
	Body *ListServersByTagRequestBody `json:"body,omitempty" xml:"body"`
}

func (o ListServersByTagRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListServersByTagRequest struct{}"
	}

	return strings.Join([]string{"ListServersByTagRequest", string(data)}, " ")
}
