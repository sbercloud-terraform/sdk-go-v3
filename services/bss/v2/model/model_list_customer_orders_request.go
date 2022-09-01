package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type ListCustomerOrdersRequest struct {

	// 订单ID。此参数不携带或携带值为空时，不作为筛选条件。 说明： 使用特殊字符进行查询的时候，请注意进行URL编码转换，如“%”的转码应为“%25”。
	OrderId *string `json:"order_id,omitempty" xml:"order_id"`

	// 客户账号ID。您可以调用查询客户列表接口获取customer_id。此参数不携带或携带值为空时，不作为筛选条件。
	CustomerId *string `json:"customer_id,omitempty" xml:"customer_id"`

	// 订单创建开始时间。UTC时间，格式：yyyy-MM-dd'T'HH:mm:ss'Z'，如“2019-05-06T08:05:01Z”。其中，HH范围是0～23，mm和ss范围是0～59。此参数不携带或携带值为空时，不作为筛选条件。
	CreateTimeBegin *string `json:"create_time_begin,omitempty" xml:"create_time_begin"`

	// 订单创建结束时间。UTC时间，格式：yyyy-MM-dd'T'HH:mm:ss'Z'，如“2019-05-06T08:05:01Z”。其中，HH范围是0～23，mm和ss范围是0～59。此参数不携带或携带值为空时，不作为筛选条件。
	CreateTimeEnd *string `json:"create_time_end,omitempty" xml:"create_time_end"`

	// 云服务类型编码，例如OBS的云服务类型编码为“hws.service.type.obs”。您可以调用查询云服务类型列表接口获取。此参数不携带或携带值为空时，不作为筛选条件。
	ServiceTypeCode *string `json:"service_type_code,omitempty" xml:"service_type_code"`

	// 订单状态：1：待审核3：处理中4：已取消5：已完成6：待支付9：待确认此参数不携带或携带值为空时，不作为筛选条件。
	Status *int32 `json:"status,omitempty" xml:"status"`

	// 订单类型：1：开通2：续订3：变更4：退订11：按需转包年/包月13：试用14：转商用15：费用调整此参数不携带或携带值为空时，不作为筛选条件。
	OrderType *string `json:"order_type,omitempty" xml:"order_type"`

	// 每次查询的订单数量，默认值为10。
	Limit *int32 `json:"limit,omitempty" xml:"limit"`

	// 偏移量，从0开始。默认值为0。 说明： offset用于分页处理，如不涉及分页，请使用默认值0。offset表示相对于满足条件的第一个数据的偏移量。如offset = 1，则返回满足条件的第二个数据至最后一个数据。例如，满足查询条件的结果共10条数据，limit取值为10，offset取值为1，则返回的数据为2~10，第一条数据不返回。
	Offset *int32 `json:"offset,omitempty" xml:"offset"`

	// 查询的订单列表排序。支持按照创建时间进行排序，带-表示倒序。创建时间：升序为createTime，倒序为-createTime。此参数不携带或携带值为空时，不作为筛选条件。
	OrderBy *string `json:"order_by,omitempty" xml:"order_by"`

	// 订单支付开始时间。UTC时间，格式：yyyy-MM-dd'T'HH:mm:ss'Z'，如“2019-05-06T08:05:01Z”。其中，HH范围是0～23，mm和ss范围是0～59。此参数不携带或携带值为空时，不作为筛选条件。
	PaymentTimeBegin *string `json:"payment_time_begin,omitempty" xml:"payment_time_begin"`

	// 订单支付结束时间。UTC时间，格式：yyyy-MM-dd'T'HH:mm:ss'Z'，如“2019-05-06T08:05:01Z”。其中，HH范围是0～23，mm和ss范围是0～59。此参数不携带或携带值为空时，不作为筛选条件。
	PaymentTimeEnd *string `json:"payment_time_end,omitempty" xml:"payment_time_end"`

	// 云经销商ID。获取方法请参见[查询云经销商列表](https://support.huaweicloud.com/api-bpconsole/espp_00003.html)。华为云总经销商（一级经销商）查询云经销商的客户订单列表时，需要携带该参数；否则只能查询自己客户的订单列表。
	IndirectPartnerId *string `json:"indirect_partner_id,omitempty" xml:"indirect_partner_id"`
}

func (o ListCustomerOrdersRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListCustomerOrdersRequest struct{}"
	}

	return strings.Join([]string{"ListCustomerOrdersRequest", string(data)}, " ")
}
