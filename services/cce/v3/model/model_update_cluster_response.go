package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type UpdateClusterResponse struct {

	// API类型，固定值“Cluster”或“cluster”，该值不可修改。
	Kind *string `json:"kind,omitempty" xml:"kind"`

	// API版本，固定值“v3”，该值不可修改。
	ApiVersion *string `json:"apiVersion,omitempty" xml:"apiVersion"`

	Metadata *ClusterMetadata `json:"metadata,omitempty" xml:"metadata"`

	Spec *ClusterSpec `json:"spec,omitempty" xml:"spec"`

	Status         *ClusterStatus `json:"status,omitempty" xml:"status"`
	HttpStatusCode int            `json:"-"`
}

func (o UpdateClusterResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateClusterResponse struct{}"
	}

	return strings.Join([]string{"UpdateClusterResponse", string(data)}, " ")
}
