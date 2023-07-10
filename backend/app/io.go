package app

import (
	"io/ioutil"
	"log"

	"google.golang.org/protobuf/proto"
)

func WriteProtoToFile(filepath string, note *EditorNote) error {
	n := toProto(note)
	// ...

	// Write the new address book back to disk.
	out, err := proto.Marshal(n)
	if err != nil {
		log.Fatalln("Failed to encode address book:", err)
	}
	if err := ioutil.WriteFile("raw", out, 0644); err != nil {
		log.Fatalln("Failed to write address book:", err)
	}
}

func ReadProtoFromFile(filepath string) ([][]byte, error) {
	// Read the existing address book.
	in, err := ioutil.ReadFile("raw")
	if err != nil {
		log.Fatalln("Error reading file:", err)
	}
	n := &Note{}
	if err := proto.Unmarshal(in, n); err != nil {
		log.Fatalln("Failed to parse address book:", err)
	}
}
