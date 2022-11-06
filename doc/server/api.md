# server api document

- GET    /health
- GET    /community/:community_id
- PUT    /community/:community_id
- DELETE /community/:community_id
- POST   /community/
- GET    /user/:id
- GET    /user/logout
- PUT    /user/update
- DELETE /user/:id
- POST   /user/signup
- POST   /user/login
- GET    /community/:community_id/recipe
- POST   /community/:community_id/recipe
- POST   /community/:community_id/recipe/step
- POST   /community/:community_id/recipe/spice
- POST   /post/create
- PUT    /post/update
- GET    /post/data/:id
- GET    /post/img/:id
- DELETE /post/:id

## GET    /health
ヘルスチェック

response

```json
{
    "health": "good"
}
```

<br>

## GET    /community/:community_id

指定したidのコミュニティを取得する

response

```json
{
    "data": {
        "id": 1,
        "name": "c1",
        "img_url": "https://hoge.png",
        "content": "hello!",
        "Recipe": null,
        "User": null
    }
}
```

<br>

## PUT    /community/:community_id

指定したidの情報を更新する

request body/json

```json
{
    "name": "c2",
    "img_url": "https://hoge.png",
    "content": "hello!"
}
```

response

```json
{
    "data": {
        "id": 1,
        "name": "c2",
        "img_url": "https://hoge.png",
        "content": "hello!",
        "Recipe": null,
        "User": null
    }
}
```

## DELETE /community/:community_id
指定したidのコミュニティを削除する

rsponse

削除できてもできなくても以下のレスポンスがでる

```json
{
    "ok": "success delete community ( id : 1 )"
}
```

## POST /community
新しいcomunityの作成をする

request body/json

```json
{
    "name": "hoge",
    "img_url": "https://hoge.png",
    "content": "hello!"
}
```

```json
{
    "data": {
        "id": 6,
        "name": "hoge",
        "img_url": "https://hoge.png",
        "content": "hello!",
        "Recipe": null,
        "User": null
    }
}
```

<br>

## GET /user/:id

指定したidのユーザーを取得する

response

```json
{
    "data": {
        "id": 1,
        "name": "hoge",
        "mail": "hoge@mail.com",
        "password": ""
    }
}
```


- GET    /user/logout
- PUT    /user/update
- DELETE /user/:id
- POST   /user/login

## POST /user/signup
ユーザー登録をする

request body/json
```json
{
    "name":"hoge",
    "mail":"hoge@mail.com",
    "password":"hoge"
}
```

response
```json
{
    "data": {
        "id": 1,
        "name": "hoge",
        "mail": "hoge@mail.com",
        "password": ""
    }
}
```

```
_cookie hogehogehogehogehogehogeho
```

## POST /user/login
ログインをする

request body/json
```json
{
    "name":"hoge",
    "mail":"hoge@mail.com",
    "password":"hoge"
}
```

response
```json
{
    "login": "ok"
}
```

```
_cookie hogehogehogehogehogehogeho
```

## GET /user/logout
ログアウトをする

response
```json
{
    "logout": "ok"
}
```

## PUT /user/update
ユーザー情報を更新する

request body/json
```json
{
    "id": 1,
    "name":"hoge",
    "mail":"hoge@mail.com",
    "password":"hoge"
}
```

response
```json
{
    "data": {
        "id": 1,
        "name": "hoge",
        "mail": "hoge@mail.com",
        "password": ""
    }
}
```


## DELETE /user/:id

指定したidのユーザーを削除する

response

```json
{
    "delete": "ok"
}
```
<br>

## GET /community/:community_id/recipe
指定したコミュニティのレシピ情報を取得する

response

```json
{
    "data": {
        "id": 1,
        "name": "ramen",
        "img_url": "https://hoge.png",
        "recipe_steps": [
            {
                "id": 1,
                "number": 1,
                "content": "step1 : hoge"
            },
            {
                "id": 2,
                "number": 2,
                "content": "step2 : fuga"
            },
            {
                "id": 3,
                "number": 3,
                "content": "step13: piyo"
            }
        ],
        "spices": [
            {
                "id": 1,
                "name": "sato"
            },
            {
                "id": 2,
                "name": "sio"
            },
            {
                "id": 3,
                "name": "su"
            },
            {
                "id": 4,
                "name": "shoyu"
            },
            {
                "id": 5,
                "name": "miso"
            }
        ]
    }
}
```



## POST   /community/:community_id/recipe
指定したidのコミュニティのレシピ情報を登録します
レシピのステップや調味料は後から登録します

request body/json

```json
{
    "data": {
        "id": 2,
        "name": "ramen",
        "img_url": "https://hoge.png",
        "recipe_steps": null,
        "spices": null
    }
}
```

response

```json
{
    "data": {
        "id": 2,
        "name": "ramen",
        "img_url": "https://hoge.png",
        "recipe_steps": null,
        "spices": null
    }
}
```

<br>

## POST   /community/:community_id/recipe/step

request body/json
```json
{
    "number": 1,
    "content": "step1 : hoge"
}
```

response
```json
{
    "data": {
        "id": 4,
        "number": 1,
        "content": "step1 : hoge"
    }
}
```

## POST   /community/:community_id/recipe/spice

request body/json
```json
{
    "name": "hoge"
}
```

response
```json
{
    "data": {
        "id": 6,
        "name": "hoge"
    }
}
```

<br>


## POST /post/create
投稿を登録をする

request body/json
```
{
    "content":"hoge",
    "user_id":1,
    "community_id":1,
    "img":hogehoge.jpg,
}
```

response
```json
画像
```

## PUT /post/update
投稿の情報を更新する

request body/json
```
{
    "id":1,
    "content":"hoge",
    "user_id":1,
    "community_id":1,
    "img":hogehoge.jpg,
}
```

response
```json
画像
```

## GET /post/img/:id
idから投稿の画像を取得する


response
```json
画像
```

## GET /post/data/:id
idから投稿のテキストなどのデータを取得する


response
```json
{
    "data": {
        "id": 1,
        "community_id": 1,
        "img": null,
        "content": "test",
        "user_id": 2,
        "User": null
    }
}
```


## DELETE /post/:id
投稿を削除する


response
```json
{
   "ok": "success delete community ( id : :id )"
}
```