`multipass launch -n k8s -c 2 -m 4G -d 20G`

# Instalar o micro-k8s:

`multipass exec k8s -- sudo snap install microk8s --classic --channel=1.18/stable`

# Incluir o usuario ubuntu para usar o microk8s

`multipass exec k8s -- sudo usermod -a -G microk8s ubuntu`

# Dar permissao, colocar como proprietario do arquivo

`multipass exec k8s -- sudo chown -f -R ubuntu ~/.kube`

# Testando

Fazendo deploy do nginx:

`multipass exec k8s -- /snap/bin/microk8s.kubectl create deployment nginx --image=nginx`

`multipass exec k8s -- /snap/bin/microk8s.kubectl get pods`

# Acessando kubernetes remotamente da maquina virtual

`multipass exec k8s -- /snap/bin/microk8s.kubectl config view --raw`

# Copia as configuracoes de acsso do kubectl, e edita o config local. Depois se executa o comando local:

`kubectl get pods`

`kubectl get nodes`

`multipass restart k8s`






