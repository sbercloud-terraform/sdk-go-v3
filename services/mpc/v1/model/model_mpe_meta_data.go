package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type MpeMetaData struct {

	// 封装类型。
	PackType *string `json:"pack_type,omitempty" xml:"pack_type"`

	// 视频时长。
	Duration *float64 `json:"duration,omitempty" xml:"duration"`

	// 视频大小。
	VideoSize *int64 `json:"video_size,omitempty" xml:"video_size"`

	// 视频宽度。
	Width *int32 `json:"width,omitempty" xml:"width"`

	// 视频高度。
	Height *int32 `json:"height,omitempty" xml:"height"`

	// 码率。
	BitRate *int32 `json:"bit_rate,omitempty" xml:"bit_rate"`

	// 音频码率。
	AudioBitRate *int32 `json:"audio_bit_rate,omitempty" xml:"audio_bit_rate"`

	// 帧率。  取值范围：0或[5,60]，0表示自适应。  单位：帧每秒。  > 若设置的帧率不在取值范围内，则自动调整为0，若设置的帧率高于片源帧率，则自动调整为片源帧率。
	FrameRate *int32 `json:"frame_rate,omitempty" xml:"frame_rate"`

	// 编码类型名称。
	CodecName *string `json:"codec_name,omitempty" xml:"codec_name"`

	// 音频编码类型。
	AudioCodecName *string `json:"audio_codec_name,omitempty" xml:"audio_codec_name"`

	// 声道数。
	Channels *int32 `json:"channels,omitempty" xml:"channels"`

	// 采样率。
	Sample *int32 `json:"sample,omitempty" xml:"sample"`

	// 是否音频。
	IsAudio *bool `json:"is_audio,omitempty" xml:"is_audio"`
}

func (o MpeMetaData) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "MpeMetaData struct{}"
	}

	return strings.Join([]string{"MpeMetaData", string(data)}, " ")
}
