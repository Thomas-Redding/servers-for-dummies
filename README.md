# servers-for-dummies

https://cloud.google.com/appengine/docs/standard/go/quickstart



https://cloud.google.com/appengine/docs/standard/go/googlecloudstorageclient/app-engine-cloud-storage-sample

https://cloud.google.com/appengine/docs/standard/python/googlecloudstorageclient/app-engine-cloud-storage-sample

https://cloud.google.com/appengine/docs/standard/python/tools/using-libraries-python-27

## Creating a Google Cloud Project

1. Navigate to the [Google Cloud Console](https://console.cloud.google.com) and create a "NEW PROJECT".
2. Follow the workflow.
3. Activate the default Cloud Storage buckets???

## Installing Dependencies

### Google Cloud

[Download](https://cloud.google.com/sdk/docs/) the Google Cloud SDK

### Golang

### Python

```bash
cd python
# Install Python's Google App Engine Cloud Storage Client module.
pip install -t lib -r requirements.txt
```



### Deploying

```bash
# Deploy the services.
gcloud app deploy default/app.yaml golang/app.yaml python/app.yaml
# Direct each subdomain to its service.
gcloud app deploy dispatch.yaml
```





## Useful Information

[Quickstart Instructions](https://cloud.google.com/appengine/docs/standard/go/quickstart)

```bash
gcloud components update                     // update gcloud

gcloud projects list                         // see list of projects
gcloud config set project redding-dev        // change project

gcloud app deploy app.yaml go/app.yaml       // deploy project/subdomain to internet
gcloud app browse                            // navigate to subdomain

// run on localhost
dev_appserver.py app.yaml --default_gcs_bucket_name redding-dev.appspot.com
```









