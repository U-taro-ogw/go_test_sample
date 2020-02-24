from flask import Flask, request, jsonify
from redis import Redis

app = Flask(__name__)
app.config['JSON_AS_ASCII'] = False
redis = Redis(host='redis', port=6379, db=1)

@app.route('/')
def hello():
    return 'Hello My App'


# redis疎通確認
@app.route('/redis_hits')
def redis_hits():
    redis.incr('hits')
    return 'Hello World! I have been seen %s times.' % redis.get('hits')


@app.route('/api_info', methods=['GET'])
def search():
    jwt_token = request.headers.get("Authorization")
    auth = redis.get(jwt_token)
    if not auth:
        return jsonify({"error": "Unauthorized"})

    ret_dict = {
        "api_name": "flask_api_one",
        "info": {
            "language": "python",
            "framework": "flask"
        }
    }
    return jsonify(ret_dict)

if __name__ == "__main__":
    app.run(host="0.0.0.0", debug=True)
