/* --------------------------------------------------------------
Project:    Kubernetes APIs test
Purpose:
Author:     Ho-Jung Kim (godmode2k@hotmail.com)
Date:       Since May 10, 2023
Filename:   rpc_server_json_apis.go

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
 - https://golang.org/pkg/net/rpc/
 - https://pkg.go.dev/database/sql
 - https://pkg.go.dev/github.com/mattn/go-sqlite3

Dependencies:
 - $ go get -u github.com/go-sql-driver/mysql
 - $ go get github.com/mattn/go-sqlite3
-------------------------------------------------------------- */
package rpc_server



//! Header
// ---------------------------------------------------------------

import (
    "fmt"
    "log"
    //"time"
    "encoding/json"
    //"runtime"
    //"regexp"

    // HTTP JSON-RPC
    //"net"
    "net/http"
    //"net/rpc"

    // $ go get -u github.com/go-sql-driver/mysql
    //"database/sql"
    //_ "github.com/go-sql-driver/mysql"
    //_ "github.com/mattn/go-sqlite3"

    "test_apis_golang/types"
)



//! Definition
// --------------------------------------------------------------------




//! Implementation
// --------------------------------------------------------------------

func (t *LocalDB) JSONRPC_dummy_test(
    request *http.Request,
    rpc_args *types.Req_JSONRPC_DummyArgs_st,
    response *string,
) error {
    fmt.Println( "JSONRPC_dummy_test()" )
    *response = fmt.Sprintf( "JSONRPC_dummy_test(): Dummy = %d", rpc_args.Dummy )

    return nil
}

func (t *LocalDB) JSONRPC_test(
    request *http.Request,
    rpc_args *types.Req_JSONRPC_Args_st,
    response *string,
) error {
    fmt.Println( "JSONRPC_test()" )

    var _result []types.Req_JSONRPC_Args_st
    //var result_str string
    //*response = result_str

    //_result[0] = types.Result_Args_st { Dummy: 10, Req_xxx: 11, Req_yyy: "test" }
    res_rpc_args := types.Req_JSONRPC_Args_st {
        Dummy: rpc_args.Dummy, Req_xxx: rpc_args.Req_xxx, Req_yyy: rpc_args.Req_yyy }

    _result = append( _result, res_rpc_args )

    result, err_marshal := json.Marshal( _result )
    if err_marshal != nil {
        panic( err_marshal.Error() )
    }

    //var json_map_arr = make( map[string]interface{} )
    //json_map_arr["res"] = _result
    //
    //result, err_marshal := json.Marshal( json_map_arr )
    //if err_marshal != nil {
    //    panic( err_marshal.Error() )
    //}


    *response = string(result)

    return nil
}

func (t *LocalDB) JSONRPC_kubernetes_request_service(
    request *http.Request,
    rpc_args *types.Req_JSONRPC_k8s_request_service_Args_st,
    //response *string,
    response *map[string]interface{},
) error {
    fmt.Println( "JSONRPC_kubernetes_request_service()" )
    fmt.Println( "request = ", rpc_args )


    /*
    _result := types.Result_block {}
    {
        var tmp_results []types.Req_JSONRPC_k8s_request_service_Args_st
        tmp_results = append( tmp_results, *rpc_args )
        tmp_result, err_marshal := json.Marshal( tmp_results )
        if err_marshal != nil {
            panic( err_marshal.Error() )
        }
        fmt.Println( "JSONRPC_kubernetes_request_service(): request = \n" + string(tmp_result) )

        _result.Result = tmp_results
    }

    result, err_marshal := json.Marshal( _result )
    if err_marshal != nil {
        panic( err_marshal.Error() )
    }

    *response = string(result)
    */


    //var _result types.Result
    var _result map[string]interface{}

    var rpc_args_map map[string]interface{}
    rpc_args_json, err_marshal := json.Marshal( rpc_args )
    if err_marshal != nil { log.Fatal( "json.Marshal: ", err_marshal ) }
    err_marshal = json.Unmarshal( rpc_args_json, &rpc_args_map )
    if err_marshal != nil { log.Fatal( "json.Unmarshal: ", err_marshal ) }

    k8s_apis__request_container_service( rpc_args_map, &_result )

/*
    var result []byte
    result, err_marshal = json.Marshal( _result )
    if err_marshal != nil {
        //panic( err_marshal.Error() )
        log.Fatal( "json.Marshal: ", err_marshal )
    }

    // Note: prefixed { "result": {...}, "error": {...}, ... }
    *response = string(result)
*/

    //*response = types.JsonPrettyPrint_Map( _result, false )

    *response = _result

    return nil
}

