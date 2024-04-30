# Kubernetes integration
This directory contains everything needed to run stos mock on microk8s

## docker-export.sh
This script is used for exporting locally built images from docker compose to internal microk8s registry

## deploy.sh
Just a quicker way to deploy the deployments included in \*.yaml files

## Accessing the system
Currently, you can access the evaluator by making a get request to [localhost/generate](http://localhost/generate?n=1)
There is also a kubernetes dashboard available either in the end of `deploy.sh` run or by running `microk8s dashboard-proxy`

## Cheat sheet while running the system
`microk8s helm upgrade logstash logstash-parent` - update pipelines/configs of logstash     
`microk8s helm uninstall logstash / install logstash logstash-parent` - reset and apply changes in logstash
