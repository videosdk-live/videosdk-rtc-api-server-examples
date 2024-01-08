# Video SDK RTC rust API server example

## Requirements

- Rust
- Cargo
- libssl-dev

## Getting started

1. Clone the repo

   ```sh
   $ git clone https://github.com/videosdk-live/videosdk-rtc-api-server-examples.git
   $ cd rust
   ```

2. Update the api key and secret values in the `src/handlers.rs` file with the ones generated from the developer console.

   ```
   videosdk_api_key=''
   videosdk_secret_key=''
   VIDEOSDK_API_ENDPOINT=https://api.videosdk.live/v2
   ```

3. Run the server

   ```sh
   $ cargo run
   ```

### More info

Visit, [videosdk.live](https://www.videosdk.live/) to know more about VideoSDK and generate API key & secret.
