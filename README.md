# Picus Golang Bootcamp|Final Project

## Basket Service API with Golang

Basket Service API Server provides a simple basket server that provides endpoints which specified as a list in "Service Endpoints" section and the feature of these endpoints are detailed in the Swagger documentation provided as a "basketService.yaml" file in the "docs" folder.  

### Technologies
* Gorm
* JWT
* Gin
* Postgres

## DB GORM Models

### Category

```
type Category struct {
	gorm.Model
	CategoryName string `json:"category_name"`
	Products     []Product
}
```

### Product
```
type Product struct {
	gorm.Model
	ProductName      string  `json:"product_name"`
	CategoryID       uint    `json:"category_id"`
	ProductStock     int     `json:"product_stock"`
	ProductPrice     float64 `json:"product_price"`
	ProductStockCode string  `json:"product_stock_code"`
	ProductType      string  `json:"product_type"`
	ProductColor     string  `json:"product_color"`
	ProductSize      string  `json:"product_size"`
	ProductMaterial  string  `json:"product_material"`
	Gender           string  `json:"gender"`
	Category         *Category
}
```

### User

```
type User struct {
	gorm.Model
	Email    string        `json:"email"`
	Password string        `json:"password"`
	Role     string        `json:"role"`
	Items    []*BasketItem `json:"items"`
}
```

### BasketItem

```
type BasketItem struct {
	gorm.Model
	ProductID uint `json:"product_id"`
	Quantity  int  `json:"quantity"`
	UserID    uint `json:"user_id"`
	Product   *Product
}
```

## Service Endpoints 

* BasePath: "http://localhost:8080/api/v1/basket-service-api"
* Endpoints features of Basket Service applications, such as body shema and responses, are detailed in the Swagger documentation provided as a "basketService.yaml" file in the "docs" folder. 

### BasePath/category:
```
POST BasePath/category/upload:
```
> The categories of the service should be in ".csv" file format and uploaded as form-Data. Note: An example test file provided as "categories.csv" file inside the "test_file" folder 

```
GET BasePath/category/list:
```
> List the categories in the service

### BasePath/product:

```
POST BasePath/product/create:
```
> Add product to the service

```
GET BasePath/product/list:
```
> List all products in the service

```
GET BasePath/product/:id:
```
> Gets the products from database in the service with given product id

```
GET BasePath/product/searchByProductName/:name
```
> Gets the products from database in the service with given product name

```
GET BasePath/product/searchByProductID/:id
```
> Searches and gets the products from database in the service with given product id

```
DELETE BasePath/product/deleteProductByID/:id
```
> Deletes the products from database in the service with given product id

```
PUT BasePath/product/deleteProductByID/:id
```
> Updates the products in database in the service with given product id

### BasePath/basket:

```
POST BasePath/basket/add:
```
> Add item to the basket in the service

```
GET BasePath/basket/list/:id
```
> List all items in the basket with given user id

```
DELETE BasePath/basket/deleteProductByID/:id
```
> Deletes the products from the basket database in the service with given product id

```
PUT BasePath/basket/deleteProductByID/:id
```
> Updates the products in the basket database in the service with given product id

### BasePath/user:

```
POST BasePath/user/signUp:
```
> Create new user, Note: Role field should be set 'Admin' or 'User'"

```
POST BasePath/user/login:
```
> User can be login to server and returns JWT Token