
require 'jwt'
require 'securerandom'
require 'net/http'
require 'uri'
require 'net/https'



class MainController < ApplicationController
	skip_forgery_protection
	$VIDEOSDK_API_KEY = ""
	$VIDEOSDK_SECRET_KEY = ""
	$VIDEOSDK_API_ENDPOINT= "https://api.videosdk.live"

	def generateToken()
		now = Time.now
		exp = now + 86400
		payload = {
			version: 2,
			apikey: $VIDEOSDK_API_KEY,
			permissions: ["allow_join", "allow_mod"],
			roles: ["CRAWLER", "PUBLISHER"],
			iat: now.to_i,
    			exp: exp.to_i
		}
		# IMPORTANT: set nil as password parameter
		token = JWT.encode(payload, $VIDEOSDK_SECRET_KEY, 'HS256')
		json_content={token:token}
		render:json=>json_content
	end

	def createMeeting()
		# parse request body to json;
		jsonObj=JSON.parse(request.body.read());
		parsedToken=jsonObj["token"];

		uri = URI.parse($VIDEOSDK_API_ENDPOINT + "/api/meetings")
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

		uri = URI.parse($VIDEOSDK_API_ENDPOINT +"/api/meetings/"+meetingId)
		#create http request
		http = Net::HTTP.new(uri.host, uri.port);
		http.use_ssl=true;
		request = Net::HTTP::Post.new(uri.request_uri)\
		#set Authorization header
		request["Authorization"]=parsedToken;
		res = http.request(request);

		render:json=>res.body

		
	end
end
