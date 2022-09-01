package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// Response Object
type UpdateUserResponse struct {

	// 用户名。
	AccessKey *string `json:"access_key,omitempty" xml:"access_key"`

	// 密钥。
	SecretKey *string `json:"secret_key,omitempty" xml:"secret_key"`

	// IP白名单。
	WhiteRemoteAddress *string `json:"white_remote_address,omitempty" xml:"white_remote_address"`

	// 是否为管理员。
	Admin *bool `json:"admin,omitempty" xml:"admin"`

	// 默认的主题权限。
	DefaultTopicPerm *UpdateUserResponseDefaultTopicPerm `json:"default_topic_perm,omitempty" xml:"default_topic_perm"`

	// 默认的消费组权限。
	DefaultGroupPerm *UpdateUserResponseDefaultGroupPerm `json:"default_group_perm,omitempty" xml:"default_group_perm"`

	// 特殊的主题权限。
	TopicPerms *[]UserTopicPerms `json:"topic_perms,omitempty" xml:"topic_perms"`

	// 特殊的消费组权限。
	GroupPerms     *[]UserGroupPerms `json:"group_perms,omitempty" xml:"group_perms"`
	HttpStatusCode int               `json:"-"`
}

func (o UpdateUserResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateUserResponse struct{}"
	}

	return strings.Join([]string{"UpdateUserResponse", string(data)}, " ")
}

type UpdateUserResponseDefaultTopicPerm struct {
	value string
}

type UpdateUserResponseDefaultTopicPermEnum struct {
	PUB     UpdateUserResponseDefaultTopicPerm
	SUB     UpdateUserResponseDefaultTopicPerm
	PUB_SUB UpdateUserResponseDefaultTopicPerm
	DENY    UpdateUserResponseDefaultTopicPerm
}

func GetUpdateUserResponseDefaultTopicPermEnum() UpdateUserResponseDefaultTopicPermEnum {
	return UpdateUserResponseDefaultTopicPermEnum{
		PUB: UpdateUserResponseDefaultTopicPerm{
			value: "PUB",
		},
		SUB: UpdateUserResponseDefaultTopicPerm{
			value: "SUB",
		},
		PUB_SUB: UpdateUserResponseDefaultTopicPerm{
			value: "PUB|SUB",
		},
		DENY: UpdateUserResponseDefaultTopicPerm{
			value: "DENY",
		},
	}
}

func (c UpdateUserResponseDefaultTopicPerm) Value() string {
	return c.value
}

func (c UpdateUserResponseDefaultTopicPerm) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *UpdateUserResponseDefaultTopicPerm) UnmarshalJSON(b []byte) error {
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

type UpdateUserResponseDefaultGroupPerm struct {
	value string
}

type UpdateUserResponseDefaultGroupPermEnum struct {
	PUB     UpdateUserResponseDefaultGroupPerm
	SUB     UpdateUserResponseDefaultGroupPerm
	PUB_SUB UpdateUserResponseDefaultGroupPerm
	DENY    UpdateUserResponseDefaultGroupPerm
}

func GetUpdateUserResponseDefaultGroupPermEnum() UpdateUserResponseDefaultGroupPermEnum {
	return UpdateUserResponseDefaultGroupPermEnum{
		PUB: UpdateUserResponseDefaultGroupPerm{
			value: "PUB",
		},
		SUB: UpdateUserResponseDefaultGroupPerm{
			value: "SUB",
		},
		PUB_SUB: UpdateUserResponseDefaultGroupPerm{
			value: "PUB|SUB",
		},
		DENY: UpdateUserResponseDefaultGroupPerm{
			value: "DENY",
		},
	}
}

func (c UpdateUserResponseDefaultGroupPerm) Value() string {
	return c.value
}

func (c UpdateUserResponseDefaultGroupPerm) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *UpdateUserResponseDefaultGroupPerm) UnmarshalJSON(b []byte) error {
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
