package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type ListCustomerBillsMonthlyBreakDownRequest struct {

	// 语言。en_US：英文。zh_CN：中文
	XLanguage *string `json:"X-Language,omitempty" xml:"X-Language"`

	// 查询分摊成本的月份，东八区时间，格式：YYYY-MM。
	SharedMonth string `json:"shared_month" xml:"shared_month"`

	// 云服务类型编码，例如OBS的云服务类型编码为“hws.service.type.obs”。您可以调用[查询云服务类型列表](https://support.huaweicloud.com/api-oce/zh-cn_topic_0000001256679455.html)接口获取。此参数不携带或携带值为空时，不作为筛选条件；携带值为空串时，作为筛选条件。
	ServiceTypeCode *string `json:"service_type_code,omitempty" xml:"service_type_code"`

	// 资源类型编码，例如ECS的VM为“hws.resource.type.vm”。您可以调用[查询资源类型列表](https://support.huaweicloud.com/api-oce/zh-cn_topic_0000001256519451.html)接口获取。此参数不携带或携带值为空时，不作为筛选条件；携带值为空串时，作为筛选条件。
	ResourceTypeCode *string `json:"resource_type_code,omitempty" xml:"resource_type_code"`

	// 云服务区编码，例如：“cn-north-1”。具体请参见[地区和终端节点](https://developer.huaweicloud.com/endpoint)对应云服务的“区域”列的值。此参数不携带或携带值为空时，不作为筛选条件。
	RegionCode *string `json:"region_code,omitempty" xml:"region_code"`

	// 计费模式。1：包年/包月3：按需10：预留实例此参数不携带或携带值为空时，不作为筛选条件。
	ChargingMode *int32 `json:"charging_mode,omitempty" xml:"charging_mode"`

	// 账单类型。1：消费-新购2：消费-续订3：消费-变更4：退款-退订5：消费-使用8：消费-自动续订9：调账-补偿14：消费-服务支持计划月末扣费16：调账-扣费18：消费-按月付费20：退款-变更 此参数不携带或携带值为空时，不作为筛选条件。
	BillType *int32 `json:"bill_type,omitempty" xml:"bill_type"`

	// 偏移量，从0开始。默认值为0。 说明： offset用于分页处理，如不涉及分页，请使用默认值0。offset表示相对于满足条件的第一个数据的偏移量。如offset = 1，则返回满足条件的第二个数据至最后一个数据。例如，满足查询条件的结果共10条数据，limit取值为10，offset取值为1，则返回的数据为2~10，第一条数据不返回。
	Offset *int32 `json:"offset,omitempty" xml:"offset"`

	// 每次查询的数量限制。默认值为10。
	Limit *int32 `json:"limit,omitempty" xml:"limit"`

	// 资源ID。此参数不携带或携带值为空时，不作为筛选条件。
	ResourceId *string `json:"resource_id,omitempty" xml:"resource_id"`

	// 资源名称。此参数不携带或携带值为空时，不作为筛选条件。
	ResourceName *string `json:"resource_name,omitempty" xml:"resource_name"`

	// 企业项目标识（企业项目ID）。default项目对应ID：0未归集（表示该云服务不支持企业项目管理能力）项目对应ID：-1其余项目对应ID获取方法请参见[如何获取企业项目ID](https://support.huaweicloud.com/usermanual-em/zh-cn_topic_0126101490.html)。此参数不携带或携带值为空时，不作为筛选条件。
	EnterpriseProjectId *string `json:"enterprise_project_id,omitempty" xml:"enterprise_project_id"`

	// 查询资源消费记录的方式。oneself：客户自己sub_customer：企业子客户all：客户自己和企业子客户默认为all，如果没有企业子客户，取值为all时查询的是客户自己的资源消费记录。此参数不携带，不作为筛选条件。
	Method *string `json:"method,omitempty" xml:"method"`

	// 企业子账号ID。此参数不携带或携带值为空或携带值为空串时，不作为筛选条件。 说明： 如果method取值不为sub_customer，则该参数无效。如果method取值为sub_customer，则该参数不能为空。
	SubCustomerId *string `json:"sub_customer_id,omitempty" xml:"sub_customer_id"`
}

func (o ListCustomerBillsMonthlyBreakDownRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListCustomerBillsMonthlyBreakDownRequest struct{}"
	}

	return strings.Join([]string{"ListCustomerBillsMonthlyBreakDownRequest", string(data)}, " ")
}
