package main

import (
	"fmt"
	"log"
	"os"

	"github.com/353solutions/building/pb"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

func main() {

	request := pb.Building{
		Name:     "test",
		Area:     10,
		Timezone: "asia/calcutta",
	}
	fmt.Println(&request)

	//Proto marshall
	data, err := proto.Marshal(&request)
	if err != nil {
		log.Fatalf("error: %s", err)
	}

	//error check
	var req2 pb.Building
	if err := proto.Unmarshal(data, &req2); err != nil {
		log.Fatalf("error: %s", err)
	}
	fmt.Println(&req2)

	//Json comparison
	fmt.Println("proto size:", len(data))

	//Mainly used for time conversion
	jdata, err := protojson.Marshal(&request)
	if err != nil {
		log.Fatalf("error: %s", err)
	}

	fmt.Println("json size:", len(jdata))
	os.Stdout.Write(jdata)

}
