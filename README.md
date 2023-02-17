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
- GoMigrate

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

    ![register](assets/register.png)
    
    **Register User**

    ![Code Verification](assets/verification_email.png)
    
    **Code Verification**

    ![Activation User](assets/activation_user.png)
    
    **Activation User**

  
  * Proses autentikasi atau login menggunakan token `jwt`. Token `jwt` ini dijuga dipakai sebagai autentikasi setiap API yang diakses

    ![Login](assets/login.png)
   
    **Login**

2. User yang terdaftar dapat membuat posting, edit dan menghapus posting.

    ![Create Post](assets/create_post.png)

    **Create Post**

    ![Create Post Failed](assets/create_post_failed.png)
    
    **Create Post Failed**
    
    ![Update Post](assets/update_post.png)
    
    **Update Post**
    
    ![Delete Post](assets/delete_post.png)
    
    **Delete Post**


3. Terdapat fitur komentar disetiap posting, yang dapat memberi komentar hanya user yang terdaftar. Fitur komentar tersebut dapat dimoderasi atau tidak tergantung user menentukan apakah pada posting tersebut dapat dimoderasi atau tanpa moderasi. Fitur komentar dapat dikomentari ulang dengan kedalaman 1;

    ![Create Comment](assets/create_comment.png)

    **Create Comment**
