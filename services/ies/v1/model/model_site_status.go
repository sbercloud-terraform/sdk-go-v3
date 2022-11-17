package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// 站点状态： - initial：待部署 - deploying：部署中 - available：可用 - unavailable：不可用
type SiteStatus struct {
	value string
}

type SiteStatusEnum struct {
	INITIAL     SiteStatus
	DEPLOYING   SiteStatus
	AVAILABLE   SiteStatus
	UNAVAILABLE SiteStatus
}

func GetSiteStatusEnum() SiteStatusEnum {
	return SiteStatusEnum{
		INITIAL: SiteStatus{
			value: "initial",
		},
		DEPLOYING: SiteStatus{
			value: "deploying",
		},
		AVAILABLE: SiteStatus{
			value: "available",
		},
		UNAVAILABLE: SiteStatus{
			value: "unavailable",
		},
	}
}

func (c SiteStatus) Value() string {
	return c.value
}

func (c SiteStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *SiteStatus) UnmarshalJSON(b []byte) error {
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
