package model

import (
	"encoding/json"

	"strings"
)

type UpdateCoverByThumbnailReq struct {
	// 截图文件的URL。 需要根据媒资ID调用[查询媒资详细信息](https://support.huaweicloud.com/api-vod/vod_04_0202.html)接口获取媒资的截图文件URL。

	ThumbnailUrl string `json:"thumbnailUrl"`
}

func (o UpdateCoverByThumbnailReq) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "UpdateCoverByThumbnailReq struct{}"
	}

	return strings.Join([]string{"UpdateCoverByThumbnailReq", string(data)}, " ")
}
