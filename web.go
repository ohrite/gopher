package main

import (
  "os"
  "log"
  "net"
  "strings"
)

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
  log.Printf(path)
  log.Printf(strings.Join(arguments, "&"))

  writeWithErrorLog(c, "/\r\n")
  response(c, "i", "Tacos are great!", "null", "(FALSE)", "0")
  response(c, "i", "", "null", "(FALSE)", "0")
  response(c, "i", "", "null", "(FALSE)", "0")
  response(c, "i", "I really like them.", "null", "(FALSE)", "0")
  response(c, "i", "You are at: " + path, "null", "(FALSE)", "0")
  writeWithErrorLog(c, ".\r\n")
}

func extractRequest(c net.Conn) (string, []string)  {
  buf := make([]byte, 4096)

  n, err := c.Read(buf)
  if (err != nil || n == 0) {
      log.Fatal(err)
      c.Close()
  }

  parts := strings.Split(string(buf), "\t")
  return parts[0], parts[1:]
}

func handleConnection(c net.Conn) {
  path, arguments := extractRequest(c)
  consumeRequest(c, path, arguments)
  c.Close()
  log.Printf("Connection from %v closed.", c.RemoteAddr())
}

func main() {
  port := os.Getenv("PORT")

  if (len(port) == 0) {
    port = "7070"
  }

  ln, err := net.Listen("tcp", ":" + port)
  log.Printf("Server open on localhost:" + port)

  if err != nil {
    panic(err)
  }

  for {
    conn, err := ln.Accept()

    if err != nil {
      panic(err)
      continue
    }

    go handleConnection(conn)
  }
}
