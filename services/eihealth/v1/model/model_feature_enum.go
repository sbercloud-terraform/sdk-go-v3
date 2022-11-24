package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

type FeatureEnum struct {
	value string
}

type FeatureEnumEnum struct {
	TOOL FeatureEnum
}

func GetFeatureEnumEnum() FeatureEnumEnum {
	return FeatureEnumEnum{
		TOOL: FeatureEnum{
			value: "TOOL",
		},
	}
}

func (c FeatureEnum) Value() string {
	return c.value
}

func (c FeatureEnum) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *FeatureEnum) UnmarshalJSON(b []byte) error {
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
