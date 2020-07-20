#admin portal

### run in local:
cd colis/portals/admin
env AUTH_CLUSTERIP=127.0.0.1 env SESSION_DB_HOST=127.0.0.1:27017 env SESSION_DB_NAME=cuahang env SESSION_DB_USER=cuahang env SESSION_DB_PASS=cuahang1234@  go run admin.go 

### run in docker:
docker build -t tidusant/colis-portal-admin . && docker run -p 8081:8080 --env AUTH_CLUSTERIP=192.168.0.105 --name colis-portal-admin tidusant/colis-portal-admin 