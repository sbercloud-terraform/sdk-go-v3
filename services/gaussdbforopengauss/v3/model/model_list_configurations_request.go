package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// Request Object
type ListConfigurationsRequest struct {

	// 语言
	XLanguage *ListConfigurationsRequestXLanguage `json:"X-Language,omitempty" xml:"X-Language"`

	// 索引位置，偏移量。从第一条数据偏移offset条数据后开始查询，默认为0（偏移0条数据，表示从第一条数据开始查询），必须为数字，不能为负数。
	Offset *int32 `json:"offset,omitempty" xml:"offset"`

	// 查询记录数。默认为100，不能为负数，最小值为1，最大值为100。
	Limit *int32 `json:"limit,omitempty" xml:"limit"`
}

func (o ListConfigurationsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListConfigurationsRequest struct{}"
	}

	return strings.Join([]string{"ListConfigurationsRequest", string(data)}, " ")
}

type ListConfigurationsRequestXLanguage struct {
	value string
}

type ListConfigurationsRequestXLanguageEnum struct {
	ZH_CN ListConfigurationsRequestXLanguage
	EN_US ListConfigurationsRequestXLanguage
}

func GetListConfigurationsRequestXLanguageEnum() ListConfigurationsRequestXLanguageEnum {
	return ListConfigurationsRequestXLanguageEnum{
		ZH_CN: ListConfigurationsRequestXLanguage{
			value: "zh-cn",
		},
		EN_US: ListConfigurationsRequestXLanguage{
			value: "en-us",
		},
	}
}

func (c ListConfigurationsRequestXLanguage) Value() string {
	return c.value
}

func (c ListConfigurationsRequestXLanguage) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListConfigurationsRequestXLanguage) UnmarshalJSON(b []byte) error {
	myConverter := converter.StringConverterFactory("string")
	if myConverter != nil {
		val, err := myConverter.CovertStringToInterface(strings.Trim(string(b[:]), "\""))
		if err == nil {
			c.value = val.(string)
			return nil
		}
		return err
	} else {
		return errors.New("convert enum data to string error")
	}
}
