All possible endpoints:
```
POST -> /add
GET -> /getbyid/:id
GET -> /getbytitle
```

Here some examples

Example of POST /add
```
{
    "name": "titlehere",
    "description": "deschere",
    "price": 999.999,
    "resource_link": "https://example.com/",
    "user_id": 1
}
```


Example of GET /getbyid/5
```
{
    "product_id": 5,
    "name": "Ice Cream",
    "description": "Source code of the game",
    "price": 100,
    "category_id": 1,
    "user_id": 1
}
```

Example of GET /getbytitle?title=Cream
```
[
    {
        "product_id": 3,
        "name": "Ice Cream",
        "description": "Source code of the game",
        "price": 100,
        "category_id": 1,
        "user_id": 1
    },
    {
        "product_id": 4,
        "name": "Ice Cream",
        "description": "Source code of the game",
        "price": 100,
        "category_id": 1,
        "user_id": 1
    },

]
```
