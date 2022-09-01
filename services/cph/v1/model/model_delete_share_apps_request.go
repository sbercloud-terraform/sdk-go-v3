package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type DeleteShareAppsRequest struct {
	Body *DeleteShareAppsRequestBody `json:"body,omitempty" xml:"body"`
}

func (o DeleteShareAppsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteShareAppsRequest struct{}"
	}

	return strings.Join([]string{"DeleteShareAppsRequest", string(data)}, " ")
}
