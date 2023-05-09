package design

import (
	. "goa.design/goa/v3/dsl"
)

var _ = API("PubSubServer REST API", func() {
	Title("PubSubServer REST API")
	Description("REST API for PubSubServer")
	Server("PubSubServer", func() {
		Host("localhost", func() {
			URI("http://localhost:8080")
		})
	})
})

var _ = Service("PubSubServer", func() {

	Description("The PubSubServer service performs operations on messages.")

	Method("publish", func() {
		Payload(func() {
			Attribute("roomID", String, func() {
				Description("aaa")
			})
		})
	})

})
