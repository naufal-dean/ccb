import requests
from app import app


@app.route('/', methods=['GET'])
def home():
    return 'Hello world!'


@app.route('/client', methods=['GET'])
def client():
    res = requests.get(app.config.get('SERVER_SERVICE_URL') + '/server')
    return 'Client endpoint! ' + res.text
