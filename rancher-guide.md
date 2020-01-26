# Creating Resource Group and Ubuntu VM

```
$ echo ${AZ_GROUP_NAME:=rancher};echo ${AZ_REGION:=eastus} 
$ az vm
$ az vm create -g $AZ_GROUP -n rancher --image ubuntults --data-disk-sizes-gb 20 --size Standard_DS2_v2 --ssh-key-values @~/.ssh/id_rsa_ubuntu.pub --admin-username rancher
```
