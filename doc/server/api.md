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

