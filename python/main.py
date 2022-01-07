from flask import Flask, request, jsonify
import jwt
import json
import datetime
import requests
from werkzeug.datastructures import Authorization

VIDEOSDK_API_KEY = ""
VIDEOSDK_SECRET_KEY = ""
VIDEOSDK_API_ENDPOINT = "https://api.videosdk.live"

app = Flask(__name__)


@app.route('/get-token')
def generateToken():
    expiration_in_seconds = 600
    expiration = datetime.datetime.now() + datetime.timedelta(seconds=expiration_in_seconds)
    token = jwt.encode(payload={
        'exp': expiration,
        'apikey': VIDEOSDK_API_KEY,
        'permissions': ["allow_join", "allow_mod"],
    }, key=VIDEOSDK_SECRET_KEY, algorithm: "HS256")

    return jsonify(token=token)

if __name__ == '__main__':
    print(generateToken())


@app.route('/create-meeting', methods=['POST'])
def createMeeting():
    obj = request.get_json()
    res = requests.post(VIDEOSDK_API_ENDPOINT + "/api/meetings",
                        headers={"Authorization": obj["token"]})
    return res.json()


@app.route('/validate-meeting/<string:meetingId>', methods=['POST'])
def validateMeeting(meetingId):
    print(meetingId)
    obj = request.get_json()
    res = requests.post(VIDEOSDK_API_ENDPOINT + "/api/meetings/" +
                        meetingId, headers={"Authorization": obj["token"]})
    return res.json()
