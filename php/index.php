<?php
require __DIR__ . '/vendor/autoload.php';

use Firebase\JWT\JWT;
use Klein\Klein as Route;

/** Your API key and secret */
$VIDEOSDK_API_KEY = "YOUR_API_KEY";
$VIDEOSDK_SECRET_KEY = "YOUR_SECRET_KEY";
$VIDEOSDK_API_ENDPOINT = "https://api.videosdk.live/v2";

$route = new Route();

$route->respond('GET', '/get-token', function ()
{

    header("Content-type: application/json; charset=utf-8");

    $issuedAt = new DateTimeImmutable();
    $expire = $issuedAt->modify('+24 hours')->getTimestamp();

    $payload = (object)[];

    $payload->apikey = $GLOBALS['VIDEOSDK_API_KEY'];
    $payload->permissions = array(
        "allow_join",
        "allow_mod"
    );
    $payload->iat = $issuedAt->getTimestamp();
    $payload->exp = $expire;
    // $payload->version = 2; //OPTIONAL  // For accessing the v2 API, For passing roomId, participantId or roles parameters in payload it is required to pass.
    // $payload->roomId = `2kyv-gzay-64pg`; //OPTIONAL // To provide customised access control, you can make the token applicable to a particular room by including the roomId in the payload.
    // $payload->participantId = `lxvdplwt`; //OPTIONAL  // You can include the participantId in the payload to limit the token's access to a particular participant.
    // $payload->roles =  array(
    //    "crawler",
    //    "rtc"
    // ); //OPTIONAL // crawler role is only for accessing v2 API, you can not use this token for running the Meeting/Room. rtc is only allow for running the Meeting / Room, you can not use server-side APIs.

    //OPTIONALLY add the version, roles, roomId, and peerId if you wish to use this token for joining the meeeting
    //with a particular roomId or participantId
    $roomId = $_GET['roomId'] ?? null;
    $peerId = $_GET['peerId'] ?? null;

    if ($roomId || $peerId) {
        $payload->version = 2;
        $payload->roles = ["rtc"];
    }

    if ($roomId) {
        $payload->roomId = $roomId;
    }

    if ($peerId) {
        $payload->participantId = $peerId;
    }

    $jwt = JWT::encode($payload, $GLOBALS['VIDEOSDK_SECRET_KEY']);

    return json_encode(array(
        "token" => $jwt
    ));
});

$route->respond('POST', '/create-meeting', function ()
{

    header("Content-type: application/json; charset=utf-8");

    $data = json_decode(file_get_contents('php://input') , true);

    $token = $data["token"];

    $curl = curl_init();

    curl_setopt_array($curl, array(
        CURLOPT_URL => $GLOBALS['VIDEOSDK_API_ENDPOINT'] . '/rooms',
        CURLOPT_HTTP_VERSION => CURL_HTTP_VERSION_1_1,
        CURLOPT_RETURNTRANSFER => true,
        CURLOPT_ENCODING => '',
        CURLOPT_MAXREDIRS => 10,
        CURLOPT_TIMEOUT => 0,
        CURLOPT_FOLLOWLOCATION => true,
        CURLOPT_HTTP_VERSION => CURL_HTTP_VERSION_1_1,
        CURLOPT_CUSTOMREQUEST => 'POST',
        CURLOPT_HTTPHEADER => array(
            'Authorization: ' . $token,
            'Content-Type: application/json'
        ) ,
    ));

    $response = curl_exec($curl);

    curl_close($curl);
    return $response;
});

$route->respond('GET', '/validate-meeting/[:meetingId]', function ($request)
{

    header("Content-type: application/json; charset=utf-8");

    $meetingId = $request->meetingId;

    $data = json_decode(file_get_contents('php://input') , true);

    $token = $data["token"];

    $curl = curl_init();

    curl_setopt_array($curl, array(
        CURLOPT_URL => $GLOBALS['VIDEOSDK_API_ENDPOINT'] . '/rooms/validate/' . $meetingId,
        CURLOPT_RETURNTRANSFER => true,
        CURLOPT_ENCODING => '',
        CURLOPT_MAXREDIRS => 10,
        CURLOPT_TIMEOUT => 0,
        CURLOPT_FOLLOWLOCATION => true,
        CURLOPT_HTTP_VERSION => CURL_HTTP_VERSION_1_1,
        CURLOPT_CUSTOMREQUEST => 'GET',
        CURLOPT_HTTPHEADER => array(
            'Authorization: ' . $token,
            'Content-Type: application/json'
        ) ,
    ));

    $response = curl_exec($curl);

    curl_close($curl);
    return $response;
});

$route->respond('POST', '/start-recording', function ()
{

    header("Content-type: application/json; charset=utf-8");

    $data = json_decode(file_get_contents('php://input') , true);

    $token = $data["token"];
    $roomId = $data["roomId"];

    $curl = curl_init();

    $body = array(
        "roomId" => $roomId,
        "config" => array(
            "layout" => array(
                "type" => "GRID",
                "priority" => "SPEAKER",
                "gridSize" => 1
            ) ,
            "theme" => "DARK",
            "mode" => "video-and-audio",
            "quality" => "low"
        ) ,
    );

    curl_setopt_array($curl, array(
        CURLOPT_URL => $GLOBALS['VIDEOSDK_API_ENDPOINT'] . '/recordings/start',
        CURLOPT_HTTP_VERSION => CURL_HTTP_VERSION_1_1,
        CURLOPT_RETURNTRANSFER => true,
        CURLOPT_ENCODING => '',
        CURLOPT_MAXREDIRS => 10,
        CURLOPT_TIMEOUT => 0,
        CURLOPT_FOLLOWLOCATION => true,
        CURLOPT_HTTP_VERSION => CURL_HTTP_VERSION_1_1,
        CURLOPT_CUSTOMREQUEST => 'POST',
        CURLOPT_HTTPHEADER => array(
            'Authorization: ' . $token,
            'Content-Type: application/json'
        ) ,
        CURLOPT_POSTFIELDS => json_encode($body) ,
    ));

    $response = curl_exec($curl);

    curl_close($curl);
    return $response;
});

$route->respond('POST', '/stop-recording', function ()
{

    header("Content-type: application/json; charset=utf-8");

    $data = json_decode(file_get_contents('php://input') , true);

    $token = $data["token"];
    $roomId = $data["roomId"];

    $curl = curl_init();

    $body = array(
        "roomId" => $roomId,
    );

    curl_setopt_array($curl, array(
        CURLOPT_URL => $GLOBALS['VIDEOSDK_API_ENDPOINT'] . '/recordings/end',
        CURLOPT_HTTP_VERSION => CURL_HTTP_VERSION_1_1,
        CURLOPT_RETURNTRANSFER => true,
        CURLOPT_ENCODING => '',
        CURLOPT_MAXREDIRS => 10,
        CURLOPT_TIMEOUT => 0,
        CURLOPT_FOLLOWLOCATION => true,
        CURLOPT_HTTP_VERSION => CURL_HTTP_VERSION_1_1,
        CURLOPT_CUSTOMREQUEST => 'POST',
        CURLOPT_HTTPHEADER => array(
            'Authorization: ' . $token,
            'Content-Type: application/json'
        ) ,
        CURLOPT_POSTFIELDS => json_encode($body) ,
    ));

    $response = curl_exec($curl);

    curl_close($curl);
    return $response;
});

$route->dispatch();

