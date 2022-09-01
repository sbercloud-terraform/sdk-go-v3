package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 集群列表对象。
type ClusterInfo struct {

	// 集群ID。
	Id string `json:"id" xml:"id"`

	// 集群名称
	Name string `json:"name" xml:"name"`

	// 集群状态，有效值包括：  - CREATING：创建中 - AVAILABLE：可用 - UNAVAILABLE：不可用 - CREATION FAILED：创建失败 - FROZEN：已冻结
	Status string `json:"status" xml:"status"`

	// 数据仓库版本
	Version string `json:"version" xml:"version"`

	// 集群上次修改时间，格式为 ISO8601：YYYY-MM-DDThh:mm:ssZ
	Updated string `json:"updated" xml:"updated"`

	// 集群创建时间，格式为 ISO8601：YYYY-MM-DDThh:mm:ssZ。
	Created string `json:"created" xml:"created"`

	// 集群服务端口，取值范围8000~30000，默认值：8000
	Port int32 `json:"port" xml:"port"`

	// 集群的内网连接信息
	Endpoints []Endpoints `json:"endpoints" xml:"endpoints"`

	// 集群实例
	Nodes []Nodes `json:"nodes" xml:"nodes"`

	// 集群标签
	Tags []Tags `json:"tags" xml:"tags"`

	// 管理员用户名
	UserName string `json:"user_name" xml:"user_name"`

	// 节点数量
	NumberOfNode int32 `json:"number_of_node" xml:"number_of_node"`

	// 事件数
	RecentEvent int32 `json:"recent_event" xml:"recent_event"`

	// 可用区
	AvailabilityZone string `json:"availability_zone" xml:"availability_zone"`

	// 企业项目ID。值为0表示默认企业项目“default”
	EnterpriseProjectId string `json:"enterprise_project_id" xml:"enterprise_project_id"`

	// 节点类型
	NodeType string `json:"node_type" xml:"node_type"`

	// 虚拟私有云ID
	VpcId string `json:"vpc_id" xml:"vpc_id"`

	// 子网ID
	SubnetId string `json:"subnet_id" xml:"subnet_id"`

	PublicIp *PublicIp `json:"public_ip" xml:"public_ip"`

	// 集群的公网连接信息，如果未指定，则默认不使用公网连接信息。
	PublicEndpoints []PublicEndpoints `json:"public_endpoints" xml:"public_endpoints"`

	// 任务信息，由key、value组成。key值为正在进行的任务，value值为正在进行任务的进度。key值的有效值包括：  - GROWING：扩容中 - RESTORING：恢复中 - SNAPSHOTTING：快照中 - REPAIRING: 修复中 - CREATING: 创建中
	ActionProgress map[string]string `json:"action_progress" xml:"action_progress"`

	// “可用”集群状态的子状态，有效值包括：  - NORMAL：正常 - READONLY：只读 - REDISTRIBUTING：重分布中 - REDISTRIBUTION-FAILURE：重分布失败 - UNBALANCED：非均衡 - UNBALANCED | READONLY：非均衡，只读 - DEGRADED：节点故障 - DEGRADED | READONLY：节点故障，只读 - DEGRADED | UNBALANCED：节点故障，非均衡 - UNBALANCED | REDISTRIBUTING：非均衡，重分布中 - UNBALANCED | REDISTRIBUTION-FAILURE：非均衡，重分布失败 - READONLY | REDISTRIBUTION-FAILURE：只读，重分布失败 - UNBALANCED | READONLY | REDISTRIBUTION-FAILURE：非均衡，只读，重分布失败 - DEGRADED | REDISTRIBUTION-FAILURE：节点故障，重分布失败 - DEGRADED | UNBALANCED | REDISTRIBUTION-FAILURE：节点故障，非均衡，只读，重分布失败 - DEGRADED | UNBALANCED | READONLY | REDISTRIBUTION-FAILURE：节点故障，非均衡，只读，重分布失败 - DEGRADED | UNBALANCED | READONLY：节点故障，非均衡，只读
	SubStatus string `json:"sub_status" xml:"sub_status"`

	// 集群管理任务，有效值包括：  - UNFREEZING：解冻中 - FREEZING：冻结中 - RESTORING：恢复中 - SNAPSHOTTING：快照中 - GROWING：扩容中 - REBOOTING：重启中 - SETTING_CONFIGURATION：安全设置配置中 - CONFIGURING_EXT_DATASOURCE：MRS连接配置中 - DELETING_EXT_DATASOURCE：删除MRS连接 - REBOOT_FAILURE：重启失败 - RESIZE_FAILURE：扩容失败
	TaskStatus string `json:"task_status" xml:"task_status"`

	// 安全组ID
	SecurityGroupId string `json:"security_group_id" xml:"security_group_id"`

	FailedReasons *FailedReason `json:"failed_reasons,omitempty" xml:"failed_reasons"`
}

func (o ClusterInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ClusterInfo struct{}"
	}

	return strings.Join([]string{"ClusterInfo", string(data)}, " ")
}
