protoc -I ../../../ -I . --go_out=. --go_opt=module=github.com/eos/protocol/openmeeting/meeting meeting.proto
protoc -I ../../../ -I . --go-grpc_out=. --go-grpc_opt=module=github.com/eos/prtocol/openmeeting/meeting meeting.proto 
