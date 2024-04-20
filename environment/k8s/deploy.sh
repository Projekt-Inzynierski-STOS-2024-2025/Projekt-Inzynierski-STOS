#/bin/bash
# Commands used for deployment of stos mock to k8s. Assumes microk8s is installed
#
alias mkctl="microk8s kubectl"
mkctl apply -f ./evaluator-deployment.yaml
mkctl apply -f ./evaluator-deployment.yaml
mkctl expose deployment/evaluator
microk8s dashboard-proxy
