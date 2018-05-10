# etcdstorage-apiserver

Implements: https://groups.google.com/d/msg/kubernetes-sig-api-machinery/rHEoQ8cgYwk/iglsNeBwCgAJ

It provides an API for handling `etcd` endpoints.

## Purpose

This API server provides the implementation for exposing namespace-proxied `etcd` server as described in implementation.

## Compatibility

HEAD of this repo will match HEAD of k8s.io/apiserver, k8s.io/apimachinery, and k8s.io/client-go.

## Running `etcdstorage-apiserver`

Steps for running the API server is similar to steps for running [`sample-apiserver`](https://github.com/kubernetes/sample-apiserver#running-it-stand-alone), so check it out for running API server.

Once API server is running, you can use the following command to list all `etcd` servers:

```
http --verify=no --cert client.crt --cert-key client.key \
        https://localhost:8443/apis/etcdstorage.k8s.io/v1alpha1/namespaces/default/etcdstorages
```

To create an `etcdstorage`, you can use the following command:

```
http --verify no -j --cert-key client.key --cert client.crt https://localhost:8443/apis/etcdstorage.k8s.io/v1alpha1/namespaces/default/flunders < <(python -c 'import sys, yaml, json; json.dump(yaml.load(sys.stdin), sys.stdout, indent=4)' < artifacts/etcdstorages/01-etcdstorage.yml)
```
