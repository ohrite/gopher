package gopher

import (
  "net"
  "bufio"
  "net/textproto"
)

type Response struct {
  responseLines ResponseLines
}

func (response *Response) AddResponseLine(line *ResponseLine) {
  response.responseLines.AddResponseLine(line)
}

func (response *Response) WriteResponse(connection net.Conn) {
  writer := response.newResponseWriter(connection)
  response.responseLines.WriteResponseLines(writer)
  dotWriter := writer.DotWriter()
  dotWriter.Close()
}

func (response *Response) newResponseWriter(connection net.Conn) (*textproto.Writer) {
  buffer := bufio.NewWriter(connection)
  return textproto.NewWriter(buffer)
}
