# Social Media REST API

REST API sederhana untuk simulasi aplikasi media sosial dengan fitur user, post, dan like menggunakan Golang dan Gin framework.

## 🚀 Cara Menjalankan

```bash
# Masuk ke folder users
cd users

# Jalankan server
go run social-media-api.go
```

Server akan berjalan di `http://localhost:8080`

## 📋 Daftar Endpoint

### 👤 User Endpoints

| Method | Endpoint | Deskripsi |
|--------|----------|-----------|
| POST | `/users` | Registrasi user baru |
| GET | `/users` | Ambil semua user |
| GET | `/users/:id` | Ambil profil user |
| PUT | `/users/:id` | Update profil user |
| DELETE | `/users/:id` | Hapus user |

### 📝 Post Endpoints

| Method | Endpoint | Deskripsi |
|--------|----------|-----------|
| POST | `/posts` | Buat postingan baru |
| GET | `/posts` | Ambil semua post |
| GET | `/posts/:id` | Ambil detail post |
| GET | `/users/:id/posts` | Ambil semua post dari satu user |
| DELETE | `/posts/:id` | Hapus post |

### ❤️ Like Endpoints

| Method | Endpoint | Deskripsi |
|--------|----------|-----------|
| POST | `/likes` | User menyukai post |
| GET | `/posts/:id/likes` | Lihat siapa saja yang like post tertentu |
| GET | `/users/:id/likes` | Lihat semua like dari seorang user |

## 📊 Struktur Data

### User
```json
{
  "id": "string",
  "username": "string",
  "email": "string",
  "bio": "string"
}
```

### Post
```json
{
  "id": "string",
  "user_id": "string",
  "content": "string",
  "created_at": "string"
}
```

### Like
```json
{
  "id": "string",
  "user_id": "string",
  "post_id": "string"
}
```

## 📝 Format Response

Semua response menggunakan format yang konsisten:

```json
{
  "message": "string",
  "data": {},
  "error": null
}
```

## 🔧 Contoh Penggunaan

### 1. Registrasi User Baru
```bash
curl -X POST http://localhost:8080/users \
  -H "Content-Type: application/json" \
  -d '{
    "username": "alice_wonder",
    "email": "alice@example.com",
    "bio": "Exploring wonderland"
  }'
```

**Response:**
```json
{
  "message": "User created successfully",
  "data": {
    "id": "3",
    "username": "alice_wonder",
    "email": "alice@example.com",
    "bio": "Exploring wonderland"
  },
  "error": null
}
```

### 2. Buat Post Baru
```bash
curl -X POST http://localhost:8080/posts \
  -H "Content-Type: application/json" \
  -d '{
    "user_id": "1",
    "content": "Just learned about REST APIs!"
  }'
```

**Response:**
```json
{
  "message": "Post created successfully",
  "data": {
    "id": "3",
    "user_id": "1",
    "content": "Just learned about REST APIs!",
    "created_at": "2025-07-15T10:30:00Z"
  },
  "error": null
}
```

### 3. Like Post
```bash
curl -X POST http://localhost:8080/likes \
  -H "Content-Type: application/json" \
  -d '{
    "user_id": "2",
    "post_id": "1"
  }'
```

**Response:**
```json
{
  "message": "Like created successfully",
  "data": {
    "id": "1",
    "user_id": "2",
    "post_id": "1"
  },
  "error": null
}
```

### 4. Ambil Semua User
```bash
curl -X GET http://localhost:8080/users
```

**Response:**
```json
{
  "message": "Users retrieved successfully",
  "data": [
    {
      "id": "1",
      "username": "john_doe",
      "email": "john@example.com",
      "bio": "Hello world!"
    },
    {
      "id": "2",
      "username": "jane_smith",
      "email": "jane@example.com",
      "bio": "Love coding!"
    }
  ],
  "error": null
}
```

### 5. Ambil Post dari User Tertentu
```bash
curl -X GET http://localhost:8080/users/1/posts
```

### 6. Ambil Like dari Post Tertentu
```bash
curl -X GET http://localhost:8080/posts/1/likes
```

### 7. Update User
```bash
curl -X PUT http://localhost:8080/users/1 \
  -H "Content-Type: application/json" \
  -d '{
    "username": "john_doe_updated",
    "email": "john.updated@example.com",
    "bio": "Updated bio!"
  }'
```

### 8. Hapus Post
```bash
curl -X DELETE http://localhost:8080/posts/1
```

## ✅ Validasi

### User
- `username`: Tidak boleh kosong dan harus unik
- `email`: Tidak boleh kosong dan harus unik

### Post
- `content`: Tidak boleh kosong
- `user_id`: Harus ada di daftar user

### Like
- `user_id`: Harus ada di daftar user
- `post_id`: Harus ada di daftar post
- Satu user hanya boleh like satu post satu kali

## 🚨 Error Handling

### Status Code
- `200`: Success
- `201`: Created
- `400`: Bad Request (validation error)
- `404`: Not Found
- `500`: Internal Server Error

### Contoh Error Response
```json
{
  "message": "Validation failed",
  "data": null,
  "error": "Username already exists"
}
```

## 🔋 Fitur Khusus

1. **Username & Email Unik**: Sistem memastikan tidak ada duplikasi
2. **Cascade Delete**: Saat user dihapus, semua post dan like-nya juga terhapus
3. **Double Like Prevention**: User tidak bisa like post yang sama dua kali
4. **Timestamp**: Setiap post memiliki timestamp otomatis
5. **Referential Integrity**: Validasi foreign key untuk user_id dan post_id

## 💡 Catatan Penting

- Data disimpan di memory (tidak persisten)
- ID dihasilkan secara otomatis menggunakan counter
- Created timestamp menggunakan format RFC3339
- Semua endpoint menggunakan format response yang konsisten
- Penghapusan user akan menghapus semua post dan like terkait
