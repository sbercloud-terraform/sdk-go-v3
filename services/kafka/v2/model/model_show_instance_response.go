package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// Response Object
type ShowInstanceResponse struct {

	// 实例名称。
	Name *string `json:"name,omitempty" xml:"name"`

	// 引擎。
	Engine *string `json:"engine,omitempty" xml:"engine"`

	// 版本。
	EngineVersion *string `json:"engine_version,omitempty" xml:"engine_version"`

	// 实例描述。
	Description *string `json:"description,omitempty" xml:"description"`

	// 实例规格。
	Specification *string `json:"specification,omitempty" xml:"specification"`

	// 消息存储空间，单位：GB。
	StorageSpace *int32 `json:"storage_space,omitempty" xml:"storage_space"`

	// Kafka实例的分区数量。
	PartitionNum *string `json:"partition_num,omitempty" xml:"partition_num"`

	// 已使用的消息存储空间，单位：GB。
	UsedStorageSpace *int32 `json:"used_storage_space,omitempty" xml:"used_storage_space"`

	// 实例连接IP地址。
	ConnectAddress *string `json:"connect_address,omitempty" xml:"connect_address"`

	// 实例连接端口。
	Port *int32 `json:"port,omitempty" xml:"port"`

	// 实例的状态。 [详细状态说明见[实例状态说明](https://support.huaweicloud.com/api-kafka/kafka-api-180514012.html)。](tag:hws)[详细状态说明见[实例状态说明](https://support.huaweicloud.com/intl/zh-cn/api-kafka/kafka-api-180514012.html)。](tag:hws_hk)
	Status *string `json:"status,omitempty" xml:"status"`

	// 实例ID。
	InstanceId *string `json:"instance_id,omitempty" xml:"instance_id"`

	// 资源规格标识。   - dms.instance.kafka.cluster.c3.mini：Kafka实例的基准带宽为100MByte/秒。   - dms.instance.kafka.cluster.c3.small.2：Kafka实例的基准带宽为300MByte/秒。   - dms.instance.kafka.cluster.c3.middle.2：Kafka实例的基准带宽为600MByte/秒。   - dms.instance.kafka.cluster.c3.high.2：Kafka实例的基准带宽为1200MByte/秒。
	ResourceSpecCode *string `json:"resource_spec_code,omitempty" xml:"resource_spec_code"`

	// 付费模式，1表示按需计费，0表示包年/包月计费。
	ChargingMode *int32 `json:"charging_mode,omitempty" xml:"charging_mode"`

	// VPC ID。
	VpcId *string `json:"vpc_id,omitempty" xml:"vpc_id"`

	// VPC的名称。
	VpcName *string `json:"vpc_name,omitempty" xml:"vpc_name"`

	// 完成创建时间。  格式为时间戳，指从格林威治时间 1970年01月01日00时00分00秒起至指定时间的偏差总毫秒数。
	CreatedAt *string `json:"created_at,omitempty" xml:"created_at"`

	// 用户ID。
	UserId *string `json:"user_id,omitempty" xml:"user_id"`

	// 用户名。
	UserName *string `json:"user_name,omitempty" xml:"user_name"`

	// 实例访问用户名。
	AccessUser *string `json:"access_user,omitempty" xml:"access_user"`

	// 订单ID，只有在包周期计费时才会有order_id值，其他计费方式order_id值为空。
	OrderId *string `json:"order_id,omitempty" xml:"order_id"`

	// 维护时间窗开始时间，格式为HH:mm:ss。
	MaintainBegin *string `json:"maintain_begin,omitempty" xml:"maintain_begin"`

	// 维护时间窗结束时间，格式为HH:mm:ss。
	MaintainEnd *string `json:"maintain_end,omitempty" xml:"maintain_end"`

	// 实例是否开启公网访问功能。 - true：开启 - false：未开启
	EnablePublicip *bool `json:"enable_publicip,omitempty" xml:"enable_publicip"`

	// Kafka实例的Kafka Manager连接地址。
	ManagementConnectAddress *string `json:"management_connect_address,omitempty" xml:"management_connect_address"`

	// 是否开启安全认证。 - true：开启 - false：未开启
	SslEnable *bool `json:"ssl_enable,omitempty" xml:"ssl_enable"`

	// 是否开启双向认证。
	SslTwoWayEnable *bool `json:"ssl_two_way_enable,omitempty" xml:"ssl_two_way_enable"`

	// 是否能够证书替换。
	CertReplaced *bool `json:"cert_replaced,omitempty" xml:"cert_replaced"`

	// 企业项目ID。
	EnterpriseProjectId *string `json:"enterprise_project_id,omitempty" xml:"enterprise_project_id"`

	// 实例扩容时用于区分老实例与新实例。 - true：新创建的实例，允许磁盘动态扩容不需要重启。 - false：老实例
	IsLogicalVolume *bool `json:"is_logical_volume,omitempty" xml:"is_logical_volume"`

	// 实例扩容磁盘次数，如果超过20次则无法扩容磁盘。
	ExtendTimes *int32 `json:"extend_times,omitempty" xml:"extend_times"`

	// 是否打开kafka自动创建topic功能。   - true：开启   - false：关闭
	EnableAutoTopic *bool `json:"enable_auto_topic,omitempty" xml:"enable_auto_topic"`

	// 实例类型：集群，cluster。
	Type *ShowInstanceResponseType `json:"type,omitempty" xml:"type"`

	// 产品标识。
	ProductId *string `json:"product_id,omitempty" xml:"product_id"`

	// 安全组ID。
	SecurityGroupId *string `json:"security_group_id,omitempty" xml:"security_group_id"`

	// 租户安全组名称。
	SecurityGroupName *string `json:"security_group_name,omitempty" xml:"security_group_name"`

	// 子网ID。
	SubnetId *string `json:"subnet_id,omitempty" xml:"subnet_id"`

	// 实例节点所在的可用区，返回“可用区ID”。
	AvailableZones *[]string `json:"available_zones,omitempty" xml:"available_zones"`

	// 总共消息存储空间，单位：GB。
	TotalStorageSpace *int32 `json:"total_storage_space,omitempty" xml:"total_storage_space"`

	// 实例公网连接IP地址。当实例开启了公网访问，实例才包含该参数。
	PublicConnectAddress *string `json:"public_connect_address,omitempty" xml:"public_connect_address"`

	// 存储资源ID。
	StorageResourceId *string `json:"storage_resource_id,omitempty" xml:"storage_resource_id"`

	// IO规格。
	StorageSpecCode *string `json:"storage_spec_code,omitempty" xml:"storage_spec_code"`

	// 服务类型。
	ServiceType *string `json:"service_type,omitempty" xml:"service_type"`

	// 存储类型。
	StorageType *string `json:"storage_type,omitempty" xml:"storage_type"`

	// 消息老化策略。
	RetentionPolicy *ShowInstanceResponseRetentionPolicy `json:"retention_policy,omitempty" xml:"retention_policy"`

	// Kafka公网开启状态。
	KafkaPublicStatus *string `json:"kafka_public_status,omitempty" xml:"kafka_public_status"`

	// kafka公网访问带宽。
	PublicBandwidth *int32 `json:"public_bandwidth,omitempty" xml:"public_bandwidth"`

	// 登录Kafka Manager的用户名。
	KafkaManagerUser *string `json:"kafka_manager_user,omitempty" xml:"kafka_manager_user"`

	// 是否开启消息收集功能。
	EnableLogCollection *bool `json:"enable_log_collection,omitempty" xml:"enable_log_collection"`

	// 跨VPC访问信息。
	CrossVpcInfo *string `json:"cross_vpc_info,omitempty" xml:"cross_vpc_info"`

	// 是否开启ipv6。
	Ipv6Enable *bool `json:"ipv6_enable,omitempty" xml:"ipv6_enable"`

	// IPv6的连接地址。
	Ipv6ConnectAddresses *[]string `json:"ipv6_connect_addresses,omitempty" xml:"ipv6_connect_addresses"`

	// 是否开启转储。新规格产品暂不支持开启转储。
	ConnectorEnable *bool `json:"connector_enable,omitempty" xml:"connector_enable"`

	// 转储任务ID。
	ConnectorId *string `json:"connector_id,omitempty" xml:"connector_id"`

	// 是否开启Kafka rest功能。
	RestEnable *bool `json:"rest_enable,omitempty" xml:"rest_enable"`

	// Kafka rest连接地址。
	RestConnectAddress *string `json:"rest_connect_address,omitempty" xml:"rest_connect_address"`

	// kafka公网访问带宽。待删除版本。
	PublicBoundwidth *int32 `json:"public_boundwidth,omitempty" xml:"public_boundwidth"`

	// 是否开启消息查询功能。
	MessageQueryInstEnable *bool `json:"message_query_inst_enable,omitempty" xml:"message_query_inst_enable"`

	// 是否开启VPC明文访问。
	VpcClientPlain *bool `json:"vpc_client_plain,omitempty" xml:"vpc_client_plain"`

	// Kafka实例支持的特性功能。
	SupportFeatures *string `json:"support_features,omitempty" xml:"support_features"`

	// 是否开启消息轨迹功能。
	TraceEnable *bool `json:"trace_enable,omitempty" xml:"trace_enable"`

	// 是否开启代理。
	AgentEnable *bool `json:"agent_enable,omitempty" xml:"agent_enable"`

	// 租户侧连接地址。
	PodConnectAddress *string `json:"pod_connect_address,omitempty" xml:"pod_connect_address"`

	// 是否开启磁盘加密。
	DiskEncrypted *bool `json:"disk_encrypted,omitempty" xml:"disk_encrypted"`

	// Kafka实例私有连接地址。
	KafkaPrivateConnectAddress *string `json:"kafka_private_connect_address,omitempty" xml:"kafka_private_connect_address"`

	// 云监控版本。
	CesVersion *string `json:"ces_version,omitempty" xml:"ces_version"`

	// 区分实例什么时候开启的公网访问：true,actived,closed,false。
	PublicAccessEnabled *string `json:"public_access_enabled,omitempty" xml:"public_access_enabled"`

	// 节点数。
	NodeNum *int32 `json:"node_num,omitempty" xml:"node_num"`

	// 是否启用新规格计费。
	NewSpecBillingEnable *bool `json:"new_spec_billing_enable,omitempty" xml:"new_spec_billing_enable"`

	// 节点数量。
	BrokerNum *int32 `json:"broker_num,omitempty" xml:"broker_num"`

	// 标签列表。
	Tags *[]TagEntity `json:"tags,omitempty" xml:"tags"`

	// 是否为容灾实例。
	DrEnable       *bool `json:"dr_enable,omitempty" xml:"dr_enable"`
	HttpStatusCode int   `json:"-"`
}

func (o ShowInstanceResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowInstanceResponse struct{}"
	}

	return strings.Join([]string{"ShowInstanceResponse", string(data)}, " ")
}

type ShowInstanceResponseType struct {
	value string
}

type ShowInstanceResponseTypeEnum struct {
	SINGLE  ShowInstanceResponseType
	CLUSTER ShowInstanceResponseType
}

func GetShowInstanceResponseTypeEnum() ShowInstanceResponseTypeEnum {
	return ShowInstanceResponseTypeEnum{
		SINGLE: ShowInstanceResponseType{
			value: "single",
		},
		CLUSTER: ShowInstanceResponseType{
			value: "cluster",
		},
	}
}

func (c ShowInstanceResponseType) Value() string {
	return c.value
}

func (c ShowInstanceResponseType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ShowInstanceResponseType) UnmarshalJSON(b []byte) error {
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

type ShowInstanceResponseRetentionPolicy struct {
	value string
}

type ShowInstanceResponseRetentionPolicyEnum struct {
	TIME_BASE      ShowInstanceResponseRetentionPolicy
	PRODUCE_REJECT ShowInstanceResponseRetentionPolicy
}

func GetShowInstanceResponseRetentionPolicyEnum() ShowInstanceResponseRetentionPolicyEnum {
	return ShowInstanceResponseRetentionPolicyEnum{
		TIME_BASE: ShowInstanceResponseRetentionPolicy{
			value: "time_base",
		},
		PRODUCE_REJECT: ShowInstanceResponseRetentionPolicy{
			value: "produce_reject",
		},
	}
}

func (c ShowInstanceResponseRetentionPolicy) Value() string {
	return c.value
}

func (c ShowInstanceResponseRetentionPolicy) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ShowInstanceResponseRetentionPolicy) UnmarshalJSON(b []byte) error {
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
