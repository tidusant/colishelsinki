# Application for project colis helsinki

This application for service to purpose of practice the microservices with golang and mongodb.
Also for making website for colis helsinki 

### Tool use:

#### proto to generate grpc protocol buffers
- quick start: https://grpc.io/docs/languages/go/quickstart/
- download source: https://github.com/grpc/grpc-go
- install:  ( cd ../../cmd/protoc-gen-go-grpc && go install . )
- generate:
> protoc --go_out=Mgrpc/service_config/service_config.proto=/internal/proto/grpc_service_config:.   --go-grpc_out=Mgrpc/service_config/service_config.proto=/internal/proto/grpc_service_config:.   --go_opt=paths=source_relative   --go-grpc_opt=paths=source_relative ./protoc/app.proto

#### Mongodb on kubernetes
- https://kubernetes.io/blog/2017/01/running-mongodb-on-kubernetes-with-statefulsets/
- statefulset: https://kubernetes.io/docs/concepts/workloads/controllers/statefulset/
- stand alone on local https://medium.com/@dilipkumar/standalone-mongodb-on-kubernetes-cluster-19e7b5896b27


