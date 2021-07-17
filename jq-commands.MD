# JSON TOOLS
## jq Versus JMESPath

1. [jq Manual](https://stedolan.github.io/jq/manual/v1.5/#Builtinoperatorsandfunctions)
1. [jq queries](https://michaelheap.com/extract-keys-using-jq/)
1. [Extract Keys](https://www.fabian-keller.de/blog/5-useful-jq-commands-parse-json-cli/)
1. [JMESPath.org](https://jmespath.org/examples.html)

## Azure 

### List Azure Locations
```
az account list-locations -o table
```

### Filter by Default Subscription and Get subscription ID and Tennant ID
```
az account list | jq -r '.[] | select(.isDefault == true)| .id, .tenantId'
```

### Clean all the resources or the ones you want (jq filters) in an azure subscription
```
az group list | jq ".[].name" | xargs -I % sh -c 'az group delete -y --no-wait -n %'
```

### Active Account
````sh
az account list | jq -r ".[] | select(.isDefault == true) | {name: .name, tenantId: .tenantId, id: .id} "
````

### Change Active Account
````
az account set -s [name or Id]
````

### Create a Cluster
````
az aks create --resource-group mesh-cluster-rg --name mesh-cluster --node-count 2
````

# AWS
# GCP

