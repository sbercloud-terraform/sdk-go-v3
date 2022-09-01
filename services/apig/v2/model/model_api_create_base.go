package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

type ApiCreateBase struct {

	// API名称。  长度为3 ~ 64位的字符串，字符串由中文、英文字母、数字、下划线组成，且只能以英文或中文开头。 > 中文字符必须为UTF-8或者unicode编码。
	Name string `json:"name" xml:"name"`

	// API类型 - 1：公有API - 2：私有API
	Type ApiCreateBaseType `json:"type" xml:"type"`

	// API的版本
	Version *string `json:"version,omitempty" xml:"version"`

	// API的请求协议 - HTTP - HTTPS - BOTH：同时支持HTTP和HTTPS
	ReqProtocol ApiCreateBaseReqProtocol `json:"req_protocol" xml:"req_protocol"`

	// API的请求方式
	ReqMethod ApiCreateBaseReqMethod `json:"req_method" xml:"req_method"`

	// 请求地址。可以包含请求参数，用{}标识，比如/getUserInfo/{userId}，支持 * % - _ . 等特殊字符，总长度不超过512，且满足URI规范。 > 需要服从URI规范。
	ReqUri string `json:"req_uri" xml:"req_uri"`

	// API的认证方式 - NONE：无认证 - APP：APP认证 - IAM：IAM认证 - AUTHORIZER：自定义认证
	AuthType ApiCreateBaseAuthType `json:"auth_type" xml:"auth_type"`

	AuthOpt *AuthOpt `json:"auth_opt,omitempty" xml:"auth_opt"`

	// 是否支持跨域 - TRUE：支持 - FALSE：不支持
	Cors *bool `json:"cors,omitempty" xml:"cors"`

	// API的匹配方式 - SWA：前缀匹配 - NORMAL：正常匹配（绝对匹配） 默认：NORMAL
	MatchMode *ApiCreateBaseMatchMode `json:"match_mode,omitempty" xml:"match_mode"`

	// 后端类型 - HTTP：web后端 - FUNCTION：函数工作流 - MOCK：模拟的后端
	BackendType ApiCreateBaseBackendType `json:"backend_type" xml:"backend_type"`

	// API描述。字符长度不超过255 > 中文字符必须为UTF-8或者unicode编码。
	Remark *string `json:"remark,omitempty" xml:"remark"`

	// API所属的分组编号
	GroupId string `json:"group_id" xml:"group_id"`

	// API请求体描述，可以是请求体示例、媒体类型、参数等信息。字符长度不超过20480 > 中文字符必须为UTF-8或者unicode编码。
	BodyRemark *string `json:"body_remark,omitempty" xml:"body_remark"`

	// 正常响应示例，描述API的正常返回信息。字符长度不超过20480 > 中文字符必须为UTF-8或者unicode编码。
	ResultNormalSample *string `json:"result_normal_sample,omitempty" xml:"result_normal_sample"`

	// 失败返回示例，描述API的异常返回信息。字符长度不超过20480 > 中文字符必须为UTF-8或者unicode编码。
	ResultFailureSample *string `json:"result_failure_sample,omitempty" xml:"result_failure_sample"`

	// 前端自定义认证对象的ID
	AuthorizerId *string `json:"authorizer_id,omitempty" xml:"authorizer_id"`

	// 标签。  支持英文，数字，下划线，且只能以英文开头。支持输入多个标签，不同标签以英文逗号分割。
	Tags *[]string `json:"tags,omitempty" xml:"tags"`

	// 分组自定义响应ID
	ResponseId *string `json:"response_id,omitempty" xml:"response_id"`

	// 集成应用ID  暂不支持
	RomaAppId *string `json:"roma_app_id,omitempty" xml:"roma_app_id"`

	// API绑定的自定义域名  暂不支持
	DomainName *string `json:"domain_name,omitempty" xml:"domain_name"`

	// 标签  待废弃，优先使用tags字段
	Tag *string `json:"tag,omitempty" xml:"tag"`

	// 请求内容格式类型：  application/json application/xml multipart/form-date text/plain  暂不支持
	ContentType *ApiCreateBaseContentType `json:"content_type,omitempty" xml:"content_type"`

	MockInfo *ApiMockCreate `json:"mock_info,omitempty" xml:"mock_info"`

	FuncInfo *ApiFuncCreate `json:"func_info,omitempty" xml:"func_info"`

	// API的请求参数列表
	ReqParams *[]ReqParamBase `json:"req_params,omitempty" xml:"req_params"`

	// API的后端参数列表
	BackendParams *[]BackendParamBase `json:"backend_params,omitempty" xml:"backend_params"`

	// mock策略后端列表
	PolicyMocks *[]ApiPolicyMockCreate `json:"policy_mocks,omitempty" xml:"policy_mocks"`

	// 函数工作流策略后端列表
	PolicyFunctions *[]ApiPolicyFunctionCreate `json:"policy_functions,omitempty" xml:"policy_functions"`
}

func (o ApiCreateBase) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ApiCreateBase struct{}"
	}

	return strings.Join([]string{"ApiCreateBase", string(data)}, " ")
}

type ApiCreateBaseType struct {
	value int32
}

type ApiCreateBaseTypeEnum struct {
	E_1 ApiCreateBaseType
	E_2 ApiCreateBaseType
}

func GetApiCreateBaseTypeEnum() ApiCreateBaseTypeEnum {
	return ApiCreateBaseTypeEnum{
		E_1: ApiCreateBaseType{
			value: 1,
		}, E_2: ApiCreateBaseType{
			value: 2,
		},
	}
}

func (c ApiCreateBaseType) Value() int32 {
	return c.value
}

func (c ApiCreateBaseType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ApiCreateBaseType) UnmarshalJSON(b []byte) error {
	myConverter := converter.StringConverterFactory("int32")
	if myConverter != nil {
		val, err := myConverter.CovertStringToInterface(strings.Trim(string(b[:]), "\""))
		if err == nil {
			c.value = val.(int32)
			return nil
		}
		return err
	} else {
		return errors.New("convert enum data to int32 error")
	}
}

type ApiCreateBaseReqProtocol struct {
	value string
}

type ApiCreateBaseReqProtocolEnum struct {
	HTTP  ApiCreateBaseReqProtocol
	HTTPS ApiCreateBaseReqProtocol
	BOTH  ApiCreateBaseReqProtocol
}

func GetApiCreateBaseReqProtocolEnum() ApiCreateBaseReqProtocolEnum {
	return ApiCreateBaseReqProtocolEnum{
		HTTP: ApiCreateBaseReqProtocol{
			value: "HTTP",
		},
		HTTPS: ApiCreateBaseReqProtocol{
			value: "HTTPS",
		},
		BOTH: ApiCreateBaseReqProtocol{
			value: "BOTH",
		},
	}
}

func (c ApiCreateBaseReqProtocol) Value() string {
	return c.value
}

func (c ApiCreateBaseReqProtocol) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ApiCreateBaseReqProtocol) UnmarshalJSON(b []byte) error {
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

type ApiCreateBaseReqMethod struct {
	value string
}

type ApiCreateBaseReqMethodEnum struct {
	GET     ApiCreateBaseReqMethod
	POST    ApiCreateBaseReqMethod
	PUT     ApiCreateBaseReqMethod
	DELETE  ApiCreateBaseReqMethod
	HEAD    ApiCreateBaseReqMethod
	PATCH   ApiCreateBaseReqMethod
	OPTIONS ApiCreateBaseReqMethod
	ANY     ApiCreateBaseReqMethod
}

func GetApiCreateBaseReqMethodEnum() ApiCreateBaseReqMethodEnum {
	return ApiCreateBaseReqMethodEnum{
		GET: ApiCreateBaseReqMethod{
			value: "GET",
		},
		POST: ApiCreateBaseReqMethod{
			value: "POST",
		},
		PUT: ApiCreateBaseReqMethod{
			value: "PUT",
		},
		DELETE: ApiCreateBaseReqMethod{
			value: "DELETE",
		},
		HEAD: ApiCreateBaseReqMethod{
			value: "HEAD",
		},
		PATCH: ApiCreateBaseReqMethod{
			value: "PATCH",
		},
		OPTIONS: ApiCreateBaseReqMethod{
			value: "OPTIONS",
		},
		ANY: ApiCreateBaseReqMethod{
			value: "ANY",
		},
	}
}

func (c ApiCreateBaseReqMethod) Value() string {
	return c.value
}

func (c ApiCreateBaseReqMethod) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ApiCreateBaseReqMethod) UnmarshalJSON(b []byte) error {
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

type ApiCreateBaseAuthType struct {
	value string
}

type ApiCreateBaseAuthTypeEnum struct {
	NONE       ApiCreateBaseAuthType
	APP        ApiCreateBaseAuthType
	IAM        ApiCreateBaseAuthType
	AUTHORIZER ApiCreateBaseAuthType
}

func GetApiCreateBaseAuthTypeEnum() ApiCreateBaseAuthTypeEnum {
	return ApiCreateBaseAuthTypeEnum{
		NONE: ApiCreateBaseAuthType{
			value: "NONE",
		},
		APP: ApiCreateBaseAuthType{
			value: "APP",
		},
		IAM: ApiCreateBaseAuthType{
			value: "IAM",
		},
		AUTHORIZER: ApiCreateBaseAuthType{
			value: "AUTHORIZER",
		},
	}
}

func (c ApiCreateBaseAuthType) Value() string {
	return c.value
}

func (c ApiCreateBaseAuthType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ApiCreateBaseAuthType) UnmarshalJSON(b []byte) error {
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

type ApiCreateBaseMatchMode struct {
	value string
}

type ApiCreateBaseMatchModeEnum struct {
	SWA    ApiCreateBaseMatchMode
	NORMAL ApiCreateBaseMatchMode
}

func GetApiCreateBaseMatchModeEnum() ApiCreateBaseMatchModeEnum {
	return ApiCreateBaseMatchModeEnum{
		SWA: ApiCreateBaseMatchMode{
			value: "SWA",
		},
		NORMAL: ApiCreateBaseMatchMode{
			value: "NORMAL",
		},
	}
}

func (c ApiCreateBaseMatchMode) Value() string {
	return c.value
}

func (c ApiCreateBaseMatchMode) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ApiCreateBaseMatchMode) UnmarshalJSON(b []byte) error {
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

type ApiCreateBaseBackendType struct {
	value string
}

type ApiCreateBaseBackendTypeEnum struct {
	HTTP     ApiCreateBaseBackendType
	FUNCTION ApiCreateBaseBackendType
	MOCK     ApiCreateBaseBackendType
}

func GetApiCreateBaseBackendTypeEnum() ApiCreateBaseBackendTypeEnum {
	return ApiCreateBaseBackendTypeEnum{
		HTTP: ApiCreateBaseBackendType{
			value: "HTTP",
		},
		FUNCTION: ApiCreateBaseBackendType{
			value: "FUNCTION",
		},
		MOCK: ApiCreateBaseBackendType{
			value: "MOCK",
		},
	}
}

func (c ApiCreateBaseBackendType) Value() string {
	return c.value
}

func (c ApiCreateBaseBackendType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ApiCreateBaseBackendType) UnmarshalJSON(b []byte) error {
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

type ApiCreateBaseContentType struct {
	value string
}

type ApiCreateBaseContentTypeEnum struct {
	APPLICATION_JSON    ApiCreateBaseContentType
	APPLICATION_XML     ApiCreateBaseContentType
	MULTIPART_FORM_DATE ApiCreateBaseContentType
	TEXT_PLAIN          ApiCreateBaseContentType
}

func GetApiCreateBaseContentTypeEnum() ApiCreateBaseContentTypeEnum {
	return ApiCreateBaseContentTypeEnum{
		APPLICATION_JSON: ApiCreateBaseContentType{
			value: "application/json",
		},
		APPLICATION_XML: ApiCreateBaseContentType{
			value: "application/xml",
		},
		MULTIPART_FORM_DATE: ApiCreateBaseContentType{
			value: "multipart/form-date",
		},
		TEXT_PLAIN: ApiCreateBaseContentType{
			value: "text/plain",
		},
	}
}

func (c ApiCreateBaseContentType) Value() string {
	return c.value
}

func (c ApiCreateBaseContentType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ApiCreateBaseContentType) UnmarshalJSON(b []byte) error {
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
