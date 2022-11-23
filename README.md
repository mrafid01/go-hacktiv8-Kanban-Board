# Hacktiv Final Project 3
<p align="justify">Kanban Board is an application for project management. User in this application will be able to add their tasks to the categories provided by an admin.</p>

# Installation
Requires [Golang](https://go.dev/dl/) and [MySQL](https://dev.mysql.com/downloads/installer/)

Config the .env first to connect into database

- **Clone repository**
```sh
git clone https://github.com/mrafid01/go-hacktiv8-Kanban-Board.git
```
- **Change directory**
```sh
cd go-hacktiv8-Kanban-Board
```
- **Run "main.go" file**
```sh
go run main.go
```

# Endpoint
## User
### Create Admin Account
- Method: **`POST`**
- Endpoint:
```
http://localhost:8080/users/admin
```
- Request:
    - Request Body:
    ```json
    {
        "full_name": "string",
        "email": "string",
        "password": "string"
    }
    ```
- Response Body:
  - Status: 201,
  - Message: "created",
  - Body:
  ```json
  {
    "data": {
        "id": "integer",
        "full_name": "string",
        "email": "string",
        "created_at": "date"
  	}
  }
  ```

### Create User Account
- Method: **`POST`**
- Endpoint:
```
http://localhost:8080/users/register
```
- Request:
    - Request Body:
    ```json
    {
        "full_name": "string",
        "email": "string",
        "password": "string"
    }
    ```
- Response Body:
  - Status: 201,
  - Message: "created",
  - Body:
  ```json
  {
    "data": {
        "id": "integer",
        "full_name": "string",
        "email": "string",
        "created_at": "date"
  	}
  }
  ```
<p align="justify">Notes: Untuk endpoint ini, role dari data user akan otomatis menjadi member. Boleh langsung diharcode di controllernya sebelum disimpan ke dalam database.</p>

### Login User
- Method: **`POST`**
- Endpoint:
```
http://localhost:8080/users/login
```
- Request:
    - Request Body:
    ```json
    {
        "email": "string",
        "password": "string"
    }
    ```
- Response Body:
  - Status: 200,
  - Message: "ok",
  - Body:
  ```json
  {
    "data": {
        "token": "jwt string"
  	}
  }
  ```
<p align="justify">Notes: Pada endpoint ini, wajib melakukan logika user login yang dimana harus melakukan pengecekan email dan password user. Pengecekan password wajib dilakukan dengan bantuan library/package <b>Bcrypt</b>.</p>

### Update Data User
- Method: **`PUT`**
- Endpoint:
```
http://localhost:8080/users/update-account
```
- Request:
    - Headers: Authorization (Bearer Token)
    - Request Body:
    ```json
    {
        "full_name": "string",
        "email": "string"
    }
    ```
- Response Body:
  - Status: 200,
  - Message: "ok",
  - Body:
  ```json
  {
    "data": {
        "id": "integer",
        "full_name": "string",
        "email": "string",
        "updated_at": "date"
  	}
  }
  ```
<p align="justify">Notes: Endpoint ini memerlukan proses autentikasi. Proses autentikasi wajib dilakukan dengan package/library <b>JsonWebToken</b>. Endpoint ini berguna untuk user mengupdate data dirinya.</p>

### Delete Account
- Method: **`DELETE`**
- Endpoint:
```
http://localhost:8080/users/delete-account
```
- Request:
    - Headers: Authorization (Bearer Token)
- Response Body:
  - Status: 200,
  - Message: "ok",
  - Body:
  ```json
  {
    "data": {
        "id": "integer",
        "type": "string",
        "created_at": "date"
  	}
  }
  ```
<p align="justify">Notes: Endpoint ini memerlukan proses autentikasi. Proses autentikasi wajib dilakukan dengan package/library <b>JsonWebToken</b>. Endpoint ini berguna untuk user menghapus akunnya.</p>


## Categories
<p align="justify">Notes: Seluruh endpoint untuk mengakses endpoint categories <b>selain</b> method GET memerlukan proses autentikasi menggunakan package JsonWebToken dan memerlukan proses autorisasi. Autorisasi diperlukan karena yang boleh mengakses endpoint categories <b>selain</b> method GET adalah user dengan role admin.</p>

### Create Categories Type (must an admin)
- Method: **`POST`**
- Endpoint:
```
http://localhost:8080/categories
```
- Request:
    - Headers: Authorization (Bearer Token)
    - Request Body:
    ```json
    {
        "type": "string"
    }
    ```
- Response Body:
  - Status: 201,
  - Message: "created",
  - Body:
  ```json
  {
    "data": {
        "id": "integer",
        "full_name": "string",
        "email": "string",
        "created_at": "date"
  	}
  }
  ```

### Get All Categories Type
- Method: **`GET`**
- Endpoint:
```
http://localhost:8080/categories
```
- Request:
    - Headers: Authorization (Bearer Token)
- Response Body:
  - Status: 200,
  - Message: "ok",
  - Body:
  ```json
  {
    "data": {
        "id": "integer",
        "type": "string",
        "updated_at": "date",
        "created_at": "date",
        "Tasks": {
            "id": "integer",
            "title": "string",
            "description": "string",
            "user_id": "integer",
            "category_id": "integer",
            "created_at": "date",
            "updated_at": "date"
                }
            }
  }
  ```

### Update Categories Type (must an admin)
- Method: **`PATCH`**
- Endpoint:
```
http://localhost:8080//categories/:categoryId
```
- Request:
    - Headers: Authorization (Bearer Token)
    - Params: categoryId (integer)
    - Request Body:
    ```json
    {
        "type": "string"
    }
    ```
- Response Body:
  - Status: 200,
  - Message: "ok",
  - Body:
  ```json
  {
    "data": {
        "id": "integer",
        "type": "string",
        "updated_at": "date"
  	}
  }
  ```

### Delete Categories Type (must an admin)
- Method: **`DELETE`**
- Endpoint:
```
http://localhost:8080//categories/:categoryId
```
- Request:
    - Headers: Authorization (Bearer Token)
    - Params: categoryId (integer)
- Response Body:
  - Status: 200,
  - Message: "ok",
  - Body:
  ```json
  {
    "data": {
        "message": "Category has been successfully deleted"
  	}
  }
  ```

**Notes: Seluruh endpoint untuk mengakses endpoint products memerlukan proses autentikasi menggunakan package JsonWebToken.**


## Tasks
### Create Task
- Method: **`POST`**
- Endpoint:
```
http://localhost:8080/tasks
```
- Request:
    - Headers: Authorization (Bearer Token)
    - Request Body:
    ```json
    {
        "title": "string",
        "description": "string",
        "category_id": "integer"
    }
    ```
- Response Body:
  - Status: 201,
  - Message: "created",
  - Body:
  ```json
  {
    "data": {
        "id": "integer",
        "title": "string",
        "status": "boolean",
        "description": "string",
        "user_id": "integer",
        "category_id": "integer",
        "created_at": "date"
  	}
  }
  ```
<p align="justify">Notes: Pada endpoint ini, harus dilakukan pengecekkan jika data category dengan <b>id</b> yang diberikan pada request body dengan field <b>categoryId</b> ada atau tidak pada database. Jika ada maka boleh disimpan ke database namun jika tidak ada maka harus melempar error. Kemudian untuk field status nilai awalnya akan otomatis menjadi false. Boleh langsung di hardcode di controllernya.</p>

### Get Tasks
- Method: **`GET`**
- Endpoint:
```
http://localhost:8080/tasks
```
- Request:
    - Headers: Authorization (Bearer Token)
- Response Body:
  - Status: 200,
  - Message: "ok",
  - Body:
  ```json
  {
    "data": {
        "id": "integer",
        "title": "string",
        "status": "boolean",
        "description": "string",
        "user_id": "integer",
        "category_id": "integer",
        "created_at": "date",
        "User": {
                "id": "integer",
                "email": "string",
                "full_name": "string"
                }
  	  }
  }
  ```

### Update Title & Description Task
- Method: **`PUT`**
- Endpoint:
```
http://localhost:8080/tasks/:taskID
```
- Request:
    - Headers: Authorization (Bearer Token)
    - Params: taskId (integer)
    - Request Body:
    ```json
    {
        "title": "string",
        "description": "string"
    }
    ```
- Response Body:
  - Status: 200,
  - Message: "ok",
  - Body:
  ```json
  {
    "data": {
        "id": "integer",
        "title": "string",
        "description": "string",
        "status": "boolean",
        "user_id": "integer",
        "category_id": "integer",
        "updated_at": "date"
  	}
  }
  ```
Notes: Pada endpoint ini, perlu dilakukan proses autorisasi yang dimana user hanya boleh mengupdate task milikinya sendiri.

### Update Status Task
- Method: **`PATCH`**
- Endpoint:
```
http://localhost:8080/tasks/update-status/:taskID
```
- Request:
    - Headers: Authorization (Bearer Token)
    - Params: taskId (integer)
    - Request Body:
    ```json
    {
        "status": "boolean"
    }
    ```
- Response Body:
  - Status: 200,
  - Message: "ok",
  - Body:
  ```json
  {
    "data": {
        "id": "integer",
        "title": "string",
        "description": "string",
        "status": "boolean",
        "user_id": "integer",
        "category_id": "integer",
        "updated_at": "date"
  	}
  }
  ```
Notes: Pada endpoint ini, perlu dilakukan proses autorisasi yang dimana user hanya boleh mengupdate task milikinya sendiri.

### Update Category Task
- Method: **`PATCH`**
- Endpoint:
```
http://localhost:8080/tasks/update-category/:taskID
```
- Request:
    - Headers: Authorization (Bearer Token)
    - Params: taskId (integer)
    - Request Body:
    ```json
    {
        "category_id": "integer"
    }
    ```
- Response Body:
  - Status: 200,
  - Message: "ok",
  - Body:
  ```json
  {
    "data": {
        "id": "integer",
        "title": "string",
        "description": "string",
        "status": "boolean",
        "user_id": "integer",
        "category_id": "integer",
        "updated_at": "date"
  	}
  }
  ```
<p align="justify">Notes: Pada endpoint ini, perlu dilakukan proses autorisasi yang dimana user hanya boleh mengupdate category id dari task milikinya sendiri. Lalu perlu dilakukan pengecekan jika category dengan id yang dinput oleh user ada di dalam database, maka proses dapat dilanjut. Namun jika tidak ada maka akan langsung melempar error.</p>

### Delete Task
- Method: **`DELETE`**
- Endpoint:
```
http://localhost:8080/tasks/:taskID
```
- Request:
    - Headers: Authorization (Bearer Token)
    - Params: taskId (integer)
- Response Body:
  - Status: 200,
  - Message: "ok",
  - Body:
  ```json
  {
    "data": {
        "message": "Task has been successfully deleted"
  	}
  }
  ```
Notes: Pada endpoint ini, perlu dilakukan proses autorisasi yang dimana user hanya boleh menghapus task miliknya sendiri.


# Group 7
1. **[Alrico Rizki Wibowo](https://github.com/alrico11)**   - GLNG-KS04-017
2. **[Ricky Khairul Faza](https://github.com/rickyfazaa)**  - GLNG-KS04-022
3. **[Muhammad Rafid](https://github.com/mrafid01)**  - GLNG-KS04-024
