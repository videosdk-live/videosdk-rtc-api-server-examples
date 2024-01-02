using System;
using System.Collections.Generic;
using System.Linq;
using System.Threading.Tasks;
using System.Web;
using Flurl;
using Flurl.Http;
using JWT.Algorithms;
using JWT.Builder;
using Microsoft.AspNetCore.Http;
using Microsoft.AspNetCore.Mvc;

// For more information on enabling MVC for empty projects, visit https://go.microsoft.com/fwlink/?LinkID=397860

namespace jwt_example.Controllers
{
    public class HelloController : Controller
    {

        string VIDEOSDK_API_ENDPOINT = "https://api.videosdk.live/v2";
        string VIDEOSDK_API_KEY = "";
        string VIDEOSDK_SECRET_KEY = "";

        [HttpGet]
        [Route("get-token")]
        public JsonResult GetToken([FromBody] GetBody getBody, [FromQuery] string roomId, [FromQuery] string peerId)
        {

            var token = JwtBuilder.Create()
                      .WithAlgorithm(new HMACSHA256Algorithm()) // symmetric
                      .WithSecret(VIDEOSDK_SECRET_KEY)
                      .AddClaim("exp", DateTimeOffset.UtcNow.AddHours(1).ToUnixTimeSeconds())
                      .AddClaim("iat", DateTimeOffset.UtcNow.ToUnixTimeSeconds())
                      .AddClaim("apikey", VIDEOSDK_API_KEY)
                      .AddClaim("permissions", new string[2] { "allow_join", "allow_mod" })
                      .Encode();
                      // .AddClaim("version", 2) //OPTIONAL  // For accessing the v2 API, For passing roomId, participantId or roles parameters in payload it is required to pass.
                      // .AddClaim("roomId", "2kyv-gzay-64pg") //OPTIONAL // To provide customised access control, you can make the token applicable to a particular room by including the roomId in the payload.
                      // .AddClaim("participantId", "lxvdplwt")  //OPTIONAL  // You can include the participantId in the payload to limit the token's access to a particular participant.
                      // .AddClaim("roles", new string[2] { "crawler", "rtc" }); //OPTIONAL // crawler role is only for accessing v2 API, you can not use this token for running the Meeting/Room. rtc is only allow for running the Meeting / Room, you can not use server-side APIs.

                      //OPTIONALLY add the version, roles, roomId, and peerId if you wish to use this token for joining the meeeting
                      //with a particular roomId or participantId
                      if (!string.IsNullOrEmpty(roomId) || !string.IsNullOrEmpty(peerId))
                      {
                          token = JwtBuilder.Create(token)
                              .AddClaim("version", "2")
                              .AddClaim("roles", new string[1] { "rtc" } )
                              .Encode();
                      }
              
                      if (!string.IsNullOrEmpty(roomId))
                      {
                          token = JwtBuilder.Create(token)
                              .AddClaim("roomId", roomId)
                              .Encode();
                      }
              
                      if (!string.IsNullOrEmpty(peerId))
                      {
                          token = JwtBuilder.Create(token)
                              .AddClaim("participantId", peerId)
                              .Encode();
                      }


            return Json(new GetBody() { token = token });
        }

        [HttpPost]
        [Route("create-meeting")]
        public async Task<JsonResult> CreateMeetingAsync([FromBody] GetBody getBody)
        {
            string uri = VIDEOSDK_API_ENDPOINT + "/rooms";

            var response = await uri
                .WithHeader("Authorization", getBody.token)
                .PostAsync()
                .ReceiveJson();

            return Json(response);
        }

        [HttpPost]
        [Route("validate-meeting/{meetingId}")]
        public async Task<JsonResult> ValidateMeetingAsync([FromBody] GetBody getBody, string meetingId)
        {
            string uri = VIDEOSDK_API_ENDPOINT + "/rooms/validate/" + meetingId;

            var response = await uri
               .WithHeader("Authorization", getBody.token)
               .GetAsync()
               .ReceiveJson();

            return Json(response);
        }
    }

    public class GetBody
    {
        public string token { get; set; }
    }
}
