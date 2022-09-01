package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type CouponQuotaV2 struct {

	// 优惠券额度ID。
	QuotaId *string `json:"quota_id,omitempty" xml:"quota_id"`

	// 优惠券额度的类型： 0：代金券额度1：现金券额度
	QuotaType *int32 `json:"quota_type,omitempty" xml:"quota_type"`

	// 创建时间。 UTC时间，格式：yyyy-MM-dd'T'HH:mm:ss'Z'，如“2019-05-06T08:05:01Z”。
	CreateTime *string `json:"create_time,omitempty" xml:"create_time"`

	// 最后一次更新时间。 UTC时间，格式：yyyy-MM-dd'T'HH:mm:ss'Z'，如“2019-05-06T08:05:01Z”。
	LastUpdateTime *string `json:"last_update_time,omitempty" xml:"last_update_time"`

	// 优惠券额度的值，精确到小数点后2位。
	QuotaValue *float64 `json:"quota_value,omitempty" xml:"quota_value"`

	// 优惠券额度的状态： 0：正常3：失效（过期失效和人工设置失效）4：额度调整中（伙伴可以查看该额度，但不能使用该额度发放优惠券）5：冻结
	QuotaStatus *int32 `json:"quota_status,omitempty" xml:"quota_status"`

	// 剩余的优惠券额度，精确到小数点后2位。
	Balance *float64 `json:"balance,omitempty" xml:"balance"`

	// 面额单位。 1：元。
	MeasureId *int32 `json:"measure_id,omitempty" xml:"measure_id"`

	// 币种。 CNY：人民币
	Currency *string `json:"currency,omitempty" xml:"currency"`

	// 生效时间。 UTC时间，格式：yyyy-MM-dd'T'HH:mm:ss'Z'，如“2019-05-06T08:05:01Z”。
	EffectiveTime *string `json:"effective_time,omitempty" xml:"effective_time"`

	// 失效时间。 UTC时间，格式：yyyy-MM-dd'T'HH:mm:ss'Z'，如“2019-05-06T08:05:01Z”。
	ExpireTime *string `json:"expire_time,omitempty" xml:"expire_time"`

	// 优惠券额度上的限制属性，具体参见表2。
	LimitInfos *[]QuotaLimitInfo `json:"limit_infos,omitempty" xml:"limit_infos"`
}

func (o CouponQuotaV2) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CouponQuotaV2 struct{}"
	}

	return strings.Join([]string{"CouponQuotaV2", string(data)}, " ")
}
