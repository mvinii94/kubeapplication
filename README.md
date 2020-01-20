# kubeapp

### Prerequisites

#### 1. Create an IAM Role "KubeAdm" with the following

##### Trust relantionship

```json
{
    "Version": "2012-10-17",
    "Statement": [
        {
            "Effect": "Allow",
            "Principal": {
                "Service": "cloudformation.amazonaws.com"
            },
            "Action": "sts:AssumeRole"
        },
        {
            "Sid": "allowAdminUser",
            "Effect": "Allow",
            "Principal": {
                "AWS": "<YourIAMUserOrRoleArn>"
            },
            "Action": "sts:AssumeRole"
        }
    ]
}
```

##### Policy

```json
{
    "Version": "2012-10-17",
    "Statement": [
        {
            "Effect": "Allow",
            "Action": "*",
            "Resource": "*"
        }
    ]
}
```

#### Create am IAM Role "EksWorkshopCodeBuildKubectlRole" with the following

##### Trust relantionship

```json
{
    "Version": "2012-10-17",
    "Statement": [
        {
            "Effect": "Allow",
            "Principal": {
                "AWS": "arn:aws:iam::<AWS_ACCOUNT_ID>:root"
            },
            "Action": "sts:AssumeRole"
        }
    ]
}
```

##### Policy

```json
{
    "Version": "2012-10-17",
    "Statement": [
        {
            "Effect": "Allow",
            "Action": "eks:Describe*",
            "Resource": "*"
        }
    ]
}
```
