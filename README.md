# technical_test_juke
Project ini adalah REST API sederhana menggunakan **Go (Gin)** dengan **MySQL**, dijalankan via **Docker Compose**. API ini menyediakan CRUD untuk tabel `employee`.

---

1. **Clone repository**
```bash
git clone https://github.com/<username>/technical_test_juke.git
cd technical_test_juke
```
2. **Bangun dan jalankan Docker**
```bash
docker compose up -d --build
```
Go API listen di port 8080 dengan akses http://localhost:8080
3. **Cek container**
```bash
docker ps
```
4. **setup database**
pilih database dan create
```bash
USE juke;
CREATE TABLE employee (
    id INT PRIMARY KEY AUTO_INCREMENT,
    name TEXT,
    email TEXT,
    position TEXT,
    salary DOUBLE,
    createdAt DATETIME DEFAULT CURRENT_TIMESTAMP
);
```

**Dokumentasi Swagger**
```bash
http://localhost:8080/swagger/index.html
```
**stop project**
```bash
docker compose down


