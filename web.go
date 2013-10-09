package main

import (
  "os"
  "log"
  "net"
  "strings"
)

import _ "github.com/joho/godotenv/autoload"

func writeWithErrorLog(c net.Conn, data string) {
  _, err := c.Write([]byte(data))
  if (err != nil) {
    log.Fatal(err)
  }
}

func writeWithTab(c net.Conn, data string) {
  writeWithErrorLog(c, data + "\t")
}

func response(c net.Conn, code string, title string, path string, server string, port string) {
  writeWithErrorLog(c, code)
  writeWithTab(c, title)
  writeWithTab(c, path)
  writeWithTab(c, server)
  writeWithTab(c, port)
  writeWithErrorLog(c, "\r\n")
}

func consumeRequest(c net.Conn, path string, arguments []string) {
  writeWithErrorLog(c, "/\r\n")
  response(c, "i", "Tacos are great!", "null", "(FALSE)", "0")
  response(c, "i", "", "null", "(FALSE)", "0")
  response(c, "i", "", "null", "(FALSE)", "0")
  response(c, "i", "I really like them.", "null", "(FALSE)", "0")
  response(c, "i", "You are at: " + path, "null", "(FALSE)", "0")
  writeWithErrorLog(c, ".\r\n")
}

func extractRequest(c net.Conn) (string, []string, error)  {
  buf := make([]byte, 4096)

  n, err := c.Read(buf)
  log.Printf("Read %v bytes", n)

  if (err != nil || n == 0) {
    log.Fatal("Read error in %v", buf)
    return "", nil, err
  }

  parts := strings.Split(string(buf), "\t")
  return parts[0], parts[1:], nil
}

func handleConnection(c net.Conn) {
  path, arguments, err := extractRequest(c)

  if (err != nil) {
    log.Fatal("Request error: %v", err)
  } else {
    consumeRequest(c, path, arguments)
  }

  c.Close()
  log.Printf("Connection from %v closed.", c.RemoteAddr())
}

func listenOrDie(port string) (net.Listener) {
  ln, err := net.Listen("tcp", ":" + port)

  if err != nil {
    panic(err)
  }

  return ln
}

func acceptOrWhatever(ln net.Listener) (net.Conn) {
  conn, err := ln.Accept()

  if err != nil {
    log.Printf("Hit accept error %v", err)
  } else {
    log.Printf("Connection from %v opened.", conn.RemoteAddr())
  }

  return conn
}

func main() {
  port := os.Getenv("PORT")
  ln := listenOrDie(port)

  log.Printf("Listening on gopher://localhost:%v", port)

  for {
    conn := acceptOrWhatever(ln)
    if conn != nil {
      go handleConnection(conn)
    }
  }
}
