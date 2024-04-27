Kubernetes Cluster is composed of:

Kubernetes Control Plane: responsavel por orquestrar o cluster kubernetes. Boa pratica ter mais de um para ter resiliencia. Dentro dele existe:

* Kube Api Server - atraves dele que ocorre a comunicação com o cluster kubernetes.
* ETCD - banco chave/valor, onde armazeno os dados referentes ao cluster. Nao é acessado diretamente, é acessado pelo Kube Api Server
* Kube Scheduler - a especificação do que eu quero executar no kubernetes é analisada, para decidir em qual nó vai ser executada, se ele vai ser executado, se o kubernetes tem capacidade de executar, tudo isso é gerenciado pelo kube scheduler
* Kube Controller Manager - Responsavel por gerenciar todos os controladores do cluster kubernetes - como replica set, admission controller.


Kubernetes Nodes: responsavel por executar os containers. Composto por:

* Kubelet - Responsável por comunicar com o Kubernetes Control Plane, utilizando o Kube Api Server, e verificado o estado do nó
* Kube Proxy - Responsável por comunicação entre os nós
* Container Runtime - Responsável por executar os containers, como o docker.

