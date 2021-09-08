## Prepare Consul Certification

- [ ] https://learn.hashicorp.com/tutorials/consul/associate-review?in=consul/certification
- [ ] https://github.com/zealvora/hashicorp-certified-consul


## Lessons

- [ ] https://docs.google.com/document/d/1ihvMYHhihZ67KJoxgIG9thLd83HQM9U7ExHdhpsPMO4/edit?usp=sharing
- [ ] https://docs.google.com/document/d/13LBrElx4-n2nNwVzj3-O-SJGJDhlMZOaTOdu0o1zqAU/edit?usp=sharing


## Install Consul CentOS

```
{
    yum install -y yum-utils
    yum-config-manager --add-repo https://rpm.releases.hashicorp.com/RHEL/hashicorp.repo
    yum -y install consul
}
````
