package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type CommentUserV4 struct {

	// 发表评论用户id
	UserNumId *int32 `json:"user_num_id,omitempty" xml:"user_num_id"`

	// 发表评论用户名称
	UserName *string `json:"user_name,omitempty" xml:"user_name"`

	// 发表评论用户昵称
	NickName *string `json:"nick_name,omitempty" xml:"nick_name"`
}

func (o CommentUserV4) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CommentUserV4 struct{}"
	}

	return strings.Join([]string{"CommentUserV4", string(data)}, " ")
}
