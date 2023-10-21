package net

import "encoding/binary"

func NewPacketReader(buffer []byte) PacketReader {
	return PacketReader{
		buffer: buffer,
		offset: 0,
	}
}

type PacketReader struct {
	buffer []byte
	offset int
}

func (r *PacketReader) ReadByte() (byte, error) {
	ret := r.buffer[r.offset]
	r.offset++
	return ret, nil
}

func (r *PacketReader) ReadUInt32() (uint32, error) {
	ret := binary.LittleEndian.Uint32(r.buffer[r.offset : r.offset+4])
	r.offset += 4
	return ret, nil
}

func (r *PacketReader) ReadString() (string, error) {
	stringlen, err := r.ReadUInt32()
	if err != nil {
		return "", err
	}

	ret := string(r.buffer[r.offset : r.offset+int(stringlen)])
	r.offset += int(stringlen)
	return ret, nil
}
