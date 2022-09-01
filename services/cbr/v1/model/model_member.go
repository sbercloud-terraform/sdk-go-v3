package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

//
type Member struct {

	// 共享状态
	Status MemberStatus `json:"status" xml:"status"`

	// 共享时间，例如:\"2020-02-05T10:38:34.209782\"
	CreatedAt *string `json:"created_at,omitempty" xml:"created_at"`

	// 更新时间，例如:\"2020-02-05T10:38:34.209782\"
	UpdatedAt *string `json:"updated_at,omitempty" xml:"updated_at"`

	// 备份副本id
	BackupId *string `json:"backup_id,omitempty" xml:"backup_id"`

	// 接受的共享备份副本注册的镜像id
	ImageId *string `json:"image_id,omitempty" xml:"image_id"`

	// 接受备份共享的项目id
	DestProjectId *string `json:"dest_project_id,omitempty" xml:"dest_project_id"`

	// 目标端接受共享备份的存储库id
	VaultId *string `json:"vault_id,omitempty" xml:"vault_id"`

	// 共享记录id
	Id *string `json:"id,omitempty" xml:"id"`
}

func (o Member) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Member struct{}"
	}

	return strings.Join([]string{"Member", string(data)}, " ")
}

type MemberStatus struct {
	value string
}

type MemberStatusEnum struct {
	PENDING  MemberStatus
	ACCEPTED MemberStatus
	REJECTED MemberStatus
}

func GetMemberStatusEnum() MemberStatusEnum {
	return MemberStatusEnum{
		PENDING: MemberStatus{
			value: "pending",
		},
		ACCEPTED: MemberStatus{
			value: "accepted",
		},
		REJECTED: MemberStatus{
			value: "rejected",
		},
	}
}

func (c MemberStatus) Value() string {
	return c.value
}

func (c MemberStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *MemberStatus) UnmarshalJSON(b []byte) error {
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
