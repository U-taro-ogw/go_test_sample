from flask import Flask, request, jsonify
from redis import Redis

app = Flask(__name__)
app.config['JSON_AS_ASCII'] = False
redis = Redis(host='redis', port=6379, db=1)

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
        return jsonify({"error": "Unauthorized"}), 401

    ret_dict = {
        "api_name": "flask_api_two",
        "info": {
            "language": "python",
            "framework": "flask"
        }
    }
    return jsonify(ret_dict), 200

if __name__ == "__main__":
    app.run(host="0.0.0.0", debug=True, port=6000)
