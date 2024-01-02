from flask import Flask, request, jsonify
import jwt
import json
import datetime
import requests
from werkzeug.datastructures import Authorization

VIDEOSDK_API_KEY = ""
VIDEOSDK_SECRET_KEY = ""
VIDEOSDK_API_ENDPOINT = "https://api.videosdk.live/v2"

app = Flask(__name__)


@app.route('/get-token')
def generateToken():
    expiration_in_seconds = 600
    expiration = datetime.datetime.now() + datetime.timedelta(seconds=expiration_in_seconds)
    payload={
        'exp': expiration,
        'apikey': VIDEOSDK_API_KEY,
        'permissions': ["allow_join", "allow_mod"],
        # 'version' : 2, //OPTIONAL  // For accessing the v2 API, For passing roomId, participantId or roles parameters in payload it is required to pass.
        # 'roomId': `2kyv-gzay-64pg`,//OPTIONAL // To provide customised access control, you can make the token applicable to a particular room by including the roomId in the payload.
        # 'participantId': `lxvdplwt`,//OPTIONAL  // You can include the participantId in the payload to limit the token's access to a particular participant.
        # 'roles': ["crawler", "rtc"],//OPTIONAL // crawler role is only for accessing v2 API, you can not use this token for running the Meeting/Room. rtc is only allow for running the Meeting / Room, you can not use server-side APIs.
    }
    

    # OPTIONALLY add the version, roles, roomId, and peerId if you wish to use this token for joining the meeeting
    # with a particular roomId or participantId
    roomId = request.args.get('roomId', '')
    peerId = request.args.get('peerId', '')

    if roomId or peerId:
        payload['version'] = 2
        payload['roles'] = ["rtc"]

    if roomId:
        payload['roomId'] = roomId

    if peerId:
        payload['participantId'] = peerId    

    token = jwt.encode(key=VIDEOSDK_SECRET_KEY, algorithm= "HS256").decode('UTF-8')

    return jsonify(token=token)

if __name__ == '__main__':
    print(generateToken())


@app.route('/create-meeting', methods=['POST'])
def createMeeting():
    obj = request.get_json()
    res = requests.post(VIDEOSDK_API_ENDPOINT + "/rooms",
                        headers={"Authorization": obj["token"]})
    return res.json()


@app.route('/validate-meeting/<string:meetingId>', methods=['GET'])
def validateMeeting(meetingId):
    print(meetingId)
    obj = request.get_json()
    res = requests.get(VIDEOSDK_API_ENDPOINT + "/rooms/validate" +
                        meetingId, headers={"Authorization": obj["token"]})
    return res.json()
