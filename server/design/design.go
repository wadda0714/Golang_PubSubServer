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
			Attribute("roomName", String, "roomName to publish", func() {
				Example("room1")
			})
		})

		Result(String)

		HTTP(func() {
			GET("/publish/{roomName}")
			Response(StatusOK)
		})

	})

	Method("subscribe", func() {
		Payload(func() {
			Attribute("roomName", String, "roomName to subscribe", func() {
				Example("room1")
			})

		})

		Result(String)

		HTTP(func() {
			GET("/subscribe/{roomName}")
			Response(StatusOK)
		})

	})

	Method("sendMessage", func() {
		Payload(func() {
			Attribute("roomName", String, "roomName to publish", func() {
				Example("room1")
			})
			Attribute("message", String, "message to publish", func() {
				Example("hello")
			})
		})
		Result(String)

		HTTP(func() {
			GET("/SendMessage/{roomName}/{message}")
			Response(StatusOK)
		})
	})

})
