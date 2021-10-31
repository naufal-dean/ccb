from app import app


@app.route('/', methods=['GET'])
def home():
    return 'Hello world!'


@app.route('/server', methods=['GET'])
def server():
    return 'Server endpoint!'
