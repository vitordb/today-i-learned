# Creating a cluster

`kind create cluster`

 ``kind create cluster --name mycluster``   //cluster custom name

# Testing deploying nginx

`kubectl create deployment nging --image=nginx`

`kubectl get pods`

`kubectl port-forward pod/nginx-7854ff8877-f2jwh 8034:80`

# Creating a cluster with multinodes

