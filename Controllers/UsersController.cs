using AccountsAPI.DTOs;
using AccountsAPI.Services;
using Microsoft.AspNetCore.Mvc;
using AccountsAPI.Models;

namespace AccountsAPI.Controllers
{
    [ApiController]
    [Route("[controller]")]
    public class UsersController : Controller
    {
        private readonly UserService userService;
        private readonly JwtTokenService jwtTokenService;
        private readonly string password;

        public UsersController(UserService userService, JwtTokenService jwtTokenService,
          string password)
        {
            this.userService = userService;
            this.jwtTokenService = jwtTokenService;
            this.password = password;
        }

        [HttpPost("login")]         //Identification by login details(a crypto wallet unique id)
        public async Task<IActionResult> Login([FromBody] UserLoginDto userLoginDto)
        {
            if(!IsRequestAuthorized()) { return Unauthorized("The password is missing or invalid"); }


            var user = await userService.GetUserByWalletIdAsync(userLoginDto.WalletId);

            if(user == null)    // If first login ever
            {
                user = await userService.AddNewUserAsync(userLoginDto.WalletId);
            }

            var jwtToken = jwtTokenService.GenerateToken(user.UserId.ToString());
            
            Response.Headers.Add("Authorization", $"Bearer {jwtToken}");
            return Ok("Successful");
        }


        [HttpGet("{walletId}")]                 //Getting a user by its wallet
        public ActionResult<User> GetUserByWalletId(string walletId)
        {
            User user = userService.GetUserByWalletId(walletId);

            if(user == null) { return NotFound("Wrong wallet"); }

            return Ok(user);
        
        }

        private bool IsRequestAuthorized() //Checking decrypted password if attached
                                          //if the request is authorized
                                          //in order to ensure secured connection
        {
            string header = Request.Headers["Authorization"];

            if (string.IsNullOrEmpty(header) || !header.StartsWith("Bearer "))
            {
                Console.WriteLine("error 1");   //Internal logs
                return false;
            }

            var recievedPassword = header.Substring("Bearer ".Length).Trim();
            if (string.IsNullOrWhiteSpace(recievedPassword))
            {
                Console.WriteLine("error 2");   //Internal logs
                return false;
            }

            return recievedPassword.Equals(password);
        }

    }
}
        