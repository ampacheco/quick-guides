# Linkerd Creation

```sh
az aks create --resource-group mesh-cluster-rg --name mesh-cluster --node-count 2
az aks list
az aks get-credentials --resource-group mesh-cluster-rg --name mesh-cluster

# kubctl aliasing and configuration
{
    alias k=kubectl
    k completion && k completion bash && . <(k completion bash)
}

# Check Cluster
k get nodes

# Install Linkerd and it to the path and configure it for autocomplete
curl -sL https://run.linkerd.io/install | sh
export PATH=$PATH:/home/alejandro/.linkerd2/bin
. <(linkerd completion bash)

# Check command is in path and working well
linkerd

# check if cluster if ready for linkerd
linkerd check --pre

# Preview Linkerd yaml output
linkerd install

# Deploy Linkerd
linkerd install | k apply -f -

# Check everything run as expected
k get ns
k -n linkerd get all
linkerd check

# expose dashboard via Ingress
cat << EOF >> dashboard.yaml
apiVersion: v1
kind: Secret
type: Opaque
metadata:
  name: web-ingress-auth
  namespace: linkerd
data:
  auth: YWRtaW46JGFwcjEkbjdDdTZnSGwkRTQ3b2dmN0NPOE5SWWpFakJPa1dNLgoK
---
apiVersion: extensions/v1beta1
kind: Ingress
metadata:
  name: web-ingress
  namespace: linkerd
  annotations:
    kubernetes.io/ingress.class: 'nginx'
    nginx.ingress.kubernetes.io/upstream-vhost: $service_name.$namespace.svc.cluster.local:8084
    nginx.ingress.kubernetes.io/configuration-snippet: |
      proxy_set_header Origin "";
      proxy_hide_header l5d-remote-ip;
      proxy_hide_header l5d-server-id;
    nginx.ingress.kubernetes.io/auth-type: basic
    nginx.ingress.kubernetes.io/auth-secret: web-ingress-auth
    nginx.ingress.kubernetes.io/auth-realm: 'Authentication Required'
spec:
  rules:
    - host: dashboard.example.com
      http:
        paths:
          - backend:
              serviceName: linkerd-web
              servicePort: 8084

EOF

# deploy ingress rule
k apply -f dashboard.yaml

# Teardown your env/
az group delete -n mesh-cluster-rg -y --no-wait
```
