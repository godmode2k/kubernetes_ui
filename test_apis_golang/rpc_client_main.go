/* --------------------------------------------------------------
Project:    Kubernetes APIs test
Purpose:
Author:     Ho-Jung Kim (godmode2k@hotmail.com)
Date:       Since May 10, 2023
Filename:   rpc_client_main.go

Last modified:  May 22, 2023
License:

*
* Copyright (C) 2023 Ho-Jung Kim (godmode2k@hotmail.com)
*
* Licensed under the Apache License, Version 2.0 (the "License");
* you may not use this file except in compliance with the License.
* You may obtain a copy of the License at
*
*      http://www.apache.org/licenses/LICENSE-2.0
*
* Unless required by applicable law or agreed to in writing, software
* distributed under the License is distributed on an "AS IS" BASIS,
* WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
* See the License for the specific language governing permissions and
* limitations under the License.
*
-----------------------------------------------------------------
Note:
-----------------------------------------------------------------
Reference:
 - https://pkg.go.dev/database/sql

Dependencies:
 -


1. Build:
        $ go build rpc_client_main.go
    or
        $ go run rpc_client_main.go
-------------------------------------------------------------- */
package main



//! Header
// ---------------------------------------------------------------

import (
    "fmt"
    "log"
    //"time"
    //"encoding/json"

    // HTTP RPC
    //"net"
    "net/http"
    //"net/rpc"

    // HTTP JSON-RPC
    "bytes"
    gorilla_json "github.com/gorilla/rpc/json"

    //"test_apis_golang/types"
)



//! Definition
// --------------------------------------------------------------------

//var SERVER_ADDRESS = "127.0.0.1"
//var SERVER_PORT = ""
//var SERVER = SERVER_ADDRESS + ":" + SERVER_PORT
//var URL = "http://" + SERVER_ADDRESS + ":" + SERVER_PORT
//var HTTP_RPC_SERVER_HOST_PORT = ":1234" // Internal
var HTTP_JSONRPC_SERVER_HOST_PORT = ":8890" // External



//! Implementation
// --------------------------------------------------------------------

//func json_rpc_request(api string, args *types.Req_JSONRPC_Args_st, url string) string {
func json_rpc_request(api string, args interface{}, url string) string {
    var result_str string
    message, err := gorilla_json.EncodeClientRequest( api, args )

    if err != nil {
        log.Fatalf("%s", err)
    }
    req, err := http.NewRequest("POST", url, bytes.NewBuffer(message))
    if err != nil {
        log.Fatalf("%s", err)
    }
    req.Header.Set("Content-Type", "application/json")
    client_jsonrpc := new(http.Client)
    resp, err := client_jsonrpc.Do(req)
    if err != nil {
        log.Fatalf("http.Client.Do(): Error: URL = %s, %s", url, err)
    }
    defer resp.Body.Close()

    //fmt.Println( "resp = ", resp )
    //fmt.Println( "resp = ", resp.Body )
    err = gorilla_json.DecodeClientResponse(resp.Body, &result_str)
    if err != nil {
        log.Fatalf("DecodeClientResponse(): Error: %s", err)
    }

    return result_str
}



func main() {
    // HTTP RPC Server
    /*
    client, err := rpc.DialHTTP("tcp", HTTP_RPC_SERVER_HOST_PORT)
    if err != nil {
        log.Fatal("dialing:", err)
    }
    */

    /*
    // Synchronous call
    args := &server.Args{7,8}
    var reply int
    err = client.Call("Arith.Multiply", args, &reply)
    if err != nil {
        log.Fatal("arith error:", err)
    }
    fmt.Printf("Arith: %d*%d=%d", args.A, args.B, reply)


    // Asynchronous call
    quotient := new(Quotient)
    divCall := client.Go("Arith.Divide", args, quotient, nil)
    replyCall := <-divCall.Done // will be equal to divCall
    // check errors, print, etc.
    */



    var api string
    //var err interface{}


    // Synchronous call
    //var result []types.Fetch_transactions_st
    var result_str string

    //query := "select id, name from foo limit 1"
    //OFFSET := uint(0)

    //err = client.Call( api, uint(0), &result_str )
    //err = client.Call( "LocalDB.", OFFSET, &result )
    //if err != nil {
    //    log.Fatal( "RPC: error: ", err )
    //}
    //fmt.Println( "RPC: result = \n", result )

    //for i := 0; i < len(result); i++ {
    //    fmt.Printf( "[%d] = %s, \n", i, result[i] )
    //}
    //fmt.Println( "\n\n" )


    // Test
    /*
    api = "LocalDB.RPC_test"
    args := &types.Req_RPC_Args_st { Dummy: 10, }
    err = client.Call( api, args, &result_str )
    if err != nil {
        log.Fatal( "RPC: error: ", err )
    }
    fmt.Println( api + ": result = \n", result_str )
    */



    // ------------------------------------------




    // HTTP JSON-RPC Server
    url := "http://localhost" + HTTP_JSONRPC_SERVER_HOST_PORT + "/rpc"


    // Test
    /*
    api = "LocalDB.JSONRPC_test"
    //args = &types.Req_JSONRPC_Args_st { Dummy: 10, }
    args_jsonrpc := &types.Req_JSONRPC_Args_st { Dummy: 10, Req_xxx: 11, Req_yyy: "test" }
    //args_jsonrpc := map[string]interface{} { "dummy": 10, "req_xxx": 11, "req_yyy": "test" }
    result_str = json_rpc_request( api, args_jsonrpc, url )
    fmt.Println( api + ": result = \n", result_str )
    */


    // cURL example
    /*
    $ curl -v -X POST -H "Content-Type: application/json" -d '{"method":"LocalDB.JSONRPC_kubernetes_request_service","params":[{"request":{"req":"container_list"}}],"id":5251029047958439656}' http://127.0.0.1:8890/rpc

    $ curl -v -X POST -H "Content-Type: application/json" -d '{"method":"LocalDB.JSONRPC_kubernetes_request_service","params":[{"request":{"container_ports":["80","8080","3306"],"deployment_image":"10.0.2.5:5000/test_img_apm:1.0","deployment_name":"test100-img","req":"create_instance","service_ports":[{"name":"apache","node_port":"0","port":"80","target_port":"80"},{"name":"tomcat","node_port":"0","port":"8080","target_port":"8080"},{"name":"mariadb","node_port":"0","port":"3306","target_port":"3306"}]}}],"id":5251029047958439656}' http://127.0.0.1:8890/rpc

    $ curl -v -X POST -H "Content-Type: application/json" -d '{"method":"LocalDB.JSONRPC_kubernetes_request_service","params":[{"request":{"req":"container_delete", "deployment_name": "test100-img-1684747326", "deployment_service_name": "test100-img-1684747326-service", "deployment_ingress_name": "test100-img-1684747326-ingress"}}],"id":5251029047958439656}' http://127.0.0.1:8890/rpc
    */


    api = "LocalDB.JSONRPC_kubernetes_request_service"
    /*
    args_jsonrpc_k8s := &types.Req_JSONRPC_k8s_request_service_Args_st {
        //Jsonrpc: "",
        //Result: "",
    }
    args_jsonrpc_k8s_request_params := &types.Req_JSONRPC_k8s_request_service_params_st {
        // "create_instance", "image_list", "container_list", "container_delete"
        Req: "create_instance",

        // "test1-img"
        Deployment_name: "test1-img",

        // "docker_image_ip:5000/test_img:1.0"
        Deployment_image: "docker_image_ip:5000/test_img:1.0",

        // [8080, 3306]
        //Container_ports: [8080, 3306],

        // "test1-img-service"
        Deployment_service_name: "test1-img-service",

        //[
        //    {"name": "tomcat", "node_port": 0, "port": 8080, "target_port": 8080},
        //    {"name": "mariadb", "node_port": 0, "port": 3306, "target_port": 3306}
        //]
        //Service_ports: []interface{},
    }
    var service_ports []interface{}
    service_port :=
        types.Req_JSONRPC_k8s_service_ports_Args_st { Name: "tomcat", Node_port: 0, Port: 8080, Target_port: 8080 }
    service_ports = append( service_ports, service_port )
    service_port =
        types.Req_JSONRPC_k8s_service_ports_Args_st { Name: "mariadb", Node_port: 0, Port: 3306, Target_port: 3306 }
    service_ports = append( service_ports, service_port )

    args_jsonrpc_k8s_request_params.Service_ports = service_ports
    args_jsonrpc_k8s.Request = args_jsonrpc_k8s_request_params
    */


    args_jsonrpc_k8s := map[string]interface{} {
        //"jsonrpc": "",
        //"result": "",

        "request": map[string]interface{} {
            // "create_instance", "image_list", "container_list", "container_delete"
            "req": "container_list",
        },

        /*
        "request": map[string]interface{} {
            // "create_instance", "image_list", "container_list", "container_delete"
            "req": "create_instance",

            // "test1-img"
            "deployment_name": "test100-img",

            // "docker_image_ip:5000/test_img:1.0"
            "deployment_image": "10.0.2.5:5000/test_img_apm:1.0",

            // ["80", "8080", "3306"]
            "container_ports": []string { "80", "8080", "3306" },

            // "test1-img-service"
            // Deployment_name + "-" + "timestamp" + "-" + service"
            //"deployment_service_name": "test1-img-service",

            //[
            //    { "name": "apache", "node_port": "0", "port": "80", "target_port": "80" },
            //    { "name": "tomcat", "node_port": "0", "port": "8080", "target_port": "8080" },
            //    { "name": "mariadb", "node_port": "0", "port": "3306", "target_port": "3306" }
            //]
            "service_ports": []map[string]interface{} {
                { "name": "apache", "node_port": "0", "port": "80", "target_port": "80" },
                { "name": "tomcat", "node_port": "0", "port": "8080", "target_port": "8080" },
                { "name": "mariadb", "node_port": "0", "port": "3306", "target_port": "3306" },
            },
        },
        */

        /*
        "request": map[string]interface{} {
            // "create_instance", "image_list", "container_list", "container_delete"
            "req": "container_delete",
            "deployment_name": "test100-img-1684747326",
            "deployment_service_name": "test100-img-1684747326-service",
            "deployment_ingress_name": "test100-img-1684747326-ingress",
        },
        */
    }

    result_str = json_rpc_request( api, args_jsonrpc_k8s, url )
    fmt.Println( api + ": result = \n", result_str )
}


