
require 'jwt'
require 'securerandom'
require 'net/http'
require 'uri'
require 'net/https'



class MainController < ApplicationController
	skip_forgery_protection
	$VIDEOSDK_API_KEY = ""
	$VIDEOSDK_SECRET_KEY = ""
	$VIDEOSDK_API_ENDPOINT= "https://api.videosdk.live/v2"

	def generateToken()
		now = Time.now
		exp = now + 86400
		payload = {
			apikey: $VIDEOSDK_API_KEY,
			permissions: ["allow_join", "allow_mod"],
			iat: now.to_i,
    		exp: exp.to_i,
			# version : 2, #OPTIONAL  # For accessing the v2 API, For passing roomId, participantId or roles parameters in payload it is required to pass.
            # roomId : `2kyv-gzay-64pg`,#OPTIONAL # To provide customised access control, you can make the token applicable to a particular room by including the roomId in the payload.
            # participantId : `lxvdplwt`,#OPTIONAL  # You can include the participantId in the payload to limit the token's access to a particular participant.
            # roles : ["crawler", "rtc"],#OPTIONAL # crawler role is only for accessing v2 API, you can not use this token for running the Meeting/Room. rtc is only allow for running the Meeting / Room, you can not use server-side APIs.
		}

		# OPTIONALLY add the version, roles, roomId, and peerId if you wish to use this token for joining the meeeting
        # with a particular roomId or participantId
		roomId = params[:roomId]
		peerId = params[:peerId]
	
		if roomId || peerId
		  payload[:version] = 2
		  payload[:roles] = ["rtc"]
		end
	
		if roomId
		  payload[:roomId] = roomId
		end
	
		if peerId
		  payload[:participantId] = peerId
		end

		# IMPORTANT: set nil as password parameter
		token = JWT.encode(payload, $VIDEOSDK_SECRET_KEY, 'HS256')
		json_content={token:token}
		render:json=>json_content
	end

	def createMeeting()
		# parse request body to json;
		jsonObj=JSON.parse(request.body.read());
		parsedToken=jsonObj["token"];

		uri = URI.parse($VIDEOSDK_API_ENDPOINT + "/rooms")
		#create http request
		http = Net::HTTP.new(uri.host, uri.port);
		http.use_ssl=true;
		request = Net::HTTP::Post.new(uri.request_uri)\
		#set Authorization header
		request["Authorization"]=parsedToken;
		res = http.request(request);

		render:json=>res.body

	end

	def validateMeeting()
		# request body excepts JSON Body
		jsonObj=JSON.parse(request.body.read());
		parsedToken=jsonObj["token"];
		meetingId= request.params["meetingId"];

		uri = URI.parse($VIDEOSDK_API_ENDPOINT +"/rooms/validate/"+meetingId)
		#create http request
		http = Net::HTTP.new(uri.host, uri.port);
		http.use_ssl=true;
		request = Net::HTTP::Get.new(uri.request_uri)
		#set Authorization header
		request["Authorization"]=parsedToken;
		res = http.request(request);

		render:json=>res.body

		
	end
end
