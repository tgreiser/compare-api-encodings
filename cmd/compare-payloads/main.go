package main

import (
	"encoding/json"
	"encoding/xml"
	"flag"
	"io/ioutil"
	"log"
	"time"

	"github.com/golang/protobuf/proto"
	"github.com/tgreiser/compare-api-encodings/lib"
	"labix.org/v2/mgo/bson"
)


// compare the different payload sizes for binary API data tranfer
func main() {
	flag.Parse()
	n := flag.Arg(0)
	if (n == "") { n = "photo.jpg" }
	b, err := ioutil.ReadFile(n)
	if err != nil {
		log.Fatalln("Unable to find: ", n, err)
	}

	log.Printf("Reading %v - %v bytes", n, len(b))

	pb := &lib.PbFile{
		Name: n,
		Data: b,
	}

	f := &lib.File{
		Name: n,
		Data: b,
	}

	marshal(func() ([]byte, error) { return proto.Marshal(pb) }, "output.pb")
	marshal(func() ([]byte, error) { return bson.Marshal(f) }, "output.bson")
	marshal(func() ([]byte, error) { return json.Marshal(f) }, "output.json")
	marshal(func() ([]byte, error) { return xml.Marshal(f) }, "output.xml")
}

type mfi func() ([]byte, error)
func marshal(mFunc mfi, outfile string) {
	t0 := time.Now()
	out, err := mFunc()
	if err != nil {
		log.Fatalln("Failed to marshal: ", err)
	}
	t1 := time.Now()

	err = ioutil.WriteFile(outfile, out, 0777)
	if err != nil {
		log.Fatalln("Failed to write output: ", err)
	}
	log.Printf("Encoding %v - %v bytes - time: %v", outfile, len(out), t1.Sub(t0).String())
}
