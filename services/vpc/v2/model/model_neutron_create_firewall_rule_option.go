package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

//
type NeutronCreateFirewallRuleOption struct {

	// 功能说明：网络ACL规则名称 取值范围：0-255个字符
	Name *string `json:"name,omitempty" xml:"name"`

	// 功能说明：网络ACL规则描述 取值范围：0-255个字符
	Description *string `json:"description,omitempty" xml:"description"`

	// 功能说明：IP协议 取值范围：支持TCP,UDP,ICMP, ICMPV6或者ip协议号（0-255）
	Protocol *string `json:"protocol,omitempty" xml:"protocol"`

	// 功能说明：对通过网络ACL的流量执行的操作 取值范围：DENY（拒绝）/ALLOW（允许）
	Action *NeutronCreateFirewallRuleOptionAction `json:"action,omitempty" xml:"action"`

	// 功能说明：IP协议版本
	IpVersion *int32 `json:"ip_version,omitempty" xml:"ip_version"`

	// 功能说明：目的地址或者CIDR
	DestinationIpAddress *string `json:"destination_ip_address,omitempty" xml:"destination_ip_address"`

	// 功能说明：目的端口号或者一段端口范围
	DestinationPort *string `json:"destination_port,omitempty" xml:"destination_port"`

	// 功能说明：源地址或者CIDR
	SourceIpAddress *string `json:"source_ip_address,omitempty" xml:"source_ip_address"`

	// 功能说明：源端口号或者一段端口范围
	SourcePort *string `json:"source_port,omitempty" xml:"source_port"`

	// 功能说明：是否使能网络ACL防火墙规则。
	Enabled *bool `json:"enabled,omitempty" xml:"enabled"`
}

func (o NeutronCreateFirewallRuleOption) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "NeutronCreateFirewallRuleOption struct{}"
	}

	return strings.Join([]string{"NeutronCreateFirewallRuleOption", string(data)}, " ")
}

type NeutronCreateFirewallRuleOptionAction struct {
	value string
}

type NeutronCreateFirewallRuleOptionActionEnum struct {
	DENY  NeutronCreateFirewallRuleOptionAction
	ALLOW NeutronCreateFirewallRuleOptionAction
}

func GetNeutronCreateFirewallRuleOptionActionEnum() NeutronCreateFirewallRuleOptionActionEnum {
	return NeutronCreateFirewallRuleOptionActionEnum{
		DENY: NeutronCreateFirewallRuleOptionAction{
			value: "DENY",
		},
		ALLOW: NeutronCreateFirewallRuleOptionAction{
			value: "ALLOW",
		},
	}
}

func (c NeutronCreateFirewallRuleOptionAction) Value() string {
	return c.value
}

func (c NeutronCreateFirewallRuleOptionAction) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *NeutronCreateFirewallRuleOptionAction) UnmarshalJSON(b []byte) error {
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
