## Prepare Consul Certification

- [ ] https://learn.hashicorp.com/tutorials/consul/associate-review?in=consul/certification
- [ ] https://github.com/zealvora/hashicorp-certified-consul


## Lessons

- [ ] [Getting Starting with Consul](https://docs.google.com/document/d/1ihvMYHhihZ67KJoxgIG9thLd83HQM9U7ExHdhpsPMO4/edit?usp=sharing)
- [ ] [Service Discovery](https://docs.google.com/document/d/13LBrElx4-n2nNwVzj3-O-SJGJDhlMZOaTOdu0o1zqAU/edit?usp=sharing)
- [ ] [Dynamic Application Configuration](https://docs.google.com/document/d/160IftpUJCNIOXFnpXwGvOFrIYmKsebkO9SQuaIvZ1JY/edit?usp=sharing)
- [ ] [Security](https://docs.google.com/document/d/1_1I2XXD_0cB0yHa3S5RXpH-FPV0O2N26QCOF6Vy1VUE/edit?usp=sharing)
- [ ] [Infrastructure & High Availability](https://docs.google.com/document/d/1aieGjl9fLDprfMrX2IbK4hxmgPGAzC5bYItwz8J5ohA/edit?usp=sharing)
- [ ] [Consul Enterprise](https://docs.google.com/document/d/1eHN3ls7hVAyjUO-1ggUWMRe833IpDFXF8hGvxuwX0GI/edit)


## Preparation

📖 [Link to Notes](https://docs.google.com/document/d/1m98VuzrK-C5kTZWTbYhNrCDbWl4pSWROYCMxjd9yJV8/edit?usp=sharing)


## Install Consul CentOS

```
{
    yum install -y yum-utils
    yum-config-manager --add-repo https://rpm.releases.hashicorp.com/RHEL/hashicorp.repo
    yum -y install consul
}
````
