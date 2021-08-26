# GCP DevOps

## Create a Repo
````
gcloud source repos create REPO_DEMO
````

## Clone a Repo
````
gcloud source repos clone REPO_DEMO
````

# Deploy with Kubernetes

## Get Sample Code

````
gsutil -m cp -r gs://spls/gsp053/orchestrate-with-kubernetes .
cd orchestrate-with-kubernetes/kubernetes
````

## Create Cluster
````
gcloud container clusters create bootcamp --num-nodes 5 --scopes "https://www.googleapis.com/auth/projecthosting,storage-rw"
````

## Canary Deployment

![image](https://user-images.githubusercontent.com/877636/130889019-122399cf-464e-46f4-bbd9-9a69e06353ef.png)


## Blue-Green Deployment 

![image](https://user-images.githubusercontent.com/877636/130888973-26e71d37-9d27-461e-97d1-c8b884894007.png)



### Referencias: 

- [ ] https://google.qwiklabs.com/games/1879/labs/10085
- [ ] https://cloud.google.com/certification/
- [ ] https://google.qwiklabs.com/games/1879/labs/10086



