package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type RunImageMediaTaggingResponse struct {
	Result         *ImageMediaTaggingResponseResult `json:"result,omitempty"`
	HttpStatusCode int                              `json:"-"`
}

func (o RunImageMediaTaggingResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RunImageMediaTaggingResponse struct{}"
	}

	return strings.Join([]string{"RunImageMediaTaggingResponse", string(data)}, " ")
}
