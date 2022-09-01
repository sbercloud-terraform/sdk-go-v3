package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type ImportTrafficRequest struct {
	Body *ImportTrafficRequestBody `json:"body,omitempty" xml:"body"`
}

func (o ImportTrafficRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ImportTrafficRequest struct{}"
	}

	return strings.Join([]string{"ImportTrafficRequest", string(data)}, " ")
}
