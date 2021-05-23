package main

import (
	"io/ioutil"
	"log"
	"os"

	pb "github.com/ganpatagarwal/protobuf-go/status"
	"google.golang.org/protobuf/proto"
)

func createStatus(fqdn, ip string, state bool) *pb.Status {
	return &pb.Status{
		Fqdn:   fqdn,
		Ip:     ip,
		Status: state,
	}
}

func main() {
	if len(os.Args) != 2 {
		log.Fatalf("Usage:  %s STATUS_RECORD_FILE\n", os.Args[0])
	}
	fname := os.Args[1]

	// Read the existing records.
	in, err := ioutil.ReadFile(fname)
	if err != nil {
		if os.IsNotExist(err) {
			log.Printf("%s: File not found.  Creating new file.\n", fname)
		} else {
			log.Fatalln("Error reading file:", err)
		}
	}

	// [START marshal_proto]
	sr := &pb.StatusReport{}
	// [START_EXCLUDE]
	if err := proto.Unmarshal(in, sr); err != nil {
		log.Fatalln("Failed to parse records:", err)
	}

	st := createStatus("localhost", "127.0.0.1", true)
	sr.Status = append(sr.Status, st)

	// Write the new records back to disk.
	out, err := proto.Marshal(sr)
	if err != nil {
		log.Fatalln("Failed to encode records:", err)
	}
	if err := ioutil.WriteFile(fname, out, 0644); err != nil {
		log.Fatalln("Failed to write records:", err)
	}
	// [END marshal_proto]
}
