package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type ShowGaussMySqlProxyResponse struct {
	Proxy *MysqlProxy `json:"proxy,omitempty" xml:"proxy"`

	MasterNode *MysqlProxyNode `json:"master_node,omitempty" xml:"master_node"`

	// 只读节点信息。
	ReadonlyNodes  *[]MysqlProxyNode `json:"readonly_nodes,omitempty" xml:"readonly_nodes"`
	HttpStatusCode int               `json:"-"`
}

func (o ShowGaussMySqlProxyResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowGaussMySqlProxyResponse struct{}"
	}

	return strings.Join([]string{"ShowGaussMySqlProxyResponse", string(data)}, " ")
}
