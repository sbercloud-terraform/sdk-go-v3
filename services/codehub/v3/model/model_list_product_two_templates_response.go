package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type ListProductTwoTemplatesResponse struct {
	Error *Error `json:"error,omitempty" xml:"error"`

	Result *TemplateListInfo `json:"result,omitempty" xml:"result"`

	// 响应状态
	Status         *string `json:"status,omitempty" xml:"status"`
	HttpStatusCode int     `json:"-"`
}

func (o ListProductTwoTemplatesResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListProductTwoTemplatesResponse struct{}"
	}

	return strings.Join([]string{"ListProductTwoTemplatesResponse", string(data)}, " ")
}
