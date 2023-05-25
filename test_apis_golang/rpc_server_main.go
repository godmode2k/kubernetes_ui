/* --------------------------------------------------------------
Project:    Kubernetes APIs test
Purpose:
Author:     Ho-Jung Kim (godmode2k@hotmail.com)
Date:       Since May 10, 2023
Filename:   rpc_server_main.go

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
$ go get -u github.com/go-sql-driver/mysql


1. Build:
        $ go build rpc_server_main.go
    or
        $ go run rpc_server_main.go
-------------------------------------------------------------- */
package main



//! Header
// ---------------------------------------------------------------

import (
    "fmt"
    "log"
    //"time"

    // sync.WaitGroup
    // sync.Mutex
    "sync"
    // context.Context: context.WithCancel(context.Background())
    //"context"
    //"os"

    // test
    //"math/rand"

    // HTTP RPC
    "net"
    "net/http"
    "net/rpc"

    // HTTP JSON-RPC
    gorilla_mux "github.com/gorilla/mux"
    gorilla_rpc "github.com/gorilla/rpc"
    gorilla_json "github.com/gorilla/rpc/json"
    //gorilla_handlers "github.com/gorilla/handlers"
    rs_cors "github.com/rs/cors"

    // $ go get -u github.com/go-sql-driver/mysql
    //"database/sql"
    //_ "github.com/go-sql-driver/mysql"
    //_ "github.com/mattn/go-sqlite3"

    //"test_apis_golang/types"
    "test_apis_golang/server"

    //"reflect"
)



//! Definition
// --------------------------------------------------------------------

//var SERVER_ADDRESS = "127.0.0.1"
//var SERVER_PORT = ""
//var SERVER = SERVER_ADDRESS + ":" + SERVER_PORT
//var URL = "http://" + SERVER_ADDRESS + ":" + SERVER_PORT
//var DB_SERVER_ADDRESS = "127.0.0.1:3306"
//var DB_NAME = "<db_name>"
//var DB_LOGIN_USERNAME = "root"
//var DB_LOGIN_PASSWORD = "<password>"
//var gDB *sql.DB
//var DB_MEMORY_NAME = "memory__<db_mame>"
//var gDBMemory *sql.DB


// for HTTP RPC Server
var HTTP_RPC_SERVER_HOST_PORT = ":1234" // Internal
//var HTTP_JSONRPC_SERVER_HOST_PORT = ":1235" // External
var HTTP_JSONRPC_SERVER_HOST_PORT = ":8890" // External
var g_localdb = new( rpc_server.LocalDB )
var UPDATES_INTERVAL = int(30) // 30 seconds
var gWG sync.WaitGroup
var gChan = make( chan uint8, 1 )
var _E_CHAN__DONE = uint8(0)
var _E_CHAN__CANCEL = uint8(1)


//var gMutex sync.Mutex
// mutex.Lock()
// defer mutex.Unlock()




//! Implementation
// --------------------------------------------------------------------
func __init() bool {
    fmt.Println( "main: __init(): initialize..." )

/*
    // DB (MySQL or MariaDB, ...)
    db, err := sql.Open( "mysql", DB_LOGIN_USERNAME + ":" + DB_LOGIN_PASSWORD + "@tcp(" + DB_SERVER_ADDRESS + ")/" + DB_NAME )

    if err != nil {
        panic( err.Error() )
    }

    gDB = db



    // In-memory DB (SQLite)
    //db_sqlite, err := sql.Open("sqlite3", "./localdb_sqlite.db")
    //db_sqlite, err := sql.Open("sqlite3", "file::memory:?mode=memory&cache=shared")
    //db_sqlite_filename := randomString(16) // func creates random string
    db_sqlite_filename := "localdb_sqlite3.db"
    db_sqlite, err := sql.Open( "sqlite3", fmt.Sprintf("file:%s?mode=memory&cache=shared", db_sqlite_filename) )
    if err != nil {
        //log.Fatal(err)
        panic( err.Error() )
    }
    gDBMemory = db_sqlite


    g_localdb.Db = gDB
    g_localdb.DbMemory = gDBMemory

*/


    // Initialize Kubernetes APIs
    g_localdb.K8s = new( rpc_server.K8s_APIS )
    g_localdb.K8s.Init_func()

    // Global pointer for access to DB and K8s APIs
    rpc_server.SetLocalDB( g_localdb )

    return true
}


// --------------------------------------------------------------------



//func run_worker_cache(ctx context.Context, ch chan int) {
//    for {
//        select {
//        case <-ctx.Done():
//            fmt.Println( "run_worker_cache()", "context: Done" )
//            close( ch )
//            gWG.Done()
//            break
//        case <-ch:
//            fmt.Println( "run_worker_cache()", "chan: ", <-ch )
//            // ch: 0, 1, 2, ...
//        }
//    }
//
//    fmt.Println( "run_worker_cache()", "finished..." )
//}


// Goroutine
/*
func run_worker_cache() {
    if g_localdb == nil {
        panic( "DB object == NULL" )
    }

    fmt.Println( "run_worker_cache()", "Starting caching..." )


    // SEE: var g_localdb = new( rpc_server.LocalDB )
    //g_localdb.Db_memory_update_txns_all_mixed()

    for {
        fmt.Println( "run_worker_cache()", "Updating txid..." )
        g_localdb.Db_memory_update_txns_all_mixed()

        fmt.Println( "run_worker_cache()", "Updating blocks..." )
        g_localdb.Db_memory_update_blocks_info()

        fmt.Println( "run_worker_cache()", "Updating balances..." )
        g_localdb.Db_memory_update_balances()

        fmt.Println()

        //time.Sleep( time.Second * time.Duration(UPDATES_INTERVAL) )
        time.Sleep( time.Millisecond * 1000 * time.Duration(UPDATES_INTERVAL) )
    } // for ()



//    var is_done = false
//    for {
//        select {
//        case <-gChan:
//            switch <-gChan {
//            case _E_CHAN__CANCEL:
//                fmt.Println( "CHAN: CANCEL:" )
//                is_done = true
//                break
//            case _E_CHAN__DONE:
//                fmt.Println( "CHAN: DONE:" )
//                is_done = true
//                break
//            }
//
//            if  is_done == true {
//                break
//            }
//        default:
//            fmt.Println( "CHAN: Waiting..." )
//        }
//
//        if  is_done == true {
//            break
//        }
//
//    }

    fmt.Println( "run_worker_cache()", "finished..." )

    gWG.Done()
}
*/

// Goroutine
func run_http_rpc_server() {
    if g_localdb == nil {
        panic( "Local Service Object == NULL" )
    }

    fmt.Println( "Starting HTTP RPC Server..." )


    // SEE: var g_localdb = new( rpc_server.LocalDB )

    rpc.Register( g_localdb )
    rpc.HandleHTTP()

    l, e := net.Listen( "tcp", HTTP_RPC_SERVER_HOST_PORT )
    if e != nil {
        log.Fatal( "listen error:", e )
    }

    //go http.Serve( l, nil )
    http.Serve( l, nil )
}

// Goroutine
func run_http_jsonrpc_server() {
    if g_localdb == nil {
        panic( "Local Service Object == NULL" )
    }

    fmt.Println( "Starting HTTP JSON-RPC Server..." )


    // SEE: var g_localdb = new( rpc_server.LocalDB )

    /*
    rpc.Register( g_localdb )
    rpc.HandleHTTP()

    l, e := net.Listen( "tcp", ":1234" )
    if e != nil {
        log.Fatal( "listen error:", e )
    }

    //go http.Serve( l, nil )
    http.Serve( l, nil )
    */

    _rpc := gorilla_rpc.NewServer()
    _rpc.RegisterCodec( gorilla_json.NewCodec(), "application/json" )
    _rpc.RegisterCodec( gorilla_json.NewCodec(), "application/json;charset=utf-8" )
    _rpc.RegisterService( g_localdb, "" )
    _router := gorilla_mux.NewRouter()
    _router.Handle( "/rpc",  _rpc )

    //http.Header.Set( "Connection", "close" )

    // Without CORS
    //http.ListenAndServe( HTTP_JSONRPC_SERVER_HOST_PORT,  _router )


    {
        // CORS
        // SEE: https://pkg.go.dev/github.com/gorilla/handlers#CORS
        // Default
        //
        //http.ListenAndServe( HTTP_JSONRPC_SERVER_HOST_PORT, gorilla_handlers.CORS()(_router) )


        // Canonical Host
        // CanonicalHost is HTTP middleware that re-directs requests to the canonical domain.
        // It accepts a domain and a status code (e.g. 301 or 302) and re-directs clients to this domain.
        // Note: If the provided domain is considered invalid by url.Parse or 
        // otherwise returns an empty scheme or host, clients are not re-directed.
        //
        //_canonical_host := gorilla_handlers.CanonicalHost( "http://www.example.org", 302 )
        //http.ListenAndServe( HTTP_JSONRPC_SERVER_HOST_PORT, _canonical_host(_router) )


        /*
        // Source: https://stackoverflow.com/questions/40985920/making-golang-gorilla-cors-handler-work
        //headersOk := gorilla_handlers.AllowedHeaders([]string{"X-Requested-With"})
        //originsOk := gorilla_handlers.AllowedOrigins([]string{os.Getenv("ORIGIN_ALLOWED")})
        //
        // (works)
        headersOk := gorilla_handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type"})
        originsOk := gorilla_handlers.AllowedOrigins([]string{"*"})
        methodsOk := gorilla_handlers.AllowedMethods([]string{"GET", "HEAD", "POST", "PUT", "OPTIONS"})
        //
        http.ListenAndServe( HTTP_JSONRPC_SERVER_HOST_PORT,
                            gorilla_handlers.CORS(originsOk, headersOk, methodsOk)(_router) )
        */


        ///*
        // (works)
        // RS CORS
        // Source: https://stackoverflow.com/questions/40985920/making-golang-gorilla-cors-handler-work
        cors := rs_cors.New( rs_cors.Options {
            AllowedOrigins: []string {
                "http://127.0.0.1:8890",
                // Browser (127.0.0.1:9010) <->
                // Request 127.0.0.1:8890 (VirtualBox Network Port Forward) <-> 10.0.2.4:8890
                "http://127.0.0.1:9010",
            },
            //AllowCredentials: true,
        })
        cors_handler := cors.Handler( _router )
        http.ListenAndServe( HTTP_JSONRPC_SERVER_HOST_PORT, cors_handler )
        //*/
    }
}

// Goroutine
//func run_txns_fetcher_main() {
//    rpc_server.Txns_fetcher_main_func()
//}

func run_init() {
    rpc_server.Init_func()
}

func run_release() {
    rpc_server.Release_func()
}



// --------------------------------------------------------------------



func main() {
    //fmt.Println( "HOST: " + URL )

    // Initialize Database
    if __init() != true {
        fmt.Println( "main(): init...", "false" )
            return
    }

    // Release Database
    //if gDB != nil {
    //    defer gDB.Close()
    //}
    //if gDBMemory != nil {
    //    defer gDBMemory.Close()
    //}


    // Initialize
    run_init()

    var GOROUTINE_TOTAL = 2
    gWG.Add( GOROUTINE_TOTAL )
    //ctx, cancel := context.WithCancel( context.Background() )
    //ctx = context.WithValue( ctx, "key", "val" )

    // Start Data caching
    //go run_worker_cache()

    // Start HTTP RPC Server
    go run_http_rpc_server()

    // Start HTTP JSON-RPC Server
    go run_http_jsonrpc_server()


    gWG.Wait()



    // Release
    run_release()
}

