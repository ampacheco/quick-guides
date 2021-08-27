# Creating Alerts

## Create Init
````
cat <<EOF > cloud-init.txt
#cloud-config
package_upgrade: true
packages:
- stress
runcmd:
- sudo stress --cpu 1
EOF
````



## Create VM
````
az vm create \
    --resource-group learn-cae37754-7eec-4db1-95cc-8b4cb5636c91 \
    --name vm1 \
    --location eastUS \
    --image UbuntuLTS \
    --custom-data cloud-init.txt \
    --generate-ssh-keys
````

# Creating Alerts using CLI :Emoiji:

## Get VMID
````
VMID=$(az vm show \
        --resource-group learn-cae37754-7eec-4db1-95cc-8b4cb5636c91 \
        --name vm1 \
        --query id \
        --output tsv)
````


## Creating the Alert
````
az monitor metrics alert create \
    -n "Cpu80PercentAlert" \
    --resource-group learn-cae37754-7eec-4db1-95cc-8b4cb5636c91 \
    --scopes $VMID \
    --condition "max percentage CPU > 80" \
    --description "Virtual machine is running at or greater than 80% CPU utilization" \
    --evaluation-frequency 1m \
    --window-size 1m \
    --severity 3
````
