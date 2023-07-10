package app

import (
	"io/ioutil"
	"log"

	"google.golang.org/protobuf/proto"
)

// read to proto
func WriteProtoToFile(filepath string, note *Note) error {

	// Write the new address book back to disk.
	out, err := proto.Marshal(note)
	if err != nil {
		log.Fatalln("Failed to encode address book:", err)
	}
	if err := ioutil.WriteFile("raw", out, 0644); err != nil {
		log.Fatalln("Failed to write address book:", err)
	}

	return nil
}

// read from proto
func ReadProtoFromFile(filepath string) *Note {
	// Read the existing address book.
	in, err := ioutil.ReadFile("raw")
	if err != nil {
		log.Fatalln("Error reading file:", err)
	}
	n := &Note{}
	if err := proto.Unmarshal(in, n); err != nil {
		log.Fatalln("Failed to parse address book:", err)
	}
	return n

}
