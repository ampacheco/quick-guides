## Prepare Consul Certification

- [ ] https://learn.hashicorp.com/tutorials/consul/associate-review?in=consul/certification
- [ ] https://github.com/zealvora/hashicorp-certified-consul


## Install Consul CentOS

```
{
    yum install -y yum-utils
    yum-config-manager --add-repo https://rpm.releases.hashicorp.com/RHEL/hashicorp.repo
    yum -y install consul
}
````
