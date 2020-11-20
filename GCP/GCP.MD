# Google Cloud API for Anthos Mesh


[GCP Lab](https://www.qwiklabs.com/focuses/8459?locale=es&parent=catalog) | ¿Cómo instalar Anthos Service Mesh en Google Kubernetes Engine?

```BASH
gcloud services enable \
    container.googleapis.com \
    compute.googleapis.com \
    monitoring.googleapis.com \
    logging.googleapis.com \
    cloudtrace.googleapis.com \
    meshca.googleapis.com \
    meshtelemetry.googleapis.com \
    meshconfig.googleapis.com \
    iamcredentials.googleapis.com \
    anthos.googleapis.com \
    gkeconnect.googleapis.com \
    gkehub.googleapis.com \
    cloudresourcemanager.googleapis.com    
```

```BASH
kubectl auth can-i '*' '*' --all-namespaces
```
yes

