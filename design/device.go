package design

import (
	. "github.com/goadesign/goa/design"
	. "github.com/goadesign/goa/design/apidsl"
)

var _ = Resource("device", func() {

	DefaultMedia(DeviceMedia)

	BasePath("/")

	Action("list", func() {
		Routing(GET("/devices"))
		Description("List devices")
		Response(OK, func() {
			Media(CollectionOf(DeviceMedia, func() {
				View("default")
				View("tiny")
			}))
		})
		Response(BadRequest, ErrorMedia)
	})

	Action("read", func() {
		Routing(GET("/devices/:id"))
		Description("read device")
		Params(func() {
			Param("id", Integer, "Device id")
		})
		Response(OK, DeviceMedia)
		Response(NotFound)
		Response(BadRequest, ErrorMedia)
	})

	Action("create", func() {
		Routing(POST("/devices"))
		Description("Create new device")
		Payload(CreateDevicePayload)
		Response(Created, "/devices/[0-9]+")
		Response(BadRequest, ErrorMedia)
	})

	Action("update", func() {
		Routing(PUT("/devices/:id"))
		Description("Update device")
		Params(func() {
			Param("id", Integer, "Device id")
		})
		Payload(UpdateDevicePayload)
		Response(NoContent)
		Response(NotFound)
		Response(BadRequest, ErrorMedia)
	})

	Action("delete", func() {
		Routing(DELETE("/devices/:id"))
		Params(func() {
			Param("id", Integer, "Device id")
		})
		Response(NoContent)
		Response(NotFound)
		Response(BadRequest, ErrorMedia)
	})
})

// CreateDevicePayload is supplied when creating a new device
var CreateDevicePayload = Type("CreateDevicePayload", func() {
	Attribute("name")
	Required("name")
})

// DeviceMedia that a device belongs to
var DeviceMedia = MediaType("application/vnd.device+json", func() {

	Description("A device")

	Attributes(func() {
		Attribute("id", Integer, "Device id")
		Attribute("href", String, "API href of device", func() {
			Example("/devices/1")
		})
		Attribute("name", String, "Name", func() {
			MinLength(1)
			MaxLength(255)
		})
		Required("id", "name")
	})

	View("default", func() {
		Attribute("id")
		Attribute("href")
		Attribute("name")
	})

	View("tiny", func() {
		Attribute("id")
		Attribute("href")
		Attribute("name")
	})

	View("link", func() {
		Attribute("id")
		Attribute("href")
	})
})
