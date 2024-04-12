using AccountsAPI.DbContexts;
using AccountsAPI.Models;
using Microsoft.EntityFrameworkCore;

namespace AccountsAPI.Services
{
    public class UserService
    {
        private readonly LazureDbContext _dbContext;

        public UserService(LazureDbContext dbContext)
        {
            _dbContext = dbContext;
        }

        public async Task<User> AddNewUserAsync(string walletId)    //Adding and saving the new user
        {
            var userToSave = new User
            {
                WalletId = walletId
            };

            try
            {
                await _dbContext.AddAsync(userToSave);
                await _dbContext.SaveChangesAsync();
                return userToSave;
            }
            catch (Exception ex)
            {
                Console.Error.WriteLine(ex);
                throw; 
            }
        }


        public async Task<User> GetUserByWalletIdAsync(string walletId)
        {
            return await _dbContext.Users.FirstOrDefaultAsync(u => u.WalletId == walletId);
        }

        public User GetUserByWalletId(string walletId)  //Getting a user by its wallet
        {
            User? user = _dbContext.Users.FirstOrDefault(u => u.WalletId == walletId);
            return user;
        }

    }
}