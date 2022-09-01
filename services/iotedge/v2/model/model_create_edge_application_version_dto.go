package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

type CreateEdgeApplicationVersionDto struct {

	// 应用版本
	Version string `json:"version" xml:"version"`

	// 应用描述
	Description *string `json:"description,omitempty" xml:"description"`

	// 应用集成的边缘升得快版本
	SdkVersion *string `json:"sdk_version,omitempty" xml:"sdk_version"`

	// 应用部署类型，分为docker容器部署类型和process进程部署类型，兼容之前数据，此字段可以为空，为空情况为docker类型
	DeployType *CreateEdgeApplicationVersionDtoDeployType `json:"deploy_type,omitempty" xml:"deploy_type"`

	// 是否允许部署多实例
	DeployMultiInstance *bool `json:"deploy_multi_instance,omitempty" xml:"deploy_multi_instance"`

	ContainerSettings *ContainerSettingsDto `json:"container_settings" xml:"container_settings"`

	LivenessProbe *ProbeDto `json:"liveness_probe,omitempty" xml:"liveness_probe"`

	ReadinessProbe *ProbeDto `json:"readiness_probe,omitempty" xml:"readiness_probe"`

	// 架构
	Arch *interface{} `json:"arch" xml:"arch"`

	// 启动命令
	Command *interface{} `json:"command,omitempty" xml:"command"`

	// 启动参数
	Args *interface{} `json:"args,omitempty" xml:"args"`

	// 应用输出路由端点
	Outputs *interface{} `json:"outputs,omitempty" xml:"outputs"`

	// 应用输入路由
	Inputs *interface{} `json:"inputs,omitempty" xml:"inputs"`

	// 应用实现的服务列表
	Services *interface{} `json:"services,omitempty" xml:"services"`
}

func (o CreateEdgeApplicationVersionDto) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateEdgeApplicationVersionDto struct{}"
	}

	return strings.Join([]string{"CreateEdgeApplicationVersionDto", string(data)}, " ")
}

type CreateEdgeApplicationVersionDtoDeployType struct {
	value string
}

type CreateEdgeApplicationVersionDtoDeployTypeEnum struct {
	DOCKER  CreateEdgeApplicationVersionDtoDeployType
	PROCESS CreateEdgeApplicationVersionDtoDeployType
}

func GetCreateEdgeApplicationVersionDtoDeployTypeEnum() CreateEdgeApplicationVersionDtoDeployTypeEnum {
	return CreateEdgeApplicationVersionDtoDeployTypeEnum{
		DOCKER: CreateEdgeApplicationVersionDtoDeployType{
			value: "docker",
		},
		PROCESS: CreateEdgeApplicationVersionDtoDeployType{
			value: "process",
		},
	}
}

func (c CreateEdgeApplicationVersionDtoDeployType) Value() string {
	return c.value
}

func (c CreateEdgeApplicationVersionDtoDeployType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *CreateEdgeApplicationVersionDtoDeployType) UnmarshalJSON(b []byte) error {
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
