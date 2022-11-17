package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type VarsStructurePrimitiveTypeHolder struct {

	// HCL支持参数，即，同一个模板可以给予不同的参数而达到不同的效果。  * var_structure可以允许客户提交最简单的字符串类型的参数  * RF支持vars_structure，vars_body和vars_uri，如果他们中声名了同一个变量，将报错400  * vars_structure中的值只支持简单的字符串类型，如果需要使用其他类型，需要用户自己在HCL引用时转换， 或者用户可以使用vars_uri、vars_body，vars_uri和vars_body中支持HCL支持的各种类型以及复杂结构  * 如果vars_structure过大，可以使用vars_uri  * 注意：vars中不应该传递任何敏感信息，RF会直接明文使用、log、展示、存储对应的vars
	VarsStructure *[]VarsStructure `json:"vars_structure,omitempty"`
}

func (o VarsStructurePrimitiveTypeHolder) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "VarsStructurePrimitiveTypeHolder struct{}"
	}

	return strings.Join([]string{"VarsStructurePrimitiveTypeHolder", string(data)}, " ")
}
