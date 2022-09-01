package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type ShowApiVersionResponse struct {
	Version        *ApiVersionResponse `json:"version,omitempty" xml:"version"`
	HttpStatusCode int                 `json:"-"`
}

func (o ShowApiVersionResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowApiVersionResponse struct{}"
	}

	return strings.Join([]string{"ShowApiVersionResponse", string(data)}, " ")
}
