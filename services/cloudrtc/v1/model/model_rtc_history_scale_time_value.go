package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type RtcHistoryScaleTimeValue struct {

	// 采样时间。日期格式按照ISO8601表示法，并使用UTC时间。格式为YYYY-MM-DD
	Date *string `json:"date,omitempty" xml:"date"`

	// 通话人数，指总的uid个数
	UserCount *int64 `json:"user_count,omitempty" xml:"user_count"`

	// 通话人次，指总的session个数
	SessionCount *int64 `json:"session_count,omitempty" xml:"session_count"`

	// 房间数
	RoomCount *int64 `json:"room_count,omitempty" xml:"room_count"`

	// 最大同时在线人数
	MaxOnlineUserCount *int64 `json:"max_online_user_count,omitempty" xml:"max_online_user_count"`

	// 最大同时在线房间数
	MaxOnlineRoomCount *int64 `json:"max_online_room_count,omitempty" xml:"max_online_room_count"`

	// 音视频通话总时长，单位秒
	CommunicationDuration *int64 `json:"communication_duration,omitempty" xml:"communication_duration"`

	// 视频通话总时长，单位秒
	VideoCommunicationDuration *int64 `json:"video_communication_duration,omitempty" xml:"video_communication_duration"`

	// 音频通话总时长，单位秒
	AudioCommunicationDuration *int64 `json:"audio_communication_duration,omitempty" xml:"audio_communication_duration"`
}

func (o RtcHistoryScaleTimeValue) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RtcHistoryScaleTimeValue struct{}"
	}

	return strings.Join([]string{"RtcHistoryScaleTimeValue", string(data)}, " ")
}
