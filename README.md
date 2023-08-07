## Task Terakhir VIX Fullstack Developer BTPNS

### Getting Started With Go Backend App
1. Lakukan konfigurasi file ```.env``` sesuai file ```.env.example```
2. Pastikan sudah melakukan instalasi go dengan menjalankan command ```go version```
3. Jalankan command ```go run main.go```

### Tech Stack
1. Go
2. Gin
3. Gorm
4. JWT
5. MySQL

### Documentation

#### List Entities 
1. [User](#User)
2. [Photo](#Photo)

#### User

#### Photo
1. ##### Get All Photos ```GET METHOD```
    endpoint : ```/photos```   
    json response : 
    ``` 
    {
    "data": [
        {
            "id": 1,
            "title": "how to be software engineer",
            "caption": "grind codeforces, leetocde lul",
            "photo_url": "letsgo",
            "user_id": 1
        },
        {
            "id": 2,
            "title": "competitive programmer 30",
            "caption": "grind codeforces, leetocde tlx lul",
            "photo_url": "letsgo",
            "user_id": 1
        },
    ],
    "message": "successfully fetch photos",
    "status": 200
    }
    ```

2. ##### Get Photos By User ID ```GET METHOD```
    endpoint : ```/photos?user={number}```    
    json response : 
    ``` 
    {
    "data": [
        {
            "id": 1,
            "title": "how to be software engineer",
            "caption": "grind codeforces, leetocde lul",
            "photo_url": "letsgo",
            "user_id": 1
        },
        {
            "id": 2,
            "title": "competitive programmer 30",
            "caption": "grind codeforces, leetocde tlx lul",
            "photo_url": "letsgo",
            "user_id": 1
        },
    ],
    "message": "successfully fetch photos",
    "status": 200
    }
    ```