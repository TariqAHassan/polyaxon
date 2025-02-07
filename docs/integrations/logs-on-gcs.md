---
title: "Logs on GCS"
meta_title: "Google GCS"
meta_description: "Using Google Cloud Storage GCS buckets to organize your jobs, builds, and experiment logs. Polyaxon allows users to manage all logs generated by jobs, builds, and experiments containers in Google Cloud Storage GCS."
custom_excerpt: "Google Cloud Storage is a RESTful online file storage web service for storing and accessing data on Google Cloud Platform infrastructure. The service combines the performance and scalability of Google's cloud with advanced security and sharing capabilities."
image: "../../content/images/integrations/gcs.png"
author:
  name: "Polyaxon"
  slug: "Polyaxon"
  website: "https://polyaxon.com"
  twitter: "polyaxonAI"
  github: "polyaxon"
tags: 
  - logging
  - storage
featured: false
visibility: public
status: published
---

Polyaxon allows users to manage all logs generated by jobs, builds, and experiments containers in Google Cloud Storage (GCS).

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

## Use the secret name and secret key in your logs persistence definition

```yaml
persistence:
  logs:
    store: gcs
    bucket: gs://[BUCKET-NAME]
    secret: [SECRET-NAME]
    secretKey: [SECRET-KEY]
```

e.g.

```yaml
persistence:
  logs:
    store: gcs
    bucket: gs://logs-bucket
    secret: gcs-secret
    secretKey: gcs-secret.json
```
