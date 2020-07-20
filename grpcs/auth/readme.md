# gRPC auth

### run in local:
cd colis/grpcs/auth
env CHADMIN_DB_HOST=127.0.0.1:27017 env CHADMIN_DB_NAME=cuahang env CHADMIN_DB_USER=cuahang env CHADMIN_DB_PASS=cuahang1234@ env PORT=32001 go run admin.go 

### run in docker:
docker build -t tidusant/colis-grpc-auth . && docker run -p 32001:8901 --env CLUSTERIP=127.0.0.1 --name colis-grpc-auth tidusant/colis-grpc-auth  