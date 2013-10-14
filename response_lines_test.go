package gopher_test

import (
	. "github.com/ohrite/gopher"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

  "net/textproto"
  "bytes"
  "bufio"
)

var _ = Describe("ResponseLines", func() {
  var (
    responseLine *ResponseLine
    responseLines *ResponseLines
  )

  BeforeEach(func() {
    responseLines = new(ResponseLines)
    responseLine = NewCommentResponseLine("dentistry")
  })

  Describe("Len()", func() {
    Context("when no response lines have been added", func() {
      It("is zero", func() {
        Expect(responseLines.Len()).To(Equal(0))
      })
    })

    Context("after a response line has been added", func() {
      BeforeEach(func() {
        responseLines.AddResponseLine(responseLine)
      })

      It("is zero", func() {
        Expect(responseLines.Len()).To(Equal(1))
      })
    })
  })

  Describe("WriteResponseLines()", func() {
    var (
      buffer bytes.Buffer
      writer *textproto.Writer
    )

    BeforeEach(func() {
      writer = textproto.NewWriter(bufio.NewWriter(&buffer))
      responseLines.AddResponseLine(responseLine)
    })

    It("writes the response line", func() {
      responseLines.WriteResponseLines(writer)
      Expect(buffer.String()).To(ContainSubstring(responseLine.String()))
    })
  })
})
