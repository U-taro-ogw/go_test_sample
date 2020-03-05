from flask import Flask, request, jsonify
from redis import Redis
from time import sleep
# import requests

app = Flask(__name__)
app.config['JSON_AS_ASCII'] = False
redis = Redis(host='redis', port=6379, db=1, decode_responses=True)

# redis疎通確認
@app.route('/redis_hits')
def redis_hits():
    redis.incr('hits')
    return 'Hello World! I have been seen %s times.' % redis.get('hits')

@app.route('/api_info', methods=['GET'])
def search():
    jwt_token = request.headers.get("Authorization")
    if not jwt_token:
        return jsonify({"error": "Unauthorized"}), 401
    auth = redis.get(jwt_token)
    if not auth:
        return jsonify({"error": "Unauthorized"}), 401

    api_response = {
        "api_name": "flask_api_two",
        "info": {
            "language": "python",
            "framework": "flask"
        }
    }
    return jsonify(api_response), 200

@app.route('/sleep_api', methods=['GET'])
def sleep_api():
    sleep(10)
    api_response = {
        "api_name": "flask_api_two",
        "sleep_time": 10
    }
    return jsonify(api_response), 200

if __name__ == "__main__":
    app.run(host="0.0.0.0", debug=True, port=6000)
