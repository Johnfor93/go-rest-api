# RESTful API Contact

Aplikasi ini merupakan aplikasi dengan RESTful API untuk data kontak. Aplikasi ini dibangun dengan menggunakan bahasa Golang serta database menggunakan MYSQL. Dengan aplikasi ini, URI yang bisa diakses yaitu

Database URL disimpan pada file .env dengan nama variabel database_uri

```
    database_uri=username:password@tcp(ip:port)/databaseName?parseTime=true
```

| No  | HTTP METHOD | URI              | Deskripsi                                                                         |
| --- | ----------- | ---------------- | --------------------------------------------------------------------------------- |
| 1   | GET         | /api/contacts    | Menampilkan semua data yang ada di database.                                      |
| 2   | GET         | /api/contact/:id | Menampilkan data yang ada di database sesuai id yang diberikan pada parameter URI |
| 3   | POST        | /api/contact     | Menambahkan data ke dalam database                                                |
| 4   | PUT         | /api/contact/:id | Mengubah data pada database untuk id data yang bersesuaian                        |
| 5   | DELETE      | /api/contact/    | Menghapus data pada database                                                      |

## 1. GET /api/contacts

Untuk endpoint dari URI ini akan menampilkan semua data yang ada di dalam database. URI ini dapat menerima 3 variabel query yaitu

- filter yang digunakan untuk memfilter data berdasarkan jenis kelamin. Nilai yang valid untuk variabel query ini adalah "male" atau "female"
- limit yang digunakan untuk membatasi banyak data yang ditampilkan. Jika limit tidak ada, maka nilai default dari limit ini akan di set 100
- offset yang digunakan untuk melewatkan banyaknya data yang ditampilkan. Jika offset tidak ada, maka nilai default dari offset ini akan di set 0

| No  | Query Variabel | value              | Default Value |
| --- | -------------- | ------------------ | ------------- |
| 1   | filter         | "male" or "female" | -             |
| 2   | limit          | number             | 100           |
| 3   | offset         | number             | 0             |

## 2. GET /api/contact/:id

URI ini akan menampilkan data sesuai dengan ID yang diberikan. Adapun nilai ID harus bertipe int

## 3. POST /api/contact/

URI ini akan menambahkan data ke dalam database. Data yang akan ditambahkan akan dimasukkan kedalam body request. Adapun variabel dan nilai yang perlu ditambahkan ke dalam body request antara lain:

| No  | Body Request | Value              |
| --- | ------------ | ------------------ |
| 1   | name         | string             |
| 2   | gender       | "male" or "female" |
| 3   | phone        | string             |
| 4   | email        | string             |

## 4. PUT /api/contact/:id

URI ini akan mengubah data database sesuai dengan id yang ada pada paramter URI. Data yang akan diupdate akan dimasukkan kedalam body request seperti yang telah dijelaskan pada point ke 3.

## 5. DELETE /api/contact/

URI ini akan menghapus data dari sebuah database sesuai dengan id yang dikirimkan melalui body request.
