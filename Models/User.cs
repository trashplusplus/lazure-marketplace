using System.ComponentModel.DataAnnotations;
using System.ComponentModel.DataAnnotations.Schema;

namespace AccountsAPI.Models
{
    public class User
    {
        [Key]
        [DatabaseGenerated(DatabaseGeneratedOption.Identity)]
        [Column("user_id")]
        public int? UserId { get; set; }
        [Column("wallet_id")]
        public string WalletId { get; set; }
        [Column("username")]
        public string? Username { get; set; }
    }
}
