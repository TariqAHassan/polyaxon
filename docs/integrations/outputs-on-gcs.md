---
title: "Outputs on GCS"
meta_title: "Google GCS"
meta_description: "Using Google Cloud Storage GCS buckets to organize your jobs outputs and experiment artifacts. Polyaxon allows users to connect to one or multiple buckets on Google Cloud Storage GCS to store job outputs and experiment artifacts."
custom_excerpt: "Google Cloud Storage is a RESTful online file storage web service for storing and accessing data on Google Cloud Platform infrastructure. The service combines the performance and scalability of Google's cloud with advanced security and sharing capabilities."
image: "../../content/images/integrations/gcs.png"
author:
  name: "Polyaxon"
  slug: "Polyaxon"
  website: "https://polyaxon.com"
  twitter: "polyaxonAI"
  github: "polyaxon"
tags: 
  - artifacts-store
  - storage
featured: false
visibility: public
status: published
---

Polyaxon allows users to connect to one or multiple buckets on Google Cloud Storage (GCS) to store job outputs and experiment artifacts.

## Create an Google cloud storage bucket

You should create a google cloud storage bucket (e.g. plx-data), and ensure the account that 
will be interacting with this bucket has the required read and write permissions. A key
for this account should be created and downloaded in JSON format. 

If you would like to create a dedicated service account, ensure that it is:

  * assigned a role that grants it read and write from google cloud storage
  * that the permissions on the bucket itself grant read and write to the service account.
    (Note that inherited permissions may not suffice and that a role with read and 
    write access may have to be assigned to the bucket manually.)

Note: This is just a tutorial. Please consult with your devops team for
information on acceptable data permissions protocols at your organization
prior to deployment.

## Create a secret on Kubernetes

You should then create a secret with this access keys information on Kubernetes on the same namespace as Polyaxon deployment:

`kubectl create secret generic gcs-secret --from-file=gcs-secret.json=path/to/gcs-key.json -n polyaxon`

## Use the secret name and secret key in your outputs persistence definition

```yaml
persistence:
  outputs:
    [OUTPUTS-NAME-TO-USE]:
      store: gcs
      bucket: gs://[BACKET-NAME]
      secret: [SECRET-NAME]
      secretKey: [SECRET-KEY]
```

e.g.

```yaml
persistence:
  outputs:
    outputs:
      store: gcs
      bucket: gs://outputs-bucket
      secret: gcs-secret
      secretKey: gcs-key.json
```

## Update/Install Polyaxon deployment

You can now [install](/setup/kubernetes/)/[upgrade](/setup/kubernetes/#upgrade-polyaxon) Polyaxon with access the outputs on GCS.

## Storing outputs and artifacts in your experiments/jobs

You can use [polyaxon-client](/references/polyaxon-client-python/) to access the outputs in your jobs/experiments.

Polyaxon client does not bundle by default the google cloud storage requirements to keep the client lightweight:

```bash
pip install polyaxon-client[gcs]
``` 

or to have more control over the version of GCS storage:

```bash
pip install polyaxon-client
pip install google-cloud-storage
``` 

In your experiment/job definition, you can add this step to be available during the run:

```yaml
build:
  ...
  build_steps:
    ...
    - pip3 install polyaxon-client[gcs]
```

## Using the outputs store from tracking

In your experiment/job, Polyaxon exposes the secret related to the outputs as well as the outputs [path](/references/polyaxon-tracking-api/in-cluster/#get-outputs-path) scheduled for the run as an an env var,  
and provides an interface to get an authenticated client for each one of these Paths.

```python
from polyaxon_client.tracking import Experiment

experiment = Experiment()
...
experiment.outputs_store.upload_file(file_path)
experiment.outputs_store.upload_dir(dir_path)
``` 

## Tensorflow

If you are using Tensorflow, you won't need to do any further configuration since Tensorflow can natively use GCS, 
Polyaxon will automatically set the required environment variables so that Tensorflow can use the bucket. 
