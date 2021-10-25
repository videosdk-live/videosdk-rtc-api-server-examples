# Video SDK RTC ruby API server example

## Requirements

- Ruby (ruby-dev or ruby-devel)
- Ruby bundler (ruby-bundler)
- sqlite3 (libsqlite3-dev)

## Getting started

1. Clone the repo

   ```sh
   $ git clone https://github.com/videosdk-live/videosdk-rtc-api-server-examples.git
   $ cd ruby
   ```

2. Update the api key and secret values in the `main_controller.rb` file with the ones generated from the developer console.

   ```
   VIDEOSDK_API_KEY=''
   VIDEOSDK_SECRET_KEY=''
   VIDEOSDK_API_ENDPOINT=https://api.videosdk.live
   ```

3. Install dependencies

   ```sh
   $  bundle install
   ```

4. Run the server

   ```sh
   $ rails s
   ```

### More info

Visit, [videosdk.live](https://www.videosdk.live/) to know more about VideoSDK and generate API key & secret.
