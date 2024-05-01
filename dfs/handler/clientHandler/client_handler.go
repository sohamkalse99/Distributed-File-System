package clientHandler

import (
	"encoding/binary"
	"net"

	"google.golang.org/protobuf/proto"
)

type ClientHandler struct {
	conn net.Conn
}

func NewClientHandler(conn net.Conn) *ClientHandler {
	clientHandler := &ClientHandler{
		conn: conn,
	}

	return clientHandler
}

func (clientHandler *ClientHandler) readN(buff []byte) error {
	byteRead := uint64(0)

	for byteRead < uint64(len(buff)) {
		n, error := clientHandler.conn.Read(buff[byteRead:])

		if error != nil {
			return error
		}

		byteRead += uint64(n)
	}

	return nil
}

func (clientHandler *ClientHandler) writeN(buff []byte) error {
	byteWrite := uint64(0)

	for byteWrite < uint64(len(buff)) {
		n, error := clientHandler.conn.Write(buff[byteWrite:])

		if error != nil {
			return error
		}

		byteWrite += uint64(n)
	}

	return nil
}

func (clientHandler *ClientHandler) Send(fileopns *FileOpns) error {
	serialized, err := proto.Marshal(fileopns)

	if err != nil {
		return err
	}

	prefix := make([]byte, 8)
	binary.LittleEndian.PutUint64(prefix, uint64(len(serialized)))
	clientHandler.writeN(prefix)
	clientHandler.writeN(serialized)

	return nil
}

func (clientHandler *ClientHandler) Receive() (*FileOpns, error) {
	// serialized, err := proto.Marshal(wrapper)
	prefix := make([]byte, 8)
	clientHandler.readN(prefix)

	payloadSize := binary.LittleEndian.Uint64(prefix)

	payload := make([]byte, payloadSize)

	clientHandler.readN(payload)

	fileopns := &FileOpns{}
	err := proto.Unmarshal(payload, fileopns)

	return fileopns, err
}

func (clientHandler *ClientHandler) Close() {
	clientHandler.conn.Close()
}
