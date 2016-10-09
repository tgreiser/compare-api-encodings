# compare-api-encodings
Golang examples and comparisons of various encoding formats for transmitting binary data to an API endpoint.

We will compare two binary efficient formats - [protobuf](https://developers.google.com/protocol-buffers/docs/overview) and [bson](http://bsonspec.org/), with two standard API 
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

### Conclusions

In my non-scientific testing:

- bson encoding ran the fastest
- protobuf was the smallest, beating bson by a hair
- XML is considered harmful

Binary encoded formats can hurt human readability, but they are optimized for processing and transmission. If you are dealing with binary data, human readability is probably not a large concern. Either bson or protobuf would be a great choice for high performance APIs that need to handle binary data. If you are curious, you can open the output files in a hex editor to check out how the data is encoded.
