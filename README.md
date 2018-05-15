# ec2-instance-resource

Tracks EC2 instances. All filters documented for [`ec2 describe-instances`](https://docs.aws.amazon.com/cli/latest/reference/ec2/describe-instances.html) should be supported.

## Source Configuration

* `access_key_id`: *Required.* The AWS access key to use.

* `filters`: *Optional.* The filters to apply when describing instances.

* `region`: *Required.* The AWS region.

* `secret_access_key`: *Required.* The AWS secret access key.

* `session_token`: *Optional.* The AWS STS session token to use.

### Example

Resource configuration for describing instances in the `eu-west-1c` availability zone.

``` yaml
resource_types:
- name: ec2-instance-resource
  type: docker-image
  source:
    repository: lloydg/ec2-instance-resource

resources:
- name: ec2-instances
  type: ec2-instance-resource
  source:
    access_key_id: <ACCESS_KEY_ID>
    filters:
      - name: availability-zone
        values: [eu-west-1c]
    region: eu-west-1
    secret_access_key: <SECRET_ACCESS_KEY>

jobs:
- name: instances
  plan:
    - get: ec2-instances
```

## Behavior

### `check`: Check for EC2 instances.

Checks for EC2 instances that match the supplied filters.

### `in`: Fetch EC2 instance IDs.

Writes the IDs of EC2 instances that match the supplied filters to a file named `ec2-instances`.

### `out`: Not implemented.
