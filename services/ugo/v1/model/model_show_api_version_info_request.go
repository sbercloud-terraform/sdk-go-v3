package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// Request Object
type ShowApiVersionInfoRequest struct {

	// 需要查询的API版本号。
	ApiVersion ShowApiVersionInfoRequestApiVersion `json:"api_version"`
}

func (o ShowApiVersionInfoRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowApiVersionInfoRequest struct{}"
	}

	return strings.Join([]string{"ShowApiVersionInfoRequest", string(data)}, " ")
}

type ShowApiVersionInfoRequestApiVersion struct {
	value string
}

type ShowApiVersionInfoRequestApiVersionEnum struct {
	V1 ShowApiVersionInfoRequestApiVersion
}

func GetShowApiVersionInfoRequestApiVersionEnum() ShowApiVersionInfoRequestApiVersionEnum {
	return ShowApiVersionInfoRequestApiVersionEnum{
		V1: ShowApiVersionInfoRequestApiVersion{
			value: "v1",
		},
	}
}

func (c ShowApiVersionInfoRequestApiVersion) Value() string {
	return c.value
}

func (c ShowApiVersionInfoRequestApiVersion) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ShowApiVersionInfoRequestApiVersion) UnmarshalJSON(b []byte) error {
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
