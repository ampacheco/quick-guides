# Cloud 9 Development Env Configuration Recommendations

- Change AWS Temporary Managed Credentials to IAM Role

- [Resize the Disk](https://docs.aws.amazon.com/cloud9/latest/user-guide/move-environment.html)

- Update AWS CLI and Install jq

```
sudo pip install --upgrade awscli && hash -r
sudo yum -y install jq gettext bash-completion moreutils
echo 'yq() {
  docker run --rm -i -v "${PWD}":/workdir mikefarah/yq yq "$@"
}' | tee -a ~/.bashrc && source ~/.bashrc

export ACCOUNT_ID=$(aws sts get-caller-identity --output text --query Account)
export AWS_REGION=$(curl -s 169.254.169.254/latest/dynamic/instance-identity/document | jq -r '.region')

test -n "$AWS_REGION" && echo AWS_REGION is "$AWS_REGION" || echo AWS_REGION is not set

echo "export ACCOUNT_ID=${ACCOUNT_ID}" | tee -a ~/.bash_profile
echo "export AWS_REGION=${AWS_REGION}" | tee -a ~/.bash_profile
aws configure set default.region ${AWS_REGION}
aws configure get default.region

```

- Validate IAM Role
```
aws sts get-caller-identity --query Arn | grep eksworkshop-admin -q && echo "IAM role valid" || echo "IAM role NOT valid"
```

- Setting Ammplify Environment

```
npm install -g @aws-amplify/cli@4.13.2
```
