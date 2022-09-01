package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type Geo struct {

	// 纬度。
	Latitude float32 `json:"latitude,omitempty" xml:"latitude"`

	// 经度。
	Longitude float32 `json:"longitude,omitempty" xml:"longitude"`

	// 城市编码。
	CityCode *string `json:"city_code,omitempty" xml:"city_code"`

	// 国家简码ISO 3166-1 alpha-2，例如：CN、US、DE、IT、SG。
	CountryCode *string `json:"country_code,omitempty" xml:"country_code"`
}

func (o Geo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Geo struct{}"
	}

	return strings.Join([]string{"Geo", string(data)}, " ")
}
