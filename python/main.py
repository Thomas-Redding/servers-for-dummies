import webapp2

from FileSystem import FileSystem

class MainPage(webapp2.RequestHandler):
  def get(self):
    fileSystem = FileSystem()
    fileSystem.write("test-file-python.txt", "Lorem ipsum dol...")
    fileContents = fileSystem.read("test-file-python.txt")
    self.response.headers['Content-Type'] = 'text/plain'
    self.response.write('Hello World (Python)!\n')
    self.response.write(self.request.path + "\n")
    self.response.write(fileContents)

app = webapp2.WSGIApplication([('/*', MainPage)], debug=True)
