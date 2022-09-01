package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type MysqlInstanceInfoDetail struct {

	// 实例ID。
	Id string `json:"id" xml:"id"`

	// 创建的实例名称。
	Name string `json:"name" xml:"name"`

	// 租户在某一region下的project ID。
	ProjectId string `json:"project_id" xml:"project_id"`

	// 实例状态。 取值： 值为“BUILD”，表示实例正在创建。 值为“ACTIVE”，表示实例正常。 值为“FAILED”，表示实例创建失败。 值为“FROZEN”，表示实例冻结。 值为“MODIFYING”，表示实例正在扩容。 值为“REBOOTING”，表示实例正在重启。 值为“RESTORING”，表示实例正在恢复。 值为“SWITCHOVER”，表示实例正在主备切换。 值为“MIGRATING”，表示实例正在迁移。 值为“BACKING UP”，表示实例正在进行备份。 值为“MODIFYING DATABASE PORT”，表示实例正在修改数据库端口。值为“STORAGE FULL”，表示实例磁盘空间满。
	Status *string `json:"status,omitempty" xml:"status"`

	// 数据库端口号。
	Port *string `json:"port,omitempty" xml:"port"`

	// 实例备注
	Alias *string `json:"alias,omitempty" xml:"alias"`

	// 实例类型，取值为“Cluster”。
	Type *string `json:"type,omitempty" xml:"type"`

	ChargeInfo *MysqlInstanceChargeInfo `json:"charge_info,omitempty" xml:"charge_info"`

	// 节点个数。
	NodeCount *int32 `json:"node_count,omitempty" xml:"node_count"`

	Datastore *MysqlDatastoreWithKernelVersion `json:"datastore,omitempty" xml:"datastore"`

	// 备份空间使用大小，单位为GB。
	BackupUsedSpace *float64 `json:"backup_used_space,omitempty" xml:"backup_used_space"`

	// 创建时间，格式为\"yyyy-mm-ddThh:mm:ssZ\"。 其中，T指某个时间的开始；Z指时区偏移量，例如北京时间偏移显示为+0800。说明：创建时返回值为空，数据库实例创建成功后该值不为空。
	Created *string `json:"created,omitempty" xml:"created"`

	// 更新时间，格式与\"created\"字段对应格式完全相同。说明：创建时返回值为空，数据库实例创建成功后该值不为空。
	Updated *string `json:"updated,omitempty" xml:"updated"`

	// 实例的写内网IP。
	PrivateWriteIps *[]string `json:"private_write_ips,omitempty" xml:"private_write_ips"`

	// 实例的公网IP。
	PublicIps *string `json:"public_ips,omitempty" xml:"public_ips"`

	// 默认用户名。
	DbUserName *string `json:"db_user_name,omitempty" xml:"db_user_name"`

	// 虚拟私有云ID。
	VpcId *string `json:"vpc_id,omitempty" xml:"vpc_id"`

	// 子网的网络ID信息。
	SubnetId *string `json:"subnet_id,omitempty" xml:"subnet_id"`

	// 安全组ID。
	SecurityGroupId *string `json:"security_group_id,omitempty" xml:"security_group_id"`

	// 实例创建的模板ID，或者应用到实例的最新参数组模板ID。
	ConfigurationId *string `json:"configuration_id,omitempty" xml:"configuration_id"`

	BackupStrategy *MysqlBackupStrategy `json:"backup_strategy,omitempty" xml:"backup_strategy"`

	Nodes *[]MysqlInstanceNodeInfo `json:"nodes,omitempty" xml:"nodes"`

	// 企业项目ID。
	EnterpriseProjectId *string `json:"enterprise_project_id,omitempty" xml:"enterprise_project_id"`

	// 时区。
	TimeZone *string `json:"time_zone,omitempty" xml:"time_zone"`

	// 可用区模式，单可用区single或多可用区multi。
	AzMode *string `json:"az_mode,omitempty" xml:"az_mode"`

	// 主可用区。
	MasterAzCode *string `json:"master_az_code,omitempty" xml:"master_az_code"`

	// 可维护时间窗，为UTC时间。
	MaintenanceWindow *string `json:"maintenance_window,omitempty" xml:"maintenance_window"`

	// 实例标签。
	Tags *[]MysqlTags `json:"tags,omitempty" xml:"tags"`

	// 专属资源池ID，只有数据库实例属于专属资源池才会返回该参数。
	DedicatedResourceId *string `json:"dedicated_resource_id,omitempty" xml:"dedicated_resource_id"`

	Proxies *[]MysqlProxyInfo `json:"proxies,omitempty" xml:"proxies"`
}

func (o MysqlInstanceInfoDetail) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "MysqlInstanceInfoDetail struct{}"
	}

	return strings.Join([]string{"MysqlInstanceInfoDetail", string(data)}, " ")
}
