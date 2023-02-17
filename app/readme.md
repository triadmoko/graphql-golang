# PROJECT CONFIG
## Prerequisite
- Golang 1.19+
- GORM
- Postgresql
- Docker
- Docker Compose
- Google Workspace For Email
- Authentication JWT 
- Gqlgen 
- Gin Gonic 

## ENV
```env
PORT=

DB_HOST=
DB_USER=
DB_PASS=
DB_NAME=
DB_PORT=
SSLMODE=

JWT_SECRET=
CONTEXT_TIME_OUT=

CONFIG_SMTP_HOST=
CONFIG_SMTP_PORT=
CONFIG_AUTH_EMAIL=
CONFIG_AUTH_PASSWORD=
```

# GRAPHQL TEST CASE


Membuat API untuk aplikasi Blog yang memiliki fitur:

1. Login dan registrasi user dengan ketentuan:

  * Ketika user mendaftar, status user `Terdaftar` dan belum dapat melakukan login kemudian ada notifikasi email sebagai proses verifikasi registrasi kemudian status user menjadi `Aktif`
  
  * Proses autentikasi atau login menggunakan token `jwt`. Token `jwt` ini dijuga dipakai sebagai autentikasi setiap API yang diakses



