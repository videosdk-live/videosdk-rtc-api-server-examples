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

        string VIDEOSDK_API_ENDPOINT = "https://api.videosdk.live";
        string VIDEOSDK_API_KEY = "";
        string VIDEOSDK_SECRET_KEY = "";

        [HttpGet]
        [Route("get-token")]
        public JsonResult GetToken()
        {

            var token = JwtBuilder.Create()
                      .WithAlgorithm(new HMACSHA256Algorithm()) // symmetric
                      .WithSecret(VIDEOSDK_SECRET_KEY)
                      .AddClaim("exp", DateTimeOffset.UtcNow.AddHours(1).ToUnixTimeSeconds())
                      .AddClaim("iat", DateTimeOffset.UtcNow.ToUnixTimeSeconds())
                      .AddClaim("apikey", VIDEOSDK_API_KEY)
                      .AddClaim("permissions", new string[2] { "allow_join", "allow_mod" })
                      .Encode();


            return Json(new GetBody() { token = token });
        }

        [HttpPost]
        [Route("create-meeting")]
        public async Task<JsonResult> CreateMeetingAsync([FromBody] GetBody getBody)
        {
            string uri = VIDEOSDK_API_ENDPOINT + "/api/meetings";

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
            string uri = VIDEOSDK_API_ENDPOINT + "/api/meetings/" + meetingId;

            var response = await uri
               .WithHeader("Authorization", getBody.token)
               .PostAsync()
               .ReceiveJson();

            return Json(response);
        }
    }

    public class GetBody
    {
        public string token { get; set; }
    }
}
