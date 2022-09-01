package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type ShowClobDetailResponse struct {

	// clob详情
	ClobString     *string `json:"clob_string,omitempty" xml:"clob_string"`
	HttpStatusCode int     `json:"-"`
}

func (o ShowClobDetailResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowClobDetailResponse struct{}"
	}

	return strings.Join([]string{"ShowClobDetailResponse", string(data)}, " ")
}
