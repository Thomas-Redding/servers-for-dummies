import os
import cloudstorage
from google.appengine.api import app_identity

class FileSystem:
  def __init__(self, bucketName=''):
    if bucketName == '':
      bucketName = os.environ.get('BUCKET_NAME', app_identity.get_default_gcs_bucket_name())
    self._bucketName = bucketName

  def write(self, filePath, data):
    self.writeFancy(filePath, data, {})
  def writeFancy(self, filePath, data, metadata):
    filename = '/' + self._bucketName + '/' + filePath
    writeRetryParams = cloudstorage.RetryParams(backoff_factor=1.1)
    with cloudstorage.open(filename, 'w', content_type='text/plain', options=metadata, retry_params=writeRetryParams) as cloudstorageFile:
      cloudstorageFile.write(data)

  def read(self, filePath):
    filename = '/' + self._bucketName + '/' + filePath
    rtn = None
    with cloudstorage.open(filename) as cloudstorageFile:
      rtn = cloudstorageFile.read()
    return rtn
