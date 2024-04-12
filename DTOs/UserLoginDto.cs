using System.ComponentModel.DataAnnotations;

namespace AccountsAPI.DTOs
{
    public class UserLoginDto
    {
        [Required]
        public string WalletId { get; set; }
    }
}
