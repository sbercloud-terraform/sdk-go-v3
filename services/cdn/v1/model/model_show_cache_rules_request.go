package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type ShowCacheRulesRequest struct {

	// 加速域名ID。获取方法请参见查询加速域名。
	DomainId string `json:"domain_id" xml:"domain_id"`

	// 当用户开启企业项目功能时，该参数生效，表示查询资源所属项目，\"all\"表示所有项目。注意：当使用子账号调用接口时，该参数必传。
	EnterpriseProjectId *string `json:"enterprise_project_id,omitempty" xml:"enterprise_project_id"`
}

func (o ShowCacheRulesRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowCacheRulesRequest struct{}"
	}

	return strings.Join([]string{"ShowCacheRulesRequest", string(data)}, " ")
}
