/*
 * BCS
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 */

package model

import (
	"encoding/json"

	"strings"
)

// couchDB信息
type CreateRequestBodyCouchdbInfo struct {
	// couchDB用户名
	UserName *string `json:"user_name,omitempty"`
	// couchDB密码
	Password *string `json:"password,omitempty"`
}

func (o CreateRequestBodyCouchdbInfo) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "CreateRequestBodyCouchdbInfo struct{}"
	}

	return strings.Join([]string{"CreateRequestBodyCouchdbInfo", string(data)}, " ")
}
