package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type ShowRecordSetByZoneResponse struct {
	Links *PageLink `json:"links,omitempty" xml:"links"`

	Recordsets *[]ShowRecordSetByZoneResp `json:"recordsets,omitempty" xml:"recordsets"`

	Metadata       *Metedata `json:"metadata,omitempty" xml:"metadata"`
	HttpStatusCode int       `json:"-"`
}

func (o ShowRecordSetByZoneResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowRecordSetByZoneResponse struct{}"
	}

	return strings.Join([]string{"ShowRecordSetByZoneResponse", string(data)}, " ")
}
