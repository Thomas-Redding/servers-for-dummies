# servers-for-dummies

https://cloud.google.com/appengine/docs/standard/go/quickstart

### Deploying

```bash
gcloud app deploy default/app.yaml golang/app.yaml python/app.yaml
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









