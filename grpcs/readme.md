#How to generate protoc for using gRPC
### install protoc tool
https://grpc.io/docs/protoc-installation/

### run generate proto
cd colis/grpcs
protoc --go_out=Mgrpc/service_config/service_config.proto=/internal/proto/grpc_service_config:.   --go-grpc_out=Mgrpc/service_config/service_config.proto=/internal/proto/grpc_service_config:.   --go_opt=paths=source_relative   --go-grpc_opt=paths=source_relative ./protoc/app.proto 