using AccountsAPI.Models;
using Microsoft.EntityFrameworkCore;


namespace AccountsAPI.DbContexts
{
    public class LazureDbContext : DbContext
    {
        public DbSet<User> Users { get; set; }
        public LazureDbContext(DbContextOptions<LazureDbContext> options) : base(options)
        {
        }
        protected override void OnModelCreating(ModelBuilder modelBuilder)
        {
            modelBuilder.Entity<User>().ToTable("users"); // In order to have a correct name mapping

        }


    }
}
