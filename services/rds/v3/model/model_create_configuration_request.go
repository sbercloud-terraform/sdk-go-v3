package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// Request Object
type CreateConfigurationRequest struct {

	// 语言
	XLanguage *CreateConfigurationRequestXLanguage `json:"X-Language,omitempty" xml:"X-Language"`

	Body *ConfigurationForCreation `json:"body,omitempty" xml:"body"`
}

func (o CreateConfigurationRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateConfigurationRequest struct{}"
	}

	return strings.Join([]string{"CreateConfigurationRequest", string(data)}, " ")
}

type CreateConfigurationRequestXLanguage struct {
	value string
}

type CreateConfigurationRequestXLanguageEnum struct {
	ZH_CN CreateConfigurationRequestXLanguage
	EN_US CreateConfigurationRequestXLanguage
}

func GetCreateConfigurationRequestXLanguageEnum() CreateConfigurationRequestXLanguageEnum {
	return CreateConfigurationRequestXLanguageEnum{
		ZH_CN: CreateConfigurationRequestXLanguage{
			value: "zh-cn",
		},
		EN_US: CreateConfigurationRequestXLanguage{
			value: "en-us",
		},
	}
}

func (c CreateConfigurationRequestXLanguage) Value() string {
	return c.value
}

func (c CreateConfigurationRequestXLanguage) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *CreateConfigurationRequestXLanguage) UnmarshalJSON(b []byte) error {
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
