# Video SDK RTC GO API server example

## Requirements

- Golang

## Getting started

1. Clone the repo

   ```sh
   $ git clone https://github.com/videosdk-live/videosdk-rtc-api-server-examples.git
   $ cd go
   ```

2. Update the api key and secret values in the `main.go` file with the ones generated from the developer console.

   ```
   VIDEOSDK_API_KEY=''
   VIDEOSDK_SECRET_KEY=''
   VIDEOSDK_API_ENDPOINT="https://api.videosdk.live"
   ```

3. Install dependencies

   ```sh
   $ go install
   ```

4. Run the server

   ```sh
   $ go run main.go
   ```

### More info

Visit, [videosdk.live](https://www.videosdk.live/) to know more about VideoSDK and generate API key & secret.
