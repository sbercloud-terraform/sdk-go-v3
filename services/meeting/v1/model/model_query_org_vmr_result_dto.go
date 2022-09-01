package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 组织查询的vmr列表，不越权显示vmr的来宾密码，主席密码等
type QueryOrgVmrResultDto struct {

	// 唯一标识。 说明：对应会议管理->创建会议接口中的vmrID。
	Id *string `json:"id,omitempty" xml:"id"`

	// 云会议室ID。 说明：对应会议管理->创建会议接口中当vmrIDType等于0（固定ID）时返回数据的conferenceID 。
	VmrId *string `json:"vmrId,omitempty" xml:"vmrId"`

	// 云会议室名称。
	VmrName *string `json:"vmrName,omitempty" xml:"vmrName"`

	// 云会议室套餐名称。
	VmrPkgName *string `json:"vmrPkgName,omitempty" xml:"vmrPkgName"`

	// 云会议室套餐会议并发方数。
	VmrPkgParties *int32 `json:"vmrPkgParties,omitempty" xml:"vmrPkgParties"`

	// 最大观众与会方数（仅网络研讨会有效）
	MaxAudienceParties *int32 `json:"maxAudienceParties,omitempty" xml:"maxAudienceParties"`

	Member *IdMarkDto `json:"member,omitempty" xml:"member"`

	Device *IdMarkDto `json:"device,omitempty" xml:"device"`

	// 云会议室状态。 * 0：正常 * 1：冻结 * 2：未分配
	Status *int32 `json:"status,omitempty" xml:"status"`

	// 到期时间，utc时间戳
	ExpireDate *int64 `json:"expireDate,omitempty" xml:"expireDate"`

	// 按次资源转商后，商用规格最大观众与会方数（仅网络研讨会有效）
	CommercialMaxAudienceParties *int32 `json:"commercialMaxAudienceParties,omitempty" xml:"commercialMaxAudienceParties"`
}

func (o QueryOrgVmrResultDto) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "QueryOrgVmrResultDto struct{}"
	}

	return strings.Join([]string{"QueryOrgVmrResultDto", string(data)}, " ")
}
