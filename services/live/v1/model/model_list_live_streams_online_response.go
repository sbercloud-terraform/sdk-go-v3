package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type ListLiveStreamsOnlineResponse struct {

	// 总条页数
	TotalPage *int64 `json:"total_page,omitempty" xml:"total_page"`

	// 总条目数
	TotalNum *int64 `json:"total_num,omitempty" xml:"total_num"`

	// 偏移量
	Offset *int64 `json:"offset,omitempty" xml:"offset"`

	// 每页条目数
	Limit *int64 `json:"limit,omitempty" xml:"limit"`

	// 请求唯一标识
	RequestId *string `json:"request_id,omitempty" xml:"request_id"`

	// 推流统计
	Streams        *[]OnlineInfo `json:"streams,omitempty" xml:"streams"`
	HttpStatusCode int           `json:"-"`
}

func (o ListLiveStreamsOnlineResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListLiveStreamsOnlineResponse struct{}"
	}

	return strings.Join([]string{"ListLiveStreamsOnlineResponse", string(data)}, " ")
}
