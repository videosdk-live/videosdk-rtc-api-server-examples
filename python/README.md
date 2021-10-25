# Video SDK RTC Python API server example

## Requirements

- Python 3
- pip

## Getting started

1. Clone the repo

   ```sh
   $ git clone https://github.com/videosdk-live/videosdk-rtc-api-server-examples.git
   $ cd python
   ```

2. Update the api key and secret values in the `main.py` file with the ones generated from the developer console.

   ```
   VIDEOSDK_API_KEY=''
   VIDEOSDK_SECRET_KEY=''
   VIDEOSDK_API_ENDPOINT=https://api.videosdk.live
   ```

3. Setup virtual environment

   ```sh
   $ python -m venv env
   $ source env/bin/activate
   $ pip install -r requirements.txt
   ```

4. Run the server

   ```sh
   $ export FLASK_APP=main
   $ flask run
   ```

### More info

Visit, [videosdk.live](https://www.videosdk.live/) to know more about VideoSDK and generate API key & secret.
