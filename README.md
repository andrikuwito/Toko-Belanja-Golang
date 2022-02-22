# Toko-Belanja-Golang
# Project 4 - Golang Hacktiv8 X Kampus Merdeka

Scalable Web Services with Golang

Kelompok : 
- Andri Nur Hidayatulloh
- Andri Kuwito

### List library : 
- [Gin Gionic](https://github.com/gin-gonic/gin) - Web Framework
- [Gorm](https://gorm.io) - GORM
- [Go Validator](https://github.com/asaskevich/govalidator) - Go Validator


### Cara Penggunaan : 

* ##### URL : http://fast-beach-63234.herokuapp.com/


* #### Endpoint List : 
    * ##### User : 
        * #### Register
            
            [POST]```http://fast-beach-63234.herokuapp.com/users/register```
            
            body :

            ```json
            {
                "full_name": "Andri Nur H",
                "password": "Bismillah",
                "email": "andriandri@gmail.com"
            }
            ```

            response :
            ```json
            {
                "status": "success",
                "data": {
                    "id": 5,
                    "full_name": "Andri Nur Hidayatulloh",
                    "email": "andribis14@gmail.com",
                    "balance": 0,
                    "created_at": "2022-02-21T10:37:39.989Z"
                }
            }
            ```
            
        * #### Login
            
            [POST]```http://fast-beach-63234.herokuapp.com/users/login```
            
            body :

            ```json
            {
                "email": "andri@gmail.com",
                "password": "Bismillah"
            }
            ```

            response :
            ```json
            {
                "status": "ok",
                "data": {
                    "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpZF91c2VyIjo0fQ.mqRz49vGePkvKGjQH04O9NamzqR-y_za5t5dvSqtMG0"
                }
            }
            ```
            
        * #### Top Up Balance Account
            
            [POST]```http://fast-beach-63234.herokuapp.com/users/topup```
            
            headers :

            ```
            {
                "Authorization": "Bearer {{token}}"
            }
            ```
            
            body :

            ```json
            {
                "balance": 40000000
            }
            ```

            response :
            ```json
            {
                "status": "ok",
                "data": {
                    "message": "Your balance has been successfully updated to Rp. 120040000"
                }
            }
            ```
            
    * ##### Transactions : 
        * #### Add Transactions
            
            [POST]```http://fast-beach-63234.herokuapp.com/transactions```
            
            headers :
            ```
            {
                "Authorization": "Bearer {{token}}"
            }
            ```
            
            body :
            ```json
            {
                "product_id": 3,
                "quantity": 10
            }
            ```

            response :
            ```json
            {
                "status": "success",
                "data": {
                    "message": "You have successfully purchased the product",
                    "transaction_bill": {
                        "total_price": 300000,
                        "quantity": 10,
                        "product_title": "Ayam Goreng"
                    }
                }
            }
            ```
            
        * #### My Transactions
            
            [GET]```http://fast-beach-63234.herokuapp.com/transactions/my-transactions```
            
            headers :

            ```
            {
                "Authorization": "Bearer {{token}}"
            }
            ```

            response :
            ```json
            {
                "status": "success",
                "data": [
                    {
                        "id": 2,
                        "product_id": 3,
                        "user_id": 4,
                        "quantity": 10,
                        "total_price": 300000,
                        "product": {
                            "id": 3,
                            "title": "Ayam Goreng",
                            "price": 30000,
                            "stock": 10,
                            "category_id": 1,
                            "created_at": "2022-01-23T12:35:39Z",
                            "updated_at": "2022-02-21T10:47:49Z"
                        },
                        "created_at": "2022-02-11T13:00:39Z",
                        "updated_at": "2022-02-11T13:00:39Z"
                    },
                    {
                        "id": 7,
                        "product_id": 3,
                        "user_id": 4,
                        "quantity": 10,
                        "total_price": 300000,
                        "product": {
                            "id": 3,
                            "title": "Ayam Goreng",
                            "price": 30000,
                            "stock": 10,
                            "category_id": 1,
                            "created_at": "2022-01-23T12:35:39Z",
                            "updated_at": "2022-02-21T10:47:49Z"
                        },
                        "created_at": "2022-02-11T13:06:33Z",
                        "updated_at": "2022-02-11T13:06:33Z"
                    },
                    {
                        "id": 8,
                        "product_id": 3,
                        "user_id": 4,
                        "quantity": 10,
                        "total_price": 300000,
                        "product": {
                            "id": 3,
                            "title": "Ayam Goreng",
                            "price": 30000,
                            "stock": 10,
                            "category_id": 1,
                            "created_at": "2022-01-23T12:35:39Z",
                            "updated_at": "2022-02-21T10:47:49Z"
                        },
                        "created_at": "2022-02-11T13:18:49Z",
                        "updated_at": "2022-02-11T13:18:49Z"
                    },
                    {
                        "id": 9,
                        "product_id": 3,
                        "user_id": 4,
                        "quantity": 10,
                        "total_price": 300000,
                        "product": {
                            "id": 3,
                            "title": "Ayam Goreng",
                            "price": 30000,
                            "stock": 10,
                            "category_id": 1,
                            "created_at": "2022-01-23T12:35:39Z",
                            "updated_at": "2022-02-21T10:47:49Z"
                        },
                        "created_at": "2022-02-11T13:30:10Z",
                        "updated_at": "2022-02-11T13:30:10Z"
                    },
                    {
                        "id": 10,
                        "product_id": 3,
                        "user_id": 4,
                        "quantity": 10,
                        "total_price": 300000,
                        "product": {
                            "id": 3,
                            "title": "Ayam Goreng",
                            "price": 30000,
                            "stock": 10,
                            "category_id": 1,
                            "created_at": "2022-01-23T12:35:39Z",
                            "updated_at": "2022-02-21T10:47:49Z"
                        },
                        "created_at": "2022-02-11T13:31:16Z",
                        "updated_at": "2022-02-11T13:31:16Z"
                    },
                    {
                        "id": 11,
                        "product_id": 3,
                        "user_id": 4,
                        "quantity": 10,
                        "total_price": 300000,
                        "product": {
                            "id": 3,
                            "title": "Ayam Goreng",
                            "price": 30000,
                            "stock": 10,
                            "category_id": 1,
                            "created_at": "2022-01-23T12:35:39Z",
                            "updated_at": "2022-02-21T10:47:49Z"
                        },
                        "created_at": "2022-02-11T13:33:33Z",
                        "updated_at": "2022-02-11T13:33:33Z"
                    },
                    {
                        "id": 12,
                        "product_id": 3,
                        "user_id": 4,
                        "quantity": 10,
                        "total_price": 300000,
                        "product": {
                            "id": 3,
                            "title": "Ayam Goreng",
                            "price": 30000,
                            "stock": 10,
                            "category_id": 1,
                            "created_at": "2022-01-23T12:35:39Z",
                            "updated_at": "2022-02-21T10:47:49Z"
                        },
                        "created_at": "2022-02-11T13:35:57Z",
                        "updated_at": "2022-02-11T13:35:57Z"
                    },
                    {
                        "id": 13,
                        "product_id": 3,
                        "user_id": 4,
                        "quantity": 10,
                        "total_price": 300000,
                        "product": {
                            "id": 3,
                            "title": "Ayam Goreng",
                            "price": 30000,
                            "stock": 10,
                            "category_id": 1,
                            "created_at": "2022-01-23T12:35:39Z",
                            "updated_at": "2022-02-21T10:47:49Z"
                        },
                        "created_at": "2022-02-11T13:36:17Z",
                        "updated_at": "2022-02-11T13:36:17Z"
                    },
                    {
                        "id": 14,
                        "product_id": 3,
                        "user_id": 4,
                        "quantity": 10,
                        "total_price": 300000,
                        "product": {
                            "id": 3,
                            "title": "Ayam Goreng",
                            "price": 30000,
                            "stock": 10,
                            "category_id": 1,
                            "created_at": "2022-01-23T12:35:39Z",
                            "updated_at": "2022-02-21T10:47:49Z"
                        },
                        "created_at": "2022-02-11T13:38:23Z",
                        "updated_at": "2022-02-11T13:38:23Z"
                    },
                    {
                        "id": 15,
                        "product_id": 3,
                        "user_id": 4,
                        "quantity": 10,
                        "total_price": 300000,
                        "product": {
                            "id": 3,
                            "title": "Ayam Goreng",
                            "price": 30000,
                            "stock": 10,
                            "category_id": 1,
                            "created_at": "2022-01-23T12:35:39Z",
                            "updated_at": "2022-02-21T10:47:49Z"
                        },
                        "created_at": "2022-02-11T13:39:59Z",
                        "updated_at": "2022-02-11T13:39:59Z"
                    },
                    {
                        "id": 16,
                        "product_id": 3,
                        "user_id": 4,
                        "quantity": 10,
                        "total_price": 300000,
                        "product": {
                            "id": 3,
                            "title": "Ayam Goreng",
                            "price": 30000,
                            "stock": 10,
                            "category_id": 1,
                            "created_at": "2022-01-23T12:35:39Z",
                            "updated_at": "2022-02-21T10:47:49Z"
                        },
                        "created_at": "2022-02-15T17:59:51Z",
                        "updated_at": "2022-02-15T17:59:51Z"
                    },
                    {
                        "id": 17,
                        "product_id": 3,
                        "user_id": 4,
                        "quantity": 10,
                        "total_price": 300000,
                        "product": {
                            "id": 3,
                            "title": "Ayam Goreng",
                            "price": 30000,
                            "stock": 10,
                            "category_id": 1,
                            "created_at": "2022-01-23T12:35:39Z",
                            "updated_at": "2022-02-21T10:47:49Z"
                        },
                        "created_at": "2022-02-21T10:47:49Z",
                        "updated_at": "2022-02-21T10:47:49Z"
                    }
                ]
            }
            ```
            
        * #### User Transactions
            
            [PATCH]```http://fast-beach-63234.herokuapp.com/transactions/user-transactions```
            
            headers :
            ```
            {
                "Authorization": "Bearer {{token}}"
            }
            ```
            response :
            ```json
            {
                "status": "success",
                "data": [
                    {
                        "id": 1,
                        "product_id": 3,
                        "user_id": 1,
                        "quantity": 20,
                        "total_price": 600000,
                        "product": {
                            "id": 3,
                            "title": "Ayam Goreng",
                            "price": 30000,
                            "stock": 10,
                            "category_id": 1,
                            "created_at": "2022-01-23T12:35:39Z",
                            "updated_at": "2022-02-21T10:47:49Z"
                        },
                        "user": {
                            "id": 1,
                            "full_name": "Vita",
                            "email": "vita@gmail.com",
                            "password": "$2a$04$ycPWwBfAR1ZC7m/Flk2TQOmsehuhdUyMwp.ZBdqk8zEKTPVDXWLIK",
                            "role": "customer",
                            "balance": 60000,
                            "created_at": "2022-01-23T11:12:12Z",
                            "updated_at": "2022-01-23T11:27:57Z"
                        },
                        "created_at": "2022-01-23T12:39:37Z",
                        "updated_at": "2022-01-23T12:39:37Z"
                    },
                    {
                        "id": 2,
                        "product_id": 3,
                        "user_id": 4,
                        "quantity": 10,
                        "total_price": 300000,
                        "product": {
                            "id": 3,
                            "title": "Ayam Goreng",
                            "price": 30000,
                            "stock": 10,
                            "category_id": 1,
                            "created_at": "2022-01-23T12:35:39Z",
                            "updated_at": "2022-02-21T10:47:49Z"
                        },
                        "user": {
                            "id": 4,
                            "full_name": "Andri Nur Hidayatulloh",
                            "email": "andribis13@gmail.com",
                            "password": "$2a$04$n.CY5tK2EUDgvuYTNoCuPOu6tZSwB0SokrLXYoa9mHS01EHaBd6hS",
                            "role": "customer",
                            "balance": 120040000,
                            "created_at": "2022-01-30T07:21:19Z",
                            "updated_at": "2022-02-21T10:45:01Z"
                        },
                        "created_at": "2022-02-11T13:00:39Z",
                        "updated_at": "2022-02-11T13:00:39Z"
                    },
                    {
                        "id": 7,
                        "product_id": 3,
                        "user_id": 4,
                        "quantity": 10,
                        "total_price": 300000,
                        "product": {
                            "id": 3,
                            "title": "Ayam Goreng",
                            "price": 30000,
                            "stock": 10,
                            "category_id": 1,
                            "created_at": "2022-01-23T12:35:39Z",
                            "updated_at": "2022-02-21T10:47:49Z"
                        },
                        "user": {
                            "id": 4,
                            "full_name": "Andri Nur Hidayatulloh",
                            "email": "andribis13@gmail.com",
                            "password": "$2a$04$n.CY5tK2EUDgvuYTNoCuPOu6tZSwB0SokrLXYoa9mHS01EHaBd6hS",
                            "role": "customer",
                            "balance": 120040000,
                            "created_at": "2022-01-30T07:21:19Z",
                            "updated_at": "2022-02-21T10:45:01Z"
                        },
                        "created_at": "2022-02-11T13:06:33Z",
                        "updated_at": "2022-02-11T13:06:33Z"
                    },
                    {
                        "id": 8,
                        "product_id": 3,
                        "user_id": 4,
                        "quantity": 10,
                        "total_price": 300000,
                        "product": {
                            "id": 3,
                            "title": "Ayam Goreng",
                            "price": 30000,
                            "stock": 10,
                            "category_id": 1,
                            "created_at": "2022-01-23T12:35:39Z",
                            "updated_at": "2022-02-21T10:47:49Z"
                        },
                        "user": {
                            "id": 4,
                            "full_name": "Andri Nur Hidayatulloh",
                            "email": "andribis13@gmail.com",
                            "password": "$2a$04$n.CY5tK2EUDgvuYTNoCuPOu6tZSwB0SokrLXYoa9mHS01EHaBd6hS",
                            "role": "customer",
                            "balance": 120040000,
                            "created_at": "2022-01-30T07:21:19Z",
                            "updated_at": "2022-02-21T10:45:01Z"
                        },
                        "created_at": "2022-02-11T13:18:49Z",
                        "updated_at": "2022-02-11T13:18:49Z"
                    },
                    {
                        "id": 9,
                        "product_id": 3,
                        "user_id": 4,
                        "quantity": 10,
                        "total_price": 300000,
                        "product": {
                            "id": 3,
                            "title": "Ayam Goreng",
                            "price": 30000,
                            "stock": 10,
                            "category_id": 1,
                            "created_at": "2022-01-23T12:35:39Z",
                            "updated_at": "2022-02-21T10:47:49Z"
                        },
                        "user": {
                            "id": 4,
                            "full_name": "Andri Nur Hidayatulloh",
                            "email": "andribis13@gmail.com",
                            "password": "$2a$04$n.CY5tK2EUDgvuYTNoCuPOu6tZSwB0SokrLXYoa9mHS01EHaBd6hS",
                            "role": "customer",
                            "balance": 120040000,
                            "created_at": "2022-01-30T07:21:19Z",
                            "updated_at": "2022-02-21T10:45:01Z"
                        },
                        "created_at": "2022-02-11T13:30:10Z",
                        "updated_at": "2022-02-11T13:30:10Z"
                    },
                    {
                        "id": 10,
                        "product_id": 3,
                        "user_id": 4,
                        "quantity": 10,
                        "total_price": 300000,
                        "product": {
                            "id": 3,
                            "title": "Ayam Goreng",
                            "price": 30000,
                            "stock": 10,
                            "category_id": 1,
                            "created_at": "2022-01-23T12:35:39Z",
                            "updated_at": "2022-02-21T10:47:49Z"
                        },
                        "user": {
                            "id": 4,
                            "full_name": "Andri Nur Hidayatulloh",
                            "email": "andribis13@gmail.com",
                            "password": "$2a$04$n.CY5tK2EUDgvuYTNoCuPOu6tZSwB0SokrLXYoa9mHS01EHaBd6hS",
                            "role": "customer",
                            "balance": 120040000,
                            "created_at": "2022-01-30T07:21:19Z",
                            "updated_at": "2022-02-21T10:45:01Z"
                        },
                        "created_at": "2022-02-11T13:31:16Z",
                        "updated_at": "2022-02-11T13:31:16Z"
                    },
                    {
                        "id": 11,
                        "product_id": 3,
                        "user_id": 4,
                        "quantity": 10,
                        "total_price": 300000,
                        "product": {
                            "id": 3,
                            "title": "Ayam Goreng",
                            "price": 30000,
                            "stock": 10,
                            "category_id": 1,
                            "created_at": "2022-01-23T12:35:39Z",
                            "updated_at": "2022-02-21T10:47:49Z"
                        },
                        "user": {
                            "id": 4,
                            "full_name": "Andri Nur Hidayatulloh",
                            "email": "andribis13@gmail.com",
                            "password": "$2a$04$n.CY5tK2EUDgvuYTNoCuPOu6tZSwB0SokrLXYoa9mHS01EHaBd6hS",
                            "role": "customer",
                            "balance": 120040000,
                            "created_at": "2022-01-30T07:21:19Z",
                            "updated_at": "2022-02-21T10:45:01Z"
                        },
                        "created_at": "2022-02-11T13:33:33Z",
                        "updated_at": "2022-02-11T13:33:33Z"
                    },
                    {
                        "id": 12,
                        "product_id": 3,
                        "user_id": 4,
                        "quantity": 10,
                        "total_price": 300000,
                        "product": {
                            "id": 3,
                            "title": "Ayam Goreng",
                            "price": 30000,
                            "stock": 10,
                            "category_id": 1,
                            "created_at": "2022-01-23T12:35:39Z",
                            "updated_at": "2022-02-21T10:47:49Z"
                        },
                        "user": {
                            "id": 4,
                            "full_name": "Andri Nur Hidayatulloh",
                            "email": "andribis13@gmail.com",
                            "password": "$2a$04$n.CY5tK2EUDgvuYTNoCuPOu6tZSwB0SokrLXYoa9mHS01EHaBd6hS",
                            "role": "customer",
                            "balance": 120040000,
                            "created_at": "2022-01-30T07:21:19Z",
                            "updated_at": "2022-02-21T10:45:01Z"
                        },
                        "created_at": "2022-02-11T13:35:57Z",
                        "updated_at": "2022-02-11T13:35:57Z"
                    },
                    {
                        "id": 13,
                        "product_id": 3,
                        "user_id": 4,
                        "quantity": 10,
                        "total_price": 300000,
                        "product": {
                            "id": 3,
                            "title": "Ayam Goreng",
                            "price": 30000,
                            "stock": 10,
                            "category_id": 1,
                            "created_at": "2022-01-23T12:35:39Z",
                            "updated_at": "2022-02-21T10:47:49Z"
                        },
                        "user": {
                            "id": 4,
                            "full_name": "Andri Nur Hidayatulloh",
                            "email": "andribis13@gmail.com",
                            "password": "$2a$04$n.CY5tK2EUDgvuYTNoCuPOu6tZSwB0SokrLXYoa9mHS01EHaBd6hS",
                            "role": "customer",
                            "balance": 120040000,
                            "created_at": "2022-01-30T07:21:19Z",
                            "updated_at": "2022-02-21T10:45:01Z"
                        },
                        "created_at": "2022-02-11T13:36:17Z",
                        "updated_at": "2022-02-11T13:36:17Z"
                    },
                    {
                        "id": 14,
                        "product_id": 3,
                        "user_id": 4,
                        "quantity": 10,
                        "total_price": 300000,
                        "product": {
                            "id": 3,
                            "title": "Ayam Goreng",
                            "price": 30000,
                            "stock": 10,
                            "category_id": 1,
                            "created_at": "2022-01-23T12:35:39Z",
                            "updated_at": "2022-02-21T10:47:49Z"
                        },
                        "user": {
                            "id": 4,
                            "full_name": "Andri Nur Hidayatulloh",
                            "email": "andribis13@gmail.com",
                            "password": "$2a$04$n.CY5tK2EUDgvuYTNoCuPOu6tZSwB0SokrLXYoa9mHS01EHaBd6hS",
                            "role": "customer",
                            "balance": 120040000,
                            "created_at": "2022-01-30T07:21:19Z",
                            "updated_at": "2022-02-21T10:45:01Z"
                        },
                        "created_at": "2022-02-11T13:38:23Z",
                        "updated_at": "2022-02-11T13:38:23Z"
                    },
                    {
                        "id": 15,
                        "product_id": 3,
                        "user_id": 4,
                        "quantity": 10,
                        "total_price": 300000,
                        "product": {
                            "id": 3,
                            "title": "Ayam Goreng",
                            "price": 30000,
                            "stock": 10,
                            "category_id": 1,
                            "created_at": "2022-01-23T12:35:39Z",
                            "updated_at": "2022-02-21T10:47:49Z"
                        },
                        "user": {
                            "id": 4,
                            "full_name": "Andri Nur Hidayatulloh",
                            "email": "andribis13@gmail.com",
                            "password": "$2a$04$n.CY5tK2EUDgvuYTNoCuPOu6tZSwB0SokrLXYoa9mHS01EHaBd6hS",
                            "role": "customer",
                            "balance": 120040000,
                            "created_at": "2022-01-30T07:21:19Z",
                            "updated_at": "2022-02-21T10:45:01Z"
                        },
                        "created_at": "2022-02-11T13:39:59Z",
                        "updated_at": "2022-02-11T13:39:59Z"
                    },
                    {
                        "id": 16,
                        "product_id": 3,
                        "user_id": 4,
                        "quantity": 10,
                        "total_price": 300000,
                        "product": {
                            "id": 3,
                            "title": "Ayam Goreng",
                            "price": 30000,
                            "stock": 10,
                            "category_id": 1,
                            "created_at": "2022-01-23T12:35:39Z",
                            "updated_at": "2022-02-21T10:47:49Z"
                        },
                        "user": {
                            "id": 4,
                            "full_name": "Andri Nur Hidayatulloh",
                            "email": "andribis13@gmail.com",
                            "password": "$2a$04$n.CY5tK2EUDgvuYTNoCuPOu6tZSwB0SokrLXYoa9mHS01EHaBd6hS",
                            "role": "customer",
                            "balance": 120040000,
                            "created_at": "2022-01-30T07:21:19Z",
                            "updated_at": "2022-02-21T10:45:01Z"
                        },
                        "created_at": "2022-02-15T17:59:51Z",
                        "updated_at": "2022-02-15T17:59:51Z"
                    },
                    {
                        "id": 17,
                        "product_id": 3,
                        "user_id": 4,
                        "quantity": 10,
                        "total_price": 300000,
                        "product": {
                            "id": 3,
                            "title": "Ayam Goreng",
                            "price": 30000,
                            "stock": 10,
                            "category_id": 1,
                            "created_at": "2022-01-23T12:35:39Z",
                            "updated_at": "2022-02-21T10:47:49Z"
                        },
                        "user": {
                            "id": 4,
                            "full_name": "Andri Nur Hidayatulloh",
                            "email": "andribis13@gmail.com",
                            "password": "$2a$04$n.CY5tK2EUDgvuYTNoCuPOu6tZSwB0SokrLXYoa9mHS01EHaBd6hS",
                            "role": "customer",
                            "balance": 120040000,
                            "created_at": "2022-01-30T07:21:19Z",
                            "updated_at": "2022-02-21T10:45:01Z"
                        },
                        "created_at": "2022-02-21T10:47:49Z",
                        "updated_at": "2022-02-21T10:47:49Z"
                    }
                ]
            }
            ```
            
    * ##### Categories : 
        * #### Add Categories
            
            [POST]```http://fast-beach-63234.herokuapp.com/categories```
            
            headers :
            ```
            {
                "Authorization": "Bearer {{token}}"
            }
            ```
            
            body :
            ```json
            {
                "type": "Jus"
            }
            ```
            response :
            ```json
            {
                "status": "success",
                "data": {
                    "id": 6,
                    "title": "Nyuci Mobil",
                    "description": "Nyuci mobik tiap sore",
                    "status": false,
                    "user_id": 3,
                    "category_id": 2,
                    "created_at": "2022-01-22T17:39:39.923+07:00"
                }
            }
            ```
            
        * #### GET Categories
            
            [GET]```http://fast-beach-63234.herokuapp.com/categories```
            
            headers :
            ```
            {
                "Authorization": "Bearer {{token}}"
            }
            ```
            response :
            ```json
            {
                "status": "success",
                "data": [
                    {
                        "id": 2,
                        "type": "Makanan",
                        "created_at": "2022-02-15T02:03:05Z",
                        "updated_at": "2022-02-15T02:11:50Z"
                    },
                    {
                        "id": 3,
                        "type": "Minuman",
                        "created_at": "2022-02-15T02:05:33Z",
                        "updated_at": "2022-02-15T02:05:33Z"
                    },
                    {
                        "id": 4,
                        "type": "Jus",
                        "created_at": "2022-02-21T10:52:23Z",
                        "updated_at": "2022-02-21T10:52:23Z"
                    }
                ]
            }
            ```
        
        * #### Update Categoires
            
            [PATCH]```http://fast-beach-63234.herokuapp.com/categories/:id```
            
            headers :
            ```
            {
                "Authorization": "Bearer {{token}}"
            }
            ```
            
            body :
            ```json
            {
                "type": "Makanan"
            }
            ```
            response :
            ```json
            {
                "status": "ok",
                "data": {
                    "id": 2,
                    "types": "Makanan",
                    "updated_at": "2022-02-21T10:54:17Z"
                }
            }
            ```
    
        * #### Delete Categories
            
            [DELETE]```http://fast-beach-63234.herokuapp.com/categories/:id```
            
            headers :
            ```
            {
                "Authorization": "Bearer {{token}}"
            }
            ```
            
            response :
            ```json
            {
                "status": "ok",
                "data": "Category deleted"
            }
            ```

    * ##### Categories : 
        * #### Add Categories
            
            [POST]```http://fast-beach-63234.herokuapp.com/products```
            
            headers :
            ```
            {
                "Authorization": "Bearer {{token}}"
            }
            ```
            
            body :
            ```json
            {
                "title": "Mobil",
                "price": 40000,
                "stock": 50,
                "category_id": 2
            }
            ```
            response :
            ```json
            {
                "status": "success",
                "data": {
                    "id": 7,
                    "title": "Mobil",
                    "price": 40000,
                    "stock": 50,
                    "created_at": "2022-02-21T10:56:12.324Z"
                }
            }
            ```
        
        * #### GET Products
            
            [GET]```http://fast-beach-63234.herokuapp.com/products```
            
            headers :
            ```
            {
                "Authorization": "Bearer {{token}}"
            }
            ```
            
            response :
            ```json
            {
                "status": "success",
                "data": [
                    {
                        "id": 3,
                        "title": "Ayam Goreng",
                        "price": 30000,
                        "stock": 10,
                        "category_id": 1,
                        "created_at": "2022-01-23T12:35:39Z",
                        "updated_at": "2022-02-21T10:47:49Z"
                    },
                    {
                        "id": 5,
                        "title": "Mobil-mobilan Truck",
                        "price": 25000,
                        "stock": 90,
                        "category_id": 2,
                        "created_at": "2022-02-15T02:18:47Z",
                        "updated_at": "2022-02-15T02:18:47Z"
                    },
                    {
                        "id": 6,
                        "title": "Robot Mainan",
                        "price": 10000,
                        "stock": 50,
                        "category_id": 2,
                        "created_at": "2022-02-15T02:22:11Z",
                        "updated_at": "2022-02-15T17:51:51Z"
                    },
                    {
                        "id": 7,
                        "title": "Mobil",
                        "price": 40000,
                        "stock": 50,
                        "category_id": 2,
                        "created_at": "2022-02-21T10:56:12Z",
                        "updated_at": "2022-02-21T10:56:12Z"
                    }
                ]
            }
            ```
            
        * #### Update Products
            
            [PUT]```http://fast-beach-63234.herokuapp.com/products/:id```
            
            headers :
            ```
            {
                "Authorization": "Bearer {{token}}"
            }
            ```
            
            body :
            ```json
            {
                "title": "Robot Mainan",
                "price": 10000
            }
            ```
            
            response :
            ```json
            {
                "status": "ok",
                "data": {
                    "id": 6,
                    "title": "Robot Mainan",
                    "price": 10000,
                    "stock": 50,
                    "updated_at": "2022-02-21T10:59:02Z"
                }
            }
            ```
            
        * #### Delete Products
            
            [PUT]```http://fast-beach-63234.herokuapp.com/products/:id```
            
            headers :
            ```
            {
                "Authorization": "Bearer {{token}}"
            }
            ```
            
            response :
            ```json
            {
                "status": "ok",
                "data": "products deleted"
            }
            ```
