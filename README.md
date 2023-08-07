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

<hr>

### Documentation

#### List Entities 
1. [User](#User)
2. [Photo](#Photo)

#### User
1. ##### Register User ```POST METHOD```
    endpoint : ```/auth/register```    
    json request body : 
    ```
    {
    "username": "userme",
    "password": "pass123",
    "email": "userme@gmail.com"
    }
    ```
    json response : 
    ```
    {
    "message": "successfully register user",
    "status": 201
    }
    ```
    json response missing value : 
    ```
    {
    "error": "UserAuth.password: non zero value required;UserAuth.username: non zero value required;email: non zero value required",
    "message": "bad body request",
    "status": 400
    }
    ```
    json response validation password : 
    ```
    {
    "error": "UserAuth.password: pass1 does not validate as stringlength(6|255)",
    "message": "bad body request",
    "status": 400
    }
    ```
    json response validation conflict : 
    ```
    {
    "message": "username or email is already used",
    "status": 409
    }
    ```
2. ##### Login User ```GET METHOD```
    endpoint : ```/auth/login```    
    json request body : 
    ```
    {
    "username": "dvnf",
    "password": "pass123",
    }
    ```
    json response : 
    ```
    {
    "message": "user successfully login",
    "status": 202
    }
    ```
3. ##### Logout User ```GET METHOD```
    endpoint : ```/auth/logout```   
    json response : 
    ```
    {
    "message": "successfully logout",
    "status": 200,
    }
    ```
4. ##### Update User ```PUT METHOD```
    endpoint : ```/users```    
    json request body : 
    ```
    {
    "username": "dvnf2",
    "password": "pass12345",
    "email": "dvnf@gmail.com"
    }
    ```
    json response : 
    ```
    {
    "message": "successfully update data",
    "status": 200
    }
    ```
5. ##### Delete User ```DELETE METHOD```
    endpoint : ```/users```    
    json response : 
    ```
    {
    "message": "successfully delete data",
    "status": 200
    }
    ```

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

3. ##### Get Photo By ID ```POST METHOD```
    endpoint : ```/photos/:id```    
    json response : 
    ```
    {
    "data": {
        "ID": 1,
        "CreatedAt": "2023-08-07T10:02:07.666Z",
        "UpdatedAt": "2023-08-07T10:02:07.666Z",
        "DeletedAt": null,
        "user_id": 1,
        "title": "how to be software engineer",
        "caption": "grind codeforces, leetocde lul",
        "photo_url": "letsgo"
    },
    "message": "successfully fetch photo",
    "status": 200
    }
    ```

4. ##### Create New Photo ```POST METHOD```
    endpoint : ```/photos```    
    json request body :
    ```
    {
    "title": "how to be software engineer",
    "caption": "grind codeforces, leetocde lul",
    "photo_url": "letsgo"
    }
    ```
    json response : 
    ```
    {
    "data_record": {
        "ID": 3,
        "CreatedAt": "2023-08-07T20:57:08.104+07:00",
        "UpdatedAt": "2023-08-07T20:57:08.104+07:00",
        "DeletedAt": null,
        "user_id": 1,
        "title": "how to be software engineer",
        "caption": "grind codeforces, leetocde lul",
        "photo_url": "letsgo"
    },
    "message": "successfully create photo",
    "status": 201
    }
    ```
    json response if fields missing :
    ```
    {
    "error": "caption: non zero value required;photo_url: non zero value required;title: non zero value required",
    "message": "bad body request",
    "status": 400
    }
    ```

5. ##### Update Photo ```PUT METHOD```
    endpoint : ```photos/10```    
    json response : 
    ```
    {
    "data_record": {
        "ID": 0,
        "CreatedAt": "0001-01-01T00:00:00Z",
        "UpdatedAt": "2023-08-07T21:21:01.169+07:00",
        "DeletedAt": null,
        "user_id": 0,
        "title": "competitive 1",
        "caption": "grind codeforces, leetocde tlx lul",
        "photo_url": "letsgo"
    },
    "message": "successfully update data",
    "status": 200
    }
    ```
    json response failed update : 
    ```
    {
    "message": "failed to update data",
    "status": 400
    }
    ```

6. ##### Delete Photo ```DELETE METHOD```
    endpoint : ```/photos/:id```    
    json response : 
    ```
    {
    "message": "successfully delete data",
    "status": 200
    }
    ```