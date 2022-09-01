package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// Response Object
type ShowEdgeApplicationVersionResponse struct {

	// 应用ID
	EdgeAppId *string `json:"edge_app_id,omitempty" xml:"edge_app_id"`

	// 应用名称
	Name *string `json:"name,omitempty" xml:"name"`

	// 部署类型docker|process
	DeployType *string `json:"deploy_type,omitempty" xml:"deploy_type"`

	// 是否允许部署多实例
	DeployMultiInstance *bool `json:"deploy_multi_instance,omitempty" xml:"deploy_multi_instance"`

	// 应用版本
	Version *string `json:"version,omitempty" xml:"version"`

	// 应用集成的边缘升得快版本
	SdkVersion *string `json:"sdk_version,omitempty" xml:"sdk_version"`

	// 应用描述
	Description *string `json:"description,omitempty" xml:"description"`

	// 创建时间
	CreateTime *string `json:"create_time,omitempty" xml:"create_time"`

	// 最后一次修改时间
	UpdateTime *string `json:"update_time,omitempty" xml:"update_time"`

	// 应用版本状态
	State *ShowEdgeApplicationVersionResponseState `json:"state,omitempty" xml:"state"`

	LivenessProbe *ProbeDto `json:"liveness_probe,omitempty" xml:"liveness_probe"`

	ReadinessProbe *ProbeDto `json:"readiness_probe,omitempty" xml:"readiness_probe"`

	// 架构
	Arch *[]string `json:"arch,omitempty" xml:"arch"`

	// 启动命令
	Command *[]string `json:"command,omitempty" xml:"command"`

	// 启动参数
	Args *[]string `json:"args,omitempty" xml:"args"`

	ContainerSettings *ContainerSettingsDto `json:"container_settings,omitempty" xml:"container_settings"`

	// 应用输出路由端点
	Outputs *[]string `json:"outputs,omitempty" xml:"outputs"`

	// 应用输入路由
	Inputs *[]string `json:"inputs,omitempty" xml:"inputs"`

	// 应用实现的服务列表
	Services *[]string `json:"services,omitempty" xml:"services"`

	// 发布时间
	PublishTime *string `json:"publish_time,omitempty" xml:"publish_time"`

	// 下线时间
	OffShelfTime   *string `json:"off_shelf_time,omitempty" xml:"off_shelf_time"`
	HttpStatusCode int     `json:"-"`
}

func (o ShowEdgeApplicationVersionResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowEdgeApplicationVersionResponse struct{}"
	}

	return strings.Join([]string{"ShowEdgeApplicationVersionResponse", string(data)}, " ")
}

type ShowEdgeApplicationVersionResponseState struct {
	value string
}

type ShowEdgeApplicationVersionResponseStateEnum struct {
	DRAFT     ShowEdgeApplicationVersionResponseState
	PUBLISHED ShowEdgeApplicationVersionResponseState
	OFF_SHELF ShowEdgeApplicationVersionResponseState
}

func GetShowEdgeApplicationVersionResponseStateEnum() ShowEdgeApplicationVersionResponseStateEnum {
	return ShowEdgeApplicationVersionResponseStateEnum{
		DRAFT: ShowEdgeApplicationVersionResponseState{
			value: "DRAFT",
		},
		PUBLISHED: ShowEdgeApplicationVersionResponseState{
			value: "PUBLISHED",
		},
		OFF_SHELF: ShowEdgeApplicationVersionResponseState{
			value: "OFF_SHELF",
		},
	}
}

func (c ShowEdgeApplicationVersionResponseState) Value() string {
	return c.value
}

func (c ShowEdgeApplicationVersionResponseState) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ShowEdgeApplicationVersionResponseState) UnmarshalJSON(b []byte) error {
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
