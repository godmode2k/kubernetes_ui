/* --------------------------------------------------------------
Project:    Kubernetes APIs test
Purpose:
Author:     Ho-Jung Kim (godmode2k@hotmail.com)
Date:       Since May 10, 2023
Filename:   rpc_server.go

Last modified:  May 18, 2023
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
    //"log"
    //"time"
    //"encoding/json"
    //"runtime"
    //"regexp"

    // HTTP RPC
    //"net"
    //"net/http"
    //"net/rpc"

    // HTTP JSON-RPC
    //gorilla_mux "github.com/gorilla/mux"
    //gorilla_rpc "github.com/gorilla/rpc"
    //gorilla_json "github.com/gorilla/rpc/json"

    // $ go get -u github.com/go-sql-driver/mysql
    //"database/sql"
    //_ "github.com/go-sql-driver/mysql"
    //_ "github.com/mattn/go-sqlite3"

    "test_apis_golang/types"
    //"test_apis_golang/server"

    //"reflect"
)



//! Definition
// --------------------------------------------------------------------

// DB and Kubernetes APIs
type LocalDB struct {
    //Db *sql.DB
    //DbMemory *sql.DB

    K8s *K8s_APIS
}

// Global pointer for access to DB and K8s APIs
var g_localdb *LocalDB = nil
func SetLocalDB(t *LocalDB) {
    g_localdb = t
}
func GetLocalDB() *LocalDB {
    return g_localdb
}



//! Implementation
// --------------------------------------------------------------------

func (t *LocalDB) RPC_dummy_test(
    rpc_args *types.Req_RPC_Args_st,
    response *string,
) error {
    fmt.Println( "RPC_dummy_test()" )
    *response = fmt.Sprintf( "RPC_dummy_test(): Dummy = %d", rpc_args.Dummy )

    return nil
}

func (t *LocalDB) RPC_test(
    rpc_args *types.Req_RPC_Args_st,
    response *string,
) error {
    fmt.Println( "RPC_test()" )
    *response = fmt.Sprintf( "JSONRPC_dummy_test():\nDummy = %d\nReq_xxx = %d\nReq_yyy = %s",
        rpc_args.Dummy, rpc_args.Req_xxx, rpc_args.Req_yyy )

    return nil
}

// --------------------------------------------------------------------

func __init() {
    fmt.Println( "rpc_server: __init(): initialize..." )
}

func __release() {
    fmt.Println( "rpc_server: __release(): destroy..." )
}

func Init_func() {
    __init()
}

func Release_func() {
    __release()
}




