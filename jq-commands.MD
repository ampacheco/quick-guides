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



# AWS
# GCP

