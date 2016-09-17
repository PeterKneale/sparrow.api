//************************************************************************//
// API "Gateway API": Application Resource Href Factories
//
// Generated with goagen v1.0.0, command line:
// $ goagen
// --design=github.com/simplicate/mango/gateway/design
// --out=$(GOPATH)\src\github.com\Simplicate\Mango\gateway
// --version=v1.0.0
//
// The content of this file is auto-generated, DO NOT MODIFY
//************************************************************************//

package app

import (
	"fmt"
	"strings"
)

// AccountHref returns the resource href.
func AccountHref(accountID interface{}) string {
	paramaccountID := strings.TrimLeftFunc(fmt.Sprintf("%v", accountID), func(r rune) bool { return r == '/' })
	return fmt.Sprintf("/accounts/%v", paramaccountID)
}

// UserHref returns the resource href.
func UserHref(userID interface{}) string {
	paramuserID := strings.TrimLeftFunc(fmt.Sprintf("%v", userID), func(r rune) bool { return r == '/' })
	return fmt.Sprintf("/users/%v", paramuserID)
}
