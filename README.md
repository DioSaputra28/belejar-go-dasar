# Belajar Go Dasar - REST API Project

Project ini adalah implementasi REST API sederhana menggunakan Go (Golang) tanpa framework (menggunakan `net/http` standard library), namun dilengkapi dengan dokumentasi Swagger.

## Fitur

- **Manajemen Produk**: CRUD (Create, Read, Update, Delete) untuk produk.
- **Manajemen Kategori**: CRUD untuk kategori produk.
- **Swagger UI**: Dokumentasi API interaktif.
- **Health Check**: Endpoint untuk memantau status server.

## Teknologi yang Digunakan

- **Go**: Bahasa pemrograman utama.
- **net/http**: Standard library untuk HTTP server.
- **Swaggo (Swag)**: Library untuk generate Swagger documentation.

## Cara Menjalankan

1.  **Clone Repository** (jika belum):

    ```bash
    git clone <repository-url>
    cd belajar-go-dasar
    ```

2.  **Install Dependencies**:

    ```bash
    go mod tidy
    ```

3.  **Jalankan Server**:
    ```bash
    go run cmd/main.go
    ```
    Output:
    ```
    Server running di localhost:8080
    ```

## Dokumentasi API

Seluruh dokumentasi API dapat diakses secara interaktif melalui Swagger UI:

**http://localhost:8080/swagger/index.html**

### Daftar Endpoint

Berikut adalah ringkasan endpoint yang tersedia:

#### Produk (`/api/produk`)

| Method   | Endpoint           | Deskripsi                                |
| :------- | :----------------- | :--------------------------------------- |
| `GET`    | `/api/produk`      | Mendapatkan daftar semua produk          |
| `POST`   | `/api/produk`      | Menambahkan produk baru                  |
| `GET`    | `/api/produk/{id}` | Mendapatkan detail produk berdasarkan ID |
| `PUT`    | `/api/produk/{id}` | Mengupdate data produk berdasarkan ID    |
| `DELETE` | `/api/produk/{id}` | Menghapus produk berdasarkan ID          |

#### Kategori (`/api/category`)

| Method   | Endpoint             | Deskripsi                                  |
| :------- | :------------------- | :----------------------------------------- |
| `GET`    | `/api/category`      | Mendapatkan daftar semua kategori          |
| `POST`   | `/api/category`      | Menambahkan kategori baru                  |
| `GET`    | `/api/category/{id}` | Mendapatkan detail kategori berdasarkan ID |
| `PUT`    | `/api/category/{id}` | Mengupdate data kategori berdasarkan ID    |
| `DELETE` | `/api/category/{id}` | Menghapus kategori berdasarkan ID          |

#### Utilitas

| Method | Endpoint              | Deskripsi                                  |
| :----- | :-------------------- | :----------------------------------------- |
| `GET`  | `/health`             | Cek status kesehatan server (Health Check) |
| `GET`  | `/swagger/index.html` | Membuka Swagger UI                         |

## Update Dokumentasi Swagger

Jika Anda melakukan perubahan pada kode API (menambah endpoint/mengubah model) dan ingin mengupdate dokumentasi Swagger:

1.  Pastikan `swag` sudah terinstall:
    ```bash
    go install github.com/swaggo/swag/cmd/swag@latest
    ```
2.  Jalankan perintah generate di root project:
    ```bash
    swag init -g cmd/main.go --parseDependency --parseInternal
    ```
    Ini akan memperbarui folder `docs/`.
