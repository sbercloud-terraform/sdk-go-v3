package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type DeleteVirtualInterfaceResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o DeleteVirtualInterfaceResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteVirtualInterfaceResponse struct{}"
	}

	return strings.Join([]string{"DeleteVirtualInterfaceResponse", string(data)}, " ")
}
