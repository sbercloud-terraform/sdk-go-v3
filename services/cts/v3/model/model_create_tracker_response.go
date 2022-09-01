package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// Response Object
type CreateTrackerResponse struct {

	// 追踪器唯一标识。
	Id *string `json:"id,omitempty" xml:"id"`

	// 追踪器创建时间戳。
	CreateTime *int64 `json:"create_time,omitempty" xml:"create_time"`

	// 事件文件转储加密所采用的秘钥id（从KMS获取）。 当\"tracker_type\"参数值为\"system\"和\"is_support_trace_files_encryption\"参数值为“是”时，此参数为必选项。
	KmsId *string `json:"kms_id,omitempty" xml:"kms_id"`

	// 是否打开事件文件校验。当前环境仅\"tracker_type\"参数值为\"system\"时支持该功能。
	IsSupportValidate *bool `json:"is_support_validate,omitempty" xml:"is_support_validate"`

	Lts *Lts `json:"lts,omitempty" xml:"lts"`

	// 标识追踪器类型。 目前支持系统追踪器类型有管理类追踪器（system）和数据类追踪器（data）。
	TrackerType *CreateTrackerResponseTrackerType `json:"tracker_type,omitempty" xml:"tracker_type"`

	// 账号ID，参见《云审计服务API参考》“获取账号ID和项目ID”章节。
	DomainId *string `json:"domain_id,omitempty" xml:"domain_id"`

	// 项目ID。
	ProjectId *string `json:"project_id,omitempty" xml:"project_id"`

	// 标识追踪器名称，当前版本默认为“system”。
	TrackerName *string `json:"tracker_name,omitempty" xml:"tracker_name"`

	// 标识追踪器状态，包括正常（enabled），停止（disabled）和异常（error）三种状态，状态为异常时需通过明细（detail）字段说明错误来源。
	Status *CreateTrackerResponseStatus `json:"status,omitempty" xml:"status"`

	// 该参数仅在追踪器状态异常时返回，用于标识追踪器异常的原因，包括桶策略异常（bucketPolicyError），桶不存在（noBucket）和欠费或冻结（arrears）三种原因。
	Detail *string `json:"detail,omitempty" xml:"detail"`

	// 事件文件转储加密功能开关。 该参数必须与kms_id参数同时使用。 当前环境仅\"tracker_type\"参数值为\"system\"时支持该功能。
	IsSupportTraceFilesEncryption *bool `json:"is_support_trace_files_encryption,omitempty" xml:"is_support_trace_files_encryption"`

	ObsInfo *ObsInfo `json:"obs_info,omitempty" xml:"obs_info"`

	DataBucket     *DataBucketQuery `json:"data_bucket,omitempty" xml:"data_bucket"`
	HttpStatusCode int              `json:"-"`
}

func (o CreateTrackerResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateTrackerResponse struct{}"
	}

	return strings.Join([]string{"CreateTrackerResponse", string(data)}, " ")
}

type CreateTrackerResponseTrackerType struct {
	value string
}

type CreateTrackerResponseTrackerTypeEnum struct {
	SYSTEM CreateTrackerResponseTrackerType
	DATA   CreateTrackerResponseTrackerType
}

func GetCreateTrackerResponseTrackerTypeEnum() CreateTrackerResponseTrackerTypeEnum {
	return CreateTrackerResponseTrackerTypeEnum{
		SYSTEM: CreateTrackerResponseTrackerType{
			value: "system",
		},
		DATA: CreateTrackerResponseTrackerType{
			value: "data",
		},
	}
}

func (c CreateTrackerResponseTrackerType) Value() string {
	return c.value
}

func (c CreateTrackerResponseTrackerType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *CreateTrackerResponseTrackerType) UnmarshalJSON(b []byte) error {
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

type CreateTrackerResponseStatus struct {
	value string
}

type CreateTrackerResponseStatusEnum struct {
	ENABLED  CreateTrackerResponseStatus
	DISABLED CreateTrackerResponseStatus
}

func GetCreateTrackerResponseStatusEnum() CreateTrackerResponseStatusEnum {
	return CreateTrackerResponseStatusEnum{
		ENABLED: CreateTrackerResponseStatus{
			value: "enabled",
		},
		DISABLED: CreateTrackerResponseStatus{
			value: "disabled",
		},
	}
}

func (c CreateTrackerResponseStatus) Value() string {
	return c.value
}

func (c CreateTrackerResponseStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *CreateTrackerResponseStatus) UnmarshalJSON(b []byte) error {
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
