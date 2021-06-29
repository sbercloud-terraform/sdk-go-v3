package model

import (
	"encoding/json"

	"strings"
)

type CreatePreheatingAssetReq struct {
	// 已发布媒资的ID。

	AssetId *string `json:"asset_id,omitempty"`
	// 已发布媒资的播放URL列表，一次最多只能预热10个URL。

	Urls *[]string `json:"urls,omitempty"`
}

func (o CreatePreheatingAssetReq) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "CreatePreheatingAssetReq struct{}"
	}

	return strings.Join([]string{"CreatePreheatingAssetReq", string(data)}, " ")
}
