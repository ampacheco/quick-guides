# Azure 


## Filter by Default Subscription and Get subscription ID and Tennant ID
```
az account list | jq -r '.[] | select(.isDefault == true)| .id, .tenantId'
```
# AWS
# GCP
