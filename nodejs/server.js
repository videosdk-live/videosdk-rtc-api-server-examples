require("dotenv").config();

const express = require("express");
const cors = require("cors");
const morgan = require("morgan");
const { default: fetch } = require("node-fetch");
const jwt = require("jsonwebtoken");

const PORT = 9000;
const app = express();

app.use(cors());
app.use(express.json());
app.use(express.urlencoded({ extended: true }));
app.use(morgan("dev"));

//
app.get("/", (req, res) => {
  res.send("Hello World!");
});

//
app.get("/get-token", (req, res) => {
  const API_KEY = process.env.VIDEOSDK_API_KEY;
  const SECRET_KEY = process.env.VIDEOSDK_SECRET_KEY;

  const options = { expiresIn: "10m", algorithm: "HS256" };

  const { roomId, peerId } = req.body;

  let payload = {
    apikey: API_KEY,
    permissions: ["allow_join", "allow_mod"], // also accepts "ask_join"
    // version: 2, //OPTIONAL  // For accessing the v2 API, For passing roomId, participantId or roles parameters in payload it is required to pass.
    // roomId: `2kyv-gzay-64pg`, //OPTIONAL // To provide customised access control, you can make the token applicable to a particular room by including the roomId in the payload.
    // participantId: `lxvdplwt`, //OPTIONAL  // You can include the participantId in the payload to limit the token's access to a particular participant.
    // roles: ["crawler", "rtc"], //OPTIONAL // crawler role is only for accessing v2 API, you can not use this token for running the Meeting/Room. rtc is only allow for running the Meeting / Room, you can not use server-side APIs.
  };

  //OPTIONALLY add the version, roles, roomId, and peerId if you wish to use this token for joining the meeeting
  //with a particular roomId or participantId
  if (roomId || peerId) {
    payload.version = 2;
    payload.roles = ["rtc"];
  }
  if (roomId) {
    payload.roomId = roomId;
  }
  if (peerId) {
    payload.participantId = peerId;
  }

  const token = jwt.sign(payload, SECRET_KEY, options);
  res.json({ token });
});

//
app.post("/create-meeting/", (req, res) => {
  const { token, region } = req.body;
  const url = `${process.env.VIDEOSDK_API_ENDPOINT}/rooms`;
  const options = {
    method: "POST",
    headers: { Authorization: token, "Content-Type": "application/json" },
    body: JSON.stringify({ region }),
  };

  fetch(url, options)
    .then((response) => response.json())
    .then((result) => res.json(result)) // result will contain roomId
    .catch((error) => console.error("error", error));
});

//
app.post("/validate-meeting/:meetingId", (req, res) => {
  const token = req.body.token;
  const meetingId = req.params.meetingId;

  const url = `${process.env.VIDEOSDK_API_ENDPOINT}/rooms/validate/${meetingId}`;

  const options = {
    method: "GET",
    headers: { Authorization: token },
  };

  fetch(url, options)
    .then((response) => response.json())
    .then((result) => res.json(result)) // result will contain roomId
    .catch((error) => console.error("error", error));
});

//
app.listen(PORT, () => {
  console.log(`API server listening at http://localhost:${PORT}`);
});
