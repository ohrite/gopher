package gopher_test

import (
	. "github.com/minifast/gopher"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Server", func() {
  var (
    server Server
  )

  BeforeEach(func() {
    server = Server{
      Address: "localhost:666",
    }
  })

  Describe("URL()", func() {
    It("points at localhost", func() {
      Expect(server.URL().Host).To(ContainSubstring("localhost"))
    })

    It("points at the port", func() {
      Expect(server.URL().Host).To(ContainSubstring(":666"))
    })

    It("references the gopher protocol", func() {
      Expect(server.URL().Scheme).To(Equal("gopher"))
    })
  })
})
