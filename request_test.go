package gopher_test

import (
  . "github.com/ohrite/gopher"
  . "github.com/onsi/ginkgo"
  . "github.com/onsi/gomega"

  "bytes"
  "io"
  "bufio"
  "net/url"
)

var _ = Describe("Request", func() {
  var (
    rawBuffer bytes.Buffer
    request *Request
  )

  Describe("NewRequest()", func() {
    var (
      urlObj *url.URL
    )

    BeforeEach(func() {
      urlObj = &url.URL{}
      request = NewRequest(urlObj, "hey it's a body")
    })

    It("sets the url", func() {
      Expect(request.URL).To(Equal(urlObj))
    })

    It("sets the body", func() {
      Expect(request.Body).To(Equal("hey it's a body"))
    })
  })


  Describe("ReadRequest()", func() {
    BeforeEach(func() {
      rawBuffer.Reset()
      rawBuffer.WriteString("/tacos\tson")
      request, _ = ReadRequest(bufio.NewReader(&rawBuffer))
    })

    It("reads the path into the url", func() {
      Expect(request.URL.Path).To(Equal("/tacos"))
    })

    It("reads the parameter into the body", func() {
      Expect(request.Body).To(Equal("son"))
    })
  })

  Describe("ReadRequestString()", func() {
    Context("when there is no error", func () {
      BeforeEach(func() {
        rawBuffer.Reset()
        rawBuffer.WriteString("/sneaky  drugboat\nhidden tacos go here")
      })

      It("returns the first line of the buffer", func() {
        output, _ := ReadRequestString(bufio.NewReader(&rawBuffer))
        Expect(output).To(Equal("/sneaky  drugboat"))
      })

      It("does not have an error", func () {
        _, err := ReadRequestString(bufio.NewReader(&rawBuffer))
        Expect(err).To(BeNil())
      })
    })

    Context("when there is an unexpected end-of-file", func () {
      BeforeEach(func() {
        rawBuffer.Reset()
      })

      It("passes through an unexpected end-of-file error", func () {
        _, err := ReadRequestString(bufio.NewReader(&rawBuffer))
        Expect(err).To(Equal(io.EOF))
      })
    })
  })

  Describe("ExtractURLAndBody()", func() {
    var (
      urlObj *url.URL
      outputBody string
    )

    BeforeEach(func() {
      urlObj, outputBody = ExtractURLAndBody("/steaks\toh\tfine&ok?")
    })

    It("sets the path", func() {
      Expect(urlObj.Path).To(Equal("/steaks"))
    })

    It("sets the body", func() {
      Expect(outputBody).To(Equal("oh\tfine&ok?"))
    })
  })

  Describe("ParseRequestString()", func() {
    Context("when there is no parameter", func() {
      It("returns a blank line", func() {
        _, body := ParseRequestString("/what")
        Expect(body).To(Equal(""))
      })
    })

    Context("when there is a parameter", func() {
      It("returns the parameter", func() {
        _, body := ParseRequestString("/what\teeth")
        Expect(body).To(Equal("eeth"))
      })
    })
  })
})
