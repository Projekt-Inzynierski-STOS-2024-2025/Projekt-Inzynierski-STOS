#/bin/bash
# Script assumes that the application was built using docker compose build
docker save environment-evaluator > evaluator.tar
docker save environment-worker > worker.tar

microk8s ctr image import worker.tar
microk8s ctr image import evaluator.tar

