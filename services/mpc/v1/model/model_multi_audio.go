package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type MultiAudio struct {

	// 音轨信息
	TracksInfo *[]TracksInfo `json:"tracks_info,omitempty" xml:"tracks_info"`

	// 音频文件
	AudioFiles *[]AudioFile `json:"audio_files,omitempty" xml:"audio_files"`

	// 默认语言
	DefaultLanguage *string `json:"default_language,omitempty" xml:"default_language"`
}

func (o MultiAudio) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "MultiAudio struct{}"
	}

	return strings.Join([]string{"MultiAudio", string(data)}, " ")
}
