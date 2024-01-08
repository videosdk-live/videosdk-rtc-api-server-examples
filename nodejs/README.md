# Video SDK RTC Node JS API server example

## Requirements

- Node.js 12+
- NPM

## Getting started

1. Clone the repo

   ```sh
   $ git clone https://github.com/videosdk-live/videosdk-rtc-nodejs-sdk-example
   $ cd nodejs
   ```

2. Copy the `.env.example` file to `.env` file.

   ```sh
   $ cp .env.example .env
   ```

3. Update the api key and secret values in the `.env` file with the ones generated from the developer console.

   ```
   VIDEOSDK_API_KEY=''
   VIDEOSDK_SECRET_KEY=''
   VIDEOSDK_API_ENDPOINT=https://api.videosdk.live/v2
   ```

4. Install NPM packages

   ```sh
   $ npm install
   ```

5. Run the server

   ```sh
   $ npm run start
   ```

### More info

Visit, [videosdk.live](https://www.videosdk.live/) to know more about VideoSDK and generate API key & secret.
