# E-Commerce REST API

REST API sederhana untuk manajemen produk, sumber barang (supplier), dan transaksi penjualan menggunakan Golang dan Gin framework.

## 🚀 Cara Menjalankan

```bash
# Install dependencies
go mod tidy

# Jalankan server
go run main.go
```

Server akan berjalan di `http://localhost:8080`

## 📋 Daftar Endpoint

### 🛍️ Product Endpoints

| Method | Endpoint | Deskripsi |
|--------|----------|-----------|
| GET | `/products` | Ambil semua produk |
| GET | `/products/:id` | Ambil produk berdasarkan ID |
| POST | `/products` | Tambah produk baru |
| PUT | `/products/:id` | Update produk |
| DELETE | `/products/:id` | Hapus produk |

### 🏪 Source Endpoints

| Method | Endpoint | Deskripsi |
|--------|----------|-----------|
| GET | `/sources` | Ambil semua source |
| GET | `/sources/:id` | Ambil source berdasarkan ID |
| POST | `/sources` | Tambah source baru |
| PUT | `/sources/:id` | Update source |
| DELETE | `/sources/:id` | Hapus source |

### 💳 Transaction Endpoints

| Method | Endpoint | Deskripsi |
|--------|----------|-----------|
| POST | `/transactions` | Buat transaksi baru |
| GET | `/transactions` | Ambil semua transaksi |
| GET | `/transactions/:id` | Ambil transaksi berdasarkan ID |

## 📊 Struktur Data

### Product
```json
{
  "id": "string",
  "name": "string",
  "description": "string",
  "price": 0,
  "stock": 0,
  "source_id": "string"
}
```

### Source
```json
{
  "id": "string",
  "name": "string"
}
```

### Transaction
```json
{
  "id": "string",
  "product_id": "string",
  "quantity": 0,
  "total": 0
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

### 1. Membuat Source Baru
```bash
curl -X POST http://localhost:8080/sources \
  -H "Content-Type: application/json" \
  -d '{
    "name": "Supplier C"
  }'
```

**Response:**
```json
{
  "message": "Source created successfully",
  "data": {
    "id": "3",
    "name": "Supplier C"
  },
  "error": null
}
```

### 2. Membuat Product Baru
```bash
curl -X POST http://localhost:8080/products \
  -H "Content-Type: application/json" \
  -d '{
    "name": "Keyboard",
    "description": "Mechanical keyboard",
    "price": 500000,
    "stock": 25,
    "source_id": "1"
  }'
```

**Response:**
```json
{
  "message": "Product created successfully",
  "data": {
    "id": "3",
    "name": "Keyboard",
    "description": "Mechanical keyboard",
    "price": 500000,
    "stock": 25,
    "source_id": "1"
  },
  "error": null
}
```

### 3. Membuat Transaksi
```bash
curl -X POST http://localhost:8080/transactions \
  -H "Content-Type: application/json" \
  -d '{
    "product_id": "1",
    "quantity": 2
  }'
```

**Response:**
```json
{
  "message": "Transaction created successfully",
  "data": {
    "id": "1",
    "product_id": "1",
    "quantity": 2,
    "total": 30000000
  },
  "error": null
}
```

### 4. Mengambil Semua Produk
```bash
curl -X GET http://localhost:8080/products
```

**Response:**
```json
{
  "message": "Products retrieved successfully",
  "data": [
    {
      "id": "1",
      "name": "Laptop",
      "description": "Gaming laptop",
      "price": 15000000,
      "stock": 8,
      "source_id": "1"
    },
    {
      "id": "2",
      "name": "Mouse",
      "description": "Wireless mouse",
      "price": 250000,
      "stock": 50,
      "source_id": "2"
    }
  ],
  "error": null
}
```

### 5. Filter Produk Berdasarkan Source
```bash
curl -X GET "http://localhost:8080/products?source_id=1"
```

### 6. Update Produk
```bash
curl -X PUT http://localhost:8080/products/1 \
  -H "Content-Type: application/json" \
  -d '{
    "name": "Gaming Laptop Updated",
    "description": "High-end gaming laptop",
    "price": 18000000,
    "stock": 5,
    "source_id": "1"
  }'
```

### 7. Hapus Produk
```bash
curl -X DELETE http://localhost:8080/products/1
```

## ✅ Validasi

### Product
- `name`: Tidak boleh kosong
- `price`: Harus lebih besar dari 0
- `stock`: Harus lebih besar atau sama dengan 0
- `source_id`: Harus ada di daftar source

### Source
- `name`: Tidak boleh kosong

### Transaction
- `quantity`: Harus lebih besar dari 0
- `product_id`: Harus ada di daftar produk
- Stock produk harus mencukupi

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
  "error": "Price must be greater than 0"
}
```

## 🔋 Fitur Bonus

1. **Filter Produk**: Query parameter `source_id` untuk filter produk berdasarkan source
2. **Validasi Lengkap**: Validasi untuk semua input
3. **Konsistensi Response**: Format response yang konsisten
4. **Logger Middleware**: Logging untuk setiap request

## 📁 Struktur Project

```
e-commerce/
├───products
├───source
├───transaction
└───users
```

## 💡 Catatan Penting

- Data disimpan di memory (tidak persisten)
- Saat transaksi dibuat, stock produk akan berkurang otomatis
- ID dihasilkan secara otomatis menggunakan counter
- Semua endpoint menggunakan format response yang konsisten
