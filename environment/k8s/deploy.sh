#/bin/bash
# Commands used for deployment of stos mock to k8s. Assumes microk8s is installed
#
alias mkctl="microk8s kubectl"
# Install necessary addons
microk8s enable dns
microk8s enable dashboard
microk8s enable registry
microk8s enable ingress
# Apply deployments
microk8s kubectl apply -f ./evaluator-deployment.yaml
microk8s kubectl apply -f ./worker-deployment.yaml
# Networking
microk8s kubectl expose deployment/evaluator
microk8s kubectl apply -f ./ingress-controller.yaml
# Start dashboard
microk8s dashboard-proxy
