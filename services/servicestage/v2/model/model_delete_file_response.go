/*
 * ServiceStage
 *
 * ServiceStage的API,包括应用管理和仓库授权管理
 *
 */

package model

import (
	"encoding/json"

	"strings"
)

// Response Object
type DeleteFileResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o DeleteFileResponse) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "DeleteFileResponse struct{}"
	}

	return strings.Join([]string{"DeleteFileResponse", string(data)}, " ")
}
