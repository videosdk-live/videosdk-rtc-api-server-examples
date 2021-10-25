package com.example.demo;

import org.springframework.boot.SpringApplication;
import org.springframework.boot.autoconfigure.SpringBootApplication;
import org.springframework.web.bind.annotation.*;


//? jwt imports
import java.time.Instant;
import java.util.Date;
import java.util.HashMap;
import java.util.Map;
import java.util.UUID;
import java.util.*;
import io.jsonwebtoken.Jwts;
import io.jsonwebtoken.SignatureAlgorithm;

//? http client
import java.io.IOException;
import java.net.URI;
import java.net.http.HttpClient;
import java.net.http.HttpRequest;

import java.net.http.HttpRequest.BodyPublishers;
import java.net.http.HttpResponse;
import java.time.Duration;

import org.json.simple.JSONObject;
//? to convert json
import org.json.simple.parser.JSONParser;
import org.json.simple.parser.ParseException;


@SpringBootApplication
@RestController

public class DemoApplication {


	String VIDEOSDK_API_KEY = "";
	String VIDEOSDK_SECRET_KEY = "";
	String VIDEOSDK_API_ENDPOINT = "https://api.videosdk.live";

	private static final HttpClient httpClient = HttpClient.newBuilder()
            .version(HttpClient.Version.HTTP_1_1)
            .connectTimeout(Duration.ofSeconds(10))
            .build();

	public static void main(String[] args) {
		SpringApplication.run(DemoApplication.class, args);
	}
	@RequestMapping(
		value="/get-token",
		method=RequestMethod.GET,
		produces = "application/json"
	)
    public JSONObject generatToken() {
		Map<String, Object> payload = new HashMap<>();
    	
		payload.put("apikey", VIDEOSDK_API_KEY);
    	payload.put("permissions", new String[]{"allow_join", "allow_mod"});
    	
		String token = Jwts.builder().setClaims(payload)
    		.setExpiration(new Date(System.currentTimeMillis() + 86400 * 1000))
    		.signWith(SignatureAlgorithm.HS256,VIDEOSDK_SECRET_KEY.getBytes()).compact();
    	
		 JSONObject tokenJsonObject = new JSONObject();
    	 tokenJsonObject.put("token", token);

		return tokenJsonObject;
    }

	@RequestMapping(
		value="/create-meeting",
		method=RequestMethod.POST,
		produces = "application/json"
	)
	public String createMeeting(@RequestBody String requestBody) throws IOException, InterruptedException, ParseException {

		//JsonParser to  json 
		JSONParser parser = new JSONParser();
		String token;
		
		JSONObject jsonBody=(JSONObject) parser.parse(requestBody.toString());
		token=(String)jsonBody.get("token");

		  HttpRequest request = HttpRequest.newBuilder()
		  .POST(BodyPublishers.ofString(""))
                .uri(URI.create(VIDEOSDK_API_ENDPOINT + "/api/meetings"))
                .setHeader("Authorization", token) // add request header
                .build();

        	HttpResponse<String> response = httpClient.send(request, HttpResponse.BodyHandlers.ofString());

        	return response.body();
    	}


	@RequestMapping(
		value="/validate-meeting/{meetingId}",
		method=RequestMethod.POST,
		produces = "application/json"
	)
	public String validateMeeting(@RequestBody String requestBody,@PathVariable String meetingId) throws ParseException, IOException, InterruptedException {
        	JSONParser parser = new JSONParser();
		String token;
		
		JSONObject jsonBody=(JSONObject) parser.parse(requestBody.toString());
		token=(String)jsonBody.get("token");
		String id=meetingId;
		String url= VIDEOSDK_API_ENDPOINT+"/api/meetings/"+id;

		  HttpRequest request = HttpRequest.newBuilder()
		  .POST(BodyPublishers.ofString(""))
                .uri(URI.create(url))
                .setHeader("Authorization", token) // add request header
                .build();

        	HttpResponse<String> response = httpClient.send(request, HttpResponse.BodyHandlers.ofString());
		
        	return (response.body());
    	}


}
