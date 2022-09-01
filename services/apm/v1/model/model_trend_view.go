package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// 视图配置信息
type TrendView struct {

	// 视图类型
	ViewType *TrendViewViewType `json:"view_type,omitempty" xml:"view_type"`

	// 采集器名称
	CollectorName *string `json:"collector_name,omitempty" xml:"collector_name"`

	// 视图对应的指标集名称
	MetricSet *string `json:"metric_set,omitempty" xml:"metric_set"`

	// 图标所需展示的标题
	Title *string `json:"title,omitempty" xml:"title"`

	// 表格的方向，H：默认，表头横向，V：表头纵向
	TableDirection *TrendViewTableDirection `json:"table_direction,omitempty" xml:"table_direction"`

	// 分组
	GroupBy *string `json:"group_by,omitempty" xml:"group_by"`

	// 过滤列表模型
	Filter *string `json:"filter,omitempty" xml:"filter"`

	// 所需展示的字段列表模型列表
	FieldItemList *[]FieldItem `json:"field_item_list,omitempty" xml:"field_item_list"`

	// 跨度
	Span *bool `json:"span,omitempty" xml:"span"`

	// span字段属性
	SpanField *string `json:"span_field,omitempty" xml:"span_field"`

	// 排序
	OrderBy *string `json:"order_by,omitempty" xml:"order_by"`

	// 是否只展示最近一笔数据
	Latest *string `json:"latest,omitempty" xml:"latest"`
}

func (o TrendView) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "TrendView struct{}"
	}

	return strings.Join([]string{"TrendView", string(data)}, " ")
}

type TrendViewViewType struct {
	value string
}

type TrendViewViewTypeEnum struct {
	TREND    TrendViewViewType
	SUMTABLE TrendViewViewType
	RAWTABLE TrendViewViewType
}

func GetTrendViewViewTypeEnum() TrendViewViewTypeEnum {
	return TrendViewViewTypeEnum{
		TREND: TrendViewViewType{
			value: "trend",
		},
		SUMTABLE: TrendViewViewType{
			value: "sumtable",
		},
		RAWTABLE: TrendViewViewType{
			value: "rawtable",
		},
	}
}

func (c TrendViewViewType) Value() string {
	return c.value
}

func (c TrendViewViewType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *TrendViewViewType) UnmarshalJSON(b []byte) error {
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

type TrendViewTableDirection struct {
	value string
}

type TrendViewTableDirectionEnum struct {
	H TrendViewTableDirection
	V TrendViewTableDirection
}

func GetTrendViewTableDirectionEnum() TrendViewTableDirectionEnum {
	return TrendViewTableDirectionEnum{
		H: TrendViewTableDirection{
			value: "H",
		},
		V: TrendViewTableDirection{
			value: "V",
		},
	}
}

func (c TrendViewTableDirection) Value() string {
	return c.value
}

func (c TrendViewTableDirection) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *TrendViewTableDirection) UnmarshalJSON(b []byte) error {
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
