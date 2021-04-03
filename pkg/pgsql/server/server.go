package server

import (
	"bytes"
	"encoding/binary"
	"github.com/codenotary/immudb/pkg/logger"
	"net"
	"os"
)

const (
	PGSQLServerHost = "localhost"
	PGSQLServerPort = "3333"
	PGSQLNetwork    = "tcp"
)

type srv struct {
	Logger logger.Logger
	Host   string
	Port   string
}

type Server interface {
	Serve() error
}

func New(setters ...Option) *srv {

	// Default Options
	cli := &srv{
		Host:   "localhost",
		Port:   "5432",
		Logger: logger.NewSimpleLogger("sqlSrv", os.Stderr),
	}

	for _, setter := range setters {
		setter(cli)
	}

	return cli
}

func (s *srv) Serve() error {
	l, err := net.Listen("tcp", s.Host+":"+s.Port)
	if err != nil {
		return err
	}
	defer l.Close()

	for {
		conn, err := l.Accept()
		if err != nil {
			return err
		}
		go s.handleRequest(conn)
	}
}

func (s *srv) handleRequest(conn net.Conn) error {
	messageType := []byte(`Z`)
	message := []byte(`i`)

	messageLength := make([]byte, 4)
	binary.BigEndian.PutUint32(messageLength, uint32(len(message)))
	readyForQuery := bytes.Join([][]byte{messageType, messageLength, message}, nil)
	conn.Write(readyForQuery)

	buf := make([]byte, 1)
	reqLen, err := conn.Read(buf)
	if err != nil {
		s.Logger.Errorf("error handling request: %s", err)
	}
	println(reqLen)
	// Send a response back to person contacting us.
	conn.Write([]byte("Message received."))
	// Close the connection when you're done with it.
	conn.Close()

	return nil
}
