

```bash
export KUBECONFIG=../.k8s/civo-red-panda-kubeconfig.yaml
# Run K9S in a terminal
k9s --all-namespaces

# Create a namespace
kubectl create namespace demo --dry-run=client -o yaml | kubectl apply -f -

# Type the following commands in another terminal
kubectl apply -f deploy.yaml -n demo
kubectl describe ingress -n demo
# go to the webapp

# Change Environment variables
kubectl apply -f deploy.yaml -n demo
# go to the webapp

# Create the config map
kubectl apply -f config.yaml -n demo

# Update the deploy manifest
kubectl apply -f deploy.yaml -n demo

# Remove everything
kubectl delete namespace demo


```

