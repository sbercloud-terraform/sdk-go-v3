package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type ShowQuotasResponse struct {
	Quotas         *Quotas `json:"quotas,omitempty" xml:"quotas"`
	HttpStatusCode int     `json:"-"`
}

func (o ShowQuotasResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowQuotasResponse struct{}"
	}

	return strings.Join([]string{"ShowQuotasResponse", string(data)}, " ")
}
