package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type RunNerResponse struct {

	// 命名实体识别结果。调用失败时无此字段。
	NamedEntities *[]NamedEntity `json:"named_entities,omitempty" xml:"named_entities"`

	// 调用失败时的错误码，具体请参见错误码。调用成功时无此字段。
	ErrorCode *string `json:"error_code,omitempty" xml:"error_code"`

	// 调用失败时的错误信息。调用成功时无此字段。
	ErrorMsg       *string `json:"error_msg,omitempty" xml:"error_msg"`
	HttpStatusCode int     `json:"-"`
}

func (o RunNerResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RunNerResponse struct{}"
	}

	return strings.Join([]string{"RunNerResponse", string(data)}, " ")
}
