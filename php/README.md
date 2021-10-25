# Video SDK RTC PHP API server example

## Requirements

- PHP
- PHP Composer
- php-curl

## Getting started

1. Clone the repo

   ```sh
   $ git clone https://github.com/videosdk-live/videosdk-rtc-api-server-examples.git
   $ cd php
   ```

2. Update the api key and secret values in the `index.php` file with the ones generated from the developer console.

   ```
   VIDEOSDK_API_KEY=''
   VIDEOSDK_SECRET_KEY=''
   VIDEOSDK_API_ENDPOINT="https://api.videosdk.live"
   ```

3. Install dependencies

   ```sh
   $ composer install
   ```

4. Run the server

   ```sh
   $ php -S localhost:3000
   ```

### More info

Visit, [videosdk.live](https://www.videosdk.live/) to know more about VideoSDK and generate API key & secret.
