package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// 作业审核结果，当作业状态为succeeded时存在
type AudioModerationResultResult struct {

	// 音频审核结果是否通过。 block：包含敏感信息，不通过 pass：不包含敏感信息，通过 review：需要人工复检
	Suggestion *AudioModerationResultResultSuggestion `json:"suggestion,omitempty" xml:"suggestion"`

	// 审核详情
	Details *[]AudioModerationResultResultDetails `json:"details,omitempty" xml:"details"`

	// 音频文本内容
	AudioText *string `json:"audio_text,omitempty" xml:"audio_text"`
}

func (o AudioModerationResultResult) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AudioModerationResultResult struct{}"
	}

	return strings.Join([]string{"AudioModerationResultResult", string(data)}, " ")
}

type AudioModerationResultResultSuggestion struct {
	value string
}

type AudioModerationResultResultSuggestionEnum struct {
	BLOCK  AudioModerationResultResultSuggestion
	PASS   AudioModerationResultResultSuggestion
	REVIEW AudioModerationResultResultSuggestion
}

func GetAudioModerationResultResultSuggestionEnum() AudioModerationResultResultSuggestionEnum {
	return AudioModerationResultResultSuggestionEnum{
		BLOCK: AudioModerationResultResultSuggestion{
			value: "block",
		},
		PASS: AudioModerationResultResultSuggestion{
			value: "pass",
		},
		REVIEW: AudioModerationResultResultSuggestion{
			value: "review",
		},
	}
}

func (c AudioModerationResultResultSuggestion) Value() string {
	return c.value
}

func (c AudioModerationResultResultSuggestion) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *AudioModerationResultResultSuggestion) UnmarshalJSON(b []byte) error {
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
