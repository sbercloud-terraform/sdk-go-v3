package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type CreateDeploymentJobsParams struct {
	Function *FgsDeploymentJobsParam `json:"function,omitempty" xml:"function"`

	Cci *CciDeploymentJobsParam `json:"cci,omitempty" xml:"cci"`
}

func (o CreateDeploymentJobsParams) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateDeploymentJobsParams struct{}"
	}

	return strings.Join([]string{"CreateDeploymentJobsParams", string(data)}, " ")
}
