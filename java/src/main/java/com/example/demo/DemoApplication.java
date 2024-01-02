package com.example.demo;

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
//? jwt imports
import java.time.Instant;
import java.util.*;
import java.util.Date;
import java.util.HashMap;
import java.util.Map;
import java.util.UUID;
import org.json.simple.JSONObject;
//? to convert json
import org.json.simple.parser.JSONParser;
import org.json.simple.parser.ParseException;
import org.springframework.boot.SpringApplication;
import org.springframework.boot.autoconfigure.SpringBootApplication;
import org.springframework.web.bind.annotation.*;

@SpringBootApplication
@RestController
public class DemoApplication {

  String VIDEOSDK_API_KEY = "";
  String VIDEOSDK_SECRET_KEY = "";
  String VIDEOSDK_API_ENDPOINT = "https://api.videosdk.live/v2";

  private static final HttpClient httpClient = HttpClient
    .newBuilder()
    .version(HttpClient.Version.HTTP_1_1)
    .connectTimeout(Duration.ofSeconds(10))
    .build();

  public static void main(String[] args) {
    SpringApplication.run(DemoApplication.class, args);
  }

  @RequestMapping(
    value = "/get-token",
    method = RequestMethod.GET,
    produces = "application/json"
  )
  public JSONObject generatToken(
    @RequestParam(required = false) String roomId,
    @RequestParam(required = false) String peerId
  ) {
    Map<String, Object> payload = new HashMap<>();

    payload.put("apikey", VIDEOSDK_API_KEY);
    payload.put("permissions", new String[] { "allow_join", "allow_mod" });
    // payload.put("version", 2); //OPTIONAL  // For accessing the v2 API, For passing roomId, participantId or roles parameters in payload it is required to pass.
    // payload.put("roomId", "2kyv-gzay-64pg"); //OPTIONAL // To provide customised access control, you can make the token applicable to a particular room by including the roomId in the payload.
    // payload.put("participantId", "lxvdplwt"); //OPTIONAL  // You can include the participantId in the payload to limit the token's access to a particular participant.
    // payload.put("roles", new String[]{"crawler", "rtc"}) //OPTIONAL // crawler role is only for accessing v2 API, you can not use this token for running the Meeting/Room. rtc is only allow for running the Meeting / Room, you can not use server-side APIs.

    //OPTIONALLY add the version, roles, roomId, and peerId if you wish to use this token for joining the meeeting
    //with a particular roomId or participantId
    if (roomId != null || peerId != null) {
      payload.put("version", 2);
      payload.put("roles", new String[] { "rtc" });
    }
    if (roomId != null) {
      payload.put("roomId", roomId);
    }
    if (peerId != null) {
      payload.put("participantId", peerId);
    }
    String token = Jwts
      .builder()
      .setClaims(payload)
      .setExpiration(new Date(System.currentTimeMillis() + 86400 * 1000))
      .signWith(SignatureAlgorithm.HS256, VIDEOSDK_SECRET_KEY.getBytes())
      .compact();

    JSONObject tokenJsonObject = new JSONObject();
    tokenJsonObject.put("token", token);

    return tokenJsonObject;
  }

  @RequestMapping(
    value = "/create-meeting",
    method = RequestMethod.POST,
    produces = "application/json"
  )
  public String createMeeting(@RequestBody String requestBody)
    throws IOException, InterruptedException, ParseException {
    //JsonParser to  json
    JSONParser parser = new JSONParser();
    String token;

    JSONObject jsonBody = (JSONObject) parser.parse(requestBody.toString());
    token = (String) jsonBody.get("token");

    HttpRequest request = HttpRequest
      .newBuilder()
      .POST(BodyPublishers.ofString(""))
      .uri(URI.create(VIDEOSDK_API_ENDPOINT + "/rooms"))
      .setHeader("Authorization", token) // add request header
      .build();

    HttpResponse<String> response = httpClient.send(
      request,
      HttpResponse.BodyHandlers.ofString()
    );

    return response.body();
  }

  @RequestMapping(
    value = "/validate-meeting/{meetingId}",
    method = RequestMethod.POST,
    produces = "application/json"
  )
  public String validateMeeting(
    @RequestBody String requestBody,
    @PathVariable String meetingId
  ) throws ParseException, IOException, InterruptedException {
    JSONParser parser = new JSONParser();
    String token;

    JSONObject jsonBody = (JSONObject) parser.parse(requestBody.toString());
    token = (String) jsonBody.get("token");
    String id = meetingId;
    String url = VIDEOSDK_API_ENDPOINT + "/rooms/validate/" + id;

    HttpRequest request = HttpRequest
      .newBuilder()
      .GET(BodyPublishers.ofString(""))
      .uri(URI.create(url))
      .setHeader("Authorization", token) // add request header
      .build();

    HttpResponse<String> response = httpClient.send(
      request,
      HttpResponse.BodyHandlers.ofString()
    );

    return (response.body());
  }
}
