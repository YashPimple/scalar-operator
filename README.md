## Scalar-Operator

Used to scale up the deployments

## Getting Started

To get started with the operaor, follow the steps below:

clone repository from GitHub:
```
git clone https://github.com/YashPimple/scalar-operator.git
```

```
cd scalar-operator
```
   
```
make manifests
```

```
kubectl apply -f config/crd/bases/api.scalaroperator.io_scalars.yaml
```

then run an `nginx` deployment on your cluster

```
kubectl create deployment nginx --image=nginx
```

Run the following command to check for the number of replicas running right now

```
kubectl get pods
```

Further run the following commands to scale up the replicas

```
make run
```

Open another terminal and paste the following CMD

```
kubectl apply -f config/samples/api_v4alpha1_scaler.yaml
```

Finally you can now check the number of replicas using the following command 
```
kubectl get pods
```

## Output:

```
 yashpimple@Yashs-Air ~ % kubectl get pods
NAME                    READY   STATUS              RESTARTS   AGE
nginx-8f458dc5b-ls5wk   0/1     ContainerCreating   0          2s
nginx-8f458dc5b-nlbg7   0/1     ContainerCreating   0          2s
nginx-8f458dc5b-rk588   0/1     ContainerCreating   0          2s
nginx-8f458dc5b-tz6jl   0/1     ContainerCreating   0          2s
nginx-8f458dc5b-wz97m   1/1     Running             0          29s
```
