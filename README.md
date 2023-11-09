# dot-golang

**1. Install dependensi**

```shellscript
go mod tidy
```

**2. Buat file .env**

Masuk ke direktori projek lalu jalankan perintah berikut di terminal

```shellscript
cp env.example .env
```

Setelah itu, atur konfigurasi pada file **.env** agar aplikasi dapat berjalan dengan baik

**3. Cara menjalankan aplikasi**

```shellscript
go run cmd/main.go
```

## Design pattern

Design pattern yang digunakan adalah **Clean Architecture**. Dengan menggunakan design pattern tersebut maka aplikasi akan dipecah menjadi beberapa layer seperti service, domain, repository, dll. Pada folder **_internal_** berisi kodingan yang tidak memiliki ketergantungan dengan library pihak ketiga. Sedangkan pada folder **_external_** berisi kodingan yang memiliki ketergantungan dengan library pihak ketiga. Dengan adanya pemisahan tersebut maka kita bisa mengubah library pihak ketiga yang ada di folder _external_ tanpa mengubah kodingan pada folder _internal_. Selain itu, dengan adanya pemisahan layer maka proses pengujian dengan mock function pada kodingan yang ada di folder _internal_ menjadi lebih mudah karena tidak memiliki ketergantungan dengan library pihak ketiga.
