package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 部署计划对象
type Deployment struct {

	// 部署计划ID。
	Id *string `json:"id,omitempty" xml:"id"`

	// 部署位置信息列表
	Distribution *[]Distribution `json:"distribution,omitempty" xml:"distribution"`

	Edgecloud *DeploymentEdgecloud `json:"edgecloud,omitempty" xml:"edgecloud"`
}

func (o Deployment) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Deployment struct{}"
	}

	return strings.Join([]string{"Deployment", string(data)}, " ")
}
