package gopher

import (
  "net/textproto"
)

type ResponseLines struct{
  lines []ResponseLine
}

func (responseLines *ResponseLines) AddResponseLine(line *ResponseLine) {
  responseLines.lines = append(responseLines.lines, *line)
}

func (responseLines *ResponseLines) Len() int {
  return len(responseLines.lines)
}

func (responseLines *ResponseLines) WriteResponseLines(writer *textproto.Writer) {
  for _, line := range responseLines.lines {
    writer.PrintfLine(line.String())
  }
}
