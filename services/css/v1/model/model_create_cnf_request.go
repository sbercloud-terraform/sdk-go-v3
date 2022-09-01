package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type CreateCnfRequest struct {

	// 指定创建配置文件的集群ID。
	ClusterId string `json:"cluster_id" xml:"cluster_id"`

	Body *CreateCnfReq `json:"body,omitempty" xml:"body"`
}

func (o CreateCnfRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateCnfRequest struct{}"
	}

	return strings.Join([]string{"CreateCnfRequest", string(data)}, " ")
}
