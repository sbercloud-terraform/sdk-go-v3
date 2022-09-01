package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

//
type PostPaidServerEipBandwidth struct {

	// 功能说明：带宽大小  带宽（Mbit/s），取值范围为[1,2000]。  调整带宽时的最小单位会根据带宽范围不同存在差异。  - 小于等于300Mbit/s：默认最小单位为1Mbit/s。 - 300Mbit/s~1000Mbit/s：默认最小单位为50Mbit/s。 - 大于1000Mbit/s：默认最小单位为500Mbit/s。  > 说明： >  > 如果share_type是PER，该参数必选项；如果share_type是WHOLE并且id有值，该参数会忽略。
	Size *int32 `json:"size,omitempty" xml:"size"`

	// 带宽的共享类型。  共享类型枚举：PER，表示独享。WHOLE，表示共享。
	Sharetype PostPaidServerEipBandwidthSharetype `json:"sharetype" xml:"sharetype"`

	// 带宽的计费类型。  - 未传该字段，表示按带宽计费。 - 字段值为空，表示按带宽计费。 - 字段值为“traffic”，表示按流量计费。 - 字段为其它值，会导致创建云服务器失败。  > 说明： >  > 如果share_type是WHOLE并且id有值，该参数会忽略。
	Chargemode *string `json:"chargemode,omitempty" xml:"chargemode"`

	// 带宽ID，创建WHOLE类型带宽的弹性IP时可以指定之前的共享带宽创建。  取值范围：WHOLE类型的带宽ID。  > 说明： >  > 当创建WHOLE类型的带宽时，该字段必选。
	Id *string `json:"id,omitempty" xml:"id"`
}

func (o PostPaidServerEipBandwidth) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PostPaidServerEipBandwidth struct{}"
	}

	return strings.Join([]string{"PostPaidServerEipBandwidth", string(data)}, " ")
}

type PostPaidServerEipBandwidthSharetype struct {
	value string
}

type PostPaidServerEipBandwidthSharetypeEnum struct {
	PER   PostPaidServerEipBandwidthSharetype
	WHOLE PostPaidServerEipBandwidthSharetype
}

func GetPostPaidServerEipBandwidthSharetypeEnum() PostPaidServerEipBandwidthSharetypeEnum {
	return PostPaidServerEipBandwidthSharetypeEnum{
		PER: PostPaidServerEipBandwidthSharetype{
			value: "PER",
		},
		WHOLE: PostPaidServerEipBandwidthSharetype{
			value: "WHOLE",
		},
	}
}

func (c PostPaidServerEipBandwidthSharetype) Value() string {
	return c.value
}

func (c PostPaidServerEipBandwidthSharetype) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *PostPaidServerEipBandwidthSharetype) UnmarshalJSON(b []byte) error {
	myConverter := converter.StringConverterFactory("string")
	if myConverter != nil {
		val, err := myConverter.CovertStringToInterface(strings.Trim(string(b[:]), "\""))
		if err == nil {
			c.value = val.(string)
			return nil
		}
		return err
	} else {
		return errors.New("convert enum data to string error")
	}
}
