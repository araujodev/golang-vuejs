GCC=go
GCMD=run
GPATH=main.go

RA512=4096
RA256=2048

run:
	make build
	$(GCC) $(GCMD) $(GPATH)

build:
	make build_db

build_db: 
	rm pkg/db/db_structs.go
	go run pkg/db/install/main.go -json=./pkg/db/config.json
	mv db_structs.go pkg/db/

create_keys:
	ssh-keygen -t rsa -b $(RA512) -f keys/app.rsa
	openssl rsa -in keys/app.rsa -pubout -outform PEM -out keys/app.rsa.pub

install:
	make install_routes
	make install_db

install_encryption:
	go get -u golang.org/x/crypto/bcrypt
install_routes:
	go get -u github.com/gorilla/mux
install_db:
	go get -u github.com/go-xorm/xorm