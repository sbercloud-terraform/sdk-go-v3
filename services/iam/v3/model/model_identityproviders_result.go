/*
 * IAM
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 */

package model

import (
	"encoding/json"

	"strings"
)

//
type IdentityprovidersResult struct {
	// 身份提供商ID。
	Id string `json:"id"`
	// 身份提供商描述信息。
	Description string `json:"description"`
	// 身份提供商是否启用，true为启用，false为停用，默认为false。
	Enabled bool `json:"enabled"`
	// 身份提供商的联邦用户ID列表。
	RemoteIds []string                `json:"remote_ids"`
	Links     *IdentityprovidersLinks `json:"links"`
}

func (o IdentityprovidersResult) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "IdentityprovidersResult struct{}"
	}

	return strings.Join([]string{"IdentityprovidersResult", string(data)}, " ")
}
