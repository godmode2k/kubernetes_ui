

// Golang

$ wget https://go.dev/dl/go1.20.4.linux-amd64.tar.gz
$ sudo tar xzvf go1.20.4.linux-amd64.tar.gz -C /usr/local/
$ echo "export PATH=$PATH:/usr/local/go/bin" >> $HOME/.profile

$ mkdir test_apis_golang
$ cd test_apis_golang
$ go mod init test_apis_golang



// Dependencies

$ go get -u github.com/go-sql-driver/mysql
$ go get github.com/mattn/go-sqlite3

$ go get github.com/gorilla/mux
$ go get github.com/gorilla/rpc
$ go get github.com/gorilla/rpc/json
$ go get github.com/gorilla/handlers
$ go get github.com/rs/cors

$ go get k8s.io/client-go@latest

// dependency management (add, remove automatically)
$ go mod tidy -e


