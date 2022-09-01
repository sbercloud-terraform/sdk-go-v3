package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type KeystoneCreateUserTokenByPasswordAndMfaResponse struct {
	Token *TokenResult `json:"token,omitempty" xml:"token"`

	XSubjectToken  *string `json:"X-Subject-Token,omitempty" xml:"X-Subject-Token"`
	HttpStatusCode int     `json:"-"`
}

func (o KeystoneCreateUserTokenByPasswordAndMfaResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "KeystoneCreateUserTokenByPasswordAndMfaResponse struct{}"
	}

	return strings.Join([]string{"KeystoneCreateUserTokenByPasswordAndMfaResponse", string(data)}, " ")
}
