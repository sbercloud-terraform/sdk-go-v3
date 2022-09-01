package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

type BatchDeleteInstanceRespResults struct {

	// 操作结果。
	Result *BatchDeleteInstanceRespResultsResult `json:"result,omitempty" xml:"result"`

	// 实例ID。
	Instance *string `json:"instance,omitempty" xml:"instance"`
}

func (o BatchDeleteInstanceRespResults) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchDeleteInstanceRespResults struct{}"
	}

	return strings.Join([]string{"BatchDeleteInstanceRespResults", string(data)}, " ")
}

type BatchDeleteInstanceRespResultsResult struct {
	value string
}

type BatchDeleteInstanceRespResultsResultEnum struct {
	SUCCESS BatchDeleteInstanceRespResultsResult
	FAILED  BatchDeleteInstanceRespResultsResult
}

func GetBatchDeleteInstanceRespResultsResultEnum() BatchDeleteInstanceRespResultsResultEnum {
	return BatchDeleteInstanceRespResultsResultEnum{
		SUCCESS: BatchDeleteInstanceRespResultsResult{
			value: "success",
		},
		FAILED: BatchDeleteInstanceRespResultsResult{
			value: "failed",
		},
	}
}

func (c BatchDeleteInstanceRespResultsResult) Value() string {
	return c.value
}

func (c BatchDeleteInstanceRespResultsResult) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *BatchDeleteInstanceRespResultsResult) UnmarshalJSON(b []byte) error {
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
