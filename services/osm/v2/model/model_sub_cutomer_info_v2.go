package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type SubCutomerInfoV2 struct {

	// 手机号（匿名化）
	Mobile *string `json:"mobile,omitempty" xml:"mobile"`

	// 邮箱（匿名化）
	Email *string `json:"email,omitempty" xml:"email"`

	// 客户id
	CustomerId *string `json:"customer_id,omitempty" xml:"customer_id"`

	// 主账号id
	DomainId *string `json:"domain_id,omitempty" xml:"domain_id"`

	// 客户名称（匿名化）
	CustomerName *string `json:"customer_name,omitempty" xml:"customer_name"`

	// 国家码
	AreaCode *string `json:"area_code,omitempty" xml:"area_code"`
}

func (o SubCutomerInfoV2) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SubCutomerInfoV2 struct{}"
	}

	return strings.Join([]string{"SubCutomerInfoV2", string(data)}, " ")
}
