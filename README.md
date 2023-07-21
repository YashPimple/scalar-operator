## Scalar-Operator

Used to scale up the deployments

## Getting Started

To get started with the operaor, follow the steps below:

Clone repository from GitHub:
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

then run a `nginx` deployment on your cluster

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
