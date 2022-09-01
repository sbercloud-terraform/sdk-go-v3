package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type MuteMeetingRequest struct {

	// 会议ID。
	ConferenceID string `json:"conferenceID" xml:"conferenceID"`

	// 会控授权令牌，通过获取会控token接口获得。
	XConferenceAuthorization string `json:"X-Conference-Authorization" xml:"X-Conference-Authorization"`

	Body *RestMuteReqBody `json:"body,omitempty" xml:"body"`
}

func (o MuteMeetingRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "MuteMeetingRequest struct{}"
	}

	return strings.Join([]string{"MuteMeetingRequest", string(data)}, " ")
}
