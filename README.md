![image](https://github.com/trashplusplus/lazure-marketplace/assets/19663951/ef6565d8-144a-454c-bf19-4a681022cdda)


All possible endpoints:
```
✉️ POST -> /product
🔴 DELETE -> /product/:id
🍃 GET -> /product/:id


🍃 GET -> /catalog
🍃 GET -> /wallet/:walletId
🍃 GET -> /category
```

Here some examples

✉️ Example of POST /product
```

Header: Authorization

{
    "name": "titlehere",
    "description": "deschere",
    "price": 999.999,
    "resource_link": "https://example.com/",
    "category_id": 1,
    "user_id": 1
}
```

🔴 Example of DELETE /product/7
```

Header: Authorization

{
    "message": "[raw meat] has been deleted"
}
```

🍃 Example of GET /product/5
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

🍃 Example of GET /catalog?name=Ice

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

🍃 Example of GET /wallet/0xgj8r03k8tgoiejrg849gj0wrgjowihj

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

🍃 Example of GET /category

```
[
    {
        "category_id": 1,
        "name": "Source Code",
        "description": "This is source code"
    },
    {
        "category_id": 2,
        "name": "Assets",
        "description": "Category for asset files"
    },
    {
        "category_id": 3,
        "name": "IP (intellectual privacy)",
        "description": "Category for intellectual privacy related files"
    }
]
```
🍃 Example of GET /get-products?limit2

Creator's own products are not visible for him

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

