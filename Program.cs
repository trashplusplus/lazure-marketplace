using AccountsAPI.DbContexts;
using AccountsAPI.Services;
using Microsoft.AspNetCore.Authentication.JwtBearer;
using Microsoft.EntityFrameworkCore;
using Microsoft.IdentityModel.Tokens;
using System.Text;

var builder = WebApplication.CreateBuilder(args);

//DB CONFIG (POSTGRESQL)    
builder.Services.AddDbContext<LazureDbContext>(options =>
    options.UseNpgsql(builder.Configuration.GetConnectionString("PostgresConnection"))
    .LogTo(Console.WriteLine, LogLevel.Information).EnableSensitiveDataLogging());  //Detailed logging

// Getting key
var encryptionKey = Environment.GetEnvironmentVariable("ENCRYPTION_KEY"); 

var password = Environment.GetEnvironmentVariable("PASSWORD");
// Configure CipherService using the key and the password
builder.Services.AddSingleton(new CipherService(encryptionKey,password));

// JWT AUTH BASIC COFIG
builder.Services.AddAuthentication(JwtBearerDefaults.AuthenticationScheme)
    .AddJwtBearer(options =>
    {
        options.TokenValidationParameters = new TokenValidationParameters
        {
            ValidateIssuer = true,  
            ValidateAudience = true,
            ValidateLifetime = true,
            ValidateIssuerSigningKey = true,
            ValidIssuer = builder.Configuration["Jwt:Issuer"],
            ValidAudience = builder.Configuration["Jwt:Audience"],
            IssuerSigningKey = new SymmetricSecurityKey(Encoding.UTF8.GetBytes(builder.Configuration["Jwt:Key"]))
        };
    });

builder.Services.AddSingleton(new JwtTokenService(
    builder.Configuration["Jwt:Key"],
    builder.Configuration["Jwt:Issuer"],
    builder.Configuration["Jwt:Audience"]));

builder.Services.AddScoped<UserService>();
builder.Services.AddControllers();
builder.Services.AddEndpointsApiExplorer();
builder.Services.AddSwaggerGen();

var app = builder.Build();

if (app.Environment.IsDevelopment())
{
    app.UseSwagger();
    app.UseSwaggerUI();
}

app.UseHttpsRedirection();

app.UseAuthentication(); // Using jwt authentication
app.UseAuthorization();

app.MapControllers();

app.Run();
