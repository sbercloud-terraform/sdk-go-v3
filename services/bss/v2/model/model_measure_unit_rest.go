package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type MeasureUnitRest struct {

	// 度量单位ID。 例如：10表示GB。
	MeasureId *int32 `json:"measure_id,omitempty" xml:"measure_id"`

	// 度量单位的名称，根据查询的语言类型返回结果。 例如：GB。
	MeasureName *string `json:"measure_name,omitempty" xml:"measure_name"`

	// 度量单位名称的英文缩写。 例如：度量单位名称“GB”的英文缩写为“GB”。
	Abbreviation *string `json:"abbreviation,omitempty" xml:"abbreviation"`

	// 度量类型。 1：货币2：时长3：流量4：数量7：容量9：行数10：周期11：频率12：个数16：带宽速率17：容量时长18：查询速率19：带宽速率（1000进制）20：性能测试用量21：面积22：视频23：吞吐量26：通用资源包抵扣单位
	MeasureType *int32 `json:"measure_type,omitempty" xml:"measure_type"`
}

func (o MeasureUnitRest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "MeasureUnitRest struct{}"
	}

	return strings.Join([]string{"MeasureUnitRest", string(data)}, " ")
}
