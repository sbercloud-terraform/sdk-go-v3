package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Kafka数据源请求内容
type KafkaContentRsp struct {

	// KAFKA连接类型
	ConnectionType *string `json:"connection_type,omitempty" xml:"connection_type"`

	// Kafka实例ID
	InstanceId *string `json:"instance_id,omitempty" xml:"instance_id"`

	// Kafka的VPCEP信息包括service_id,service_name,client_port
	VpcepInfos *[]VpcepInfoRsp `json:"vpcep_infos,omitempty" xml:"vpcep_infos"`

	// Kafka的broker信息包括broker_ip, broker_port
	BrokerInfos *[]KafkaBrokerInfo `json:"broker_infos,omitempty" xml:"broker_infos"`

	AuthInfo *KafkaAuthInfo `json:"auth_info,omitempty" xml:"auth_info"`
}

func (o KafkaContentRsp) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "KafkaContentRsp struct{}"
	}

	return strings.Join([]string{"KafkaContentRsp", string(data)}, " ")
}
