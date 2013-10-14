package gopher_test

import (
	. "github.com/ohrite/gopher"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Response", func() {
  var (
    response *Response
    connection *testConnection
  )

  BeforeEach(func() {
    response = new(Response)
    connection = new(testConnection)
  })

  Describe("WriteResponse()", func() {
    Context("when there are no response lines", func(){
      BeforeEach(func() {
        response.WriteResponse(connection)
      })

      It("contains a dot-newline ending", func() {
        Expect(connection.WriteBuf.String()).To(ContainSubstring(".\r\n"))
      })
    })

    Context("when a response line has been added", func(){
      var (
        responseLine *ResponseLine
      )

      BeforeEach(func() {
        responseLine = NewCommentResponseLine("tacos")
        response.AddResponseLine(responseLine)
        response.WriteResponse(connection)
      })

      It("contains a response line", func() {
        Expect(connection.WriteBuf.String()).To(ContainSubstring(responseLine.String()))
      })
    })
  })
})
