package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type ListSyntaxConversionProgressResponse struct {

	// 对象总数。
	TotalObjectsCount *int32 `json:"total_objects_count,omitempty" xml:"total_objects_count"`

	// 完成语法转换的对象数量。
	CompletedObjectsCount *int32 `json:"completed_objects_count,omitempty" xml:"completed_objects_count"`

	// 语法转换的对象列表。
	ObjectsList    *[]DatabaseObject `json:"objects_list,omitempty" xml:"objects_list"`
	HttpStatusCode int               `json:"-"`
}

func (o ListSyntaxConversionProgressResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListSyntaxConversionProgressResponse struct{}"
	}

	return strings.Join([]string{"ListSyntaxConversionProgressResponse", string(data)}, " ")
}
