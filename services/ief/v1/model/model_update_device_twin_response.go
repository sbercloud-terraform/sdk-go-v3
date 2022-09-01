package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type UpdateDeviceTwinResponse struct {
	PropertyVisitors *ValueInPropertyVisitors `json:"property_visitors,omitempty" xml:"property_visitors"`

	Twin           *ValueInTwinResponse `json:"twin,omitempty" xml:"twin"`
	HttpStatusCode int                  `json:"-"`
}

func (o UpdateDeviceTwinResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateDeviceTwinResponse struct{}"
	}

	return strings.Join([]string{"UpdateDeviceTwinResponse", string(data)}, " ")
}
