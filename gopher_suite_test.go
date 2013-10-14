package gopher_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

  "bytes"
  "net"
  "time"
	"testing"
)

func TestGopher(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Gopher Suite")
}

type dummyAddr string

func (a dummyAddr) Network() string { return string(a) }
func (a dummyAddr) String() string { return string(a) }

type testConnection struct {
  WriteBuf bytes.Buffer
}

func (connection *testConnection) Write(byteArray []byte) (int, error) {
  return connection.WriteBuf.Write(byteArray)
}

func (*testConnection) Read([]byte) (int, error) { return 0, nil }
func (*testConnection) Close() error { return nil }
func (*testConnection) LocalAddr() net.Addr { return dummyAddr("local") }
func (*testConnection) RemoteAddr() net.Addr { return dummyAddr("remote") }
func (*testConnection) SetDeadline(time.Time) error      { return nil }
func (*testConnection) SetReadDeadline(time.Time) error  { return nil }
func (*testConnection) SetWriteDeadline(time.Time) error { return nil }
