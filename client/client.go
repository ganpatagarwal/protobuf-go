package main

import (
	"io/ioutil"
	"log"
	"os"

	pb "github.com/ganpatagarwal/protobuf-go/status"
	"google.golang.org/protobuf/proto"
)

func listRecords(sr *pb.StatusReport) {
	for _, st := range sr.Status {
		log.Println(st)
	}
}

// Main reads the entire records from a file and prints all the
// information inside.
func main() {
	if len(os.Args) != 2 {
		log.Fatalf("Usage:  %s STATUS_RECORDS_FILE\n", os.Args[0])
	}
	fname := os.Args[1]

	// [START unmarshal_proto]
	// Read the existing records.
	in, err := ioutil.ReadFile(fname)
	if err != nil {
		log.Fatalln("Error reading file:", err)
	}

	sr := &pb.StatusReport{}
	if err := proto.Unmarshal(in, sr); err != nil {
		log.Fatalln("Failed to parse records:", err)
	}
	// [END unmarshal_proto]

	listRecords(sr)
}
