package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type SearchAttendanceRecordsOfHisMeetingRequest struct {

	// 会议UUID。
	ConfUUID string `json:"confUUID" xml:"confUUID"`

	// 指定返回的记录索引。该值必须大于等于0； 默认为0。
	Offset *int32 `json:"offset,omitempty" xml:"offset"`

	// 指定返回的记录数，默认是20，最大500条。
	Limit *int32 `json:"limit,omitempty" xml:"limit"`

	// 查询用来当作关键词的字符串。
	SearchKey *string `json:"searchKey,omitempty" xml:"searchKey"`

	// 用户的UUID（已在USG注册过的）。
	UserUUID *string `json:"userUUID,omitempty" xml:"userUUID"`

	// 标识是否为第三方portal过来的请求。
	XAuthorizationType *string `json:"X-Authorization-Type,omitempty" xml:"X-Authorization-Type"`

	// 用于区分到哪个HCSO站点鉴权。
	XSiteId *string `json:"X-Site-Id,omitempty" xml:"X-Site-Id"`

	// 语言。默认简体中文。 - zh-CN: 简体中文。 - en-US: 美国英文。
	AcceptLanguage *string `json:"Accept-Language,omitempty" xml:"Accept-Language"`
}

func (o SearchAttendanceRecordsOfHisMeetingRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SearchAttendanceRecordsOfHisMeetingRequest struct{}"
	}

	return strings.Join([]string{"SearchAttendanceRecordsOfHisMeetingRequest", string(data)}, " ")
}
