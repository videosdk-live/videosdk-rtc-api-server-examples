
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
	    .post("https://api.videosdk.live/api/meetings")
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
	let mut uri=String::from("https://api.videosdk.live/api/meetings/");
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
pub async fn get_token() -> impl Responder {
	let videosdk_api_key=String::from("");
	let videosdk_secret_key=String::from("");

	let payload = Payload {
	    apikey : videosdk_api_key,
	    permissions: [String::from("allow_join"),String::from("allow_mod")],
	};

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
//METHOD::POST
pub async fn validate_meeting(meeting_id: web::Path<String>,body: web::Json<PostBody>) -> impl Responder {
	let res=validate_meeting_request(&body.token,&meeting_id).ok();//return response()
	let res_string:String=res.unwrap();
	let meeting_response:MeetingResponse=serde_json::from_str(&res_string).unwrap();
	HttpResponse::Ok().json(meeting_response)
}