# Video SDK RTC Java API server example

## Requirements

- Java JDK 11

## Getting started

1. Clone the repo

   ```sh
   $ git clone https://github.com/videosdk-live/videosdk-rtc-api-server-examples.git
   $ cd java
   ```

2. Update the api key and secret values in the `DemoApplication.java` file with the ones generated from the developer console.

   ```
   VIDEOSDK_API_KEY=''
   VIDEOSDK_SECRET_KEY=''
   VIDEOSDK_API_ENDPOINT="https://api.videosdk.live"
   ```

3. Run the server

   ```sh
   $  ./mvnw spring-boot:run
   ```

### More info

Visit, [videosdk.live](https://www.videosdk.live/) to know more about VideoSDK and generate API key & secret.
