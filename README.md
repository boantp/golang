# go-mysql-rest-api
a set of APIs to be used by front-end engineers to develop an application that store and display tax amounts, Tech stack : Go, Docker-Compose, MySQL, HTML, CSS

This API serve 3 routes at Order Collection [/order/{store_id}], Cart Collection [/cart], Tax Code Collection [/tax_code] .
- You can perform fetch List Tax Code [GET] from Tax Code Collection
- Add to cart or Create Tax Object Add Cart [POST] from Cart Collection
- And can get/view the bill List All Order Item [GET] from Order Collection, from here we use store_id=1 hardcode at GUI add cart
so we can call [/order/1] to fetch my bill

The frontend show some GUI for add cart / add tax object, status update add cart, view my bill

See [API Documentation](https://github.com/boantp/go-mysql-rest-api/blob/master/apiary.apib) on how to use it.

## Directory Structure
```
go-mysql-rest-api
    |--config                   - to initialize template and database
        |--db.go                - for initialize mysql database connection
        |--tpl.go               - for template view configuration
    |--controllers              - to store package controllers
        |--cart.go              - to handle Cart Collection [/cart]
        |--order.go             - to handle Order Collection [/order/{store_id}]
        |--tax_code.go          - to handle Tax Code Collection [/tax_code]
        |--web.go               - to handle Web/Frontend for GUI add cart["/"], GUI view bill ["order_view/:store_id]
    |--docker                   - Dockerfile for Golang at folder web, Dockerfile for MySQL at folder db
        |--db
            |--Dockerfile
        |--web
            |--Dockerfile
    |--models                    - to store package models for object and mysql query
        |--order_details.go      - for table order_details
        |--orders.go             - for total amount, total tax amount, grand total object, and for table orders
        |--tax_code.go           - for table tax_code
    |--mysql_init                - init Table shopee
        |--shopee.sql
    |--templates                 - to store html file for golang *gohtml
    |--apiary.apib               - json file docs from APIARY for API DOCS
    |--database_design.png       - DB structure and explanation
    |--docker-compose.yml        - for docker-compose service and config
    |--main.go                   

  
```

## Setup

**Steps**
1. git clone git@github.com:boantp/go-mysql-rest-api.git
2. install docker and docker-compose 
3. open terminal and run docker-compose build (service are build), docker-compose up(builds, recreates, attaches to container for service), docker-compose down (stop containers and remove containers etc) See [Docker Documentation](https://docs.docker.com/compose/reference/build/) on how to use it.

You can also follow the official [docs](https://golang.org/doc/install) of installation if you want to know the complete process.

### Project setup

1. Clone the repository in your `$GOPATH/src/` directory. If you have used the bash script for setup, your `$GOPATH` variable should point to `$HOME/go`
2. **Follow the steps 2-6 only if you have to set-up databse by yourself**. The MongoDB database is hosted on mLab free trial account for now and might expire. In that case, you'll need the steps below.
3. To run project locally, Install Mongo DB - https://www.mongodb.com/download-center?jmp=nav#community
4. After installing Mongo DB, start it's server by typing `mongod` in Terminal.
5. Open a new tab in terminal and type `mongo < dummyData.js` to insert the dummmy product data.
6. Open file `store/repository.go`, find the `SERVER` variable and replace the URL. 

```
const SERVER = "http://localhost:27017"
```
7. Last thing required to run the project, install all the go dependencies
```
// Library to handle jwt authentication 
$ go get "github.com/dgrijalva/jwt-go"

// Libraries to handle network routing
$ go get "github.com/gorilla/mux"
$ go get "github.com/gorilla/context"
$ go get "github.com/gorilla/handlers"

// mgo library for handling Mongo DB
$ go get "gopkg.in/mgo.v2"
```
Yay! Now we're ready to run the API :tada: <br>
8. Type `export PORT=8000` in Terminal and open http://localhost:8000 in your browser to see the products.

## API Documentation and Usage

It is **recommended** to install some extension to beautify JSON(like [JSON Formatter](https://chrome.google.com/webstore/detail/json-formatter/bcjindcccaagfpapjjmafapmmgkkhgoa)) if you're trying in a browser.

**Important** - Don't forget to define $PORT in your shell variables. <br>Example: `export PORT=8000`

```sh
BASE_URL = "http://localhost:$PORT"
'OR'
BASE_URL = https://gruesome-monster-22811.herokuapp.com/
```

### 1. View Products

- **Endpoint Name** - `Index`      <br>
- **Method** - `GET`               <br>
- **URL Pattern** - `/`            <br>
- **Usage** 
    - Open BASE_URL in browser
    - **Terminal/CURL**
    ```sh
    curl -X GET BASE_URL
    ```
- **Expected Response** - JSON containing all the products in database <br>
- **Example**
![Screenshot](/screenshots/All-Products.png?raw=true)

### 2. View Single Product

- **Endpoint Name** - `GetProduct`    <br>
- **Method** - `GET`                  <br>
- **URL Pattern** - `/products/{id}`  <br>
- **Usage**
    - Open BASE_URL/products/{id} in browser
    - **Terminal/CURL**
```
curl -X GET BASE_URL/products/{id} 
```
- **Expected Response** - Product with the {id} in database
- **NOTE** - There are only six(6) ids in the database, so 1 <= {id} <= 6   
- **Example**
![Screenshot](/screenshots/GetProduct-Request.png)

### 3. Search Product

- **Endpoint Name** - `SearchProduct`  <br>
- **Method** - `GET`                   <br>
- **URL Pattern** - `/Search/{query}`  <br>
- **Usage** - Browser OR curl        
- **BROWSER**
    - Open BASE_URL/Search/{query} in browser
    - **Terminal/CURL**
    ```sh
    curl -X GET BASE_URL/Search/{query}
    ```
- **Expected Response** - Products matching the search query <br>
- **Example**
![Screenshot](/screenshots/Search-Request.png)

### 4. Authentication
For **Adding**, **Updating** and **Deleting** products from database you must send a JWT token in Authentication header.

- **Endpoint Name** - `GetToken` <br>
- **Method** - `POST`            <br>
- **URL Pattern** - `/get-token` <br>
- **Usage** - CURL OR POSTMAN ONLY
    - **Terminal/CURL**
    ```sh
    curl -X POST \
    -H "Content-Type: application/json" \
    -d '{ username: "<YOUR_USERNAME>", password: "<RANDOM_PASSWORD>"}' \
    BASE_URL/get-token
    ```
- **Expected Response** - A JWT Authentication Token as shown below
- **Example**
![Screenshot](/screenshots/Authentication-Request.png)

### 5. Add Product

- **Endpoint Name** - `AddProduct` <br>
- **Method** - `POST`              <br>
- **URL Pattern** - `/AddProduct`  <br>
- **Usage** - CURL OR POSTMAN ONLY
    - **Terminal/CURL**
    ```sh
    curl -X POST \
    -H "Authorization: Bearer <ACCESS_TOKEN>" \
    -d '{ "_id": 11, 
        "title": "Memes",
        "image": "I am selling memes, hehe.",          
        "price": 1,
        "rating": 5
        }' \
    BASE_URL/AddProduct
    ```
- **Expected Response** - Addition successful without any error message. Check the logs in Terminal window which is running server. 
- **Example**
![Screenshot](/screenshots/AddProduct-Request.png)

### 6. Update Product

- **Endpoint Name** - `UpdateProduct` <br>
- **Method** - `PUT`                  <br>
- **URL Pattern** - `/UpdateProduct`  <br>
- **Usage** - CURL OR POSTMAN ONLY
    - **Terminal/CURL**
    ```sh
    curl -X PUT \
    -H "Authorization: Bearer <ACCESS_TOKEN>" \
    -d '{ "ID": 14, 
        "title": "Memes",
        "image": "I am not selling memes to you, hehe.",          
        "price": 1000,
        "rating": 5
        }' \
    BASE_URL/UpdateProduct
    ```
- **Expected Response** - Update successful without any error message. Check the logs in Terminal window which is running server. <br>
- **Example**
![Screenshot](/screenshots/Update-Request.png)

### 7. Delete Product

- **Endpoint Name** - `DeleteProduct` <br>
- **Method** - `DELETE` <br>
- **URL Pattern** - `/deleteProduct/{id}` <br>
- **Usage** - CURL OR POSTMAN ONLY
    - **Terminal/CURL**
    ```sh
    curl -X DELETE \
    -H "Authorization: Bearer <ACCESS_TOKEN>" \
    BASE_URL/deleteProduct/{id}
    ```
- **Expected Response** - Deletion successful without any error message. Check the logs in Terminal window which is running server. <br>
- **Example**
![Screenshot](/screenshots/Delete-Request.png)

## TODO
* [ ] Write unit tests to test every method
* [ ] Improve the code by proper exception handling
* [ ] Add repository badges like TravisCI, Better Code, Codacy etc.
* [ ] Create a REST API server project using this package as a boilerplate
* [ ] User and roles management
* [ ] Session management using JWT tokens


