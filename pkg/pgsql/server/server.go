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
	//startupMessage
	startupMessageLenghtB := make([]byte, 4)
	conn.Read(startupMessageLenghtB)
	protocolVersion := make([]byte, 4)
	conn.Read(protocolVersion)

	startupMessageLenght := binary.BigEndian.Uint32(startupMessageLenghtB)
	connStringLenght := startupMessageLenght - 8
	connString := make([]byte, connStringLenght)
	conn.Read(connString)
	/*scanner := bufio.NewScanner(conn)
	scanner.Split(bufio.ScanWords)
	firstWord := scanner.Text()
	println(firstWord)*/

	messageType := []byte(`R`)
	messageLength := make([]byte, 4)
	message := make([]byte, 4)
	binary.BigEndian.PutUint32(messageLength, uint32(8))
	binary.BigEndian.PutUint32(message, uint32(3))
	AuthenticationCleartextPassword := bytes.Join([][]byte{messageType, messageLength, message}, nil)
	conn.Write(AuthenticationCleartextPassword)

	pwMessageType := make([]byte, 1)
	conn.Read(pwMessageType)

	pwMessageLength := make([]byte, 4)
	conn.Read(pwMessageLength)

	password := make([]byte, int(binary.BigEndian.Uint32(pwMessageLength))-5)
	conn.Read(password)

	if string(password) == "test" {
		messageType := []byte(`R`)
		messageLength := make([]byte, 4)
		message := make([]byte, 4)
		binary.BigEndian.PutUint32(messageLength, uint32(8))
		binary.BigEndian.PutUint32(message, uint32(0))
		AuthenticationOK := bytes.Join([][]byte{messageType, messageLength, message}, nil)
		conn.Write(AuthenticationOK)

		messageType = []byte(`Z`)
		message = make([]byte, 4)
		idle := []byte(`I`)
		binary.BigEndian.PutUint32(message, uint32(5))
		ReadyForQuery := bytes.Join([][]byte{messageType, message, idle}, nil)
		conn.Write(ReadyForQuery)

		queryString := make([]byte, 500)
		conn.Read(queryString)
		println(queryString)
	}
	/*buf := make([]byte, 1)
	reqLen, err := conn.Read(buf)
	if err != nil {
		s.Logger.Errorf("error handling request: %s", err)
	}
	println(reqLen)
	// Send a response back to person contacting us.
	conn.Write([]byte("Message received."))
	// Close the connection when you're done with it.
	conn.Close()*/

	return nil
}
