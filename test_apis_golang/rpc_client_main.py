# -----------------------------------------------------------------
# Project:    Kubernetes APIs test
# Purpose:
# Author:     Ho-Jung Kim (godmode2k@hotmail.com)
# Date:       Since May 10, 2023
# Filename:   rpc_client_main.py
#
# Last modified:  May 22, 2023
# License:
#
# *
# * Copyright (C) 2023 Ho-Jung Kim (godmode2k@hotmail.com)
# *
# * Licensed under the Apache License, Version 2.0 (the "License");
# * you may not use this file except in compliance with the License.
# * You may obtain a copy of the License at
# *
# *      http://www.apache.org/licenses/LICENSE-2.0
# *
# * Unless required by applicable law or agreed to in writing, software
# * distributed under the License is distributed on an "AS IS" BASIS,
# * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
# * See the License for the specific language governing permissions and
# * limitations under the License.
# *
# -----------------------------------------------------------------
# Note:
# -----------------------------------------------------------------
# Reference:
#  - https://pkg.go.dev/database/sql
#
# Dependencies:
#  -
#
# 1. Build:
#         $ go build rpc_client_main.go
#     or
#         $ go run rpc_client_main.go
# -----------------------------------------------------------------



# Header
# --------------------------------------------------------------------
# pip3 install urllib
import urllib.request
import json



# Definition
# --------------------------------------------------------------------
HTTP_RPC_SERVER_HOST_PORT = ":1234" # Internal
HTTP_JSONRPC_SERVER_HOST_PORT = ":8890" # External



# Implementation
# --------------------------------------------------------------------

def rpc_call(url, method, args):
    data = json.dumps({
        'id': 1,
        'method': method,
        'params': [args]
        }).encode('utf8')
    req = urllib.request.Request( url, data, {'Content-Type': 'application/json'})

    #req = urllib.request.Request(url)
    #req.add_header('Authorization', 'Bearer ' + token)
    #req.add_header("Content-Type", "application/json")

    response = urllib.request.urlopen(req)
    #response = urllib.request.urlopen(req, data)

    print( "response = ", response.headers )

    _res = response.read()
    res = _res.decode("utf8")
    return json.loads(res)


if __name__ == "__main__":
    url = "http://127.0.0.1" + HTTP_JSONRPC_SERVER_HOST_PORT + "/rpc"

    #args = { 'Dummy': 1 }
    args = { 'Dummy': 1, 'Req_xxx': 11, 'Req_yyy': "test" }
    print( "args = ", args )
    #result = rpc_call( url, "LocalDB.JSONRPC_test", args )
    print( "result =" )
    #print( result )


    # cURL example
    #
    # $ curl -v -X POST -H "Content-Type: application/json" -d '{"method":"LocalDB.JSONRPC_kubernetes_request_service","params":[{"request":{"req":"container_list"}}],"id":5251029047958439656}' http://127.0.0.1:8890/rpc
    #
    # $ curl -v -X POST -H "Content-Type: application/json" -d '{"method":"LocalDB.JSONRPC_kubernetes_request_service","params":[{"request":{"container_ports":["80","8080","3306"],"deployment_image":"10.0.2.5:5000/test_img_apm:1.0","deployment_name":"test100-img","req":"create_instance","service_ports":[{"name":"apache","node_port":"0","port":"80","target_port":"80"},{"name":"tomcat","node_port":"0","port":"8080","target_port":"8080"},{"name":"mariadb","node_port":"0","port":"3306","target_port":"3306"}]}}],"id":5251029047958439656}' http://127.0.0.1:8890/rpc
    #
    # $ curl -v -X POST -H "Content-Type: application/json" -d '{"method":"LocalDB.JSONRPC_kubernetes_request_service","params":[{"request":{"req":"container_delete", "deployment_name": "test100-img-1684747326", "deployment_service_name": "test100-img-1684747326-service", "deployment_ingress_name": "test100-img-1684747326-ingress"}}],"id":5251029047958439656}' http://127.0.0.1:8890/rpc


    args = {
        #"jsonrpc": "",
        #"result": "",

        "request": {
            # "create_instance", "image_list", "container_list", "container_delete"
            "req": "container_list",
        }
    }


    """
    args = {
        #"jsonrpc": "",
        #"result": "",

        "request": {
            # "create_instance", "image_list", "container_list", "container_delete"
            "req": "create_instance",

            # "test1-img"
            "deployment_name": "test100-img",

            # "docker_image_ip:5000/test_img:1.0"
            "deployment_image": "10.0.2.5:5000/test_img_apm:1.0",

            # ["80", "8080", "3306"]
            "container_ports": [ "80", "8080", "3306" ],

            # "test1-img-service"
            # Deployment_name + "-" + "timestamp" + "-" + service"
            #"deployment_service_name": "test1-img-service",

            #[
            #    { "name": "apache", "node_port": "0", "port": "80", "target_port": "80" },
            #    { "name": "tomcat", "node_port": "0", "port": "8080", "target_port": "8080" },
            #    { "name": "mariadb", "node_port": "0", "port": "3306", "target_port": "3306" }
            #]
            "service_ports": [
                { "name": "apache", "node_port": "0", "port": "80", "target_port": "80" },
                { "name": "tomcat", "node_port": "0", "port": "8080", "target_port": "8080" },
                { "name": "mariadb", "node_port": "0", "port": "3306", "target_port": "3306" },
            ]
        }
    }
    """


    """
    args = {
        #"jsonrpc": "",
        #"result": "",

        "request": {
            # "create_instance", "image_list", "container_list", "container_delete"
            "req": "container_delete",
            "deployment_name": "test100-img-1684747326",
            "deployment_service_name": "test100-img-1684747326-service",
            "deployment_ingress_name": "test100-img-1684747326-ingress",
        }
    }
    """

    print( "args = ", args )
    result = rpc_call( url, "LocalDB.JSONRPC_kubernetes_request_service", args )
    print( "result =" )
    print( result )


