package gopher

import (
  "net/url"
  "bufio"
  "strings"
  "net/textproto"
)

type Request struct {
  URL url.URL
  Body string
}

func NewRequest(url url.URL, body string) (request *Request) {
  return &Request{
    URL: url,
    Body: body,
  }
}

func ReadRequest(buffer *bufio.Reader) (request *Request, err error) {
  requestString, err := ReadRequestString(buffer)
  if err != nil {
    return nil, err
  }

  url, body := ExtractURLAndBody(requestString)
  request = NewRequest(url, body)

  return request, nil
}

func ReadRequestString(buffer *bufio.Reader) (string, error) {
  reader := textproto.NewReader(buffer)
  return reader.ReadLine()
}

func ExtractURLAndBody(requestString string) (url url.URL, body string) {
  url.Path, body = ParseRequestString(requestString)
  return url, body
}

func ParseRequestString(requestString string) (string, body string) {
  parts := strings.Split(requestString, "\t")

  body = ""
  if len(parts) > 1 {
    body = strings.Join(parts[1:], "\t")
  }

  return parts[0], body
}
