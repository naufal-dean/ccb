import os
from flask import Flask


app = Flask(__name__)
app.config['APP_PORT'] = os.environ.get('APP_PORT', '8080')
app.config['SERVER_SERVICE_URL'] = os.environ.get('SERVER_SERVICE_URL', '')


from app import routes
