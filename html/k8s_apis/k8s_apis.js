

const SET_TYPE = 2;
let k8s_api_url = "";
let REQUEST_HTTP_METHOD = "";
let API_REQUEST_ENDPOINT = "";
let RPC_METHOD = "";

// backend: Python version
if ( SET_TYPE == 1 ) {
    k8s_api_url = "http://127.0.0.1:8888";
    REQUEST_HTTP_METHOD = "POST";
    API_REQUEST_ENDPOINT = "/k8s_request_service";
    CONTENT_TYPE = "application/x-www-form-urlencoded; charset=UTF-8";
}
// backend: Golang version
else if ( SET_TYPE == 2 ) {
    k8s_api_url = "http://127.0.0.1:8890";
    REQUEST_HTTP_METHOD = "POST";
    API_REQUEST_ENDPOINT = "/rpc";
    CONTENT_TYPE = "application/json";
    RPC_METHOD = "LocalDB.JSONRPC_kubernetes_request_service";
}



function func_menu__list() {
    console.log( "func_menu__list()" );

    var url = k8s_api_url + API_REQUEST_ENDPOINT;
    //var params = '{ "request": {"req": "container_list"} }';

    if ( SET_TYPE == 1 ) {
        var json_obj = {
            request: { req: "container_list" }
        };
    }
    else if ( SET_TYPE == 2 ) {
        var timestamp = Math.round( +new Date()/1000 );
        var json_obj = {
            method: RPC_METHOD,
            params: [{
                request: { req: "container_list" }
            }],
            id: timestamp
        };
    }

    //_ajaxSubmitCB( "#form_request", url, "POST", func_cb );
    _ajaxSubmitCB( json_obj, url, REQUEST_HTTP_METHOD, func_list_cb );
}

function func_menu__create() {
    console.log( "func_menu__create()" );

    var url = k8s_api_url + API_REQUEST_ENDPOINT;
    //var params = '{ "request": {"req": "create_instance"} }';

    let val_deployment_name = document.getElementById("deployment_name").value;
    let val_deployment_image = document.getElementById("deployment_image").value;
    //let val_deployment_service_name = document.getElementById("deployment_service_name").value;
    let val_deployment_service_name = val_deployment_name + "-service";

    let val_deployment_image_tomcat = document.getElementById("deployment_image_tomcat").value;
    let val_deployment_image_mariadb = document.getElementById("deployment_image_mariadb").value;
    //let val_deployment_service_name_tomcat = document.getElementById("deployment_service_name_tomcat").value;
    //let val_deployment_service_name_mariadb = document.getElementById("deployment_service_name_mariadb").value;

    let val_container_ports = [];
    let val_service_ports = [];


    let table = document.getElementById("table_service_ports");
    let rows_len = table.rows.length;

    console.log( "rows length = ", rows_len );

    for ( let i = 0; i < rows_len; i++ ) {
        let cells_len = table.rows[i].cells.length;
        console.log("rows = ", i );

        let val_service_ports__name = "";
        //let val_service_ports__node_port = 0;
        let val_service_ports__port = 0;
        let val_service_ports__target_port = 0;

        for ( let j = 0; j < cells_len; j++ ) {
            let cell = table.rows[i].cells[j];
            //let _html = cell.innerHTML;
            //let html = document.createElement( 'span' );
            //html.innerHTML = _html;
            //console.log( "value = ", html.getElementsByTagName("input")[0].value );

            let html = cell.getElementsByTagName("input");
            //console.log( "value = ", html );
            if ( html.length <= 0 ) continue; // <input> tag not found

            let tag = html[0];
            let name = tag.name;
            let val = tag.value;

            if ( tag.name == "service_ports__name" ) {
                val_service_ports__name = val;
            }
            //else if ( tag.name == "service_ports__node_port" ) {
            //    val_service_ports__node_port = parseInt(val);
            //    val_container_ports.push( parseInt(val) );
            //}
            else if ( tag.name == "service_ports__port" ) {
                if ( SET_TYPE == 1 ) {
                    val_service_ports__port = parseInt(val);
                }
                else if ( SET_TYPE == 2 ) {
                    val_service_ports__port = String(val);
                }
            }
            else if ( tag.name == "service_ports__target_port" ) {
                if ( SET_TYPE == 1 ) {
                    val_service_ports__target_port = parseInt(val);
                }
                else if ( SET_TYPE == 2 ) {
                    val_service_ports__target_port = String(val);
                }
            }
        }

        val_service_ports.push(
          { name: val_service_ports__name,
            //node_port: val_service_ports__node_port,
            port: val_service_ports__port,
            target_port: val_service_ports__target_port }
        );
    }



    if ( SET_TYPE == 1 ) {
        var json_obj = {
            request: { req: "create_instance",
                deployment_name: val_deployment_name,
                deployment_image: val_deployment_image,

                deployment_image_tomcat: val_deployment_image_tomcat,
                deployment_image_mariadb: val_deployment_image_mariadb,

                container_ports: val_container_ports,
                deployment_service_name: val_deployment_service_name,

                //deployment_service_name_tomcat: val_deployment_service_name_tomcat,
                //deployment_service_name_mariadb: val_deployment_service_name_mariadb,

                service_ports: val_service_ports
            }
        }
    }
    else if ( SET_TYPE == 2 ) {
        var timestamp = Math.round( +new Date()/1000 );
        var json_obj = {
            method: RPC_METHOD,
            params: [{
                request: { req: "create_instance",
                    deployment_name: val_deployment_name,
                    deployment_image: val_deployment_image,

                    deployment_image_tomcat: val_deployment_image_tomcat,
                    deployment_image_mariadb: val_deployment_image_mariadb,

                    container_ports: val_container_ports,
                    deployment_service_name: val_deployment_service_name,

                    //deployment_service_name_tomcat: val_deployment_service_name_tomcat,
                    //deployment_service_name_mariadb: val_deployment_service_name_mariadb,

                    service_ports: val_service_ports
                }
            }],
            id: timestamp
        };
    }


    //console.log( json_obj );

    //_ajaxSubmitCB( "#form_request", url, "POST", func_create_cb );
    _ajaxSubmitCB( json_obj, url, REQUEST_HTTP_METHOD, func_create_cb );
}

/*
function func_menu__delete() {
    console.log( "func_menu__delete()" );

    var url = k8s_api_url + "/k8s_request_service";
    //var params = '{ "request": {"req": "container_delete"} }';

    let val_deployment_name = document.getElementById("deployment_name").value;
    let val_deployment_image = document.getElementById("deployment_image").value;

    var json_obj = {
        request: { req: "container_delete" }
    };

    //_ajaxSubmitCB( "#form_request", url, "POST", func_cb );
    _ajaxSubmitCB( json_obj, url, "POST", func_delete_cb );
}
*/


// --------------------------------------------


function menu_list__delete_row_submit(tag_deployment, v) {
    console.log( "func_menu_list__delete_row_submit()" );

    var result = window.confirm( "DELETE: All of [" + tag_deployment + "]" + "\n\n" + "Are you sure?" );
    if ( !result ) {
        console.log( "cancel..." );
        return;
    }

    var url = k8s_api_url + API_REQUEST_ENDPOINT;
    //var params = '{ "request": {"req": "container_list"} }';



    // request delete service
    let val_service_list__deployment_name = "";
    let val_service_list__service_name = "";
    let val_service_list__ingress_name = "";
    let row_pos = v.parentNode.parentNode.rowIndex;
    let table = document.getElementById("table_service_list");
    let row = table.rows[row_pos]

    let cells_len = row.cells.length;

    for ( let i = 0; i < cells_len; i++ ) {
        let cell = row.cells[i];
        let html = cell.getElementsByTagName("input");
        //console.log( "value = ", html );
        if ( html.length <= 0 ) continue; // <input> tag not found

        let tag = html[0];
        if ( html[0].name == "service_list__pod_creation_timestamp" ) { tag = html[1]; }
        let name = tag.name;
        let val = tag.value;

        if ( tag.name == "service_list__deployment_name" ) {
            val_service_list__deployment_name = val;
        }
        else if ( tag.name == "service_list__service_name" ) {
            val_service_list__service_name = val;
        }
        else if ( tag.name == "service_list__ingress_name" ) {
            val_service_list__ingress_name = val;
        }
    }

    console.log( "deployment name = ", val_service_list__deployment_name );
    console.log( "service name = ", val_service_list__service_name );
    console.log( "ingress name = ", val_service_list__ingress_name );

    //document.getElementById("table_service_list").deleteRow( row );



    if ( SET_TYPE == 1 ) {
        var json_obj = {
            request: { req: "container_delete",
                deployment_name: val_service_list__deployment_name,
                deployment_service_name: val_service_list__service_name,
                deployment_ingress_name: val_service_list__ingress_name
            }
        };
    }
    else if ( SET_TYPE == 2 ) {
        var timestamp = Math.round( +new Date()/1000 );
        var json_obj = {
            method: RPC_METHOD,
            params: [{
                request: { req: "container_delete",
                    deployment_name: val_service_list__deployment_name,
                    deployment_service_name: val_service_list__service_name,
                    deployment_ingress_name: val_service_list__ingress_name
                }
            }],
            id: timestamp
        };
    }


    var _json_data = JSON.stringify( json_obj );
    //console.log( "request = \n" + _json_data );


    //_ajaxSubmitCB( "#form_request", url, "POST", func_cb );
    _ajaxSubmitCB( json_obj, url, REQUEST_HTTP_METHOD, func_delete_cb );
}


// --------------------------------------------

function open_window_submit(URL) {
    console.log( "func_open_window_submit()" );

    console.log( "URL = " + URL );

    window.open( URL, "_blank" ).focus();
}


// --------------------------------------------


function _ajaxSubmitCB(_json_obj, _url, _method, _func_cb) {
    var _json_data = JSON.stringify( _json_obj );
    //var request_form_data = .serialize();
    console.log( "url = \n" + _url );
    console.log( "request = \n" + _json_data );

    $.ajax({
        type: _method,
        //dataType: "json",
        //contentType: "application/json",
        contentType: CONTENT_TYPE,
        url: _url,
        cache: false,
        //data: request_form_data,
        data: _json_data,

        //async: true,
        //responseType: 'arraybuffer',

        beforeSend: function(xhr)
        {
        },
        success: function(result, status, xhr)
        {
            console.log( "result = ", result, status )
            _func_cb( result );
        },
        error: function(xhr, status, error)
        {
            console.log( "error = ", status, error );
        },
        complete: function(xhr, status) {
            console.log( "complete = ", status )
        }
    });

    /*
    var xhr = new XMLHttpRequest()
    console.log( "url = ", _url )
    xhr.open( "POST", _url, true )
    xhr.setRequestHeader("Accept", "application/json");
    xhr.setRequestHeader("Content-Type", "application/json");

    xhr.onprogress = function () {
        console.log( "PROGRESS:", xhr.responseText )
    }
    xhr.send()
    */

}

function make_input_tag(name, val) {
    //let html = '<input type="text" name="' + name + '" id="' + name + '" value="' + val + '">';
    let html = '<input type="text" value="' + val + '">';
    return html;
}

function func_list_cb(data) {
    console.log( "func_list_cb()" );
    //console.log(data);

    let _json = JSON.stringify( data );
    let _json_obj = JSON.parse( _json );
    //console.log( _json_obj );

    if ( _json_obj["result"]["ret"].toUpperCase() !== "true".toUpperCase() ) {
        console.log( "result = false" );
        return;
    }

    let _obj = _json_obj["result"]["msg"]["info"];
    console.log( _obj );

    console.log( "FIXME: size = ", _obj.length );

    let index = -1;

    let deployment = null;
    let pod = null;
    let service = null;
    let node = null;
    let ingress = null;

    let _body = document.getElementsByClassName("content_body")[0];
    let _body_list_columns = null;
    let _body_list_table = null;
    let html = "";
    let total_size = _obj.length;

    if ( _body == null ) {
        console.log( "Error: content body" );
        return;
    }

    _body_list_columns = _body.getElementsByClassName("list_columns")[0];
    if ( _body_list_columns == null ) {
        console.log( "Error: list columns" );
        return;
    }
    //! FIXME:
    //_body_list_columns.innerHTML = "total: " + String(total_size) + "<br>";
    _body_list_columns.innerHTML = "<br> <div> Index | Status | Date | Deployment | Pod | External IP" +
        " | Ingress (name, node-port, port, target-port)" +
        " | Service (name, node-port, port, target-port) </div> <br>";

    //_body_list_table = _body.getElementsByClassName("list_table")[0]; // <div>
    _body_list_table = _body.getElementsByClassName("table_service_list")[0]; // <table>
    if ( _body_list_table == null ) {
        console.log( "Error: list table" );
        return;
    }
    _body_list_table.innerHTML = "";

    for ( var i = 0; i < _obj.length; i++ ) {
        let _keys = Object.keys(_obj)[i];
        //console.log( _keys );

        for ( var key = 0; key < _keys.length; key++ ) {
            let _val = _obj[_keys[key]];

            /*
            if ( _val.hasOwnProperty("deployment") &&
                _val.hasOwnProperty("pod") &&
                _val.hasOwnProperty("node") &&
                _val.hasOwnProperty("service") ) {
            }
            else continue;
            */
            if ( !_val.hasOwnProperty("deployment") ) {
                console.log( "no such key: 'deployment'" );
                continue;
            }


            deployment = _val["deployment"];
            pod = _val["pod"];
            service = _val["service"];
            node = _val["node"];
            ingress = _val["ingress"];
            let ingress_name = ""
            let ingress_ports_html = "";
            let ingress_node_port = 0;

            console.log( "key = ", _keys[key] );
            console.log( "deployment = ", deployment );
            console.log( "pod = ", pod );
            console.log( "service = ", service );
            console.log( "node = ", node );
            console.log( "ingress = ", ingress );

            if ( _body_list_table != null ) {
                index += 1;

                //let item_deployment = JSON.stringify(deployment);
                //let item_pod = JSON.stringify(pod);
                //let item_service = JSON.stringify(service);
                //let item_node = JSON.stringify(node);

                // </div>
                /*
                html = "" +
                    make_input_tag("", deployment["name"]) + " " +
                    make_input_tag("", pod["node_name"]) + " " +
                    //pod["selector"] +
                    //node["node_name"] +
                    make_input_tag("", node["ip"]) + " ";

                    //service["selector"] +
                    //JSON.stringify(service["ports"]) +
                    //service["ports"] +


                let _keys_ports = Object.keys(service["ports"])
                //console.log( "ports...", _keys_ports );
                for ( var key_port = 0; key_port < _keys_ports.length; key_port++ ) {
                    let _val_port = service["ports"][_keys_ports[key_port]];

                    let port_list = _val_port["name"] + ": " + 
                        _val_port["node_port"] + " / " +
                        _val_port["port"] + " / " +
                        _val_port["target_port"]

                    html += make_input_tag("", port_list ) + " ";
                }

                // delete
                //_body_list_table.innerHTML += html + "<br>";
                */


                /*
                if ( deployment.hasOwnProperty("name") &&
                    pod.hasOwnProperty("node_name") &&
                    node.hasOwnProperty("ip") &&
                    service.hasOwnProperty("name") ) {
                }
                else continue;
                */


                // <table>
                let table = document.getElementById("table_service_list");
                let row = table.insertRow(0);
                //let row_details = table.insertRow(0);
                let row_child_html = "";
                row.id = deployment["name"]; // <tr id = "">
                //row_details.id = row.id + "_details";

                let cell_len = Number(0);
                let cell0 = row.insertCell(cell_len++); // Index
                let cell1 = row.insertCell(cell_len++); // Status
                let cell2 = row.insertCell(cell_len++); // CreationTimestamp
                let cell3 = row.insertCell(cell_len++); // Deployment Name
                let cell4 = row.insertCell(cell_len++); // Pod Name
                let cell5 = row.insertCell(cell_len++); // External IP
                let cell6 = row.insertCell(cell_len++); // Ingress Name
                let cell7 = row.insertCell(cell_len++); // Ingress Ports
                //let cell7; // Ingress Ports
                /*
                {
                    // cell7 (for list layout)
                    //cell_len++;

                    let target_tr = document.getElementById( row.id );
                    var new_td = document.createElement( "TD" );
                    var new_tr = document.createElement( "TR" );
                    //new_td.innerHTML = ingress_ports_html;
                    new_td.innerHTML = "<tr> </tr>";
                    new_tr.appendChild( new_td );
                    cell7 = target_tr.appendChild( new_tr );
                }
                */
                let cell8 = row.insertCell(cell_len++); // Service Name
                let cell9 = row.insertCell(cell_len++); // Service Ports

                cell0.innerHTML = '<td> <input disabled readonly type="text" size=3 name="service_list__index" value="' + String(index) + '"> </td>';
                cell1.innerHTML = '<td> <input disabled readonly type="text" size=12 name="service_list__status" value="' + pod["status"] + '"> </td>';
                //cell2.innerHTML = '<td> <input disabled readonly type="text" size=12 name="service_list__deployment_name" value="' + deployment["name"] + '"> </td>';
                //cell3.innerHTML = '<td> <input disabled readonly type="text" size=12 name="service_list__node_name" value="' + pod["node_name"] + '"> </td>';
                //cell4.innerHTML = '<td> <input disabled readonly type="text" size=12 name="service_list__node_ip" value="' + node["ip"] + '"> </td>';


                let cell2_innerHTML = "";
                //let pod_creation_timestamp = new Date( parseInt(pod["creation_timestamp"]) * 1000 ).toLocaleString()
                let pod_creation_date = new Date( parseInt(pod["creation_timestamp"]) * 1000 )
                pod_creation_timestamp = pod_creation_date.getFullYear() + "-" +
                                        ("0" + (pod_creation_date.getMonth() + 1)).slice( -2 ) + "-" +
                                        ("0" + pod_creation_date.getDate()).slice( -2 ) + " " +
                                        ("0" + pod_creation_date.getHours()).slice( -2 ) + ":" +
                                        ("0" + pod_creation_date.getMinutes()).slice( -2 ) + ":" +
                                        ("0" + pod_creation_date.getSeconds()).slice( -2 );
                cell2_innerHTML += '<table> <tr> <td> <input disabled readonly type="text" size=25 name="service_list__pod_creation_timestamp" value="' + pod_creation_timestamp + " (" + pod["creation_timestamp"] + ") " + '"> </td> </tr> </table>';
                cell2_innerHTML += '<table> <tr> <td> <input disabled readonly type="text" size=25 name="service_list__deployment_name" value="' + deployment["name"] + '"> </td> </tr> </table>';
                cell2_innerHTML += '<table> <tr> <td> <input disabled readonly type="text" size=25 name="service_list__node_name" value="' + pod["node_name"] + '"> </td> </tr> </table>';
                cell2_innerHTML += '<table> <tr> <td> <input disabled readonly type="text" size=25 name="service_list__node_ip" value="' + node["ip"] + '"> </td> </tr> </table>';
                cell2.innerHTML = cell2_innerHTML;



                //cell6.innerHTML = '<td> <input disabled readonly type="text" size=12 name="service_list__ingress_name" value="' + ingress["ingress_name"] + '"> </td>';
                //cell7.innerHTML = '<td> <input disabled readonly type="text" size=12 name="service_list__service_name" value="' + service["name"] + '"> </td>';

                if ( typeof ingress == "undefined" ) {
                    ingress_name = ""
                }
                else {
                    ingress_name = ingress["ingress_name"]

                    let _ingress_ports = ingress["ports"];
                    //console.log( "ports...", _ingress_ports );
                    let ingress_ports_name = _ingress_ports["name"];
                    let ingress_ports_node_port = _ingress_ports["node_port"];
                    ingress_node_port = ingress_ports_node_port;
                    let ingress_ports_port = _ingress_ports["port"];
                    let ingress_ports_target_port = _ingress_ports["target_port"];

                    ingress_cell0_html = '<td> <input disabled readonly type="text" size=12 name="ingress_ports_name" value="'
                        + ingress_ports_name + '"> </td>';
                    ingress_cell1_html = '<td> <input disabled readonly type="text" size=4 name="ingress_ports_node_port" value="'
                        + ingress_ports_node_port + '"> </td>';
                    ingress_cell2_html = '<td> <input disabled readonly type="text" size=4 name="ingress_ports_port" value="'
                        + ingress_ports_port + '"> </td>';
                    ingress_cell3_html = '<td> <input disabled readonly type="text" size=4 name="ingress_ports_target_port" value="'
                        + ingress_ports_target_port + '"> </td>';


                    ingress_ports_html = "<table> <tr>";
                    ingress_ports_html += ingress_cell0_html;
                    ingress_ports_html += ingress_cell1_html;
                    ingress_ports_html += ingress_cell2_html;
                    ingress_ports_html += ingress_cell3_html;
                    ingress_ports_html += "</tr> </table>";

                    //console.log( row_child_html );
                    //cell6.innerHTML = ingress_ports_html;
                }

                //cell6.innerHTML = '<td> <input disabled readonly type="text" size=12 name="service_list__ingress_name" value="' + ingress_name + '"> </td>';

                let cell6_innerHTML = "";
                cell6_innerHTML += '<table> <tr> <td> <input disabled readonly type="text" size=30 name="service_list__ingress_name" value="' + ingress_name + '"> </td> </tr> </table>';
                cell6_innerHTML += ingress_ports_html;
                cell6.innerHTML = cell6_innerHTML;


                if ( typeof service == "undefined" ) {
                    cell_len--;
                }
                else {
                    //cell8.innerHTML = '<td> <input disabled readonly type="text" size=12 name="service_list__service_name" value="' + service["name"] + '"> </td>';

                    //let fixed_rows_count = cell_len;
                    //let fixed_column_size_total = Number(0);
                    let _keys_ports = Object.keys(service["ports"])
                    //console.log( "ports...", _keys_ports );
                    for ( var key_port = 0; key_port < _keys_ports.length; key_port++ ) {
                        let _val_port = service["ports"][_keys_ports[key_port]];

                        //let port_list = _val_port["name"] + ": " + 
                        //    _val_port["node_port"] + " / " +
                        //    _val_port["port"] + " / " +
                        //    _val_port["target_port"];

                        let ports_name = _val_port["name"];
                        let ports_node_port = _val_port["node_port"];
                        let ports_port = _val_port["port"];
                        let ports_target_port = _val_port["target_port"];


                        //! FIXME
                        if ( ingress.hasOwnProperty("paths") ) {
                            ports_node_port += "," + ingress["paths"][key_port];
                        }


                        //let cell = row.insertCell(key_port + fixed_rows_count); // Service (Name, Node-Port, Port, Target-Port)
                        //cell.innerHTML = '<td> <input disabled readonly type="text" size=26 name="service_list_ports" value="' + port_list + '"> </td>';

                        let cell0 = row.insertCell( cell_len++ ); // Service: Name
                        let cell1 = row.insertCell( cell_len++ ); // Service: Node-Port
                        let cell2 = row.insertCell( cell_len++ ); // Service: Port
                        let cell3 = row.insertCell( cell_len++ ); // Service: Port, Target-Port

                        cell0_html = '<td> <input disabled readonly type="text" size=12 name="service_list_ports_name" value="'
                            + ports_name + '"> </td>';
                        cell1_html = '<td> <input disabled readonly type="text" size=14 name="service_list_ports_node_port" value="'
                            + ports_node_port + '"> </td>';
                        cell2_html = '<td> <input disabled readonly type="text" size=4 name="service_list_ports_port" value="'
                            + ports_port + '"> </td>';
                        cell3_html = '<td> <input disabled readonly type="text" size=4 name="service_list_ports_target_port" value="'
                            + ports_target_port + '"> </td>';

                        let URL = "#";
                        if ( ingress_node_port > 0 ) {
                            let HOST = window.location.host.split(":")[0];
                            URL = "http://" + HOST + ":"
                                + ingress_node_port + ports_node_port.split(",")[1];
                            url_html = '<td> <button name="service_list_open_window" onclick="open_window_submit(\''
                                + URL + '\'); return false;"> Open </button> </td>';
                        }
                        //console.log( url_html );


                        //cell0.innerHTML = cell0_html;
                        //cell1.innerHTML = cell1_html;
                        //cell2.innerHTML = cell2_html;
                        //cell3.innerHTML = cell3_html;


                        //row_child_html += "<tr>";
                        row_child_html += "<table> <tr>";
                        row_child_html += cell0_html;
                        row_child_html += cell1_html;
                        row_child_html += cell2_html;
                        row_child_html += cell3_html;
                        row_child_html += url_html;
                        row_child_html += "</tr> </table>";
                        //row_child_html += "</tr>";
                    } // for ()

                    //console.log( row_child_html );


                    //let target_tr = document.getElementById( row.id );
                    //var new_td = document.createElement( "TD" );
                    //var new_tr = document.createElement( "TR" );
                    //new_td.innerHTML = row_child_html;
                    //new_tr.appendChild( new_td );
                    //target_tr.appendChild( new_tr );

                    cell8_innerHTML = '<table> <tr> <td> <input disabled readonly type="text" size=30 name="service_list__service_name" value="' + service["name"] + '"> </td> </tr> </table>';
                    row_child_html = cell8_innerHTML + row_child_html;
                    cell8.innerHTML = row_child_html;
                }

                // DELETE button
                let cell = row.insertCell( cell_len++ );
                cell.innerHTML = '<td> <button name="service_list_delete" onclick="menu_list__delete_row_submit(\''
                    + deployment["name"] + '\', '
                    + 'this); return false;"> DELETE </button> </td>';


                // table row highlight
                row.addEventListener('click', function() {
                    var tr = $(this).closest('tr');
                    [].forEach.call(table.rows, function(e) {
                        e.classList.remove("highlight");
                    });

                    this.className += ' highlight';
                }, false);
            }
        } // for ()

        //_body_list_table.innerHTML += String(i) + " " + html + "<br>";
    } // for ()
}

function func_create_cb(data) {
    console.log( "func_create_cb()" );
    //console.log(data);

    let _json = JSON.stringify( data );
    let _json_obj = JSON.parse( _json );
    //console.log( _json_obj );

    if ( _json_obj["result"]["ret"].toUpperCase() !== "true".toUpperCase() ) {
        console.log( "result = false" );
        return;
    }
}

function func_delete_cb(data) {
    console.log( "func_delete_cb()" );
    //console.log(data);

    let _json = JSON.stringify( data );
    let _json_obj = JSON.parse( _json );
    //console.log( _json_obj );

    if ( _json_obj["result"]["ret"].toUpperCase() !== "true".toUpperCase() ) {
        console.log( "result = false" );
        return;
    }
}


//$(document).ready( function() {
//    func_menu__list();
//});


