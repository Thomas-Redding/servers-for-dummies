import os
import lib.cloudstorage
from google.appengine.api import app_identity
import webapp2

class MainPage(webapp2.RequestHandler):
  def get(self):
    bucketName = os.environ.get('BUCKET_NAME', app_identity.get_default_gcs_bucket_name())

    self.response.headers['Content-Type'] = 'text/plain'
    self.response.write('Hello World (Python)!')

  def write(self, bucketName, filePath, data):
    filename = '/' + bucketName + '/' + filePath
    writeRetryParams = cloudstorage.RetryParams(backoff_factor=1.1)
    with cloudstorage.open(filename, 'w', content_type='text/plain', options={'x-goog-meta-foo': 'foo', 'x-goog-meta-bar': 'bar'}, retry_params=writeRetryParams) as cloudstorageFile:
      cloudstorageFile.write(data)

app = webapp2.WSGIApplication([('/*', MainPage)], debug=True)
