# <a style="font-family:cursive">Hacktiv8 Final Project 3</a>
<p align="justify">Kanban Board is an application for project management. User in this application will be able to add their tasks to the categories provided by an admin.</p>

# Installation
Requires <b>[Golang](https://go.dev/dl/)</b> and <b>[MySQL](https://dev.mysql.com/downloads/installer/)</b>

Config the .env first to connect into database

- **Clone repository**
```bash
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

# Project Structure
> <i>Press to switch into the folder you want to go</i><br>
 
 📦[go-hacktiv8-Kanban-Board](https://github.com/mrafid01/go-hacktiv8-Kanban-Board)<br>
 ┣ 📂[config](https://github.com/mrafid01/go-hacktiv8-Kanban-Board/tree/master/config)<br>
 ┃ ┣ 📜[db.go](https://github.com/mrafid01/go-hacktiv8-Kanban-Board/blob/master/config/db.go)<br>
 ┃ ┗ 📜[db_test.go](https://github.com/mrafid01/go-hacktiv8-Kanban-Board/blob/master/config/db_test.go)<br>
 ┣ 📂[controller](https://github.com/mrafid01/go-hacktiv8-Kanban-Board/tree/master/controller)<br>
 ┃ ┣ 📜[category_controller.go](https://github.com/mrafid01/go-hacktiv8-Kanban-Board/blob/master/controller/category_controller.go)<br>
 ┃ ┣ 📜[task_controller.go](https://github.com/mrafid01/go-hacktiv8-Kanban-Board/blob/master/controller/task_controller.go)<br>
 ┃ ┗ 📜[user_controller.go](https://github.com/mrafid01/go-hacktiv8-Kanban-Board/blob/master/controller/user_controller.go)<br>
 ┣ 📂[helper](https://github.com/mrafid01/go-hacktiv8-Kanban-Board/tree/master/helper)<br>
 ┃ ┣ 📜[error.go](https://github.com/mrafid01/go-hacktiv8-Kanban-Board/blob/master/helper/error.go)<br>
 ┃ ┗ 📜[response.go](https://github.com/mrafid01/go-hacktiv8-Kanban-Board/blob/master/helper/response.go)<br>
 ┣ 📂[middleware](https://github.com/mrafid01/go-hacktiv8-Kanban-Board/tree/master/middleware)<br>
 ┃ ┣ 📜[jwt.go](https://github.com/mrafid01/go-hacktiv8-Kanban-Board/blob/master/middleware/jwt.go)<br>
 ┃ ┗ 📜[middleware.go](https://github.com/mrafid01/go-hacktiv8-Kanban-Board/blob/master/middleware/middleware.go)<br>
 ┣ 📂[model](https://github.com/mrafid01/go-hacktiv8-Kanban-Board/tree/master/model)<br>
 ┃ ┣ 📂[entity](https://github.com/mrafid01/go-hacktiv8-Kanban-Board/tree/master/model/entity)<br>
 ┃ ┃ ┣ 📜[category.go](https://github.com/mrafid01/go-hacktiv8-Kanban-Board/blob/master/model/entity/category.go)<br>
 ┃ ┃ ┣ 📜[task.go](https://github.com/mrafid01/go-hacktiv8-Kanban-Board/blob/master/model/entity/task.go)<br>
 ┃ ┃ ┗ 📜[user.go](https://github.com/mrafid01/go-hacktiv8-Kanban-Board/blob/master/model/entity/user.go)<br>
 ┃ ┣ 📂[input](https://github.com/mrafid01/go-hacktiv8-Kanban-Board/tree/master/model/input)<br>
 ┃ ┃ ┣ 📜[category_input.go](https://github.com/mrafid01/go-hacktiv8-Kanban-Board/blob/master/model/input/category_input.go)<br>
 ┃ ┃ ┣ 📜[task_input.go](https://github.com/mrafid01/go-hacktiv8-Kanban-Board/blob/master/model/input/task_input.go)<br>
 ┃ ┃ ┗ 📜[user_input.go](https://github.com/mrafid01/go-hacktiv8-Kanban-Board/blob/master/model/input/user_input.go)<br>
 ┃ ┗ 📂[response](https://github.com/mrafid01/go-hacktiv8-Kanban-Board/tree/master/model/response)<br>
 ┃ ┃ ┣ 📜[category_response.go](https://github.com/mrafid01/go-hacktiv8-Kanban-Board/blob/master/model/response/category_response.go)<br>
 ┃ ┃ ┣ 📜[task_response.go](https://github.com/mrafid01/go-hacktiv8-Kanban-Board/blob/master/model/response/task_response.go)<br>
 ┃ ┃ ┗ 📜[user_response.go](https://github.com/mrafid01/go-hacktiv8-Kanban-Board/blob/master/model/response/user_response.go)<br>
 ┣ 📂[repository](https://github.com/mrafid01/go-hacktiv8-Kanban-Board/tree/master/repository)<br>
 ┃ ┣ 📜[category_repository.go](https://github.com/mrafid01/go-hacktiv8-Kanban-Board/blob/master/repository/category_repository.go)<br>
 ┃ ┣ 📜[task_repository.go](https://github.com/mrafid01/go-hacktiv8-Kanban-Board/blob/master/repository/task_repository.go)<br>
 ┃ ┗ 📜[user_repository.go](https://github.com/mrafid01/go-hacktiv8-Kanban-Board/blob/master/repository/user_repository.go)<br>
 ┣ 📂[service](https://github.com/mrafid01/go-hacktiv8-Kanban-Board/tree/master/service)<br>
 ┃ ┣ 📜[category_service.go](https://github.com/mrafid01/go-hacktiv8-Kanban-Board/blob/master/service/category_service.go)<br>
 ┃ ┣ 📜[task_service.go](https://github.com/mrafid01/go-hacktiv8-Kanban-Board/blob/master/service/task_service.go)<br>
 ┃ ┗ 📜[user_service.go](https://github.com/mrafid01/go-hacktiv8-Kanban-Board/blob/master/service/user_service.go)<br>
 ┣ 📜[.env](https://github.com/mrafid01/go-hacktiv8-Kanban-Board/blob/master/.env)<br>
 ┣ 📜[go.mod](https://github.com/mrafid01/go-hacktiv8-Kanban-Board/blob/master/go.mod)<br>
 ┣ 📜[go.sum](https://github.com/mrafid01/go-hacktiv8-Kanban-Board/blob/master/go.sum)<br>
 ┣ 📜[Hacktiv8-KanbanBoard-Kelompok-7.postman_collection.json](https://documenter.getpostman.com/view/23401248/2s8YsnYcEk)<br>
 ┣ 📜[main.go](https://github.com/mrafid01/go-hacktiv8-Kanban-Board/blob/master/main.go)<br>
 ┗ 📜[README.md](https://github.com/mrafid01/go-hacktiv8-Kanban-Board/blob/master/README.md)<br>

## Postman Documentation Publish Version 🚀
**`LINK`** 
https://documenter.getpostman.com/view/23401248/2s8YsnYcEk

# Endpoint
## 1. User
### Create Admin Account
> <i>digunakan untuk membuat akun dengan role Admin.</i>
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
> <i>digunakan untuk membuat akun dengan role User member.</i>
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

### Login Account
> <i>digunakan untuk melakukan login atau autentikasi Member/Admin.</i>
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

### Update Account Data
> <i>digunakan untuk melakukan perubahan full name dan email akun.</i>
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
> <i>digunakan untuk melakukan penghapusan akun.</i>
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
        "message": "Your account has been successfully deleted"
  	}
  }
  ```
<p align="justify">Notes: Endpoint ini memerlukan proses autentikasi. Proses autentikasi wajib dilakukan dengan package/library <b>JsonWebToken</b>. Endpoint ini berguna untuk user menghapus akunnya.</p>


## 2. Categories
<p align="justify">Notes: Seluruh endpoint untuk mengakses endpoint categories <b>selain</b> method GET memerlukan proses autentikasi menggunakan package JsonWebToken dan memerlukan proses autorisasi. Autorisasi diperlukan karena yang boleh mengakses endpoint categories <b>selain</b> method GET adalah user dengan role admin.</p>

### Create Categories Type <i>(must an admin)</i>
> <i>digunakan untuk membuat tipe kategori.</i>
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
        "type": "string",
        "created_at": "date"
  	}
  }
  ```

### Get All Categories Type
> <i>digunakan untuk menampilkan semua tipe kategori.</i>
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

### Update Categories Type <i>(must an admin)</i>
> <i>digunakan untuk melakukan perubahan tipe kategori.</i>
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

### Delete Categories Type <i>(must an admin)</i>
> <i>digunakan untuk melakukan penghapusan tipe kategori.</i>
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

<p align="justify"><b>Notes: Seluruh endpoint untuk mengakses endpoint products memerlukan proses autentikasi menggunakan package JsonWebToken.</b></p>


## 3. Tasks
### Create Task
> <i>digunakan untuk membuat task baru.</i>
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
> <i>digunakan untuk menampilkan task.</i>
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

### Update Task Title & Description 
> <i>digunakan untuk melakukan pembaharuan task title dan task description.</i>
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
<p align="justify">Notes: Pada endpoint ini, perlu dilakukan proses autorisasi yang dimana user hanya boleh mengupdate task milikinya sendiri.</p>

### Update Task Status
> <i>digunakan untuk melakukan perubahan dari status task (true/false).</i>
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
<p align="justify">Notes: Pada endpoint ini, perlu dilakukan proses autorisasi yang dimana user hanya boleh mengupdate task milikinya sendiri.</p>

### Update Task Category
> <i>digunakan untuk melakukan perubahan dari kategori task.</i>
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
> <i>digunakan untuk melakukan penghapusan task.</i>
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
<p align="justify">Notes: Pada endpoint ini, perlu dilakukan proses autorisasi yang dimana user hanya boleh menghapus task miliknya sendiri.</p>
<br>

# Group 7
<pre style="font-family:verdana">

<p>1. <a href="https://github.com/alrico11"><b>Alrico Rizki Wibowo</b></a>&emsp; —&emsp;GLNG-KS04-017</p>
<p>2. <a href="https://github.com/rickyfazaa"><b>Ricky Khairul Faza</b></a>&emsp; —&emsp;GLNG-KS04-022</p>
<p>3. <a href="https://github.com/mrafid01"><b>Muhammad Rafid</b></a>&emsp; —&emsp;GLNG-KS04-024</p></pre>
