package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type Quotas struct {

	// 配额上限
	QuotaLimit *int32 `json:"quota_limit,omitempty" xml:"quota_limit"`

	// 配额类型
	Type *string `json:"type,omitempty" xml:"type"`

	// 已使用配额数
	Used *int32 `json:"used,omitempty" xml:"used"`

	// 域名所属用户的domain_id。
	UserDomainId *string `json:"user_domain_id,omitempty" xml:"user_domain_id"`
}

func (o Quotas) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Quotas struct{}"
	}

	return strings.Join([]string{"Quotas", string(data)}, " ")
}
