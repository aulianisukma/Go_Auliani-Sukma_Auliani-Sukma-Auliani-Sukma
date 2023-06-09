**API DOCUMENTATION**

**REGISTER**

Method : POST

Endpoint : users/register

Request (Body) : 

{

`    `"name" : "aulianisukma",

`    `"email" : "aulianisukma@gmail.com",

`    `"role" : "users",

`    `"password" : "12345"

}

Response :

{

`    `"message": "success create new user",

`    `"user": {

`        `"ID": 1,

`        `"CreatedAt": "2023-06-08T19:44:00.25+07:00",

`        `"UpdatedAt": "2023-06-08T19:44:00.25+07:00",

`        `"DeletedAt": **null**,

`        `"name": "aulianisukma",

`        `"email": "aulianisukma@gmail.com",

`        `"password": "12345"

`    `}

}

**LOGIN**

Method: POST

Endpoint: users/login

Request (Body):

{

`    `"email" : "aulianisukma@gmail.com",

`    `"password" : "12345"

}

Response :

{

`    `"message": "success login user",

`    `"user": {

`        `"id": 1,

`        `"name": "aulianisukma",

`        `"email": "aulianisukma@gmail.com",

`        `"token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2ODYyMzIwOTQsIm5hbWUiOiJhdWxpYW5pc3VrbWEiLCJ1c2VySUQiOjF9.\_UwfxYPrP-nWf1wAbmR6e84A6V7hjPiKBj\_GLwvd8SU"

`    `}

}

**CREATE ITEM**

Method: POST

Endpoint: /items

Request (Body):

{

`    `"category\_id": 1,

`    `"name": "samsung",

`    `"description": "available",

`    `"stock" : 5,

`    `"price": 10000000

}

Response:

{

`    `"item": {

`        `"ID": 4,

`        `"CreatedAt": "2023-06-08T20:01:55.852+07:00",

`        `"UpdatedAt": "2023-06-08T20:01:55.852+07:00",

`        `"DeletedAt": **null**,

`        `"category\_id": 1,

`        `"Category": {

`            `"ID": 1,

`            `"CreatedAt": "2023-06-08T19:55:26.548+07:00",

`            `"UpdatedAt": "2023-06-08T19:55:26.548+07:00",

`            `"DeletedAt": **null**,

`            `"category\_name": "elektronik"

`        `},

`        `"name": "samsung",

`        `"description": "available",

`        `"stock": 5,

`        `"price": 10000000

`    `},

`    `"message": "success create new item"

}

**GET ALL ITEM**

Method: GET

Endpoint: /items

Response:

{

`    `"items": [

`        `{

`            `"ID": 4,

`            `"CreatedAt": "2023-06-08T20:01:55.852+07:00",

`            `"UpdatedAt": "2023-06-08T20:01:55.852+07:00",

`            `"DeletedAt": **null**,

`            `"category\_id": 1,

`            `"Category": {

`                `"ID": 1,

`                `"CreatedAt": "2023-06-08T19:55:26.548+07:00",

`                `"UpdatedAt": "2023-06-08T19:55:26.548+07:00",

`                `"DeletedAt": **null**,

`                `"category\_name": "elektronik"

`            `},

`            `"name": "samsung",

`            `"description": "available",

`            `"stock": 5,

`            `"price": 10000000

`        `},

`        `{

`            `"ID": 6,

`            `"CreatedAt": "2023-06-08T20:42:39.881+07:00",

`            `"UpdatedAt": "2023-06-08T20:42:39.881+07:00",

`            `"DeletedAt": **null**,

`            `"category\_id": 1,

`            `"Category": {

`                `"ID": 1,

`                `"CreatedAt": "2023-06-08T19:55:26.548+07:00",

`                `"UpdatedAt": "2023-06-08T19:55:26.548+07:00",

`                `"DeletedAt": **null**,

`                `"category\_name": "elektronik"

`            `},

`            `"name": "lg",

`            `"description": "available",

`            `"stock": 5,

`            `"price": 5000000

`        `},

`        `{

`            `"ID": 7,

`            `"CreatedAt": "2023-06-08T20:43:19.253+07:00",

`            `"UpdatedAt": "2023-06-08T20:43:19.253+07:00",

`            `"DeletedAt": **null**,

`            `"category\_id": 1,

`            `"Category": {

`                `"ID": 1,

`                `"CreatedAt": "2023-06-08T19:55:26.548+07:00",

`                `"UpdatedAt": "2023-06-08T19:55:26.548+07:00",

`                `"DeletedAt": **null**,

`                `"category\_name": "elektronik"

`            `},

`            `"name": "sharp",

`            `"description": "available",

`            `"stock": 5,

`            `"price": 1000000

`        `}

`    `],

`    `"messages": "success get all products"

}

**GET ITEM BY ID**

Method: GET

Endpoint: /items/:id

Response:

{

`    `"items": {

`        `"ID": 4,

`        `"CreatedAt": "2023-06-08T20:01:55.852+07:00",

`        `"UpdatedAt": "2023-06-08T20:01:55.852+07:00",

`        `"DeletedAt": **null**,

`        `"category\_id": 1,

`        `"Category": {

`            `"ID": 1,

`            `"CreatedAt": "2023-06-08T19:55:26.548+07:00",

`            `"UpdatedAt": "2023-06-08T19:55:26.548+07:00",

`            `"DeletedAt": **null**,

`            `"category\_name": "elektronik"

`        `},

`        `"name": "samsung",

`        `"description": "available",

`        `"stock": 5,

`        `"price": 10000000

`    `},

`    `"messages": "success get item"

}

**GET ITEM BY ID CATEGORY**

Method: GET

Endpoint: /items/category/:id

Response:

{

`    `"items": [

`        `{

`            `"ID": 3,

`            `"CreatedAt": "2023-06-08T19:56:03.206+07:00",

`            `"UpdatedAt": "2023-06-08T19:56:03.206+07:00",

`            `"DeletedAt": **null**,

`            `"category\_id": 1,

`            `"Category": {

`                `"ID": 1,

`                `"CreatedAt": "2023-06-08T19:55:26.548+07:00",

`                `"UpdatedAt": "2023-06-08T19:55:26.548+07:00",

`                `"DeletedAt": **null**,

`                `"category\_name": "elektronik"

`            `},

`            `"name": "polyrton",

`            `"description": "available",

`            `"stock": 5,

`            `"price": 4000000

`        `},

`        `{

`            `"ID": 4,

`            `"CreatedAt": "2023-06-08T20:01:55.852+07:00",

`            `"UpdatedAt": "2023-06-08T20:01:55.852+07:00",

`            `"DeletedAt": **null**,

`            `"category\_id": 1,

`            `"Category": {

`                `"ID": 1,

`                `"CreatedAt": "2023-06-08T19:55:26.548+07:00",

`                `"UpdatedAt": "2023-06-08T19:55:26.548+07:00",

`                `"DeletedAt": **null**,

`                `"category\_name": "elektronik"

`            `},

`            `"name": "samsung",

`            `"description": "available",

`            `"stock": 5,

`            `"price": 10000000

`        `}

`    `],

`    `"messages": "success get item"

}

**UPDATE ITEM** 

Method: PUT

Endpoint: /items/:id

Request (Body):

{

`    `"category\_id": 1,

`    `"name": "panasonic",

`    `"description": "available",

`    `"stock" : 5,

`    `"price": 10000000

}

Response :

{

`    `"message": "success update item data",

`    `"user": {

`        `"ID": 3,

`        `"CreatedAt": "2023-06-08T19:56:03.206+07:00",

`        `"UpdatedAt": "2023-06-08T20:10:51.182+07:00",

`        `"DeletedAt": **null**,

`        `"category\_id": 1,

`        `"Category": {

`            `"ID": 1,

`            `"CreatedAt": "2023-06-08T19:55:26.548+07:00",

`            `"UpdatedAt": "2023-06-08T19:55:26.548+07:00",

`            `"DeletedAt": **null**,

`            `"category\_name": "elektronik"

`        `},

`        `"name": "panasonic",

`        `"description": "available",

`        `"stock": 5,

`        `"price": 10000000

`    `}

}

**DELETE ITEM** 

Method: DELETE

Endpoint: /items/:id

Response:

{

`    `"id": 3,

`    `"messages": "success delete product"

}










