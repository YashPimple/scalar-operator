# Scalar-Operator

Scalar-Operator is a Kubernetes Operator designed to simplify scaling up deployments effortlessly.

## Getting Started

To get started with the Scalar-Operator, follow the steps below:

1. Clone the repository from GitHub:
   ```bash
   git clone https://github.com/YashPimple/scalar-operator.git
   cd scalar-operator

2. Generate the Kubernetes manifests:
   ```bash
   make manifests

3. Apply the Custom Resource Definition (CRD):
   ```bash
   kubectl apply -f config/crd/bases/api.scalaroperator.io_scalars.yaml

4. Deploy an nginx deployment on your cluster (or any other deployment you wish to scale):
   ```bash
   kubectl create deployment nginx --image=nginx

5. Check the number of replicas running:
   ```bash
   kubectl get pods

6. Scale up the replicas using the Scalar-Operator:
   ```bash
   make run
   
7. Open another terminal and apply the Scalar custom resource to scale up the replicas:
   ```bash
   kubectl apply -f config/samples/api_v4alpha1_scaler.yaml
   
8. Finally, check the number of replicas:
   ```bash
   kubectl get pods

9. Output
   You should see an output similar to the following:
   ```bash
   yashpimple@Yashs-Air ~ % kubectl get pods
   NAME                    READY   STATUS              RESTARTS   AGE
   nginx-8f458dc5b-ls5wk   0/1     ContainerCreating   0          2s
   nginx-8f458dc5b-nlbg7   0/1     ContainerCreating   0          2s
   nginx-8f458dc5b-rk588   0/1     ContainerCreating   0          2s
   nginx-8f458dc5b-tz6jl   0/1     ContainerCreating   0          2s
   nginx-8f458dc5b-wz97m   1/1     Running             0          29s

