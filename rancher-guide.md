# Creating Resource Group and Ubuntu VM
```
{ 
  echo ${AZ_GROUP_NAME:=rancher};echo ${AZ_REGION:=eastus} 
  az group create -n ${AZ_GROUP_NAME} -l ${AZ_REGION}
  az vm create -g ${AZ_GROUP_NAME} -n ${AZ_GROUP_NAME} --image ubuntults --data-disk-sizes-gb 20 --size Standard_DS2_v2 --ssh-key-value @~/.ssh/id_rsa_ubuntu.pub --admin-username rancher
}
```

# Install Docker 
```

```


# Install Certbot
```
```




# Request Certificate
```
```

# Deploy & Run rnacher on Docker
```
```

<script src="https://gist.github.com/nisrulz/11c0d63428b108f10c83.js"></script>
