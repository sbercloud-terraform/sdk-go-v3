package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type ShowJobDetailResponse struct {

	// 作业下发人数
	AcceptJobNum *int32 `json:"accept_job_num,omitempty" xml:"accept_job_num"`

	// 作业截止时间, 日期格式：yyyy-MM-dd HH:mm:ss
	EndTime *string `json:"end_time,omitempty" xml:"end_time"`

	// 作业答案是否公布(unpublish:表示未公布答案, publish:表示已公布答案)
	IsAnswerVisibility *string `json:"is_answer_visibility,omitempty" xml:"is_answer_visibility"`

	// 作业成绩是否公布(unpublish:表示未公布成绩, publish:表示已公布成绩)
	IsScoreVisibility *string `json:"is_score_visibility,omitempty" xml:"is_score_visibility"`

	// 作业均分
	AverageScore *string `json:"average_score,omitempty" xml:"average_score"`

	// 老师手动评分人数
	ScoreJobNum *int32 `json:"score_job_num,omitempty" xml:"score_job_num"`

	// 作业提交人数
	SubmitJobNum   *int32 `json:"submit_job_num,omitempty" xml:"submit_job_num"`
	HttpStatusCode int    `json:"-"`
}

func (o ShowJobDetailResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowJobDetailResponse struct{}"
	}

	return strings.Join([]string{"ShowJobDetailResponse", string(data)}, " ")
}
