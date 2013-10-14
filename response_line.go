package gopher

import (
  "fmt"
)

type ResponseLine struct {
  Type string
  UserName string
  Selector string
  Host string
  Port string
}

func (line *ResponseLine) String() string {
  return fmt.Sprintf("%v%v\t%v\t%v\t%v", line.Type, line.UserName, line.Selector, line.Host, line.Port)
}

func NewCommentResponseLine(comment string) *ResponseLine {
  return &ResponseLine{
    Type: "i",
    UserName: comment,
    Host: "(NULL)",
    Port: "0",
  }
}

func NewDirectoryResponseLine(userName, selector, host, port string) *ResponseLine {
  return &ResponseLine{
    Type: "1",
    UserName: userName,
    Selector: selector,
    Host: host,
    Port: port,
  }
}

func NewFileResponseLine(userName, selector, host, port string) *ResponseLine {
  return &ResponseLine{
    Type: "0",
    UserName: userName,
    Selector: selector,
    Host: host,
    Port: port,
  }
}

func NewPromptResponseLine(userName, selector, host, port string) *ResponseLine {
  return &ResponseLine{
    Type: "7",
    UserName: userName,
    Selector: selector,
    Host: host,
    Port: port,
  }
}
