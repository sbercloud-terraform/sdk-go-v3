package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

type CloudWafHostItem struct {

	// 域名id
	Id *string `json:"id,omitempty" xml:"id"`

	// 域名id
	Hostid *string `json:"hostid,omitempty" xml:"hostid"`

	// 华为云区域ID，控制台创建的域名会携带此参数，api调用创建的域名此参数为空，可以通过华为云上地区和终端节点文档查询区域ID对应的中文名称
	Region *string `json:"region,omitempty" xml:"region"`

	// 域名描述信息，可选参数。
	Description *string `json:"description,omitempty" xml:"description"`

	// WAF部署模式，默认是1，目前仅支持反代模式，冗余参数
	Type *int32 `json:"type,omitempty" xml:"type"`

	// 防护域名是否使用代理   - false：不使用代理   - true：使用代理
	Proxy *bool `json:"proxy,omitempty" xml:"proxy"`

	// 创建的云模式防护域名
	Hostname *string `json:"hostname,omitempty" xml:"hostname"`

	// cname前缀
	AccessCode *string `json:"access_code,omitempty" xml:"access_code"`

	// 防护策略id
	Policyid *string `json:"policyid,omitempty" xml:"policyid"`

	// 创建防护域名的时间
	Timestamp *int64 `json:"timestamp,omitempty" xml:"timestamp"`

	// 域名防护状态：  - -1：bypass，该域名的请求直接到达其后端服务器，不再经过WAF  - 0：暂停防护，WAF只转发该域名的请求，不做攻击检测  - 1：开启防护，WAF根据您配置的策略进行攻击检测
	ProtectStatus *int32 `json:"protect_status,omitempty" xml:"protect_status"`

	// 域名接入状态，0表示未接入，1表示已接入
	AccessStatus *int32 `json:"access_status,omitempty" xml:"access_status"`

	// 是否使用独享ip   - true：使用独享ip   - false：不实用独享ip
	ExclusiveIp *bool `json:"exclusive_ip,omitempty" xml:"exclusive_ip"`

	// 付费模式，目前只支持prePaid预付款模式
	PaidType *CloudWafHostItemPaidType `json:"paid_type,omitempty" xml:"paid_type"`

	// 网站名称，对应WAF控制台域名详情中的网站名称
	WebTag *string `json:"web_tag,omitempty" xml:"web_tag"`

	Flag *Flag `json:"flag,omitempty" xml:"flag"`
}

func (o CloudWafHostItem) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CloudWafHostItem struct{}"
	}

	return strings.Join([]string{"CloudWafHostItem", string(data)}, " ")
}

type CloudWafHostItemPaidType struct {
	value string
}

type CloudWafHostItemPaidTypeEnum struct {
	PRE_PAID CloudWafHostItemPaidType
}

func GetCloudWafHostItemPaidTypeEnum() CloudWafHostItemPaidTypeEnum {
	return CloudWafHostItemPaidTypeEnum{
		PRE_PAID: CloudWafHostItemPaidType{
			value: "prePaid",
		},
	}
}

func (c CloudWafHostItemPaidType) Value() string {
	return c.value
}

func (c CloudWafHostItemPaidType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *CloudWafHostItemPaidType) UnmarshalJSON(b []byte) error {
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
