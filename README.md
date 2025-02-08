# k8s-go-user
simple restful api written in go for testing it on k8s

### Build Image From Dockerfile(alptht is dockerhub username)
docker build --tag k8s-go-user .
docker build -t alptht/k8s-go-user:multistage -f Dockerfile.multistage .

### Check Docker Images
docker image ls

### Tag Images
docker image tag k8s-go-user:latest k8s-go-user:v1.0

### Remove Tag
docker image rm k8s-go-user:v1.0

### Run Image inside Container(-detached --name and --published on ports)
docker run -d -p 8081:8081 --name k8s-go-user k8s-go-user:multistage

### List Containers
docker ps -a

### Stop Container
docker stop k8s-go-user

### Remove all non working container.
docker container prune

### Restart Container
docker restart k8s-go-user

### Login to Docker
docker login

### Push to DockerHub(alptht is dockerhub username)
docker push alptht/k8s-go-user:multistage

---

###### K8s IMPERATIVE

### Create New Deployment
kubectl create deployment k8s-go-user-deployment --image=alptht/k8s-go-user:multistage

### Show deployments
kubectl get deployment

### Edit deployment
kubectl edit deployment nginx-depl

### Show deployment as yaml
kubectl get deployment deplname -o yaml

### Create Service
kubectl expose deployment k8s-go-user-deployment --port=8081
kubectl expose deployment k8s-go-user-deployment --type=NodePort --port=8081
kubectl expose deployment k8s-go-user-deployment --type=LoadBalancer --port=8081

### List Services
kubectl get services

### Get Replicaset
kubectl get replicaset

### Delete Service
kubectl delete svc k8s-go-user-deployment

### Scale Deployment
kubectl scale deployment k8s-go-user-deployment --replicas=4

### Get Pods
kubectl get pods -o wide
kubectl get pods --watch

### Get Logs
kubectl logs podName

### Check Inside Pod
kubectl exec -it podname -- bin/bash 

### Minikube problem fix for sending request
minikube service k8s-go-user

### Describe deployment
kubectl describe deployment 

### Create a namespace
kubectl create namespace my-namespace

### Delete everything(in Minikube kubernetes service will also be deleted, but automaticly recreated)
kubectl delete all --all

---

##### K8s DECLERATIVE

### Apply deployment and service
kubectl apply -f deployment.yaml
kubectl apply -f service.yaml

### Delete deployment and service
kubectl delete -f deployment.yaml
kubectl delete -f service.yaml

### Base64 through shell
echo -n 'username' | base64
echo -n 'password' | base64