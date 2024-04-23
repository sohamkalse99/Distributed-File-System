package clientSNHandler

import (
	"encoding/binary"
	"net"

	"google.golang.org/protobuf/proto"
)

type ClientSNHandler struct {
	conn net.Conn
}

func NewClientSNHandler(conn net.Conn) *ClientSNHandler {
	clientSNHandler := &ClientSNHandler{
		conn: conn,
	}

	return clientSNHandler
}

func (clientSNHandler *ClientSNHandler) readN(buff []byte) error {
	byteRead := uint64(0)

	for byteRead < uint64(len(buff)) {
		n, error := clientSNHandler.conn.Read(buff)

		if error != nil {
			return error
		}

		byteRead += uint64(n)
	}

	return nil
}

func (clientSNHandler *ClientSNHandler) writeN(buff []byte) error {
	byteWrite := uint64(0)

	for byteWrite < uint64(len(buff)) {
		n, error := clientSNHandler.conn.Write(buff)

		if error != nil {
			return error
		}

		byteWrite += uint64(n)
	}

	return nil
}

func (clientSNHandler *ClientSNHandler) Send(chunkDetails *ChunkDetails) error {
	serialized, err := proto.Marshal(chunkDetails)

	if err != nil {
		return err
	}

	prefix := make([]byte, 8)
	binary.LittleEndian.PutUint64(prefix, uint64(len(serialized)))
	clientSNHandler.writeN(prefix)
	clientSNHandler.writeN(serialized)

	return nil
}

func (clientSNHandler *ClientSNHandler) Receive() (*ChunkDetails, error) {
	// serialized, err := proto.Marshal(wrapper)
	prefix := make([]byte, 8)
	clientSNHandler.readN(prefix)

	payloadSize := binary.LittleEndian.Uint64(prefix)

	payload := make([]byte, payloadSize)

	clientSNHandler.readN(payload)

	fileopns := &ChunkDetails{}
	err := proto.Unmarshal(payload, fileopns)

	return fileopns, err
}

func (clientSNHandler *ClientSNHandler) Close() {
	clientSNHandler.conn.Close()
}
