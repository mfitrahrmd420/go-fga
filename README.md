# FGA Kominfo Learning

## Day 8
Membuat 1st Web App (Backend Server):
Membuat user beserta order history dari user

- Membuat Aplikasi 
    - Config
        - harus connect ke DATABASE (postgres)
    - Membuat Domain
        di level ini hanya akan ada interface (method) dan struct (model/entity)
        - User
            - menambahkan user
            - get user detail
        - Order
            - menambahkan order terhadap user
            - mendapatkan semua order dari suatu user
    - Membuat Repository
        layer untuk mendapatkan data, query dan semua yang berhubungan dengan proses mendapatkan data, akan ada di layer ini
        - repo user
        - repo order
    - Membuat Usecase
        layer untuk menjalankan logic (business logic): masking data, checking data, dll
        - usecase untuk user
        - usercase untuk order
    - http server (gin gonic)
        - membuat handler
            layer untuk mendapatkan data yang diberikan oleh client pada suatu request (binding body payload, mendapatkan query paramater)
            - order
            - user
        - membuat router
            layer untuk menamakan API kita (path API kita) sehingga bisa dikenali oleh client: /v1/user, /v1/order
            - order
            - user
        - membuat documentation untuk API 
- Json package
    - apa itu json
    - kenapa json ada
    - json di golang (tag/annotation)
        - json marshall
        - json unmarshall
        - transform from map to struct
- API Documentation (Swagger)
    - Apa itu swagger
    - kenapa swagger ada
    - implementasi menggunakan swaggo