/*
 * CES
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 */

package model

import (
	"encoding/json"

	"strings"
)

// Response Object
type UpdateAlarmResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o UpdateAlarmResponse) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "UpdateAlarmResponse struct{}"
	}

	return strings.Join([]string{"UpdateAlarmResponse", string(data)}, " ")
}
