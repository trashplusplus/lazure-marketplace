# AccountsAPI

All active endpoints: 

Users Controller

GET: userscontroller/{walletId}  - Getting a user by its wallet
POST: userscontroller/login - Identification by login details(a crypto wallet unique id)

Example:
userscontroller/login
{
    "walletId":"somewallet"
}

Response:
200("Successful)

