package gopher_test

import (
	. "github.com/minifast/gopher"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("ResponseLine", func() {
  var (
    responseLine *ResponseLine
  )

  Describe("String()", func() {
    BeforeEach(func() {
      responseLine = &ResponseLine{
        Type: "type",
        UserName: "userName",
        Selector: "selector",
        Host: "host",
        Port: "port",
      }
    })

    It("contains the type", func() {
      Expect(responseLine.String()).To(ContainSubstring("type"))
    })

    It("contains the user name", func() {
      Expect(responseLine.String()).To(ContainSubstring("userName\t"))
    })

    It("contains the selector", func() {
      Expect(responseLine.String()).To(ContainSubstring("\tselector\t"))
    })

    It("contains the host", func() {
      Expect(responseLine.String()).To(ContainSubstring("\thost"))
    })

    It("contains the port", func() {
      Expect(responseLine.String()).To(ContainSubstring("\tport"))
    })
  })

  Describe("NewCommentResponseLine()", func() {
    BeforeEach(func() {
      responseLine = NewCommentResponseLine("bazinga!")
    })

    It("sets the comment type", func() {
      Expect(responseLine.Type).To(Equal("i"))
    })

    It("sets the user name", func() {
      Expect(responseLine.UserName).To(Equal("bazinga!"))
    })

    It("does not set the selector", func() {
      Expect(responseLine.Selector).To(Equal(""))
    })

    It("sets the host", func() {
      Expect(responseLine.Host).To(Equal("(NULL)"))
    })

    It("sets the port", func() {
      Expect(responseLine.Port).To(Equal("0"))
    })
  })

  Describe("NewDirectoryResponseLine()", func() {
    BeforeEach(func() {
      responseLine = NewDirectoryResponseLine("chocolate", "/great", "home", "4")
    })

    It("sets the directory type", func() {
      Expect(responseLine.Type).To(Equal("1"))
    })

    It("sets the user name", func() {
      Expect(responseLine.UserName).To(Equal("chocolate"))
    })

    It("sets the selector", func() {
      Expect(responseLine.Selector).To(Equal("/great"))
    })

    It("sets the host", func() {
      Expect(responseLine.Host).To(Equal("home"))
    })

    It("sets the port", func() {
      Expect(responseLine.Port).To(Equal("4"))
    })
  })

  Describe("NewFileResponseLine()", func() {
    BeforeEach(func() {
      responseLine = NewFileResponseLine("tacos", "/nomz", "out", "6")
    })

    It("sets the directory type", func() {
      Expect(responseLine.Type).To(Equal("0"))
    })

    It("sets the user name", func() {
      Expect(responseLine.UserName).To(Equal("tacos"))
    })

    It("sets the selector", func() {
      Expect(responseLine.Selector).To(Equal("/nomz"))
    })

    It("sets the host", func() {
      Expect(responseLine.Host).To(Equal("out"))
    })

    It("sets the port", func() {
      Expect(responseLine.Port).To(Equal("6"))
    })
  })

  Describe("NewPromptResponseLine()", func() {
    BeforeEach(func() {
      responseLine = NewPromptResponseLine("tacos", "/nomz", "out", "6")
    })

    It("sets the directory type", func() {
      Expect(responseLine.Type).To(Equal("7"))
    })

    It("sets the user name", func() {
      Expect(responseLine.UserName).To(Equal("tacos"))
    })

    It("sets the selector", func() {
      Expect(responseLine.Selector).To(Equal("/nomz"))
    })

    It("sets the host", func() {
      Expect(responseLine.Host).To(Equal("out"))
    })

    It("sets the port", func() {
      Expect(responseLine.Port).To(Equal("6"))
    })
  })
})
