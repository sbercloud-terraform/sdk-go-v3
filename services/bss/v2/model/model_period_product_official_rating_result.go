package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type PeriodProductOfficialRatingResult struct {

	// ID标识，来源于请求中的ID。
	Id *string `json:"id,omitempty" xml:"id"`

	// 包年/包月产品的ID。
	ProductId *string `json:"product_id,omitempty" xml:"product_id"`

	// 包年/包月产品的官网价。
	OfficialWebsiteAmount *float64 `json:"official_website_amount,omitempty" xml:"official_website_amount"`

	// 价格度量单位标识。 1：元
	MeasureId *int32 `json:"measure_id,omitempty" xml:"measure_id"`

	// 分期金额的官网价。  说明： 暂只支持IES产品。
	InstallmentOfficialWebsiteAmount *string `json:"installment_official_website_amount,omitempty" xml:"installment_official_website_amount"`

	// 分期付款的周期类型。 2：月  说明： 暂只支持IES产品。
	InstallmentPeriodType *int32 `json:"installment_period_type,omitempty" xml:"installment_period_type"`
}

func (o PeriodProductOfficialRatingResult) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PeriodProductOfficialRatingResult struct{}"
	}

	return strings.Join([]string{"PeriodProductOfficialRatingResult", string(data)}, " ")
}
