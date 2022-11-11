# K8S AWS Construct Library

Provides Kubernetes resources for integrating with AWS services.

## Fargate Logging

Fargate logging causes the output of pods running on EKS Farget to be sent to a logging service for storage and review.

By default, logs are written to CloudWatch Logs.

Enable Fargate logging on an EKS cluster:

```
declare const cluster: eks.FargateCluster;

const logger = new k8s_aws.FargateLogger(this, 'logger', {
    cluster: cluster,
    fargateProfiles: [
        cluster.defaultProfile
    ]
});
```

Permissions for sending logs to their configured destination are added to the Fargate profiles associated with the logger.

When adding new Fargate Profiles be sure to associate them with the logger to ensure they have sufficient permissions to write logs.

```
declare const profile: eks.FargateProfile;
declare const logger: k8s_aws.FargateLogger;

logger.addFargateProfile(profile);
```

Configure logging to write to a Kinesis Firehose delivery stream:

```
declare const cluster: eks.FargateCluster;
declare const deliveryStream: kinesis_hirehose.DeliveryStream;

const logger = new k8s_aws.FargateLogger(this, 'logger', {
    cluster: cluster,
    fargateProfiles: [
        cluster.defaultProfile
    ],
    outputs: [
        k8s_aws.FluentBitOutput.kinesisFirehose(k8s_aws.FluentBitMatch.ALL, deliveryStream);
    ]
});
```

Configure logging to write to a Kinesis data stream:

```
declare const cluster: eks.FargateCluster;
declare const stream: kinesis.Stream;

const logger = new k8s_aws.FargateLogger(this, 'logger', {
    cluster: cluster,
    fargateProfiles: [
        cluster.defaultProfile
    ],
    outputs: [
        k8s_aws.FluentBitOutput.kinesis(k8s_aws.FluentBitMatch.ALL, stream);
    ]
});
```

Configure logging to write to an OpenSearch domain:

```
declare const cluster: eks.FargateCluster;
declare const domain: opensearch.Domain;

const logger = new k8s_aws.FargateLogger(this, 'logger', {
    cluster: cluster,
    fargateProfiles: [
        cluster.defaultProfile
    ],
    outputs: [
        k8s_aws.FluentBitOutput.opensearch(k8s_aws.FluentBitMatch.ALL, domain);
    ]
});
```

Filter out log messages matching the AWS load balancer health check user agent:

```
declare const logger: k8s_aws.FargateLogger;

logger.addFilter(k8s_aws.FluentBitFilter.grep(k8s_aws.FluentBitMatch.ALL, {
    exclude: true,
    key: 'log',
    regex: 'ELB-HealthChecker'
}));
```

## Container Insights

[AWS Container Insights](https://docs.aws.amazon.com/AmazonCloudWatch/latest/monitoring/ContainerInsights.html) provides advanced diagnostic and performance metrics for your containerized applications running on AWS. For EKS cluster, Container Insights is provided by using [AWS Distro for OpenTelemetry](https://aws.amazon.com/blogs/containers/introducing-amazon-cloudwatch-container-insights-for-amazon-eks-fargate-using-aws-distro-for-opentelemetry/).

To enable Container Insights for pods running on your EKS cluster:

```
declare const cluster: eks.Cluster;

const collector = new k8s_aws.AdotCollector(this, 'adot-collector', {
    cluster: cluster
});
```

## Route 53

Enable management of Route 53 hosted zones for ingress and service hosts:

```
declare const cluster: eks.Cluster;

const manager = new k8s_aws.Route53Dns(this, 'route53-dns', {
    cluster: cluster
});
```

Only enable managment of Route 53 DNS to only records that end with `example.com`:

```
declare const manager: k8s_aws.Route53Dns;

manager.addDomainFilter('example.com');
```

Only allow management for hosted zones that are tagged with `managed-dns=enabled`:

```
declare const manager: k8s_aws.Route53Dns;

manager.addZoneTag({
    key: 'managed-dns',
    value: 'enabled'
});
```

Only allow creates and updates of DNS records and not deletes:

```
declare const cluster: eks.Cluster;

const manager = new k8s_aws.Route53Dns(this, 'route53-dns', {
    cluster: cluster,
    syncPolicy: ExternalDnsSyncPolicy.UPSERT_ONLY
});
```

## Secrets Manager

Enable synchronization of specific secret between Secrets Manager and Kubernetes:

```
declare const cluster: eks.Cluster;

const operator = new k8s_aws.ExternalSecretsOperator(this, 'external-secrets', {
    cluster: cluster
});
```

To tell the external secrets operator to synchronise a secret:

```
declase const operator: k8s_aws.ExternalSecretsOperator;
declare const secret: secretsmanager.Secret;

operator.registerSecretsManagerSecret('sychronized-secret', secret);
```

Give the secret a human friendly name in Kubernetes:

```
declase const operator: k8s_aws.ExternalSecretsOperator;
declare const secret: secretsmanager.Secret;

operator.registerSecretsManagerSecret('sychronized-secret', secret, {
    name: 'database-secret'
});
```

Only import specific JSON keys from a secret:

```
declase const operator: k8s_aws.ExternalSecretsOperator;
declare const secret: secretsmanager.Secret;

operator.registerSecretsManagerSecret('sychronized-secret', secret, {
    fields: [
        { kubernetesKey: 'username' },
        { kubernetesKey: 'password' },
    ]
});
```

Map secret fields that need to be different between Secrets Manager and Kubernetes.

```
declase const operator: k8s_aws.ExternalSecretsOperator;
declare const secret: secretsmanager.Secret;

operator.registerSecretsManagerSecret('sychronized-secret', secret, {
    fields: [
        {
            kubernetesKey: 'user',
            remoteKey: 'username',
        },
        {
            kubernetesKey: 'pass',
            remoteKey: 'password'
        },
    ]
});
```

## Echoserver

A basic Kubernetes test service that can be used for testing Kubernetes cluster integrations.

This is a simple HTTP service that listens for incoming requests and echo details of requests back to the user.

Log messages are produced for each request and provide a convenient way to test logging filter and output configurations.

To create an echoserver service:

```
declare const cluster: eks.Cluster;

const echoserver = new k8s_aws.Echoserver(this, 'echoserver', {
    cluster: cluster
});
```
