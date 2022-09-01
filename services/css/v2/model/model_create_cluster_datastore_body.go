package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 引擎类型，目前只支持elasticsearch。
type CreateClusterDatastoreBody struct {

	// CSS集群引擎版本号。详细请参考CSS[支持的集群版本](css_03_0056.xml)。
	Version string `json:"version" xml:"version"`

	// 引擎类型，支持elasticsearch。
	Type string `json:"type" xml:"type"`
}

func (o CreateClusterDatastoreBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateClusterDatastoreBody struct{}"
	}

	return strings.Join([]string{"CreateClusterDatastoreBody", string(data)}, " ")
}
