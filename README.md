# 🚀 Golang Application Development. Practice - 2

A simple and effective API for managing task lists, written in Go.

---

## 🛡 Authentication

All API requests must contain a special header for access rights verification.

| Header       | Value | Description      |
| :----------- |:------|:-----------------|
| `X-API-Key`  | `secret-key-123`    | Я умру в нищете. |

> **Error example (401 Unauthorized):** > `{"error": "unauthorized"}`

---

## 🛠 Endpoints

### 1. Get all tasks
- **URL:** `/tasks`
- **Method:** `GET`
- **Headers:** `X-API-Key: <key>`

### 2. Get task by ID
- **URL:** `/tasks/{id}`
- **Method:** `GET`
- **Headers:** `X-API-Key: <key>`

### 3. Create task
- **URL:** `/tasks`
- **Method:** `POST`
- **Headers:** `X-API-Key: <key>`
- **Body (JSON):**
  ```json
  {
    "title": "Buy milk"
  }

### 4. Update task status
- **URL:** `/tasks/{id}`
- **Method:** `PATCH`
- **Headers:** `X-API-Key: <key>`
- **Body (JSON):**
  ```json
  {
    "title": "Buy milk"
  }

---

### Request example (cURL)
 ```curl
curl -X GET http://localhost:8080/tasks \
     -H "X-API-Key: secret-key-123"
