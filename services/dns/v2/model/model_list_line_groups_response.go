package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type ListLineGroupsResponse struct {

	// 列表对象。
	Linegroups *[]CreateLineGroupsResp `json:"linegroups,omitempty"`

	Metadata       *Metedata `json:"metadata,omitempty"`
	HttpStatusCode int       `json:"-"`
}

func (o ListLineGroupsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListLineGroupsResponse struct{}"
	}

	return strings.Join([]string{"ListLineGroupsResponse", string(data)}, " ")
}
