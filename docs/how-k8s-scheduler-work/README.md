## How does pod scheduling work

This project uses the `kwok` tool to simulate a kubernetes cluster has multiple nodes.

> `kwok` will be in charge of all nodes with the annotation `kwok.x-k8s.io/node=fake`. If they carry an accurate `.spec.nodeName` field, kwok will ensure they stay in the Running state.


## Create a Cluster

```
kwokctl create cluster --name=demo
```


## Add nodes

```
kubectl apply 
```

## Delete a Cluster 

```
kwokctl delete cluster --name=demo
```