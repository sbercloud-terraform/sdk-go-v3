package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/sdktime"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CentralNetworkErRouteTableAttachment 企业路由器的路由表附件详情。
type CentralNetworkErRouteTableAttachment struct {

	// 资源ID标识符。
	Id string `json:"id"`

	// 实例名字。
	Name string `json:"name"`

	// 实例描述。不支持 <>。
	Description *string `json:"description,omitempty"`

	// 实例所属帐号ID。
	DomainId string `json:"domain_id"`

	State *CentralNetworkConnectionStateEnum `json:"state"`

	// 实例创建时间。UTC时间格式，yyyy-MM-ddTHH:mm:ss。
	CreatedAt *sdktime.SdkTime `json:"created_at"`

	// 实例更新时间。UTC时间格式，yyyy-MM-ddTHH:mm:ss。
	UpdatedAt *sdktime.SdkTime `json:"updated_at"`

	// 资源ID标识符。
	CentralNetworkId string `json:"central_network_id"`

	// 资源ID标识符。
	CentralNetworkPlaneId string `json:"central_network_plane_id"`

	// 资源ID标识符。
	GlobalConnectionBandwidthId *string `json:"global_connection_bandwidth_id,omitempty"`

	// 是否冻结
	IsFrozen bool `json:"is_frozen"`

	BandwidthType *BandwidthTypeEnum `json:"bandwidth_type"`

	// 带宽值定义，单位Mbps。
	BandwidthSize *int64 `json:"bandwidth_size,omitempty"`

	// 资源ID标识符。
	EnterpriseRouterId string `json:"enterprise_router_id"`

	// 实例所属项目ID。
	EnterpriseRouterProjectId string `json:"enterprise_router_project_id"`

	// RegionID。
	EnterpriseRouterRegionId string `json:"enterprise_router_region_id"`

	// 资源ID标识符。
	EnterpriseRouterAttachmentId *string `json:"enterprise_router_attachment_id,omitempty"`

	// 资源ID标识符。
	EnterpriseRouterTableId string `json:"enterprise_router_table_id"`

	// 站点编码定义
	EnterpriseRouterSiteCode string `json:"enterprise_router_site_code"`

	// 资源ID标识符。
	AttachedErTableId string `json:"attached_er_table_id"`

	// RegionID。
	AttachedErTableRegionId string `json:"attached_er_table_region_id"`

	// 实例所属项目ID。
	AttachedErTableProjectId string `json:"attached_er_table_project_id"`

	// 站点编码定义
	AttachedErTableSiteCode string `json:"attached_er_table_site_code"`

	// 资源ID标识符。
	AttachedErId string `json:"attached_er_id"`

	// 资源ID标识符。
	AttachedErAttachmentId *string `json:"attached_er_attachment_id,omitempty"`

	ApprovedState *ApprovedStateEnum `json:"approved_state"`

	HostedCloud *HostedCloudEnum `json:"hosted_cloud,omitempty"`

	// 审批拒绝创建企业路由表附件的原因。
	Reason *string `json:"reason,omitempty"`
}

func (o CentralNetworkErRouteTableAttachment) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CentralNetworkErRouteTableAttachment struct{}"
	}

	return strings.Join([]string{"CentralNetworkErRouteTableAttachment", string(data)}, " ")
}
