#/bin/bash
# Commands used for deployment of stos mock to k8s. Assumes microk8s is installed
#
alias mkctl="microk8s kubectl"
# Install necessary addons
microk8s enable dns
microk8s enable dashboard
microk8s enable registry
microk8s enable ingress
microk8s enable helm
# Install additional repos via helm - logstash
microk8s helm repo add elastic https://helm.elastic.co
microk8s helm dep build logstash-parent/
microk8s helm install logstash logstash-parent/

# Initialize elasticsearch deployment
microk8s kubectl create -f https://download.elastic.co/downloads/eck/2.12.1/crds.yaml
microk8s kubectl apply -f https://download.elastic.co/downloads/eck/2.12.1/operator.yaml
microk8s kubectl apply -f ./elasticsearch-config.yaml

#Initialize kibana deployment
microk8s kubectl apply -f ./kibana-config.yaml
# Apply deployments
microk8s kubectl apply -f ./evaluator-deployment.yaml
microk8s kubectl apply -f ./worker-deployment.yaml
# Networking
microk8s kubectl expose deployment/evaluator
microk8s kubectl apply -f ./ingress-controller.yaml
# Start dashboard
microk8s dashboard-proxy
