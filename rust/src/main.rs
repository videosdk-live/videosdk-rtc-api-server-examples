use actix_web::{web, App, HttpServer};


mod handlers;

#[actix_rt::main]
async fn main() -> std::io::Result<()> {
    std::env::set_var("RUST_LOG", "actix_web=debug");

    // Start http server
    HttpServer::new(move || {
        App::new()
            .route("/get-token", web::get().to(handlers::get_token))
            .route("/create-meeting", web::post().to(handlers::create_meeting))
            .route("/validate-meeting/{meetingId}", web::post().to(handlers::validate_meeting))
    })
    .bind("127.0.0.1:8080")?
    .run()
    .await
}