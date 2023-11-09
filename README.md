# dot-golang

**1. Cara menjalankan aplikasi**

```shellscript
docker compose up -d
```

**2. Akses api**

```shellscript
curl -I http://localhost:8001/blog/
```

**3. Cara menutup aplikasi**

```shellscript
docker compose down
```

## Design pattern

Design pattern yang digunakan adalah **Clean Architecture**. Dengan menggunakan design pattern tersebut maka aplikasi akan dipecah menjadi beberapa layer seperti service, domain, repository, dll. Pada folder **_internal_** berisi kodingan yang tidak memiliki ketergantungan dengan library pihak ketiga. Sedangkan pada folder **_external_** berisi kodingan yang memiliki ketergantungan dengan library pihak ketiga. Dengan adanya pemisahan tersebut maka kita bisa mengubah library pihak ketiga yang ada di folder _external_ tanpa mengubah kodingan pada folder _internal_. Selain itu, dengan adanya pemisahan layer maka proses pengujian dengan mock function pada kodingan yang ada di folder _internal_ menjadi lebih mudah karena tidak memiliki ketergantungan dengan library pihak ketiga.
