[NOT YET FINAL AND STILL UNDER PROGRESS]

# HELM CHART : Xendit-TA Service
## A. Overview

Helm chart for deploying Xendit-ta service on Kubernetes cluster. 

## B. Development
### B.1 Requirements
A. Install Helm v2.12.1: 
`$ brew install https://raw.githubusercontent.com/Homebrew/homebrew-core/7c5ab2af66f83f767db328838424b1f379bd30d4/Formula/kubernetes-helm.rb`

B. Initialize tiller: 
`$ helm init --tiller-namespace=<namespace>`
To verity, check helm version: 
`$ helm version`

C. Download chart dependencies: 
`$ helm dependency update`

### B.2 Start Application 

A. Connect to Kubenertes cluster: `$ kubectl config set-context <my_kubernetess_cluster>`

B. Run application via Makefile:
```
$ make deploy-helm
$ make deploy-helm NAMESPACE=<my_namespace>
```
