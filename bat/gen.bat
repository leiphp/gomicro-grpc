cd datamodels/protos
protoc --micro_out=../ --go_out=../ prods.proto
protoc --micro_out=../ --go_out=../ prod_service.proto
protoc-go-inject-tag -input=../prods.pb.go
cd .. && cd ..
