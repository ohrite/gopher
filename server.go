package gopher

import (
  "bufio"
  "net"
  "net/url"
)

type Server struct {
  Address string
  handler func(net.Conn, *Request)
}

func NewServer(address string) *Server {
  return &Server{
    Address: address,
  }
}

func (server *Server) URL() *url.URL {
  return &url.URL{
    Host: server.Address,
    Scheme: "gopher",
  }
}

func (server *Server) ListenAndServe(handler func(net.Conn, *Request)) (error) {
  listener, err := net.Listen("tcp", server.Address)
  if err != nil {
    return err
  }

  server.handler = handler

  return server.Serve(listener)
}

func (server *Server) Serve(listener net.Listener) (err error) {
  for {
    err = server.Accept(listener)
    if err != nil {
      break
    }
  }
  return err
}

func (server *Server) Accept(listener net.Listener) (error) {
  connection, err := listener.Accept()
  if err == nil {
    go server.HandleConnection(connection)
  }

  return err
}

func (server *Server) HandleConnection(connection net.Conn) {
  request, err := server.ConnectRequest(connection)
  if err == nil {
    defer connection.Close()
    server.HandleRequest(connection, request)
  }
}

func (server *Server) ConnectRequest(connection net.Conn) (*Request, error) {
  return ReadRequest(bufio.NewReader(connection))
}

func (server *Server) HandleRequest(connection net.Conn, request *Request) {
  server.handler(connection, request)
}
