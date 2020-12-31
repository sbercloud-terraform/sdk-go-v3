/*
 * ECS
 *
 * ECS Open API
 *
 */

package model

import (
	"encoding/json"

	"strings"
)

// This is a auto create Body Object
type BatchAddServerNicsRequestBody struct {
	// 需要添加的网卡参数列表。
	Nics []BatchAddServerNicOption `json:"nics"`
}

func (o BatchAddServerNicsRequestBody) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "BatchAddServerNicsRequestBody struct{}"
	}

	return strings.Join([]string{"BatchAddServerNicsRequestBody", string(data)}, " ")
}
