# compare-api-encodings
Golang examples and comparisons of various API formats for transmitting binary data.

We will compare two binary efficient formats - protobuf and bson, with two standard API 
formats, json and xml.

To install:

	go get -u github.com/tgreiser/compare-api-encodings

To Run:

	go run cmd/compare-payloads/main.go <path to binary file>
	2016/10/09 14:09:27 Reading /home/me/testimg.jpg - 136356 bytes
	2016/10/09 14:09:27 Encoding output.pb - 136402 bytes - time: 171.399µs
	2016/10/09 14:09:27 Encoding output.bson - 136423 bytes - time: 71.193µs
	2016/10/09 14:09:27 Encoding output.json - 181869 bytes - time: 375.829µs
	2016/10/09 14:09:27 Encoding output.xml - 298511 bytes - time: 4.627468ms

To build protobufs (only needed if you change something). NOTE - you must have protoc and protoc-gen-go.

	compare-api-encodings$ protoc -I=./lib --go_out=./lib ./lib/pb_file.proto
