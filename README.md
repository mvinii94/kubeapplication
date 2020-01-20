# kubeapp

## Setup

### Prerequisite

* [aws-cli](https://docs.aws.amazon.com/cli/latest/userguide/cli-chap-install.html)
* [kubectl](https://docs.aws.amazon.com/eks/latest/userguide/install-kubectl.html)
* [aws-iam-authenticator](https://docs.aws.amazon.com/eks/latest/userguide/install-aws-iam-authenticator.html)

#### 1. Create bucket

```bash
aws s3 mb s3://<BucketName> --region eu-west-1 --endpoint-url https://s3.eu-west-1.amazonaws.com
```

#### 2. Copy the templates to S3 Bucket

```bash
cd ~/kubeapplication

aws s3 sync ./templates s3://<BucketName>/templates/
```

#### 3. Go to the AWS console and create a CloudFormation Stack using the following template

```bash
# Under BucketName parameter, use the same bucket name that you've created
# Under KeyPair parameter, select an existent keypair (OPTIONAL)
/templates/cluster/parent.yaml
```

#### 4. While the EKS Cluster is being provisioned, create an IAM Role for CodeBuild have enough permissions to use kubectl

```bash
cd ~/kubeapplication

AWS_ACCOUNT_ID="123456789012"

TRUST="{ \"Version\": \"2012-10-17\", \"Statement\": [ { \"Effect\": \"Allow\", \"Principal\": { \"AWS\": \"arn:aws:iam::${AWS_ACCOUNT_ID}:root\" }, \"Action\": \"sts:AssumeRole\" } ] }"

echo '{ "Version": "2012-10-17", "Statement": [ { "Effect": "Allow", "Action": "eks:Describe*", "Resource": "*" } ] }' > /tmp/iam-role-policy

aws iam create-role --role-name EksWorkshopCodeBuildKubectlRole --assume-role-policy-document "${TRUST}" --output text --query 'Role.Arn'

aws iam put-role-policy --role-name EksWorkshopCodeBuildKubectlRole --policy-name eks-describe --policy-document file:///tmp/iam-role-policy
```

#### 5. After the EKS Cluster is provisioned, configure the kubectl and map the above IAM Role to the aws-auth ConfigMap for the EKS Cluster

```bash
# configure kubectl
aws eks update-kubeconfig --name KubernetesClusterDevel --alias KubernetesClusterDevel --region eu-west-1

# Update ConfigMap
ROLE="    - rolearn: arn:aws:iam::${AWS_ACCOUNT_ID}:role/EksWorkshopCodeBuildKubectlRole\n      username: build\n      groups:\n        - system:masters"

kubectl get -n kube-system configmap/aws-auth -o yaml | awk "/mapRoles: \|/{print;print \"${ROLE}\";next}1" > /tmp/aws-auth-patch.yml

kubectl patch configmap/aws-auth -n kube-system --patch "$(cat /tmp/aws-auth-patch.yml)"
```

#### 6. Also, let's enable our NodeGroup instances to join the EKS Cluster as workers

```bash
# Get the NodeInstanceRole output value from the stack kubernetesClusterDevel-EKSNodeGroup-<RandomString>
aws cloudformation describe-stacks --stack-name kubernetesClusterDevel-EKSNodeGroup-<RandomString> --region eu-west-1 --output json | jq -r '.Stacks[0].Outputs[0] | select(.OutputKey=="NodeInstanceRole") | .OutputValue'

# Download the following config map file
curl -o /tmp/aws-auth-cm.yaml https://amazon-eks.s3-us-west-2.amazonaws.com/cloudformation/2019-11-15/aws-auth-cm.yaml

# Open the file and replace the "rolearn" key value with the role arn we get from the stack output

# Apply the ConfigMap to the EKS Cluster
kubectl -f apply /tmp/aws-auth-cm.yaml

# You should be able to see the worker node being in Ready state
kubectl get nodes --watch
```

#### 7. Now we are ready to create our simple CI/CD pipeline to deploy our sample application and the EKS resources

#### Go to the AWS console and create a CloudFormation Stack using the following template

```bash
/templates/app/amazon-codepipeline.yaml
```
