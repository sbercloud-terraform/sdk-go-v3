package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type SimpleProject struct {

	// 项目名称
	ProjectName *string `json:"project_name,omitempty" xml:"project_name"`

	// 项目id
	ProjectId *string `json:"project_id,omitempty" xml:"project_id"`
}

func (o SimpleProject) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SimpleProject struct{}"
	}

	return strings.Join([]string{"SimpleProject", string(data)}, " ")
}
