//************************************************************************//
// API "Sparrow API": Application Media Types
//
// Generated with goagen v1.0.0, command line:
// $ goagen
// --design=github.com/simplicate/sparrow.api/design
// --out=C:\dev\go\src\github.com\simplicate\sparrow.api
// --version=v1.0.0
//
// The content of this file is auto-generated, DO NOT MODIFY
//************************************************************************//

package client

import (
	"github.com/goadesign/goa"
	"net/http"
	"unicode/utf8"
)

// An account (default view)
//
// Identifier: application/vnd.account+json; view=default
type Account struct {
	// Account name
	Name string `form:"name" json:"name" xml:"name"`
}

// Validate validates the Account media type instance.
func (mt *Account) Validate() (err error) {
	if mt.Name == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "name"))
	}

	if utf8.RuneCountInString(mt.Name) < 1 {
		err = goa.MergeErrors(err, goa.InvalidLengthError(`response.name`, mt.Name, utf8.RuneCountInString(mt.Name), 1, true))
	}
	if utf8.RuneCountInString(mt.Name) > 255 {
		err = goa.MergeErrors(err, goa.InvalidLengthError(`response.name`, mt.Name, utf8.RuneCountInString(mt.Name), 255, false))
	}
	return
}

// DecodeAccount decodes the Account instance encoded in resp body.
func (c *Client) DecodeAccount(resp *http.Response) (*Account, error) {
	var decoded Account
	err := c.Decoder.Decode(&decoded, resp.Body, resp.Header.Get("Content-Type"))
	return &decoded, err
}

// DecodeErrorResponse decodes the ErrorResponse instance encoded in resp body.
func (c *Client) DecodeErrorResponse(resp *http.Response) (*goa.ErrorResponse, error) {
	var decoded goa.ErrorResponse
	err := c.Decoder.Decode(&decoded, resp.Body, resp.Header.Get("Content-Type"))
	return &decoded, err
}

// A user (default view)
//
// Identifier: application/vnd.user+json; view=default
type User struct {
	// First name
	FirstName string `form:"first_name" json:"first_name" xml:"first_name"`
	// API href of user
	Href string `form:"href" json:"href" xml:"href"`
	// User id
	ID int `form:"id" json:"id" xml:"id"`
	// Last name
	LastName string `form:"last_name" json:"last_name" xml:"last_name"`
	// name
	Name *string `form:"name,omitempty" json:"name,omitempty" xml:"name,omitempty"`
}

// Validate validates the User media type instance.
func (mt *User) Validate() (err error) {
	if mt.Href == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "href"))
	}
	if mt.FirstName == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "first_name"))
	}
	if mt.LastName == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "last_name"))
	}

	if utf8.RuneCountInString(mt.FirstName) < 1 {
		err = goa.MergeErrors(err, goa.InvalidLengthError(`response.first_name`, mt.FirstName, utf8.RuneCountInString(mt.FirstName), 1, true))
	}
	if utf8.RuneCountInString(mt.FirstName) > 255 {
		err = goa.MergeErrors(err, goa.InvalidLengthError(`response.first_name`, mt.FirstName, utf8.RuneCountInString(mt.FirstName), 255, false))
	}
	if utf8.RuneCountInString(mt.LastName) < 1 {
		err = goa.MergeErrors(err, goa.InvalidLengthError(`response.last_name`, mt.LastName, utf8.RuneCountInString(mt.LastName), 1, true))
	}
	if utf8.RuneCountInString(mt.LastName) > 255 {
		err = goa.MergeErrors(err, goa.InvalidLengthError(`response.last_name`, mt.LastName, utf8.RuneCountInString(mt.LastName), 255, false))
	}
	if mt.Name != nil {
		if utf8.RuneCountInString(*mt.Name) < 1 {
			err = goa.MergeErrors(err, goa.InvalidLengthError(`response.name`, *mt.Name, utf8.RuneCountInString(*mt.Name), 1, true))
		}
	}
	if mt.Name != nil {
		if utf8.RuneCountInString(*mt.Name) > 255 {
			err = goa.MergeErrors(err, goa.InvalidLengthError(`response.name`, *mt.Name, utf8.RuneCountInString(*mt.Name), 255, false))
		}
	}
	return
}

// A user (link view)
//
// Identifier: application/vnd.user+json; view=link
type UserLink struct {
	// API href of user
	Href string `form:"href" json:"href" xml:"href"`
	// User id
	ID int `form:"id" json:"id" xml:"id"`
}

// Validate validates the UserLink media type instance.
func (mt *UserLink) Validate() (err error) {
	if mt.Href == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "href"))
	}

	return
}

// A user (tiny view)
//
// Identifier: application/vnd.user+json; view=tiny
type UserTiny struct {
	// API href of user
	Href string `form:"href" json:"href" xml:"href"`
	// User id
	ID int `form:"id" json:"id" xml:"id"`
	// name
	Name *string `form:"name,omitempty" json:"name,omitempty" xml:"name,omitempty"`
}

// Validate validates the UserTiny media type instance.
func (mt *UserTiny) Validate() (err error) {
	if mt.Href == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "href"))
	}

	if mt.Name != nil {
		if utf8.RuneCountInString(*mt.Name) < 1 {
			err = goa.MergeErrors(err, goa.InvalidLengthError(`response.name`, *mt.Name, utf8.RuneCountInString(*mt.Name), 1, true))
		}
	}
	if mt.Name != nil {
		if utf8.RuneCountInString(*mt.Name) > 255 {
			err = goa.MergeErrors(err, goa.InvalidLengthError(`response.name`, *mt.Name, utf8.RuneCountInString(*mt.Name), 255, false))
		}
	}
	return
}

// DecodeUser decodes the User instance encoded in resp body.
func (c *Client) DecodeUser(resp *http.Response) (*User, error) {
	var decoded User
	err := c.Decoder.Decode(&decoded, resp.Body, resp.Header.Get("Content-Type"))
	return &decoded, err
}

// DecodeUserLink decodes the UserLink instance encoded in resp body.
func (c *Client) DecodeUserLink(resp *http.Response) (*UserLink, error) {
	var decoded UserLink
	err := c.Decoder.Decode(&decoded, resp.Body, resp.Header.Get("Content-Type"))
	return &decoded, err
}

// DecodeUserTiny decodes the UserTiny instance encoded in resp body.
func (c *Client) DecodeUserTiny(resp *http.Response) (*UserTiny, error) {
	var decoded UserTiny
	err := c.Decoder.Decode(&decoded, resp.Body, resp.Header.Get("Content-Type"))
	return &decoded, err
}

// UserCollection is the media type for an array of User (default view)
//
// Identifier: application/vnd.user+json; type=collection; view=default
type UserCollection []*User

// Validate validates the UserCollection media type instance.
func (mt UserCollection) Validate() (err error) {
	for _, e := range mt {
		if e.Href == "" {
			err = goa.MergeErrors(err, goa.MissingAttributeError(`response[*]`, "href"))
		}
		if e.FirstName == "" {
			err = goa.MergeErrors(err, goa.MissingAttributeError(`response[*]`, "first_name"))
		}
		if e.LastName == "" {
			err = goa.MergeErrors(err, goa.MissingAttributeError(`response[*]`, "last_name"))
		}

		if utf8.RuneCountInString(e.FirstName) < 1 {
			err = goa.MergeErrors(err, goa.InvalidLengthError(`response[*].first_name`, e.FirstName, utf8.RuneCountInString(e.FirstName), 1, true))
		}
		if utf8.RuneCountInString(e.FirstName) > 255 {
			err = goa.MergeErrors(err, goa.InvalidLengthError(`response[*].first_name`, e.FirstName, utf8.RuneCountInString(e.FirstName), 255, false))
		}
		if utf8.RuneCountInString(e.LastName) < 1 {
			err = goa.MergeErrors(err, goa.InvalidLengthError(`response[*].last_name`, e.LastName, utf8.RuneCountInString(e.LastName), 1, true))
		}
		if utf8.RuneCountInString(e.LastName) > 255 {
			err = goa.MergeErrors(err, goa.InvalidLengthError(`response[*].last_name`, e.LastName, utf8.RuneCountInString(e.LastName), 255, false))
		}
		if e.Name != nil {
			if utf8.RuneCountInString(*e.Name) < 1 {
				err = goa.MergeErrors(err, goa.InvalidLengthError(`response[*].name`, *e.Name, utf8.RuneCountInString(*e.Name), 1, true))
			}
		}
		if e.Name != nil {
			if utf8.RuneCountInString(*e.Name) > 255 {
				err = goa.MergeErrors(err, goa.InvalidLengthError(`response[*].name`, *e.Name, utf8.RuneCountInString(*e.Name), 255, false))
			}
		}
	}
	return
}

// UserCollection is the media type for an array of User (tiny view)
//
// Identifier: application/vnd.user+json; type=collection; view=tiny
type UserTinyCollection []*UserTiny

// Validate validates the UserTinyCollection media type instance.
func (mt UserTinyCollection) Validate() (err error) {
	for _, e := range mt {
		if e.Href == "" {
			err = goa.MergeErrors(err, goa.MissingAttributeError(`response[*]`, "href"))
		}

		if e.Name != nil {
			if utf8.RuneCountInString(*e.Name) < 1 {
				err = goa.MergeErrors(err, goa.InvalidLengthError(`response[*].name`, *e.Name, utf8.RuneCountInString(*e.Name), 1, true))
			}
		}
		if e.Name != nil {
			if utf8.RuneCountInString(*e.Name) > 255 {
				err = goa.MergeErrors(err, goa.InvalidLengthError(`response[*].name`, *e.Name, utf8.RuneCountInString(*e.Name), 255, false))
			}
		}
	}
	return
}

// DecodeUserCollection decodes the UserCollection instance encoded in resp body.
func (c *Client) DecodeUserCollection(resp *http.Response) (UserCollection, error) {
	var decoded UserCollection
	err := c.Decoder.Decode(&decoded, resp.Body, resp.Header.Get("Content-Type"))
	return decoded, err
}

// DecodeUserTinyCollection decodes the UserTinyCollection instance encoded in resp body.
func (c *Client) DecodeUserTinyCollection(resp *http.Response) (UserTinyCollection, error) {
	var decoded UserTinyCollection
	err := c.Decoder.Decode(&decoded, resp.Body, resp.Header.Get("Content-Type"))
	return decoded, err
}
