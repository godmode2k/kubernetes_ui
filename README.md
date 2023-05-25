# Kubernetes UI


Summary
----------
> RPC Server, RPC Client and Web interface for Kubernetes </br>
> </br>
> WORK IN-PROGRESS


Environment
----------
> build all and tested on GNU/Linux

    GNU/Linux: Ubuntu 20.04_x64 LTS
    Kubernetes (Master node, Worker node x2)
    Docker (+ Docker Private Registry)
    Python: v3.8.10 (pip 20.0.2)
    Go: go1.20.4 linux/amd64
    Apache2
    cURL


Prerequisites
----------
```sh
* VirtualBox: Kubernetes
 - Master node (+ Backend: JSON RPC API Server)
 - Worker node1 (+ Docker & Private Registry)
 - Worker node2


* kubernetes
 - (optional) $HOME/.kube/config permission for user


* Docker
Install: https://docs.docker.com/engine/install/


* Docker Private Registry
$ sudo docker pull registry:latest
$ sudo docker run -d --restart=always -p 5000:5000 registry:latest

// Test Image: Ubuntu 20.04, Apache2, Tomcat9, MariaDB, ...
$ sudo docker build -f Dockerfile_test_image --tag test_img_apm:1.0 .
$ sudo docker tag test_img_apm:1.0 (Docker Private Registry IP):5000/test_img_apm:1.0
$ sudo docker push (Docker Private Registry IP):5000/test_img_apm:1.0
```


Dependencies
----------
```sh
// Python3
// for Python RPC client code test
$ sudo apt-get install python3 python3-pip


// Apache2
// for Web interface test
$ sudo apt-get install apache2


// cURL
// for JSON RPC client test
$ sudo apt-get install curl


// Golang
$ wget https://go.dev/dl/go1.20.4.linux-amd64.tar.gz
$ sudo tar xzvf go1.20.4.linux-amd64.tar.gz -C /usr/local/
$ echo "export PATH=$PATH:/usr/local/go/bin" >> $HOME/.profile


// Golang modules
$ go get -u github.com/go-sql-driver/mysql
$ go get github.com/mattn/go-sqlite3
$ go get github.com/gorilla/mux
$ go get github.com/gorilla/rpc
$ go get github.com/gorilla/rpc/json
$ go get github.com/gorilla/handlers
$ go get github.com/rs/cors
$ go get k8s.io/client-go@latest

$ go mod tidy -e
```


Run & Test
----------
```sh
// Server
(Golang)
$ go run rpc_server.go


// Clients

(Golang)
$ go run rpc_client.go

(Python)
$ python3 rpc_client.py

(cURL)
// list
$ curl -v -X POST -H "Content-Type: application/json" -d '{"method":"LocalDB.JSONRPC_kubernetes_request_service","params":[{"request":{"req":"container_list"}}],"id":5251029047958439656}' http://127.0.0.1:8890/rpc

// create
$ curl -v -X POST -H "Content-Type: application/json" -d '{"method":"LocalDB.JSONRPC_kubernetes_request_service","params":[{"request":{"container_ports":["80","8080","3306"],"deployment_image":"10.0.2.5:5000/test_img_apm:1.0","deployment_name":"test100-img","req":"create_instance","service_ports":[{"name":"apache","node_port":"0","port":"80","target_port":"80"},{"name":"tomcat","node_port":"0","port":"8080","target_port":"8080"},{"name":"mariadb","node_port":"0","port":"3306","target_port":"3306"}]}}],"id":5251029047958439656}' http://127.0.0.1:8890/rpc

// delete
$ curl -v -X POST -H "Content-Type: application/json" -d '{"method":"LocalDB.JSONRPC_kubernetes_request_service","params":[{"request":{"req":"container_delete", "deployment_name": "test100-img-1684747326", "deployment_service_name": "test100-img-1684747326-service", "deployment_ingress_name": "test100-img-1684747326-ingress"}}],"id":5251029047958439656}' http://127.0.0.1:8890/rpc

(Web)
(project home)/html/k8s_apis/index.html
$ sudo ln -s (project home)/html/k8s_apis /var/www/html/

http://127.0.0.1:80/k8s_apis



// CORS error
rpc_server_main.go

...
cors := rs_cors.New( rs_cors.Options {
    AllowedOrigins: []string {
        "http://[IP]:[Port]",  // <-- ADD here...
    },
...
```


TODO
----------
```sh
Kubernetes:
 - Update Pod, Service, ...
 - Command-line interface for Pod
 - PV(PVC): iSCSI, NFS, OpenEBS, ...

 - (optional) iptables
 - Authentication: HMAC, DB (account, secret-key, ...), ...
```


Screenshots
----------
> Web interface </br>
<img src="https://github.com/godmode2k/kubernetes_ui/raw/main/screenshot_create.png" width="50%" height="50%">
<img src="https://github.com/godmode2k/kubernetes_ui/raw/main/screenshot_list.png" width="50%" height="50%">


LICENSE
----------
```sh
Apache-2.0 license
```

