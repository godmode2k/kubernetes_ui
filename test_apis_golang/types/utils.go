/* --------------------------------------------------------------
Project:    Kubernetes APIs test
Purpose:
Author:     Ho-Jung Kim (godmode2k@hotmail.com)
Date:       Since May 10, 2023
Filename:   utils.go

Last modified:  May 24, 2023
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
    "fmt"
    "runtime"
    "path"
    "encoding/json"
    "bytes"
    "strconv"
    "strings"
)



//! Definition
// --------------------------------------------------------------------



//! Implementation
// --------------------------------------------------------------------

// Source: https://stackoverflow.com/questions/25927660/how-to-get-the-current-function-name
func Trace__() string {
    pc := make([]uintptr, 10)  // at least 1 entry needed
    runtime.Callers(2, pc)
    f := runtime.FuncForPC(pc[0])
    file, line := f.FileLine(pc[0])
    //fmt.Printf("%s:%d %s\n", file, line, f.Name())

    s := fmt.Sprintf("%s:%d %s()", path.Base(file), line, f.Name())
    return s
}

// Source: https://stackoverflow.com/questions/17640360/file-or-line-similar-in-golang
//func __function__() string {
func Function__() string {
    _, fileName, fileLine, ok := runtime.Caller(1)
    var s string
    if ok {
        //s = fmt.Sprintf("%s:%d", fileName, fileLine)
        s = fmt.Sprintf("%s:%d", path.Base(fileName), fileLine)
    } else {
        s = ""
    }
    return s
}

func JsonPrettyPrint_Map(in map[string]interface{}, indent ...bool) string {
    val, _ := json.Marshal( in )
    _indent := true
    if len(indent) > 0 { _indent = indent[0] }
    return JsonStrPrettyPrint( string(val), _indent )
}
func JsonPrettyPrint_Interface(in []interface{}, indent ...bool) string {
    val, _ := json.Marshal( in )
    _indent := true
    if len(indent) > 0 { _indent = indent[0] }
    return JsonStrPrettyPrint( string(val), _indent )
}
// Source: https://stackoverflow.com/questions/19038598/how-can-i-pretty-print-json-using-go
func JsonStrPrettyPrint(in string, indent ...bool) string {
    var out bytes.Buffer

    //_indent := "\t"
    indent_str := "    "
    use_indent := true

    if len(indent) > 0 {
        use_indent = indent[0]
        if indent[0] == false {
            indent_str = ""
        }
    }

    //err := json.Indent(&out, []byte(in), "", "\t")
    //err := json.Indent(&out, []byte(in), "", "    ")
    err := json.Indent(&out, []byte(in), "", indent_str)
    if err != nil {
        return in
    }
    //return out.String()

    if use_indent == true {
        return out.String()
    } else {
        return strings.Replace( out.String(), "\n", "", -1 )
    }
}

//int32ptr := func (i int32) *int32 { return &i }
//int64ptr := func (i int64) *int64 { return &i }
func Int32Ptr(i int32) *int32 { return &i }
func Int64Ptr(i int64) *int64 { return &i }
func Int32ToStr(i int32) string {
    return string( fmt.Sprintf("%d", i) )
}
func Int64ToStr(i int64) string {
    return string( fmt.Sprintf("%d", i) )
}
func IntToStr(i int64) string {
    return Int64ToStr( i )
}
func StrToInt32(s string) int32 {
    _int64, _ := strconv.ParseInt( s, 10, 64 )
    _int32 := int32( _int64 )
    return _int32
}
func StrToInt64(s string) int64 {
    _int64, _ := strconv.ParseInt( s, 10, 64 )
    return _int64
}
func StrPtr(s string) *string { return &s }

