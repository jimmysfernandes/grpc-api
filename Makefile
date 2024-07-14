generate-key:
	openssl genpkey -algorithm RSA -out server.key

generate-cert:
	openssl req -x509 -new -nodes -key server.key -sha256 -days 365 -out server.crt -config openssl.cnf

generate-proto:
	protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative proto/*.proto
