using AccountsAPI.DTOs;
using AccountsAPI.Services;
using Microsoft.AspNetCore.Http;
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

        public UsersController(UserService userService, JwtTokenService jwtTokenService)
        {
            this.userService = userService;
            this.jwtTokenService = jwtTokenService;
        }

        [HttpPost("login")]         //Identification by login details(a crypto wallet unique id)
        public async Task<IActionResult> Login([FromBody] UserLoginDto userLoginDto)
        {
            
            var user = await userService.GetUserByWalletIdAsync(userLoginDto.WalletId);

            if(user == null)    // If first login ever
            {
                user = await userService.AddNewUserAsync(userLoginDto.WalletId);
                Console.WriteLine(user.ToString());
            }

            var jwtToken = jwtTokenService.GenerateToken(userLoginDto.WalletId);
            return Ok(new { Token = jwtToken, UserId = user.UserId });
        }


        [HttpGet("{walletId}")]                 //Getting a user by its wallet
        public ActionResult<User> GetUserByWalletId(string walletId)
        {
            User user = userService.GetUserByWalletId(walletId);

            if(user == null) { return NotFound("Wrong wallet"); }

            return Ok(user);
        }
        
    }
}
        