/* --------------------------------------------------------------
Project:    Kubernetes APIs test
Purpose:
Author:     Ho-Jung Kim (godmode2k@hotmail.com)
Date:       Since May 10, 2023
Filename:   k8s_apis.go

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
Reference:
 - https://golang.org/pkg/net/rpc/
 - https://pkg.go.dev/database/sql
 - https://pkg.go.dev/github.com/mattn/go-sqlite3
 - https://pkg.go.dev/k8s.io/api/

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
    "time"
    // now := time.Now()
    // secs := now.Unix(), nanos := now.UnixNano(), millis := nano / 1000000
    // fmt.Println( time.Unix(secs, 0), time.Unix(0, nanos) )
    "encoding/json"
    //"runtime"
    //"regexp"
    //"reflect"
    //"strconv"
    "strings"
    "sort"

    //"bufio"
    "context"
    "flag"
    //"os"
    "path/filepath"


    // HTTP JSON-RPC
    //"net"
    //"net/http"
    //"net/rpc"


    // $ go get -u github.com/go-sql-driver/mysql
    //"database/sql"
    //_ "github.com/go-sql-driver/mysql"
    //_ "github.com/mattn/go-sqlite3"


    // Kubernetes [
    networkingv1 "k8s.io/api/networking/v1"
    appsv1 "k8s.io/api/apps/v1"
    apiv1 "k8s.io/api/core/v1"
    metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
    "k8s.io/apimachinery/pkg/util/intstr"
    "k8s.io/client-go/kubernetes"
    "k8s.io/client-go/tools/clientcmd"
    "k8s.io/client-go/util/homedir"
    //"k8s.io/client-go/util/retry"
    // Kubernetes ]

    "test_apis_golang/types"
)



//! Definition
// --------------------------------------------------------------------
type K8s_APIS struct {
    clientset *kubernetes.Clientset
}

func (t *K8s_APIS) Init_func() error {
    fmt.Println( types.Trace__() )

    var kubeconfig *string
    if home := homedir.HomeDir(); home != "" {
        kubeconfig = flag.String("kubeconfig", filepath.Join(home, ".kube", "config"), "(optional) absolute path to the kubeconfig file")
    } else {
        kubeconfig = flag.String("kubeconfig", "", "absolute path to the kubeconfig file")
    }
    flag.Parse()

    config, err := clientcmd.BuildConfigFromFlags("", *kubeconfig)
    if err != nil {
        panic(err)
    }
    clientset, err := kubernetes.NewForConfig(config)
    if err != nil {
        panic(err)
    }


    t.clientset = clientset


    t.__test()

    return nil
}

func (t *K8s_APIS) __test() {
    {
        //NAMESPACE_NULL := ""

        //t.node_list()
        //t.deployment_list( NAMESPACE_NULL )
        //t.pod_list( NAMESPACE_NULL )
        //t.service_list( NAMESPACE_NULL )
        //t.ingress_list( NAMESPACE_NULL )
    }

    {
        //NAMESPACE_DEFAULT := "default"

        //timestamp := time.Now().Unix()
        //fmt.Println( "timestamp = ", timestamp )

        /*
        // Creates Deployment
        deployment_name := "test100" + "-" + types.IntToStr( timestamp )
        deployment_image := "10.0.2.5:5000/test_img_apm:1.0"
        container_ports := []string { "80", "8080", "3306" }
        var opts interface{}
        t.create_deployment(
            NAMESPACE_DEFAULT,
            deployment_name,
            deployment_image,
            container_ports,
            opts,
        )
        */


        /*
        // Creates Service
        deployment_service_name := deployment_name + "-service"
        deployment_service_ports := []map[string]string {
            {
                "name": "apache",
                "node_port": "",
                "port": "80",
                "target_port": "80",
            },
            {
                "name": "tomcat",
                "node_port": "",
                "port": "8080",
                "target_port": "8080",
            },
            {
                "name": "mariadb",
                "node_port": "",
                "port": "3306",
                "target_port": "3306",
            },
        }
        t.create_service(
            NAMESPACE_DEFAULT,
            deployment_name,
            deployment_service_name,
            deployment_service_ports,
        )
        */


        /*
        // Ingress
        deployment_ingress_name := deployment_name + "-ingress"
        deployment_ingress_host := "test.com"
        t.create_ingress(
            NAMESPACE_DEFAULT,
            deployment_ingress_name,
            deployment_ingress_host,
            deployment_service_name,
            deployment_service_ports,
        )
        */


        /*
        // Deletes all
        deployment_name := "test100-1684308153"
        deployment_service_name := deployment_name + "-service"
        deployment_ingress_name := deployment_name + "-ingress"
        t.delete_container(
            NAMESPACE_DEFAULT,
            deployment_name,
            deployment_service_name,
            deployment_ingress_name,
        )
        */


        // Container (Pod) details List
        //username := ""
        //t.container_details_list( username )
    }
}

func (t *K8s_APIS) node_list() []interface{} {
    fmt.Println( types.Trace__() )

    nodes, err := t.clientset.CoreV1().Nodes().List( context.TODO(), metav1.ListOptions{} )
    if err != nil {
        panic( err.Error() )
    }
    //fmt.Println( nodes )
    //fmt.Println( reflect.TypeOf(nodes) )

    var res []interface{}
    for i := 0; i < len(nodes.Items); i++ {
        namespace := nodes.Items[i].ObjectMeta.Namespace
        if namespace == "kube-system" {
            continue
        }

        name := nodes.Items[i].ObjectMeta.Name
        //namespace := nodes.Items[i].ObjectMeta.Namespace
        //labels := nodes.Items[i].ObjectMeta.Labels
        //uid := nodes.Items[i].ObjectMeta.UID

        var ipaddr string
        //var ipaddr_internal string

        // Node IPs
        {
            // projectcalico
            //ipaddr = Split( nodes.Items[i].ObjectMeta.Annotations["projectcalico.org/IPv4Address"], "/" )[0]
            //ipaddr_internal = nodes.Items[i].ObjectMeta.Annotations["projectcalico.org/IPv4IPIPTunnelAddr"]

            // flannel
            ipaddr = nodes.Items[i].ObjectMeta.Annotations["flannel.alpha.coreos.com/public-ip"]
            //ipaddr_internal = ipaddr
        }


        //fmt.Println( nodes.Items[i] )
        //fmt.Println( fmt.Printf( "%s\t%s\t%s\t%s", name, ipaddr, ipaddr_internal, uid ) )


        data := map[string]string { "node_name": name, "ip": ipaddr }
        res = append( res, data )
    }

    //fmt.Println( res )
    return res
}

func (t *K8s_APIS) deployment_list(NAMESPACE string) []interface{} {
    fmt.Println( types.Trace__() )

    //NAMESPACE := "default"
    //deployments := t.clientset.AppsV1().Deployments( NAMESPACE )
    deployments, err := t.clientset.AppsV1().Deployments(NAMESPACE).List( context.TODO(), metav1.ListOptions{} )
    if err != nil {
        panic( err.Error() )
    }
    //fmt.Println( deployments )

    var res []interface{}
    for i := 0; i < len(deployments.Items); i++ {
        namespace := deployments.Items[i].ObjectMeta.Namespace
        if namespace == "kube-system" {
            continue
        }

        name := deployments.Items[i].ObjectMeta.Name


        //fmt.Println( deployments.Items[i] )
        //fmt.Println( fmt.Printf( "%s\t%s", namespace, name ) )


        data := map[string]string { "namespace": namespace, "name": name }
        res = append( res, data )
    }

    //fmt.Println( res )
    return res
}

func (t *K8s_APIS) pod_list(NAMESPACE string) []interface{} {
    fmt.Println( types.Trace__() )

    pods, err := t.clientset.CoreV1().Pods(NAMESPACE).List( context.TODO(), metav1.ListOptions{} )
    if err != nil {
        panic( err.Error() )
    }
    //fmt.Println( pods )

    var res []interface{}
    for i := 0; i < len(pods.Items); i++ {
        namespace := pods.Items[i].ObjectMeta.Namespace
        if namespace == "kube-system" {
            continue
        }

        var selector string
        _, ok := pods.Items[i].ObjectMeta.Labels["app"]
        if ok {
            selector = pods.Items[i].ObjectMeta.Labels["app"]
        } else {
            continue
        }

        var status string
        if pods.Items[i].ObjectMeta.DeletionTimestamp != nil {
            // "Pending", "Running"
            if pods.Items[i].Status.Phase == apiv1.PodPending ||
                pods.Items[i].Status.Phase == apiv1.PodRunning {
                status = "Terminating"
            } else {
                status = string( pods.Items[i].Status.Phase )
            }
        } else {
            status = string( pods.Items[i].Status.Phase )
        }

        var name string
        var pod_ip string
        var node_name string
        var image string
        var creation_timestamp string

        name = pods.Items[i].ObjectMeta.Name
        pod_ip = pods.Items[i].Status.PodIP
        node_name = pods.Items[i].Spec.NodeName
        image = pods.Items[i].Spec.Containers[0].Image
        creation_timestamp = types.Int64ToStr( pods.Items[i].ObjectMeta.CreationTimestamp.Unix() )


        //fmt.Println( pods.Items[i] )
        //fmt.Println( fmt.Printf("%s\t%s\t%s\t%s\t%s\t%s\t%s",
        //      pod_ip, namespace, name, node_name, selector, image, creation_timestamp) )


        data := map[string]string {
            "namespace": namespace,
            "node_name": node_name,
            "name": name,
            "pod_ip": pod_ip,
            "selector": selector,
            "image": image,
            "status": status,
            "creation_timestamp": creation_timestamp,
        }
        res = append( res, data )
    }

    // descending order by 'creation_timestamp'
    sort.SliceStable(
        res,
        func(i, j int) bool {
            return ( res[i].(map[string]string)["creation_timestamp"] <
                    res[j].(map[string]string)["creation_timestamp"] )
        },
    )

    //fmt.Println( res )
    return res
}

func (t *K8s_APIS) service_list(NAMESPACE string) []interface{} {
    fmt.Println( types.Trace__() )

    services, err := t.clientset.CoreV1().Services(NAMESPACE).List( context.TODO(), metav1.ListOptions{} )
    if err != nil {
        panic( err.Error() )
    }
    //fmt.Println( services )

    var res []interface{}
    for i := 0; i < len(services.Items); i++ {
        namespace := services.Items[i].ObjectMeta.Namespace
        if namespace == "kube-system" {
            continue
        }

        var selector string
        _, ok := services.Items[i].Spec.Selector["app"]
        if ok {
            selector = services.Items[i].Spec.Selector["app"]
        } else {
            continue
        }


        var cluster_ip string
        var external_ips []string
        var name string
        var ports []map[string]interface{}


        //fmt.Println( services.Items[i].Spec.Ports )
        for _, val := range services.Items[i].Spec.Ports {
            //fmt.Println( "key = ", k, ", val = ", val )
            //ports["name"] = val.Name
            //ports["node_port"] = val.NodePort
            //ports["port"] = val.Port
            // TargetPort.IntVal, TargetPort.StrVal
            //ports["target_port"] = val.TargetPort.IntVal
            //ports["protocol"] = val.Protocol

            _port := map[string]interface{} {
                "name": val.Name,
                "node_port": val.NodePort,
                "port": val.Port,

                // TargetPort.IntVal, TargetPort.StrVal
                "target_port": val.TargetPort.IntVal,

                "protocol": val.Protocol,
            }

            //fmt.Println( "name = ", val.Name )
            //fmt.Println( "node_port = ", val.NodePort )
            //fmt.Println( "port = ", val.Port )
            //fmt.Println( "target_port = ", val.TargetPort, "IntVal = ", val.TargetPort.IntVal )
            //fmt.Println( "protocol = ", val.Protocol )

            ports = append( ports, _port )
        }
        //fmt.Println( ports )


        cluster_ip = services.Items[i].Spec.ClusterIP
        external_ips = services.Items[i].Spec.ExternalIPs
        name = services.Items[i].ObjectMeta.Name


        //fmt.Println( pods.Items[i] )
        //fmt.Println( fmt.Printf( "%s\t%s\t%s\t%s\t%s", cluster_ip, external_i_ps, namespace, name, selector ) )
        //fmt.Println( ports )


        data := map[string]interface{} {
            "namespace": namespace,
            "name": name,
            "selector": selector,
            "ports": ports,
            "cluster_ip": cluster_ip,
            "external_ips": external_ips,
        }
        res = append( res, data )
    }

    //fmt.Println( res )
    return res
}

func (t *K8s_APIS) ingress_list(NAMESPACE string) []interface{} {
    fmt.Println( types.Trace__() )

    services, err_svc := t.clientset.CoreV1().Services(NAMESPACE).List( context.TODO(), metav1.ListOptions{} )
    if err_svc != nil {
        panic( err_svc.Error() )
    }
    //fmt.Println( services )

    ingress_data := map[string]interface{} { "name": "", "port": 0 }
    for i := 0; i < len(services.Items); i++ {
        name := services.Items[i].ObjectMeta.Name

        namespace := services.Items[i].ObjectMeta.Namespace
        if namespace == "kube-system" {
            continue
        }

        //var selector string
        found_ingress := false
        _, ok := services.Items[i].Spec.Selector["app"]
        if ok {
            // found, pass
            continue
        } else {
            if found_ingress == false {
                if namespace == "ingress-nginx" && name == "ingress-nginx-controller-svc" {
                    found_ingress = true
                    //fmt.Println( "ingress found..." )

                    var ports []interface{}
                    //fmt.Println( services.Items[i].Spec.Ports )
                    for _, val := range services.Items[i].Spec.Ports {
                        //fmt.Println( "key = ", k, ", val = ", val )
                        //fmt.Println( "name = ", val.Name )
                        //fmt.Println( "node_port = ", val.NodePort )
                        //fmt.Println( "port = ", val.Port )
                        //fmt.Println( "target_port = ", val.TargetPort, "IntVal = ", val.TargetPort.IntVal )
                        //fmt.Println( "protocol = ", val.Protocol )

                        port := map[string]interface{} {
                            "name": val.Name,
                            "node_port": val.NodePort,
                            "port": val.Port,
                            // TargetPort.IntVal, TargetPort.StrVal
                            "target_port": val.TargetPort.IntVal,
                            "protocol": val.Protocol,
                        }
                        ports = append( ports, port )
                    }
                    //fmt.Println( ports )

                    ingress_data["name"] = services.Items[i].ObjectMeta.Name
                    //ingress_data["name"] = ports[0]["name"]
                    ingress_data["node_port"] = ports[0].(map[string]interface{})["node_port"]
                    ingress_data["port"] = ports[0].(map[string]interface{})["port"]
                    ingress_data["target_port"] = ports[0].(map[string]interface{})["target_port"]
                    break
                }
            }
        }
    }



    ingresses, err := t.clientset.NetworkingV1().Ingresses(NAMESPACE).List( context.TODO(), metav1.ListOptions{} )
    if err != nil {
        panic( err.Error() )
    }
    //fmt.Println( ingresses )

    var res []interface{}
    for i := 0; i < len(ingresses.Items); i++ {
        name := ingresses.Items[i].ObjectMeta.Name

        namespace := ingresses.Items[i].ObjectMeta.Namespace
        if namespace == "kube-system" {
            continue
        }


        var service_name string
        var paths []interface{}


        for _, rule := range ingresses.Items[i].Spec.Rules {
            for _, path := range rule.IngressRuleValue.HTTP.Paths {
                //fmt.Println( "-->", path.Backend.Service.Name )
                service_name = path.Backend.Service.Name
                path_endpoint := path.Path
                paths = append( paths, path_endpoint )
            }
        }


        //fmt.Println( pods.Items[i] )
        //fmt.Println( fmt.Printf( "%s\t%s\t%s\t%s\t%s", cluster_ip, external_i_ps, namespace, name, selector ) )


        data := map[string]interface{} {
            "namespace": namespace,
            "ingress_name": name,
            "paths": paths,
            "service_name": service_name,
            "ports": ingress_data,
        }
        res = append( res, data )
    }

    //fmt.Println( res )
    return res
}

// container (Pod) details
func (t *K8s_APIS) container_details_list(username string) interface{} {
    fmt.Println( types.Trace__() )
    fmt.Println( "username = ", username )

    //res := map[string]interface{} { "result": "" }
    //res := make( map[string]interface{} )
    var res []interface{}

    NAMESPACE_NULL := ""

    _node_list := t.node_list()
    _deployment_list := t.deployment_list( NAMESPACE_NULL )
    _pod_list := t.pod_list( NAMESPACE_NULL )
    _service_list := t.service_list( NAMESPACE_NULL )
    _ingress_list := t.ingress_list( NAMESPACE_NULL )

    fmt.Println( "-------" )

    // TODO:
    // username from DB
    //username := username + "-"
    //fmt.Println( "username = ", username )
    username = "admin" + "-"
    USERNAME_ADMIN := "admin-"
    found_username := false
    fmt.Println( "username = ", username )


    var info []interface{}
    for _, pod := range _pod_list {
        data := map[string]interface{} { "pod": pod }

        if username == USERNAME_ADMIN {
            for _, deployment := range _deployment_list {
                if pod.(map[string]string)["selector"] == deployment.(map[string]string)["name"] {
                    data["deployment"] = deployment
                }
            }

            for _, service := range _service_list {
                if pod.(map[string]string)["selector"] ==
                        service.(map[string]interface{})["selector"].(string) {
                    data["service"] = service

                    for _, ingress := range _ingress_list {
                        if ingress.(map[string]interface{})["service_name"].(string) ==
                                service.(map[string]interface{})["name"].(string) {
                            data["ingress"] = ingress
                        }
                    }
                }
            }

            for _, node := range _node_list {
                if pod.(map[string]string)["node_name"] ==
                        node.(map[string]string)["node_name"] {
                    data["node"] = node
                }
            }
        } else {
            found_username = false
            for _, deployment := range _deployment_list {
                if pod.(map[string]string)["selector"] ==
                        deployment.(map[string]string)["name"] {
                    if strings.Contains(deployment.(map[string]string)["name"], username) == false {
                        continue
                    }
                    found_username = true
                    data["deployment"] = deployment
                }
            }

            for _, service := range _service_list {
                if pod.(map[string]string)["selector"] ==
                        service.(map[string]interface{})["selector"].(string) {
                    if strings.Contains(service.(map[string]interface{})["name"].(string), username) == false {
                        continue
                    }
                    found_username = true
                    data["service"] = service

                    for _, ingress := range _ingress_list {
                        if ingress.(map[string]interface{})["service_name"].(string) ==
                                service.(map[string]interface{})["name"].(string) {
                            data["ingress"] = ingress
                        }
                    }
                }
            }

            if found_username != true {
                continue
            }

            for _, node := range _node_list {
                if pod.(map[string]string)["node_name"] == node.(map[string]string)["node_name"] {
                    data["node"] = node
                }
            }
        }

        info = append( info, data )
    } // for
    //fmt.Println( info )


    // service has no pod and deployment
    for _, service := range _service_list {
        found_service := false
        service_data := map[string]interface{} {
            "pod": map[string]interface{} {},
            "deployment": map[string]string {"name": ""},
            "node": map[string]string{"ip":""},
        }

        for _, pod := range _pod_list {
            if len(service.(map[string]interface{})["selector"].(string)) > 0 &&
                    service.(map[string]interface{})["selector"].(string) == pod.(map[string]string)["selector"] {
                found_service = true
                break
            }
        }

        if found_service != true {
            if service.(map[string]interface{})["name"].(string) == "kubernetes" {
                continue
            }

            if username == USERNAME_ADMIN {
                // pass
            } else {
                if strings.Contains(service.(map[string]interface{})["name"].(string), username) == false {
                    continue
                }
            }

            service_data["service"] = service

            for _, ingress := range _ingress_list {
                if ingress.(map[string]interface{})["service_name"].(string) ==
                        service.(map[string]interface{})["name"].(string) {
                    service_data["ingress"] = ingress
                }
            }

            info = append( info, service_data )
        }
    } // for
    //fmt.Println( info )


    //res["result"] = info
    res = info
    //fmt.Println( res )
    //fmt.Println( types.JsonPrettyPrint_Map(res) )
    //fmt.Println( types.JsonPrettyPrint_Interface(res) )
    return res
}

// container_ports: node port(expose port): ["node port1", "node port2", ...]
// opts = { "use_command": True,
//           "command": ["/bin/sh", "-c", ""], "args": [],
//           "host_aliases": {"ip": "", "hostnames":[]}
// }
func (t *K8s_APIS) create_deployment(
    NAMESPACE string,
    deployment_name string,
    deployment_image string,
    container_ports []string,
    opts interface{},
) []interface{} {
    fmt.Println( types.Trace__() )
    fmt.Println( "namespace = ", NAMESPACE )
    fmt.Println( "deployment name = ", deployment_name )
    fmt.Println( "deployment image = ", deployment_image )
    fmt.Println( "container ports = ", container_ports )
    fmt.Println( "options = ", opts )

    var _container_ports []apiv1.ContainerPort
    for _, val := range container_ports {
        //_node_port, _ := strconv.Atoi( val )
        //_node_port_int32 := int32( _node_port )
        //_node_port_int64, _ := strconv.ParseInt( val, 10, 64 )
        //_node_port_int32 := int32( _node_port_int64 )

        //_container_port := apiv1.ContainerPort { ContainerPort: _node_port_int32 }
        _container_port := apiv1.ContainerPort { ContainerPort: types.StrToInt32(val) }

        _container_ports = append( _container_ports, _container_port )
    }



    var containers []apiv1.Container

    container := apiv1.Container {
        Name: deployment_name,
        Image: deployment_image,
        Ports: _container_ports,
    }
    containers = append( containers, container )
    //fmt.Println( containers )


    // Affinity
    {
    }


    /*
    if opts != nil && opts.(map[string]interface{})["use_command"] == true {
        command := []string { opts.(map[string]string)["command"] }
        //args := []string { opts.(map[string]string{})["args"] }
        //host_aliases := { "ip": "", "hostnames":[] }

        container := apiv1.Container {
            Name: deployment_name,
            Image: deployment_image,
            Ports: _container_ports,

            Lifecycle: &apiv1.Lifecycle {
                PostStart: &apiv1.LifecycleHandler {
                    Exec: &apiv1.ExecAction {
                        Command: command,
                    },
                },
            },
        }
    } else {
        container := apiv1.Container {
            Name: deployment_name,
            Image: deployment_image,
            Ports: _container_ports,
        }
    }
    containers = append( containers, container )
    */

    // Secrets for Docker Private Registry
    // $ sudo kubectl create secret docker-registry <new-secret-name>
    //    --docker-server=<registry-server>
    //    --docker-username=<username>
    //    --docker-password=<password>
    //    #--docker-email=<email>
    // 
    //image_pull_secrets := []apiv1.LocalObjectReference { {Name: "test_key"} }


    var template apiv1.PodTemplateSpec

    template = apiv1.PodTemplateSpec {
        ObjectMeta: metav1.ObjectMeta {
            Labels: map[string]string { "app": deployment_name },
        },
        Spec: apiv1.PodSpec {
            Containers: containers,
            //ImagePullSecrets: image_pull_secrets,
        },
    }
    fmt.Println( template )


    /*
    var pod_spec apiv1.PodSpec

    //if opts != nil && opts.(map[string]interface{})["use_command"] == true {
    if opts != nil && opts.(map[string]interface{})["use_command"] != false {
        //host_aliases := []apiv1.HostAlias { {"ip": "", "hostnames":[]} }

        host_aliases := []apiv1.HostAlias {
            {
                IP: opts.(map[string]interface{})["host_aliases"].(map[string]string)["ip"],
                Hostnames: []string {
                    opts.(map[string]interface{})["host_aliases"].(map[string]string)["hostnames"],
                },
            },
        }

        pod_spec = apiv1.PodSpec {
            Containers: containers,
            //ImagePullSecrets: image_pull_secrets,
            HostAliases: host_aliases,
        }
    } else {
        pod_spec = apiv1.PodSpec {
            Containers: containers,
            //ImagePullSecrets: image_pull_secrets,
        }
    }
    template = apiv1.PodTemplateSpec {
        ObjectMeta: metav1.ObjectMeta {
            Labels: map[string]string { "app": deployment_name },
        },
        //Spec: apiv1.PodSpec {
        //    Containers: containers,
        //    //ImagePullSecrets: image_pull_secrets,
        //},
        Spec: pod_spec,
    }
    fmt.Println( template )
    */


    int32ptr := func (i int32) *int32 { return &i }
    deployment_spec := appsv1.DeploymentSpec {
        //Replicas: 1,
        Replicas: int32ptr(1),
        Selector: &metav1.LabelSelector {
            MatchLabels: map[string]string { "app": deployment_name },
        },
        Template: template,
    }

    deployment := &appsv1.Deployment {
        //api_version = "apps/v1",
        //kind = "Deployment",
        ObjectMeta: metav1.ObjectMeta { Name: deployment_name },
        Spec: deployment_spec,
    }
    fmt.Println( deployment )


    result, err_deployment := t.clientset.AppsV1().Deployments(NAMESPACE).Create( context.TODO(), deployment, metav1.CreateOptions{} )
    if err_deployment != nil {
        panic( err_deployment.Error() )
    }
    fmt.Println( result )


    return []interface{} { result }
}

// service_name = deployment_name + "-service"
// worker_node_ips: ["ip address", "..."]
// service_ports: dict array: [{"name":"", "node_port":"", "port":"", "target_port":""}, {...}]
func (t *K8s_APIS) create_service(
    NAMESPACE string,
    deployment_name string,
    deployment_service_name string,
    //worker_node_ips []string,
    //service_ports []types.Req_JSONRPC_k8s_service_ports_Args_st,
    service_ports []map[string]interface{},
) interface{} {
    fmt.Println( types.Trace__() )
    fmt.Println( "deploymeny name = ", deployment_name )
    fmt.Println( "deploymeny service name = ", deployment_service_name )
    //fmt.Println( "worker node ips = ", worker_node_ips )
    fmt.Println( "service ports = ", service_ports )


    var _service_ports []apiv1.ServicePort

    for _, val := range service_ports {
        _name := val["name"].(string)
        //_node_port := types.StrToInt32( val["node_port"].(string) )
        _port := types.StrToInt32( val["port"].(string) )
        _target_port := types.StrToInt32( val["target_port"].(string) )

        //_node_port_int64, _ := strconv.ParseInt( _node_port, 10, 64 )
        //_node_port_int32 := int32( _node_port_int64 )
        //_port_int64, _ := strconv.ParseInt( _port, 10, 64 )
        //_port_int32 := int32( _port_int64 )
        //_target_port_int64, _ := strconv.ParseInt( _target_port, 10, 64 )
        //_target_port_int32 := int32( _target_port_int64 )

        _service_port := apiv1.ServicePort {
            Name: _name,
            //NodePort: _node_port_int32,
            Port: _port,
            TargetPort: intstr.IntOrString {
                // TargetPort.IntVal, TargetPort.StrVal
                IntVal: _target_port,
            },
        }

        _service_ports = append( _service_ports, _service_port )
    }
    fmt.Println( _service_ports )


    service := &apiv1.Service {
        //api_version = "v1",
        //kind = "Service",
        ObjectMeta: metav1.ObjectMeta { Name: deployment_service_name },
        Spec: apiv1.ServiceSpec {
            //Type = ("LoadBalancer", "NodePort", "ClusterIP"),
            //ExternalIPs: worker_node_ips, // array: []
            Selector: map[string]string { "app": deployment_name },
            //Ports: []apiv1.ServicePort {
            //    { Name: , NodePort: , Port: , TargetPort: , },
            //    { Name: , NodePort: , Port: , TargetPort: , },
            //},
            Ports: _service_ports,
        },
    }


    result, err_service := t.clientset.CoreV1().Services(NAMESPACE).Create( context.TODO(), service, metav1.CreateOptions{} )
    if err_service != nil {
        panic( err_service.Error() )
    }
    fmt.Println( result )


    return result
}

// ingress_name = deployment_name + "-ingress"
// host = ""
// namespace = DEFAULT_NAMESPACE
func (t *K8s_APIS) create_ingress(
    NAMESPACE string,
    ingress_name string,
    host string,
    service_name string,
    //service_ports []types.Req_JSONRPC_k8s_service_ports_Args_st,
    service_ports []map[string]interface{},
) interface{} {
    fmt.Println( types.Trace__() )

    fmt.Println( "namespace: ", NAMESPACE )
    fmt.Println( "ingress_name: ", ingress_name )
    fmt.Println( "host: ", host )
    //fmt.Println( "endpoint: ", endpoint )
    fmt.Println( "dst svc name: ", service_name )
    fmt.Println( "dst svc ports: ", service_ports )

    // Nginx Ingress Controller
    //- https://github.com/kubernetes/ingress-nginx
    //- https://github.com/kubernetes/ingress-nginx/blob/main/docs/deploy/index.md
    //- https://github.com/kubernetes/ingress-nginx/tree/main/deploy/static/provider/cloud
    //
    //(old)
    //$ kubectl apply -f https://raw.githubusercontent.com/kubernetes/ingress-nginx/master/deploy/static/mandatory.yaml
    //$ kubectl apply -f https://raw.githubusercontent.com/kubernetes/ingress-nginx/main/deploy/static/provider/cloud/deploy.yaml
    //
    //
    //$ kubectl get deployment --namespace=ingress-nginx
    //$ kubectl get pods --namespace=ingress-nginx
    //
    //$ kubectl expose deploy nginx-ingress-controller -name nginx-ingress-controller-svc --type=NodePort -n -ingress-nginx
    //$ kubectl get service -o wide
    //
    // create pod with ingress (pathname, ...)
    //...
    //(error: failed calling webhook "validate.nginx.ingress.kubernetes.io" ...)
    //$ kubectl delete validatingwebhookconfiguration ingress-nginx-admission
    //
    //$ kubectl get ingress
    //$ kubectl describe ingress <name>
    //
    //http://127.0.0.1:[ingress nodeport]/pathname
    //
    //reboot nodes if doesn't work
    //
    //or
    //
    //ingress-nginx: 504 timed-out error
    //.spec.externalTrafficPolicy: "Local" -> "Cluster"
    //
    //https://pwittrock.github.io/docs/tasks/access-application-cluster/create-external-load-balancer/#preserving-the-client-source-ip
    //
    //service.spec.externalTrafficPolicy
    //
    //This feature can be activated by setting externalTrafficPolicy to “Local” in the Service Configuration file.
    //{
    //    "kind": "Service",
    //        "apiVersion": "v1",
    //        "metadata": {
    //            "name": "example-service",
    //        },
    //        "spec": {
    //            "ports": [{
    //                "port": 8765,
    //                "targetPort": 9376
    //            }],
    //            "selector": {
    //                "app": "example"
    //            },
    //            "type": "LoadBalancer",
    //            "externalTrafficPolicy": "Local"
    //        }
    //}


    var paths []networkingv1.HTTPIngressPath

    for _, port := range service_ports {
        //_name := port["name"]
        //_port := port["port"].(int32)
        //_port_int64, _ := strconv.ParseInt( _port, 10, 64 )
        //_port_int32 := int32( _port_int64 )

        //_node_port := types.StrToInt32( port["node_port"].(string) )
        _port := types.StrToInt32( port["port"].(string) )
        //_target_port := types.StrToInt32( port["target_port"].(string) )


        endpoint := "/" + service_name + "/" + types.Int32ToStr(_port)

        path_type := networkingv1.PathTypeExact
        path := networkingv1.HTTPIngressPath {
            Path: endpoint, // "/test_path1"
            PathType: &path_type, // PathTypeExact, PathTypePrefix ("Exact", "Prefix")
            Backend: networkingv1.IngressBackend {
                Service: &networkingv1.IngressServiceBackend {
                    Port: networkingv1.ServiceBackendPort {
                        Number: _port, // 80
                    },
                    Name: service_name, // "test_path_service"
                },
            },
        }

        paths = append( paths, path )
    }

    ingress := &networkingv1.Ingress {
        //api_version = "networking.k8s.io/v1",
        //kind = "Ingress",
        ObjectMeta: metav1.ObjectMeta {
            Name: ingress_name,
            Annotations: map[string]string {
                // annotations or ingress_class_name
                "kubernetes.io/ingress.class": "nginx",
                "nginx.ingress.kubernetes.io/rewrite-target": "/",
                //"nginx.ingress.kubernetes.io/ssl-redirect": "false",
            },
        },
        Spec: networkingv1.IngressSpec {
            // annotations or ingress_class_name
            //IngressClassName: StrPtr("nginx"),
            Rules: []networkingv1.IngressRule {
                {
                    // http://example.com/test_path1 -> (ingress) -> (service: test_path_svc:80) -> (pod:80)
                    //

                    // connect to IP if not specified 'host' 
                    //Host: host, // "example.com"

                    IngressRuleValue: networkingv1.IngressRuleValue {
                        HTTP: &networkingv1.HTTPIngressRuleValue {
                            //
                            //paths: []interface{} {
                            //    {
                            //        networkingv1.HTTPIngressPath {
                            //            Path: endpoint, // "/test_path1"
                            //            PathType: StrPtr("Exact"),
                            //            Backend: networkingv1.V1IngressBackend {
                            //                Service: networkingv1.V1IngressServiceBackend {
                            //                    Port: networkingv1.V1ServiceBackendPort {
                            //                        Number: dst_svc_port // 80
                            //                    },
                            //                    Name: dst_svc_name // "test_path_svc"
                            //                },
                            //            },
                            //        },
                            //    },
                            //}
                            //
                            Paths: paths,
                        },
                    },
                },
            }, // Rules
        }, // Spec
    } // networkingv1.Ingress

    // Creation of the Deployment in specified namespace
    // (Can replace "default" with a namespace you may have created)
    result, err_ingress := t.clientset.NetworkingV1().Ingresses(NAMESPACE).Create( context.TODO(), ingress, metav1.CreateOptions{} )
    if err_ingress != nil {
        panic( err_ingress.Error() )
    }
    fmt.Println( result )


    return result
}

// ingress_name = ""
func (t *K8s_APIS) delete_container(
    NAMESPACE string,
    deployment_name string,
    service_name string,
    ingress_name string,
) (interface{}, interface{}, interface{}) {
    fmt.Println( types.Trace__() )

    var ret_deployment interface{}
    var ret_service interface{}
    var ret_ingress interface{}


    // delete deployment
    int64ptr := func (i int64) *int64 { return &i }
    if deployment_name != "" && len(strings.Trim(deployment_name, "")) > 0 {
        propagation_policy_foreground := metav1.DeletePropagationForeground
        delete_options := metav1.DeleteOptions {
            //PropagationPolicy: "Foreground",
            PropagationPolicy: &propagation_policy_foreground,
            GracePeriodSeconds: int64ptr(5),
        }
        deployment := t.clientset.AppsV1().Deployments( NAMESPACE )
        ret_deployment = deployment.Delete( context.TODO(), strings.Trim(deployment_name, ""), delete_options )
    } else {
        fmt.Println( "deployment_name == null" )
    }
    fmt.Println( ret_deployment )

    // delete service
    if service_name != "" && len(strings.Trim(service_name, "")) > 0 {
        delete_options := metav1.DeleteOptions {}
        service := t.clientset.CoreV1().Services( NAMESPACE )
        ret_service = service.Delete( context.TODO(), strings.Trim(service_name, ""), delete_options )
    } else {
        fmt.Println( "service_name == null" )
    }
    fmt.Println( ret_service )

    // delete ingress
    if ingress_name != "" && len(strings.Trim(ingress_name, "")) > 0 {
        delete_options := metav1.DeleteOptions {}
        ingress := t.clientset.NetworkingV1().Ingresses( NAMESPACE )
        ret_ingress = ingress.Delete( context.TODO(), strings.Trim(ingress_name, ""), delete_options )
    } else {
        fmt.Println( "ingress_name == null" )
    }
    fmt.Println( ret_ingress )


    return ret_deployment, ret_service, ret_ingress
}

// INIT = False
func (t *K8s_APIS) port_forward_iptables_add_rule(
    SRC_IP string,
    DST_IP string,
    DPORT string,
    INIT bool,
) interface{} {
    fmt.Println( types.Trace__() )
    return nil
}

func (t *K8s_APIS) port_forward_iptables_del_rule(DPORT string) interface{} {
    fmt.Println( types.Trace__() )
    return nil
}

// SRC_IP = None
// DST_IP = None
// DPORT = None
func (t *K8s_APIS) port_forward(
    flag_add bool,
    deployment_name string,
    deployment_service_name string,
    username string,
    SRC_IP string,
    DST_IP string,
    DPORT string,
) interface{} {
    fmt.Println( types.Trace__() )
    return nil
}

// SRC_IP = None
// DST_IP = None
// DPORT = None
func (t *K8s_APIS) create_port_forward(
    deployment_name string,
    deployment_service_name string,
    SRC_IP string,
    DST_IP string,
    DPORT string,
) interface{} {
    fmt.Println( types.Trace__() )
    return nil
}

// username = None
// DPORT = None
func (t *K8s_APIS) delete_port_forward(
    deployment_name string,
    deployment_service_name string,
    username string,
    DPORT string,
) interface{} {
    fmt.Println( types.Trace__() )
    return nil
}

//func (t *K8s_APIS) request_containers() []interface{} {
//    fmt.Println( types.Trace__() )
//    return nil
//}

func (t *K8s_APIS) request_container_service_test() {
    fmt.Println( types.Trace__() )
}

func (t *K8s_APIS) request_container_service(
    _req_json map[string]interface{},
    //req_json types.Req_JSONRPC_k8s_request_service_Args_st,
) map[string]interface{} {
    fmt.Println( types.Trace__() )

    NAMESPACE_DEFAULT := "default"
    timestamp := time.Now().Unix()
    fmt.Println( "timestamp = ", timestamp )


    var ret_deployment []interface{}
    var ret_deployment_service interface{}
    var ret_deployment_ingress interface{}

    req_json := _req_json["request"].(map[string]interface{})


    //
    // All-in-one
    //

    /*
    // Creates Deployment
    deployment_name := "test100" + "-" + types.IntToStr( timestamp )
    deployment_image := "10.0.2.5:5000/test_img_apm:1.0"
    container_ports := []string { "80", "8080", "3306" }
    var opts interface{}
    ret_deployment = t.create_deployment(
        NAMESPACE_DEFAULT,
        deployment_name,
        deployment_image,
        container_ports,
        opts,
    )


    // Creates Service
    deployment_service_name := deployment_name + "-service"
    deployment_service_ports := []map[string]string {
        {
            "name": "apache",
            "node_port": "",
            "port": "80",
            "target_port": "80",
        },
        {
            "name": "tomcat",
            "node_port": "",
            "port": "8080",
            "target_port": "8080",
        },
        {
            "name": "mariadb",
            "node_port": "",
            "port": "3306",
            "target_port": "3306",
        },
    }
    ret_deployment_service = t.create_service(
        NAMESPACE_DEFAULT,
        deployment_name,
        deployment_service_name,
        deployment_service_ports,
    )


    // Ingress
    deployment_ingress_name := deployment_name + "-ingress"
    deployment_ingress_host := "test.com"
    ret_deployment_ingress = t.create_ingress(
        NAMESPACE_DEFAULT,
        deployment_ingress_name,
        deployment_ingress_host,
        deployment_service_name,
        deployment_service_ports,
    )
    */



    //
    // Standalone
    //

//! NOT TESTED
/*
    deployment_service_ports := []map[string]string {
        {
            "name": "apache",
            "node_port": "",
            "port": "80",
            "target_port": "80",
        },
        {
            "name": "tomcat",
            "node_port": "",
            "port": "8080",
            "target_port": "8080",
        },
        {
            "name": "mariadb",
            "node_port": "",
            "port": "3306",
            "target_port": "3306",
        },
    }

    var service_ports_apache []map[string]string
    var service_ports_tomcat []map[string]string
    var service_ports_mariadb []map[string]string
    for _, i := range deployment_service_ports {
        if i["name"] == "apache" {
            service_ports_apache = append( service_ports_apache, i )
        }
        if i["name"] == "tomcat" {
            service_ports_tomcat = append( service_ports_tomcat, i )
        }
        if i["name"] == "mariadb" {
            service_ports_mariadb = append( service_ports_mariadb, i )
        }
    }


    //var req_json map[string]interface{}
    //req_json := make( map[string]interface{} )
    req_json := map[string]interface{} {
        //"deployment_name_mariadb": map[string]interface{} {},
        "deployment_name_mariadb": "test000",
        "deployment_image_mariadb": "127.0.0.1:5000/mariadb:1.0",
        "deployment_service_name": "test000",
    }

    // MariaDB
    deployment_name := req_json["deployment_name_mariadb"].(string) + "-" + types.IntToStr( timestamp )
    deployment_image := req_json["deployment_image_mariadb"].(string)
    deployment_service_name := req_json["deployment_service_name"].(string) + "-mariadb"
    opts := map[string]interface{} {
        // Docker
        // RUN sed -i -e 's/^bind-address/#&/' /etc/mysql/mariadb.conf.d/50-server.cnf
        // ENTRYPOINT service mysql restart && tail -f /dev/null

        "command": []string {
            "/bin/sh",
            "-c",
            "/usr/bin/sed -i -e 's/^bind-address/#&/' /etc/mysql/mariadb.conf.d/50-server.cnf" +
            " && " +
            "service mysql restart",
        },
        "args": []string {},
        "host_aliases": map[string]interface{} {},
    }
    fmt.Println( "deployment name = ", deployment_name )
    fmt.Println( "deployment image = ", deployment_image )
    fmt.Println( "deployment service name = ", deployment_service_name )
    fmt.Println( "deployment opts = ", opts )

    //create_deployment( deployment_name, deployment_image, container_ports, opts )
    //ret_mariadb := create_service( deployment_name, deployment_service_name, worker_node_ips, service_ports )
    //db_service_hostname := "database"
    //db_ip := ret_mariadb.spec.cluster_ip
    //db_port := 0
    //fmt.Println( "mariadb ip = ", db_ip )
    //
    ////for _, port := range ret_mariadb.spec.ports {
    ////    if port.name == "mariadb" {
    ////        db_port = types.StrToInt32( port.node_port )
    ////    }
    ////}
    ////fmt.Println( "mariadb port = ", db_port )


    // Tomcat
    db_service_hostname := "database"
    db_ip := "ip" // ret_mariadb.spec.cluster_ip
    deployment_name = req_json["deployment_name_tomcat"].(string) + "-" + types.IntToStr( timestamp )
    deployment_image = req_json["deployment_image_tomcat"].(string)
    deployment_service_name = req_json["deployment_service_name"].(string) + "-tomcat"
    opts = map[string]interface{} {
        // Docker
        // ENTRYPOINT /usr/share/tomcat9/bin/startup.sh && apachectl -D FOREGROUND

        "command": []string {
            "/bin/sh", "-c", "/usr/share/tomcat9/bin/startup.sh && service apache2 restart",
        },
        "args": []string {},
        "host_aliases": map[string]interface{} {
            "ip": db_ip,
            "hostnames": []string { db_service_hostname },
        },
    }
    fmt.Println( "deployment name = ", deployment_name )
    fmt.Println( "deployment image = ", deployment_image )
    fmt.Println( "deployment service name = ", deployment_service_name )
    fmt.Println( "deployment opts = ", opts )

    //create_deployment( deployment_name, deployment_image, container_ports, opts )
    //create_service( deployment_name, deployment_service_name, worker_node_ips, service_ports )
*/



    //deployment_name := "test100" + "-" + types.IntToStr( timestamp )
    //deployment_image := "10.0.2.5:5000/test_img_apm:1.0"
    //container_ports := []string { "80", "3306" }
    //var opts interface{}

    //deployment_service_name := deployment_name + "-service"
    //deployment_service_ports := []map[string]interface{}

    //deployment_ingress_name := deployment_name + "-ingress"
    //deployment_ingress_host := ""


    fmt.Println( "request = ", req_json )



    //var res []interface{}
    //res := map[string]interface{} { "result": "" }
    res := make( map[string]interface{} )

    req_type := req_json["req"].(string)
    if req_type == types.ReqType_create {
        deployment_name := req_json["deployment_name"].(string) + "-" + types.IntToStr( timestamp )
        deployment_image := req_json["deployment_image"].(string)

        var container_ports []string
        {
            _container_ports := req_json["container_ports"].([]interface{})
            for _, val := range _container_ports {
                //fmt.Println( "type = ", reflect.TypeOf(val) )
                container_ports = append( container_ports, val.(string) )
            }
        }

        var opts interface{}
        deployment_service_name := deployment_name + "-service"
        var deployment_service_ports []map[string]interface{}
        {
            _deployment_service_ports := req_json["service_ports"].([]interface{})
            for _, val := range _deployment_service_ports {
                //fmt.Println( "type = ", reflect.TypeOf(val) )

                _deployment_service_port := val.(map[string]interface{})
                _port := make( map[string]interface{} )
                for _k, _v := range _deployment_service_port {
                    //fmt.Println( _k, " = ", _v )
                    _port[_k] = _v.(string)
                }
                //fmt.Println( "port = ", _port )
                deployment_service_ports = append( deployment_service_ports, _port )
            }
        }

        deployment_ingress_name := deployment_name + "-ingress"
        deployment_ingress_host := ""

        fmt.Println( "deployment name = ", deployment_name )
        fmt.Println( "deployment image = ", deployment_image)
        fmt.Println( "container ports = ", container_ports )
        fmt.Println( "service name = ", deployment_service_name )
        fmt.Println( "service ports = ", deployment_service_ports )
        fmt.Println( "ingress name = ", deployment_ingress_name )
        fmt.Println( "host name = ", deployment_ingress_host )

        ret_deployment = t.create_deployment(
            NAMESPACE_DEFAULT,
            deployment_name,
            deployment_image,
            container_ports,
            opts,
        )

        ret_deployment_service = t.create_service(
            NAMESPACE_DEFAULT,
            deployment_name,
            deployment_service_name,
            deployment_service_ports,
        )

        ret_deployment_ingress = t.create_ingress(
            NAMESPACE_DEFAULT,
            deployment_ingress_name,
            deployment_ingress_host,
            deployment_service_name,
            deployment_service_ports,
        )


        //fmt.Println( "ret deployment = ", ret_deployment )
        //fmt.Println( "ret deployment_service = ", ret_deployment_service )
        //fmt.Println( "ret deployment_ingress = ", ret_deployment_ingress )

        //res := types.Result_block {}
        if ret_deployment != nil &&
            ret_deployment_service != nil &&
            ret_deployment_ingress != nil {
            //res = append( res, map[string]string{"result": "true"} )
            //res["result"] = "true"

            res = map[string]interface{} {
                "ret": "true",
                "msg": map[string]interface{} {
                    "info": map[string]interface{} {},
                },
            }
        } else {
            var rets []interface{}
            rets = append( rets, ret_deployment )
            rets = append( rets, ret_deployment_service )
            rets = append( rets, ret_deployment_ingress )

            res = map[string]interface{} {
                "ret": "false",
                "msg": map[string]interface{} {
                    "info": rets,
                },
            }
        }
    } else if req_type == types.ReqType_delete {
        var rets []interface{}
        var ret_deployment interface{}
        var ret_service interface{}
        var ret_ingress interface{}

        deployment_name := req_json["deployment_name"].(string)
        //deployment_service_name := deployment_name + "-service"
        //deployment_ingress_name := deployment_name + "-ingress"
        deployment_service_name := req_json["deployment_service_name"].(string)
        deployment_ingress_name := req_json["deployment_ingress_name"].(string)

        //deployment_name := "test100-1684308153"
        //deployment_service_name := deployment_name + "-service"
        //deployment_ingress_name := deployment_name + "-ingress"
        ret_deployment, ret_service, ret_ingress = t.delete_container(
            NAMESPACE_DEFAULT,
            deployment_name,
            deployment_service_name,
            deployment_ingress_name,
        )
        rets = append( rets, ret_deployment )
        rets = append( rets, ret_service )
        rets = append( rets, ret_ingress )

        res = map[string]interface{} {
            "ret": "true",
            "msg": map[string]interface{} {
                "info": rets,
            },
        }
    } else if req_type == types.ReqType_list {
        // Container (Pod) details List

        var ret_list interface{}

        username := ""
        ret_list = t.container_details_list( username )

        res = map[string]interface{} {
            "ret": "true",
            "msg": map[string]interface{} {
                "info": ret_list,
            },
        }
    }


    //fmt.Println( res )

    return res
}



//! Implementation
// --------------------------------------------------------------------

//func (t *LocalDB) test__k8s_apis__request_container_service(rpc_args *types.Req_JSONRPC_Args_st, response *string) error {
func test__k8s_apis__request_container_service(rpc_args *types.Req_JSONRPC_Args_st, response *string) error {
    fmt.Println( "test__k8s_apis__request_container_service()" )

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

//func (t *LocalDB) k8s_apis__request_container_service(
func k8s_apis__request_container_service(
    rpc_args map[string]interface{},
    //rpc_args *types.Req_JSONRPC_k8s_request_service_Args_st,
    result *map[string]interface{},
    //result *types.Result
) error {
    fmt.Println( "k8s_apis__request_container_service()" )


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
    */



    // thread
    // call: func (t *K8s_APIS) request_container_service(...)
    //GetLocalDB().K8s.request_container_service_test()
    _result := GetLocalDB().K8s.request_container_service( rpc_args )

    //_result := types.Result { Result: "ok..." }

    *result = _result

    return nil
}

