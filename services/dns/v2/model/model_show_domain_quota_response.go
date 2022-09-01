package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type ShowDomainQuotaResponse struct {
	Quotas         *[]DomainQuotaResponseQuotas `json:"quotas,omitempty" xml:"quotas"`
	HttpStatusCode int                          `json:"-"`
}

func (o ShowDomainQuotaResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowDomainQuotaResponse struct{}"
	}

	return strings.Join([]string{"ShowDomainQuotaResponse", string(data)}, " ")
}
