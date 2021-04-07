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

		////##-> dataRowDescription
		//Byte1('T')
		messageType = []byte(`T`)
		// Length of message contents in bytes, including self.
		// Int32
		messageLength = make([]byte, 4)
		binary.BigEndian.PutUint32(messageLength, uint32(1+4+4+4+2+4+2+4+2+4+2))
		// Specifies the number of fields in a row (can be zero).
		// Int16
		fieldNumb := make([]byte, 4)
		binary.BigEndian.PutUint16(fieldNumb, uint16(1))
		// The field name.
		// String
		fieldName := []byte(`first`)
		// If the field can be identified as a column of a specific table, the object ID of the table; otherwise zero.
		// Int32
		id := make([]byte, 4)
		binary.BigEndian.PutUint32(id, uint32(1))
		// If the field can be identified as a column of a specific table, the attribute number of the column; otherwise zero.
		// Int16
		attributeNumber := make([]byte, 2)
		binary.BigEndian.PutUint16(attributeNumber, uint16(1))
		// The object ID of the field's data type.
		// Int32
		objectId := make([]byte, 4)
		binary.BigEndian.PutUint32(objectId, uint32(PgTypeMap["oid"]))
		// The data type size (see pg_type.typlen). Note that negative values denote variable-width types.
		// For a fixed-size type, typlen is the number of bytes in the internal representation of the type. But for a variable-length type, typlen is negative. -1 indicates a “varlena” type (one that has a length word), -2 indicates a null-terminated C string.
		// Int16
		dataTypeSize := make([]byte, 2)
		binary.BigEndian.PutUint16(dataTypeSize, uint16(4))
		// The type modifier (see pg_attribute.atttypmod). The meaning of the modifier is type-specific.
		// atttypmod records type-specific data supplied at table creation time (for example, the maximum length of a varchar column). It is passed to type-specific input functions and length coercion functions. The value will generally be -1 for types that do not need atttypmod.
		// Int32
		typeModifier := make([]byte, 4)
		binary.BigEndian.PutUint32(typeModifier, uint32(1))
		// The format code being used for the field. Currently will be zero (text) or one (binary). In a RowDescription returned from the statement variant of Describe, the format code is not yet known and will always be zero.
		// Int16
		formatCode := make([]byte, 2)
		binary.BigEndian.PutUint16(formatCode, uint16(0))

		rowDesc := bytes.Join([][]byte{messageType, messageLength, fieldNumb, fieldName, id, attributeNumber, objectId, dataTypeSize, typeModifier, formatCode}, nil)
		conn.Write(rowDesc)

		////##-> dataRow
		// Identifies the message as a data row.
		// Byte1('D')
		messageType = []byte(`D`)
		// Length of message contents in bytes, including self.
		// Int32
		selfMessageLength := make([]byte, 4)
		binary.BigEndian.PutUint32(messageLength, uint32(16))
		// The number of column values that follow (possibly zero).
		// Int16
		columnNumb := make([]byte, 4)
		binary.BigEndian.PutUint16(columnNumb, uint16(1))

		// Next, the following pair of fields appear for each column:
		// The length of the column value, in bytes (this count does not include itself). Can be zero. As a special case, -1 indicates a NULL column value. No value bytes follow in the NULL case.
		// Int32
		valueLength := make([]byte, 4)
		binary.BigEndian.PutUint32(valueLength, uint32(4))
		// The value of the column, in the format indicated by the associated format code. n is the above length.
		// Byten
		value := make([]byte, 4)
		binary.BigEndian.PutUint32(messageLength, uint32(1))

		dataRow := bytes.Join([][]byte{messageType, selfMessageLength, columnNumb, valueLength, value}, nil)
		conn.Write(dataRow)

		//ready for query
		messageType = []byte(`Z`)
		message = make([]byte, 4)
		idle = []byte(`I`)
		binary.BigEndian.PutUint32(message, uint32(5))
		ReadyForQuery = bytes.Join([][]byte{messageType, message, idle}, nil)
		conn.Write(ReadyForQuery)
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
