package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type ShowStaticRouteResponse struct {
	Route *Route `json:"route,omitempty"`

	// 请求ID
	RequestId      *string `json:"request_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ShowStaticRouteResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowStaticRouteResponse struct{}"
	}

	return strings.Join([]string{"ShowStaticRouteResponse", string(data)}, " ")
}
