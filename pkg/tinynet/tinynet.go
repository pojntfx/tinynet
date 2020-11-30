package tinynet

import (
	"encoding/binary"
	"errors"
	"strconv"
	"strings"
	"time"
	"unsafe"

	"github.com/pojntfx/webassembly-berkeley-sockets-via-webrtc/examples/go/pkg/sockets"
)

type IP []byte

type Addr interface {
	Network() string
	String() string
}

type TCPAddr struct {
	stringAddr string

	IP   IP
	Port int
	Zone string
}

func (t *TCPAddr) Network() string {
	return "tcp"
}

func (t *TCPAddr) String() string {
	return t.stringAddr
}

func ResolveTCPAddr(network, address string) (*TCPAddr, error) {
	parts := strings.Split(address, ":")

	ip := make([]byte, 4) // xxx.xxx.xxx.xxx
	for i, part := range strings.Split(parts[0], ".") {
		innerPart, err := strconv.Atoi(part)
		if err != nil {
			return nil, errors.New("could not parse IP")
		}

		ip[i] = byte(innerPart)
	}

	port, err := strconv.Atoi(parts[1])
	if err != nil {
		return nil, errors.New("could not parse port")
	}

	return &TCPAddr{
		stringAddr: address,

		IP:   ip,
		Port: port,
		Zone: "",
	}, nil
}

type Listener interface {
	Accept() (Conn, error)

	Close() error

	Addr() Addr
}

func Listen(network, address string) (Listener, error) {
	laddr, err := ResolveTCPAddr(network, address)
	if err != nil {
		return TCPListener{}, err
	}

	return ListenTCP(network, laddr)
}

func ListenTCP(network string, laddr *TCPAddr) (*TCPListener, error) {
	// Create address
	serverAddress := sockets.SockaddrIn{
		SinFamily: sockets.PF_INET,
		SinPort:   sockets.Htons(uint16(laddr.Port)),
		SinAddr: struct{ SAddr uint32 }{
			SAddr: binary.LittleEndian.Uint32(laddr.IP),
		},
	}

	// Create socket
	serverSocket, err := sockets.Socket(sockets.PF_INET, sockets.SOCK_STREAM, 0)
	if err != nil {
		return nil, err
	}

	// Bind
	if err := sockets.Bind(serverSocket, &serverAddress); err != nil {
		return nil, err
	}

	// Listen
	if err := sockets.Listen(serverSocket, 5); err != nil {
		return nil, err
	}

	return &TCPListener{
		fd:   serverSocket,
		addr: laddr,
	}, nil
}

type TCPListener struct {
	fd   int32
	addr Addr
}

func (t TCPListener) Close() error {
	return sockets.Shutdown(t.fd, sockets.SHUT_RDWR)
}

func (t TCPListener) Addr() Addr {
	return t.addr
}

func (l TCPListener) Accept() (Conn, error) {
	conn, err := l.AcceptTCP()

	return conn, err
}

func (l *TCPListener) AcceptTCP() (*TCPConn, error) {
	clientAddress := sockets.SockaddrIn{}

	// Accept
	clientSocket, err := sockets.Accept(l.fd, &clientAddress)
	if err != nil {
		return nil, err
	}

	return &TCPConn{
		fd: clientSocket,
	}, nil
}

func Dial(network, address string) (Conn, error) {
	raddr, err := ResolveTCPAddr(network, address)
	if err != nil {
		return TCPConn{}, err
	}

	conn, err := DialTCP(network, nil, raddr) // TODO: Set laddr here
	if err != nil {
		return TCPConn{}, err
	}

	return *conn, err
}

func DialTCP(network string, laddr, raddr *TCPAddr) (*TCPConn, error) {
	// Create address
	serverAddress := sockets.SockaddrIn{
		SinFamily: sockets.PF_INET,
		SinPort:   sockets.Htons(uint16(raddr.Port)),
		SinAddr: struct{ SAddr uint32 }{
			SAddr: binary.LittleEndian.Uint32(raddr.IP),
		},
	}

	// Create socket
	serverSocket, err := sockets.Socket(sockets.PF_INET, sockets.SOCK_STREAM, 0)
	if err != nil {
		return nil, err
	}

	// Connect
	if err := sockets.Connect(serverSocket, &serverAddress); err != nil {
		return nil, err
	}

	return &TCPConn{
		fd:    serverSocket,
		laddr: laddr,
		raddr: raddr,
	}, nil
}

type Conn interface {
	Read(b []byte) (n int, err error)

	Write(b []byte) (n int, err error)

	Close() error

	LocalAddr() Addr

	RemoteAddr() Addr

	SetDeadline(t time.Time) error

	SetReadDeadline(t time.Time) error

	SetWriteDeadline(t time.Time) error
}

type TCPConn struct {
	fd int32

	laddr Addr
	raddr Addr
}

func (c TCPConn) Read(b []byte) (int, error) {
	readMsg := make([]byte, unsafe.Sizeof(b))

	n, err := sockets.Recv(c.fd, &readMsg, uint32(len(b)), 0)
	if n == 0 {
		return int(n), errors.New("client disconnected")
	}

	copy(b, readMsg)

	return int(n), err
}

func (c TCPConn) Write(b []byte) (int, error) {
	n, err := sockets.Send(c.fd, b, 0)
	if n == 0 {
		return int(n), errors.New("client disconnected")
	}

	return int(n), err
}

func (c TCPConn) Close() error {
	return sockets.Shutdown(c.fd, sockets.SHUT_RDWR)
}

func (c TCPConn) LocalAddr() Addr {
	return c.laddr
}

func (c TCPConn) RemoteAddr() Addr {
	return c.laddr
}

func (c TCPConn) SetDeadline(t time.Time) error {
	// TODO: Currently there is an infinite deadline

	return nil
}

func (c TCPConn) SetReadDeadline(t time.Time) error {
	// TODO: Currently there is an infinite deadline

	return nil
}

func (c TCPConn) SetWriteDeadline(t time.Time) error {
	// TODO: Currently there is an infinite deadline

	return nil
}
