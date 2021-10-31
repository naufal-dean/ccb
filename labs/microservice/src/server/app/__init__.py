import os
from flask import Flask


app = Flask(__name__)
app.config['APP_PORT'] = os.environ.get('APP_PORT', '8081')


from app import routes
