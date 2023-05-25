/* --------------------------------------------------------------
Project:    Kubernetes APIs test
Purpose:
Author:     Ho-Jung Kim (godmode2k@hotmail.com)
Date:       Since May 10, 2023
Filename:   request_data.go

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
-------------------------------------------------------------- */
package types



//! Header
// ---------------------------------------------------------------

import (
    //"encoding/json"
)



//! Definition
// --------------------------------------------------------------------

type Result struct {
    Jsonrpc string `json:"jsonrpc"`
    Id int `json:"id"`
    Result string `json:"result"`
}

type Result_block struct {
    Jsonrpc string `json:"jsonrpc"`
    Id int `json:"id"`
    Result map[string]interface{} `json:"result"`
}

// --------------------------------------------------------------------

type Req_JSONRPC_DummyArgs_st struct {
    Jsonrpc string `json:"jsonrpc"`
    Dummy int `json:"dummy"`
    Result string `json:"result"`
}

type Req_JSONRPC_Args_st struct {
    Jsonrpc string `json:"jsonrpc"`
    Dummy int `json:"dummy"`
    Req_xxx int `json:"req_xxx"`
    Req_yyy string `json:"req_yyy"`
    Result string `json:"result"`
}

type Req_RPC_DummyArgs_st struct {
    Dummy int
}

type Req_RPC_Args_st struct {
    Dummy int
    Req_xxx int
    Req_yyy string
}

/*
request = 
{
    "request": {
        "req": "create_instance",
        "deployment_name": "test1-img",
        "deployment_image": "10.0.2.5:5000/test_img_apm:1.0",
        "deployment_image_tomcat": "10.0.2.5:5000/test_img_tomcat:1.0",
        "deployment_image_mariadb": "10.0.2.5:5000/test_img_mariadb:1.0",
        "container_ports": [],
        "deployment_service_name": "test1-img-service",
        "service_ports": [
            {"name": "tomcat", "port": 8080, "target_port": 8080},
            {"name": "mariadb", "port": 3306, "target_port": 3306}
        ]
    }
}
*/
type Req_JSONRPC_k8s_request_service_Args_st struct {
    Jsonrpc string `json:"jsonrpc"`
    Result string `json:"result"`

    Request Req_JSONRPC_k8s_request_service_params_st `json:"request"`
}
type Req_JSONRPC_k8s_request_service_params_st struct {
    // "create_instance", "image_list", "container_list", "container_delete"
    Req string `json:"req"`

    // "test1-img"
    Deployment_name string `json:"deployment_name"`

    // "docker_image_ip:5000/test_img:1.0"
    Deployment_image string `json:"deployment_image"`

    // ["80", "8080", "3306"]
    //Container_ports []string `json:"container_ports"`
    Container_ports []string `json:"container_ports"`

    // "test1-img-<timestamp>-service"
    // Deployment_name + "-" + "timestamp" + "-" + service"
    Deployment_service_name string `json:"deployment_service_name"`
    // "test1-img-<timestamp>-ingress"
    // Deployment_name + "-" + "timestamp" + "-" + ingress"
    Deployment_ingress_name string `json:"deployment_ingress_name"`

    //[
    //    { "name": "apache", "node_port": "0", "port": "80", "target_port": "80" },
    //    { "name": "tomcat", "node_port": "0", "port": "8080", "target_port": "8080" },
    //    { "name": "mariadb", "node_port": "0", "port": "3306", "target_port": "3306" }
    //]
    Service_ports []map[string]interface{} `json:"service_ports"`
    //Service_ports []interface{} `json:"service_ports"`
    //Service_ports Req_JSONRPC_k8s_service_ports_Args_st `json:"service_ports"`
}

type Req_JSONRPC_k8s_service_ports_Args_st struct {
    Name string `json:"name"`
    Node_port int `json:"node_port"`
    Port int `json:"port"`
    Target_port int `json:"target_port"`
}

// { "use_command": True,
//   "command": ["/bin/sh", "-c", ""],
//   "args": [],
//   "host_aliases": {"ip": "", "hostnames":[]
// }
type Req_JSONRPC_k8s_init_cmds_Args_st struct {
    Use_command bool
    Command []string
    Args []string
    Host_aliases interface{}
}

const (
    ReqType_create = "create_instance"
    ReqType_delete = "container_delete"
    ReqType_list = "container_list"
)



//! Implementation
// --------------------------------------------------------------------

/*
func Fetch_transactions_json(data *Fetch_transactions_st) string {
    message, err := json.Marshal( data )
    if err != nil {
        panic( err.Error() )
    }

    return string(message)
}

func Fetch_transactions_array_json(data *[]Fetch_transactions_st) string {
    message, err := json.Marshal( data )
    if err != nil {
        panic( err.Error() )
    }

    return string(message)
}
*/


