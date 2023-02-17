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

     ![Delete Comment](assets/delete_comment.png)

    **Delete Comment**

    Comment dapat di delete ketika komentar tersebut adalaha milik user atau postingan tersebut milik user tersebut

     ![Update Comment](assets/update_comment.png)

    **Update Comment**

4. Terdapat fitur `like` pada setiap posting. Yang dapat melakukan `like` adalah user yang terdaftar.

     ![Like](assets/like_unlike.png)

    **Like and Unlike**
    Fitur like dan unlike saat pertama kali di hit akan menjadi `like` ketika di hit kedua kalinya akan otomatis menjadi `unlike` dan begitu seterusnya 

     ![Like](assets/total_like.png)

    **Total Like**

5. Terdapat fitur pemberitahuan setiap kali ada user lain yang like dan komentar pada postingan user. Notifikasi dikirmkan ke email.

    ![Mail Comment](assets/send_mail_comment.png)

    **Send Mail Comment**

     ![Mail Comment](assets/send_mail_like.png)

    **Send Mail Like**

6. Jika user `A` misalnya mengomentari postingan user lain, kemudian komentar tersebut dikomentari user lainnya, maka user `A` mendapatkan notifikasi email bahwa komentarnya dikomentari. 