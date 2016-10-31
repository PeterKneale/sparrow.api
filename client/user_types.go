//************************************************************************//
// API "Sparrow API": Application User Types
//
// Generated with goagen v1.0.0, command line:
// $ goagen
// --design=github.com/simplicate/sparrow.api/design
// --out=$(GOPATH)\src\github.com\simplicate\sparrow.api
// --version=v1.0.0
//
// The content of this file is auto-generated, DO NOT MODIFY
//************************************************************************//

package client

import "github.com/goadesign/goa"

// createUserPayload user type.
type createUserPayload struct {
	FirstName *string `form:"first_name,omitempty" json:"first_name,omitempty" xml:"first_name,omitempty"`
	LastName  *string `form:"last_name,omitempty" json:"last_name,omitempty" xml:"last_name,omitempty"`
}

// Validate validates the createUserPayload type instance.
func (ut *createUserPayload) Validate() (err error) {
	if ut.FirstName == nil {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "first_name"))
	}
	if ut.LastName == nil {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "last_name"))
	}

	return
}

// Publicize creates CreateUserPayload from createUserPayload
func (ut *createUserPayload) Publicize() *CreateUserPayload {
	var pub CreateUserPayload
	if ut.FirstName != nil {
		pub.FirstName = *ut.FirstName
	}
	if ut.LastName != nil {
		pub.LastName = *ut.LastName
	}
	return &pub
}

// CreateUserPayload user type.
type CreateUserPayload struct {
	FirstName string `form:"first_name" json:"first_name" xml:"first_name"`
	LastName  string `form:"last_name" json:"last_name" xml:"last_name"`
}

// Validate validates the CreateUserPayload type instance.
func (ut *CreateUserPayload) Validate() (err error) {
	if ut.FirstName == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "first_name"))
	}
	if ut.LastName == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "last_name"))
	}

	return
}

// updateAccountPayload user type.
type updateAccountPayload struct {
	Name *string `form:"name,omitempty" json:"name,omitempty" xml:"name,omitempty"`
}

// Publicize creates UpdateAccountPayload from updateAccountPayload
func (ut *updateAccountPayload) Publicize() *UpdateAccountPayload {
	var pub UpdateAccountPayload
	if ut.Name != nil {
		pub.Name = ut.Name
	}
	return &pub
}

// UpdateAccountPayload user type.
type UpdateAccountPayload struct {
	Name *string `form:"name,omitempty" json:"name,omitempty" xml:"name,omitempty"`
}

// updateUserPayload user type.
type updateUserPayload struct {
	FirstName *string `form:"first_name,omitempty" json:"first_name,omitempty" xml:"first_name,omitempty"`
	LastName  *string `form:"last_name,omitempty" json:"last_name,omitempty" xml:"last_name,omitempty"`
}

// Publicize creates UpdateUserPayload from updateUserPayload
func (ut *updateUserPayload) Publicize() *UpdateUserPayload {
	var pub UpdateUserPayload
	if ut.FirstName != nil {
		pub.FirstName = ut.FirstName
	}
	if ut.LastName != nil {
		pub.LastName = ut.LastName
	}
	return &pub
}

// UpdateUserPayload user type.
type UpdateUserPayload struct {
	FirstName *string `form:"first_name,omitempty" json:"first_name,omitempty" xml:"first_name,omitempty"`
	LastName  *string `form:"last_name,omitempty" json:"last_name,omitempty" xml:"last_name,omitempty"`
}
