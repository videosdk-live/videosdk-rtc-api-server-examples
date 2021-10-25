# Video SDK RTC .NET core API server example

## Requirements

- .NET core 3.1

## Getting started

1. Clone the repo

   ```sh
   $ git clone https://github.com/videosdk-live/videosdk-rtc-api-server-examples.git
   $ cd dotnet/jwt-example
   ```

2. Update the api key and secret values in the `HelloController.cs` file with the ones generated from the developer console.

   ```
   VIDEOSDK_API_KEY=''
   VIDEOSDK_SECRET_KEY=''
   VIDEOSDK_API_ENDPOINT=https://api.videosdk.live
   ```

3. Install dependencies

   ```sh
   $  dotnet restore
   ```

4. Run the server

   ```sh
   $ dotnet run
   ```

### More info

Visit, [videosdk.live](https://www.videosdk.live/) to know more about VideoSDK and generate API key & secret.
