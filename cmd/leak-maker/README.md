## leak-maker

### Create docker network

```
docker network create k8s-lab
```

### To start a k3d cluster

> https://k3d.io/v5.4.8/usage/registries/

1. create k3d registry
```
k3d registry create lab-registry.localhost --port 12345
```

2. create k3d cluster
```
k3d cluster create -p "8088:80@loadbalancer" --registry-use k3d-lab-registry.localhost:12345 --network k8s-lab lab
```

### To build image and push to registry

1. use `skaffold` to build `leak-maker` image:

```
skaffold build -f skaffold/leak-maker.yaml
```

2. verify the image upload is completed

### Setup manifest

1. setup `traefik` configuration

```
kubectl apply -f manifests/traefik
```

2. setup `prometheus` configuration

```
kubectl apply -f manifests/prometheus/setup

kubectl apply -f manifests/prometheus
```

3. setup `leak-maker` configuration

```
kubectl apply -f manifests/leak-maker
```

### Note

#### You can now use the registry like this (example):
1. create a new cluster that uses this registry
```
k3d cluster create --registry-use k3d-lab-registry.localhost:12345
```

1. tag an existing local image to be pushed to the registry
```
docker tag nginx:latest k3d-lab-registry.localhost:12345/mynginx:v0.1
```

1. push that image to the registry
```
docker push k3d-lab-registry.localhost:12345/mynginx:v0.1
```

1. run a pod that uses this image
```
kubectl run mynginx --image k3d-lab-registry.localhost:12345/mynginx:v0.1
```