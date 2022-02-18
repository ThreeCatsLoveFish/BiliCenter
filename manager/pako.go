package manager

import (
	"bytes"
	"compress/zlib"
)

func PakoDeflate(data []byte) []byte {
	var buf bytes.Buffer
	writer, _ := zlib.NewWriterLevel(&buf, 6)
	writer.Write(data)
	writer.Close()
	return buf.Bytes()
}

func PakoInflate(data []byte) []byte {
	buf := make([]byte, 1024)
	res, _ := zlib.NewReader(bytes.NewBuffer(data))
	defer res.Close()
	size, _ := res.Read(buf)
	return buf[:size]
}
