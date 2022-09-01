package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// Response Object
type ShowImageResponse struct {

	// 镜像ID。
	Id *string `json:"id,omitempty" xml:"id"`

	// 镜像名称。
	Name *string `json:"name,omitempty" xml:"name"`

	// 镜像状态。
	Status *ShowImageResponseStatus `json:"status,omitempty" xml:"status"`

	// 镜像格式。
	DiskFormat *ShowImageResponseDiskFormat `json:"disk_format,omitempty" xml:"disk_format"`

	// 最小系统盘（单位：GB），取值为40～1024GB。
	MinDisk *int32 `json:"min_disk,omitempty" xml:"min_disk"`

	// 最小内存（单位：MB），默认值为0。
	MinRam *int32 `json:"min_ram,omitempty" xml:"min_ram"`

	// 镜像所属租户ID。
	Owner *string `json:"owner,omitempty" xml:"owner"`

	// 是否受保护。
	Protected *bool `json:"protected,omitempty" xml:"protected"`

	// 可见性。
	Visibility *string `json:"visibility,omitempty" xml:"visibility"`

	// 创建时间。
	CreatedAt *string `json:"created_at,omitempty" xml:"created_at"`

	// 更新时间。
	UpdatedAt *string `json:"updated_at,omitempty" xml:"updated_at"`

	// 镜像链接信息。
	Self *string `json:"self,omitempty" xml:"self"`

	// 是否是删除的镜像，取值为true或者false。
	Deleted *bool `json:"deleted,omitempty" xml:"deleted"`

	// 镜像使用环境类型。
	VirtualEnvType *string `json:"virtual_env_type,omitempty" xml:"virtual_env_type"`

	// 删除时间，格式为UTC时间。
	DeletedAt *string `json:"deleted_at,omitempty" xml:"deleted_at"`

	// 镜像关联的任务ID。
	RelationJobId *string `json:"relation_job_id,omitempty" xml:"relation_job_id"`

	// 镜像类型。  取值范围： - gold：公有镜像； - private：私有镜像。
	Imagetype *ShowImageResponseImagetype `json:"__imagetype,omitempty" xml:"__imagetype"`

	// 镜像平台分类。
	Platform *string `json:"__platform,omitempty" xml:"__platform"`

	// 镜像系统类型。
	OsType *ShowImageResponseOsType `json:"__os_type,omitempty" xml:"__os_type"`

	// 镜像的操作系统具体版本。
	OsVersion *string `json:"__os_version,omitempty" xml:"__os_version"`

	// 是否是注册过的镜像。
	Isregistered *bool `json:"__isregistered,omitempty" xml:"__isregistered"`

	// 如果镜像支持KVM，取值为true，否则无该属性。
	SupportKvm *string `json:"__support_kvm,omitempty" xml:"__support_kvm"`

	// 如果镜像是支持KVM虚拟化平台下的GPU类型，取值为“V100_vGPU”或者“RTX5000”，否则无该属性。
	SupportKvmGpuType *string `json:"__support_kvm_gpu_type,omitempty" xml:"__support_kvm_gpu_type"`

	// 如果镜像支持AI加速，取值为true，否则无该属性。
	SupportKvmAscend310 *string `json:"__support_kvm_ascend_310,omitempty" xml:"__support_kvm_ascend_310"`

	// 如果镜像支持计算增强，取值为true，否则无该属性。
	SupportKvmHi1822Hiovs *string `json:"__support_kvm_hi1822_hiovs,omitempty" xml:"__support_kvm_hi1822_hiovs"`

	// 如果镜像为ARM架构类型，取值为true，否则无该属性。
	SupportArm *string `json:"__support_arm,omitempty" xml:"__support_arm"`

	// 镜像启动模式，取值为uefi或bios，不指定时无该属性。
	HwFirmwareType *string `json:"hw_firmware_type,omitempty" xml:"hw_firmware_type"`

	// 镜像来源。  - 来源边缘实例：instance:<实例id> - 来源IMS：ims:<镜像id>:<region id>
	DataSource *string `json:"data_source,omitempty" xml:"data_source"`

	// 如果镜像支持GPU T4类型，取值为true，否则无该属性。
	SupportGpuT4 *string `json:"__support_gpu_t4,omitempty" xml:"__support_gpu_t4"`

	OriginRegionInfo *CloudImageRegionInfo `json:"origin_region_info,omitempty" xml:"origin_region_info"`

	// 边缘区域详情。
	EdgeRegionInfo *[]EdgeImageRegionInfo `json:"edge_region_info,omitempty" xml:"edge_region_info"`
	HttpStatusCode int                    `json:"-"`
}

func (o ShowImageResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowImageResponse struct{}"
	}

	return strings.Join([]string{"ShowImageResponse", string(data)}, " ")
}

type ShowImageResponseStatus struct {
	value string
}

type ShowImageResponseStatusEnum struct {
	QUEUED  ShowImageResponseStatus
	SAVING  ShowImageResponseStatus
	DELETED ShowImageResponseStatus
	KILLED  ShowImageResponseStatus
	ACTIVE  ShowImageResponseStatus
}

func GetShowImageResponseStatusEnum() ShowImageResponseStatusEnum {
	return ShowImageResponseStatusEnum{
		QUEUED: ShowImageResponseStatus{
			value: "queued",
		},
		SAVING: ShowImageResponseStatus{
			value: "saving",
		},
		DELETED: ShowImageResponseStatus{
			value: "deleted",
		},
		KILLED: ShowImageResponseStatus{
			value: "killed",
		},
		ACTIVE: ShowImageResponseStatus{
			value: "active",
		},
	}
}

func (c ShowImageResponseStatus) Value() string {
	return c.value
}

func (c ShowImageResponseStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ShowImageResponseStatus) UnmarshalJSON(b []byte) error {
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

type ShowImageResponseDiskFormat struct {
	value string
}

type ShowImageResponseDiskFormatEnum struct {
	ZVHD2 ShowImageResponseDiskFormat
	VHD   ShowImageResponseDiskFormat
	ZVHD  ShowImageResponseDiskFormat
	RAW   ShowImageResponseDiskFormat
	QCOW2 ShowImageResponseDiskFormat
}

func GetShowImageResponseDiskFormatEnum() ShowImageResponseDiskFormatEnum {
	return ShowImageResponseDiskFormatEnum{
		ZVHD2: ShowImageResponseDiskFormat{
			value: "zvhd2",
		},
		VHD: ShowImageResponseDiskFormat{
			value: "vhd",
		},
		ZVHD: ShowImageResponseDiskFormat{
			value: "zvhd",
		},
		RAW: ShowImageResponseDiskFormat{
			value: "raw",
		},
		QCOW2: ShowImageResponseDiskFormat{
			value: "qcow2",
		},
	}
}

func (c ShowImageResponseDiskFormat) Value() string {
	return c.value
}

func (c ShowImageResponseDiskFormat) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ShowImageResponseDiskFormat) UnmarshalJSON(b []byte) error {
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

type ShowImageResponseImagetype struct {
	value string
}

type ShowImageResponseImagetypeEnum struct {
	GOLD    ShowImageResponseImagetype
	PRIVATE ShowImageResponseImagetype
}

func GetShowImageResponseImagetypeEnum() ShowImageResponseImagetypeEnum {
	return ShowImageResponseImagetypeEnum{
		GOLD: ShowImageResponseImagetype{
			value: "gold",
		},
		PRIVATE: ShowImageResponseImagetype{
			value: "private ",
		},
	}
}

func (c ShowImageResponseImagetype) Value() string {
	return c.value
}

func (c ShowImageResponseImagetype) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ShowImageResponseImagetype) UnmarshalJSON(b []byte) error {
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

type ShowImageResponseOsType struct {
	value string
}

type ShowImageResponseOsTypeEnum struct {
	LINUX   ShowImageResponseOsType
	WINDOWS ShowImageResponseOsType
	OTHER   ShowImageResponseOsType
}

func GetShowImageResponseOsTypeEnum() ShowImageResponseOsTypeEnum {
	return ShowImageResponseOsTypeEnum{
		LINUX: ShowImageResponseOsType{
			value: "Linux",
		},
		WINDOWS: ShowImageResponseOsType{
			value: "Windows",
		},
		OTHER: ShowImageResponseOsType{
			value: "Other",
		},
	}
}

func (c ShowImageResponseOsType) Value() string {
	return c.value
}

func (c ShowImageResponseOsType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ShowImageResponseOsType) UnmarshalJSON(b []byte) error {
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
