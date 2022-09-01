package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// IP池对象。 支持IPv4和IPv6
type IpPool struct {

	// 线路的ID。
	Id *string `json:"id,omitempty" xml:"id"`

	// 线路所属站点ID。
	SiteId *string `json:"site_id,omitempty" xml:"site_id"`

	// 线路标识。
	PoolId *string `json:"pool_id,omitempty" xml:"pool_id"`

	// IPv4[或IPv6](tag:hide)线路。  取值范围： - 4：IPv4线路 [- 6：IPv6线路](tag:hide)
	IpVersion *string `json:"ip_version,omitempty" xml:"ip_version"`

	Operator *Operator `json:"operator,omitempty" xml:"operator"`

	// 线路的显示名称。
	DisplayName *string `json:"display_name,omitempty" xml:"display_name"`
}

func (o IpPool) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "IpPool struct{}"
	}

	return strings.Join([]string{"IpPool", string(data)}, " ")
}
