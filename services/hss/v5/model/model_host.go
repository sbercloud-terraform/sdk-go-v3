package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type Host struct {

	// 服务器名称
	HostName *string `json:"host_name,omitempty" xml:"host_name"`

	// 服务器ID
	HostId *string `json:"host_id,omitempty" xml:"host_id"`

	// Agent ID
	AgentId *string `json:"agent_id,omitempty" xml:"agent_id"`

	// 私有IP地址
	PrivateIp *string `json:"private_ip,omitempty" xml:"private_ip"`

	// 弹性公网IP地址
	PublicIp *string `json:"public_ip,omitempty" xml:"public_ip"`

	// 所属企业项目名称
	EnterpriseProjectName *string `json:"enterprise_project_name,omitempty" xml:"enterprise_project_name"`

	// 服务器状态，包含如下4种。   - ACTIVE ：运行中。   - SHUTOFF ：关机。   - BUILDING ：创建中。   - ERROR ：故障。
	HostStatus *string `json:"host_status,omitempty" xml:"host_status"`

	// Agent状态，包含如下5种。   - not_installed ：未安装。   - online ：在线。   - offline ：离线。   - install_failed ：安装失败。   - installing ：安装中。
	AgentStatus *string `json:"agent_status,omitempty" xml:"agent_status"`

	// 安装结果，包含如下12种。   - install_succeed ：安装成功。   - network_access_timeout ：网络不通，访问超时。   - invalid_port ：无效端口。   - auth_failed ：认证错误，口令不正确。   - permission_denied ：权限错误，被拒绝。   - no_available_vpc ：没有相同VPC的agent在线虚拟机。   - install_exception ：安装异常。   - invalid_param ：参数错误。   - install_failed ：安装失败。   - package_unavailable ：安装包失效。   - os_type_not_support ：系统类型错误。   - os_arch_not_support ：架构类型错误。
	InstallResultCode *string `json:"install_result_code,omitempty" xml:"install_result_code"`

	// 主机开通的版本，包含如下6种输入。   - hss.version.null ：无。   - hss.version.basic ：基础版。   - hss.version.enterprise ：企业版。   - hss.version.premium ：旗舰版。   - hss.version.wtp ：网页防篡改版。   - hss.version.container.enterprise ：容器版。
	Version *string `json:"version,omitempty" xml:"version"`

	// 防护状态，包含如下2种。 - closed ：未防护。 - opened ：防护中。
	ProtectStatus *string `json:"protect_status,omitempty" xml:"protect_status"`

	// 系统镜像
	OsImage *string `json:"os_image,omitempty" xml:"os_image"`

	// 操作系统类型，包含如下2种。   - Linux ：Linux。   - Windows ：Windows。
	OsType *string `json:"os_type,omitempty" xml:"os_type"`

	// 操作系统位数
	OsBit *string `json:"os_bit,omitempty" xml:"os_bit"`

	// 云主机安全检测结果，包含如下4种。 - undetected ：未检测。 - clean ：无风险。 - risk ：有风险。 - scanning ：检测中。
	DetectResult *string `json:"detect_result,omitempty" xml:"detect_result"`

	// 收费模式，包含如下2种。   - packet_cycle ：包年/包月。   - on_demand ：按需。
	ChargingMode *string `json:"charging_mode,omitempty" xml:"charging_mode"`

	// 云服务资源实例ID（UUID）
	ResourceId *string `json:"resource_id,omitempty" xml:"resource_id"`

	// 是否非华为云机器
	OutsideHost *bool `json:"outside_host,omitempty" xml:"outside_host"`

	// 服务器组ID
	GroupId *string `json:"group_id,omitempty" xml:"group_id"`

	// 服务器组名称
	GroupName *string `json:"group_name,omitempty" xml:"group_name"`

	// 策略组ID
	PolicyGroupId *string `json:"policy_group_id,omitempty" xml:"policy_group_id"`

	// 策略组名称
	PolicyGroupName *string `json:"policy_group_name,omitempty" xml:"policy_group_name"`

	// 资产风险
	Asset *int32 `json:"asset,omitempty" xml:"asset"`

	// 漏洞风险
	Vulnerability *int32 `json:"vulnerability,omitempty" xml:"vulnerability"`

	// 基线风险
	Baseline *int32 `json:"baseline,omitempty" xml:"baseline"`

	// 入侵风险
	Intrusion *int32 `json:"intrusion,omitempty" xml:"intrusion"`

	// 资产重要性，包含如下4种   - important ：重要资产   - common ：一般资产   - test ：测试资产
	AssetValue *string `json:"asset_value,omitempty" xml:"asset_value"`

	// 标签列表
	Labels *[]string `json:"labels,omitempty" xml:"labels"`
}

func (o Host) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Host struct{}"
	}

	return strings.Join([]string{"Host", string(data)}, " ")
}
