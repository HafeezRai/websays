Article, Category, and Product API
This is a RESTful JSON API for managing articles, categories, and products, built using the Gin framework, text files for persistence of categories, and a PostgreSQL database for persistence of products.

Getting Started
These instructions will get you a copy of the project up and running on your local machine for development and testing purposes.

Prerequisites
Go (v1.15 or later)
PostgreSQL
Installing
Clone the repository to your local machine

Create a PostgreSQL database and configure the connection details in the init function of the main file.

Start the API:
$ go run main.go


The API should now be running on http://localhost:8000.


# Building the Docker Image

To build the Docker image, run the following command in the project directory:

$ docker build -t assessment .

# Running the Docker Container

To run the Docker container, use the following command:

$ docker run -p 8000:8000 assessment

This will run the container and map port 8000 on the host to port 8000 in the container.


API Endpoints
The following endpoints are available for managing articles, categories, and products:

Articles
Method	Endpoint	Description
GET	/articles	Retrieve a list of all articles
GET	/articles/:id	Retrieve a single article by its ID
POST	/articles	Create a new article
PUT	/articles/:id	Update an existing article
DELETE	/articles/:id	Delete an article
Categories
Method	Endpoint	Description
GET	/categories	Retrieve a list of all categories
GET	/categories/:id	Retrieve a single category by its ID
POST	/categories	Create a new category
PUT	/categories/:id	Update an existing category
DELETE	/categories/:id	Delete a category
Products
Method	Endpoint	Description
GET	/products	Retrieve a list of all products
GET	/products/:id	Retrieve a single product by its ID
POST	/products	Create a new product
PUT	/products/:id	Update an existing product
DELETE	/products/:id	Delete a product
Built With
Gin - The web framework used
PostgreSQL - The database used for persistence of products
Text files - The persistence method used for categories


