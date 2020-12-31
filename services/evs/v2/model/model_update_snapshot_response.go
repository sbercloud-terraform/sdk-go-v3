/*
 * EVS
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
type UpdateSnapshotResponse struct {
	Snapshot       *SnapshotDetails `json:"snapshot,omitempty"`
	HttpStatusCode int              `json:"-"`
}

func (o UpdateSnapshotResponse) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "UpdateSnapshotResponse struct{}"
	}

	return strings.Join([]string{"UpdateSnapshotResponse", string(data)}, " ")
}
