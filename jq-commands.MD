# Azure 

## List Azure Locations
```
az account list-locations -o table
```

## Filter by Default Subscription and Get subscription ID and Tennant ID
```
az account list | jq -r '.[] | select(.isDefault == true)| .id, .tenantId'
```

## Clean all the resources or the ones you want (jq filters) in an azure subscription
```
az group list | jq ".[].name" | xargs -I % sh -c 'az group delete -y --no-wait -n %'
```

# Active Account
````sh
az account list | jq -r ".[] | select(.isDefault == true) | {name: .name, tenantId: .tenantId, id: .id} "
````

# Change Active Account
````
az account set -s [name or Id]
````

# Create a Cluster
````
az aks create --resource-group mesh-cluster-rg --name mesh-cluster --node-count 2
````

# AWS
# GCP

