package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type ShowDefaultConfigResponse struct {

	// 是否开启L7层防护
	EnableL7 *bool `json:"enable_L7,omitempty" xml:"enable_L7"`

	// 流量分段ID，取值范围：1～9
	TrafficPosId *int64 `json:"traffic_pos_id,omitempty" xml:"traffic_pos_id"`

	// HTTP请求数分段ID，取值范围：1～15
	HttpRequestPosId *int64 `json:"http_request_pos_id,omitempty" xml:"http_request_pos_id"`

	// 清洗时访问限制分段ID，取值范围：1～8
	CleaningAccessPosId *int64 `json:"cleaning_access_pos_id,omitempty" xml:"cleaning_access_pos_id"`

	// 应用类型ID，可选取值： - 0 - 1
	AppTypeId      *int64 `json:"app_type_id,omitempty" xml:"app_type_id"`
	HttpStatusCode int    `json:"-"`
}

func (o ShowDefaultConfigResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowDefaultConfigResponse struct{}"
	}

	return strings.Join([]string{"ShowDefaultConfigResponse", string(data)}, " ")
}
