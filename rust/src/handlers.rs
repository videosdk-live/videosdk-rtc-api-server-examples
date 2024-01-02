
use actix_web::{web,Responder,Result, HttpResponse};
use serde::{Serialize,Deserialize};
use jsonwebtoken::{ encode,  EncodingKey, Header,Algorithm};


extern crate reqwest;
extern crate serde_json;



// For JS-TOKEN
#[derive(Debug, Serialize)]
struct  Payload{
    apikey:String,
    permissions:[String;2],
}

//Info Structure
#[derive(Deserialize)]
pub struct PostBody {
    token: String,
}

//token REsponse
#[derive(Serialize, Deserialize)]
struct TokenResponse {
    token: String,
}
//meetinResponse
#[derive(Serialize, Deserialize)]
struct MeetingResponse {
    meetingId: String,
}



#[tokio::main]
pub async fn create_meeting_request(token:&String)-> Result<String, Box<dyn std::error::Error>>{
	let client = reqwest::Client::new();
	let jwt_token = token;
	let res = client
	    .post("https://api.videosdk.live/v2/rooms")
	    .header("Authorization", jwt_token)
	    .send()
	    .await?;
	let body=res.text().await?;
	Ok(body)
}

#[tokio::main]
pub async fn validate_meeting_request(token: &String,meeting_id:&String)-> Result<String, Box<dyn std::error::Error>>{
	let client = reqwest::Client::new();
	let jwt_token =token;
	let mut uri=String::from("https://api.videosdk.live/v2/rooms/validate/");
	uri.push_str(meeting_id);
	let res = client
	    .post(uri)
	    .header("Authorization", jwt_token)
	    .send()
	    .await?;
	let body=res.text().await?;
	Ok(body)
}


//Generate token handler
pub async fn get_token(req_body: web::Json<RequestBody>) -> impl Responder {
	let videosdk_api_key=String::from("");
	let videosdk_secret_key=String::from("");

	let roomId = req_body.roomId;
    let peerId = req_body.peerId;

	let payload = Payload {
	    apikey : videosdk_api_key,
	    permissions: [String::from("allow_join"),String::from("allow_mod")],
		// version: 2, //OPTIONAL  // For accessing the v2 API, For passing roomId, participantId or roles parameters in payload it is required to pass.
        // roomId: `2kyv-gzay-64pg`, //OPTIONAL // To provide customised access control, you can make the token applicable to a particular room by including the roomId in the payload.
        // participantId: `lxvdplwt`, //OPTIONAL  // You can include the participantId in the payload to limit the token's access to a particular participant.
        // roles: [String::from("crawler"),String::from("rtc")], //OPTIONAL // crawler role is only for accessing v2 API, you can not use this token for running the Meeting/Room. rtc is only allow for running the Meeting / Room, you can not use server-side APIs.
	};

	// Optionally add the version, roles, roomId, and peerId if you wish to use this token for joining the meeting
    // with a particular roomId or participantId
    if roomId.is_some() || peerId.is_some() {
        payload.version = Some(2);
        payload.roles = Some(vec![String::from("rtc")]);
    }

    if let Some(room_id) = roomId {
        payload.roomId = Some(room_id);
    }

    if let Some(peer_id) = peerId {
        payload.participantId = Some(peer_id);
    }


	// let header = Header::new(Algorithm:: HS256);
	let token = encode(&Header::new(Algorithm::HS256), &payload, &EncodingKey::from_secret(videosdk_secret_key.as_ref()))
	    .map_err(|err| println!("{:?}", err)).ok();
    
	    HttpResponse::Ok().json(TokenResponse {
		token: String::from(token.unwrap()),
	    })
}


//ROUTE::/create-meeting
//METHOD::POST
pub async fn create_meeting(body: web::Json<PostBody>) -> impl Responder {
	let res=create_meeting_request(&body.token).ok();//return response()
	let res_string:String=res.unwrap();
	let meeting_response:MeetingResponse=serde_json::from_str(&res_string).unwrap();
	HttpResponse::Ok().json(meeting_response)
}

//ROUTE::/validate-meeting
//METHOD::GET
pub async fn validate_meeting(meeting_id: web::Path<String>,body: web::Json<GetBody>) -> impl Responder {
	let res=validate_meeting_request(&body.token,&meeting_id).ok();//return response()
	let res_string:String=res.unwrap();
	let meeting_response:MeetingResponse=serde_json::from_str(&res_string).unwrap();
	HttpResponse::Ok().json(meeting_response)
}