# servers-for-dummies

## Summary

This project sets ups three services (`default`, `my-go`, and `my-py`). The services run independently from each other but share storage. By default, each service serves the subdomain of the same name and `default` services all other requests. In `dispatch.yaml`, we set it up so `my-go` and `my-py` both serve an additional subdomain.

The `default` service just returns a constant string. `my-go` and `my-py` are intentionally written to be as simple and similar as possible. For each request, both services

1. Write a string to a file.
2. Read from that file.
3. Send a response with a constant string, URL path, and the data from the file.

## Setup

### Set up the Google Cloud Project

1. Navigate to the [Google Cloud Console](https://console.cloud.google.com) and create a "NEW PROJECT".
2. Follow the workflow.
3. TODO: Activate the default Cloud Storage buckets and other stuff...

### Install Dependencies

1. [Download](https://cloud.google.com/sdk/docs/) and install the Google Cloud SDK.
2. [Download](https://golang.org/dl/) and install Go.
3. TODO: Install Go dependencies
4. Install the Google App Engine Cloud Storage Client Module

```bash
cd python
pip install -t lib -r requirements.txt
```

### Deploying

```bash
gcloud components update                     # Update gcloud.
gcloud config set project redding-dev        # Change to right project.

# Deploy the services.
gcloud app deploy default/app.yaml golang/app.yaml python/app.yaml
# Direct each subdomain to its service.
gcloud app deploy dispatch.yaml

# You can't deploy more than one service to localhost without port forwarding, so choose one:
dev_appserver.py app.yaml --default_gcs_bucket_name redding-dev.appspot.com
dev_appserver.py golang/app.yaml --default_gcs_bucket_name redding-dev.appspot.com
dev_appserver.py python/app.yaml --default_gcs_bucket_name redding-dev.appspot.coms
```

## Related Links

[Quickstart Instructions](https://cloud.google.com/appengine/docs/standard/go/quickstart)

https://cloud.google.com/appengine/docs/standard/go/quickstart

https://cloud.google.com/appengine/docs/standard/go/googlecloudstorageclient/app-engine-cloud-storage-sample

https://cloud.google.com/appengine/docs/standard/python/googlecloudstorageclient/app-engine-cloud-storage-sample

https://cloud.google.com/appengine/docs/standard/python/tools/using-libraries-python-27

