package infra

import (
	"bytes"
	"compress/zlib"
	"io"
)

func PakoDeflate(data []byte) []byte {
	var buf bytes.Buffer
	writer, _ := zlib.NewWriterLevel(&buf, 6)
	writer.Write(data)
	writer.Close()
	return buf.Bytes()
}

func PakoInflate(data []byte) []byte {
	res, _ := zlib.NewReader(bytes.NewBuffer(data))
	defer res.Close()
	out, _ := io.ReadAll(res)
	return out
}
