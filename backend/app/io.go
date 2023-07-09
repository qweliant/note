package app

import (
	"encoding/binary"
	"io"
	"os"

	"google.golang.org/protobuf/proto"
)

func WriteProtoToFile(filepath string, note *EditorNote) error {
	outFile, err := os.OpenFile(filepath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		return err
	}

	convertedNote := toProto(note)
	encodedData, err := proto.Marshal(convertedNote)
	if err != nil {
		return err
	}

	// writing the length of the encoded item before writing the data
	buf := make([]byte, 4)
	binary.LittleEndian.PutUint32(buf, uint32(len(encodedData)))
	if _, err := outFile.Write(buf); err != nil {
		return err
	}

	// writing the actual transaction item to the file
	if _, err := outFile.Write(encodedData); err != nil {
		return err
	}

	return nil
}

func ReadProtoFromFile(filepath string) ([][]byte, error) {
	file, err := os.Open(filepath)
	if err != nil {
		return nil, err
	}

	var offset int64
	content := make([][]byte, 0)

	for {
		// reading the length of the encoded item before reading each item
		buf := make([]byte, 4)
		if _, err := file.ReadAt(buf, offset); err != nil {
			if err == io.EOF {
				break
			}
			return nil, err
		}
		itemSize := binary.LittleEndian.Uint32(buf)
		offset += 4

		// reading the actual encoded item
		item := make([]byte, itemSize)
		if _, err := file.ReadAt(item, offset); err != nil {
			if err == io.EOF {
				break
			}
			return nil, err
		}

		content = append(content, item)
		offset += int64(itemSize)
	}

	return content, nil
}
