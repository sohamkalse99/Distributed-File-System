package snHandler

import (
	"encoding/binary"
	"net"

	"google.golang.org/protobuf/proto"
)

type StorageNodeHandler struct {
	conn net.Conn
}

func NewStorageNodeHandler(conn net.Conn) *StorageNodeHandler {
	storageNodeHandler := &StorageNodeHandler{
		conn: conn,
	}

	return storageNodeHandler
}

func (storageNodeHandler *StorageNodeHandler) readN(buff []byte) error {
	byteRead := uint64(0)

	for byteRead < uint64(len(buff)) {
		n, error := storageNodeHandler.conn.Read(buff)

		if error != nil {
			return error
		}

		byteRead += uint64(n)
	}

	return nil
}

func (storageNodeHandler *StorageNodeHandler) writeN(buff []byte) error {
	byteWrite := uint64(0)

	for byteWrite < uint64(len(buff)) {
		n, error := storageNodeHandler.conn.Write(buff)

		if error != nil {
			return error
		}

		byteWrite += uint64(n)
	}

	return nil
}

func (storageNodeHandler *StorageNodeHandler) Send(wrapper *Wrapper) error {
	serialized, err := proto.Marshal(wrapper)

	if err != nil {
		return err
	}

	prefix := make([]byte, 8)
	binary.LittleEndian.PutUint64(prefix, uint64(len(serialized)))
	storageNodeHandler.writeN(prefix)
	storageNodeHandler.writeN(serialized)

	return nil
}

func (storageNodeHandler *StorageNodeHandler) Receive() (*Wrapper, error) {
	// serialized, err := proto.Marshal(wrapper)
	prefix := make([]byte, 8)
	storageNodeHandler.readN(prefix)

	payloadSize := binary.LittleEndian.Uint64(prefix)

	payload := make([]byte, payloadSize)

	storageNodeHandler.readN(payload)

	wrapper := &Wrapper{}
	err := proto.Unmarshal(payload, wrapper)

	return wrapper, err
}

func (storageNodeHandler *StorageNodeHandler) Close() {
	storageNodeHandler.conn.Close()
}
