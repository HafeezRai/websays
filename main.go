package main

import (
	"fmt"
    "io/ioutil"
	"strconv"
    "encoding/json"
    "net/http"
	"github.com/gin-gonic/gin"
    "github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
    
)


type Article struct {
	ID      int    `json:"id"`
	Title   string `json:"title"`
	Content string `json:"content"`
}
var articles []Article


type Category struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

var categories []Category


type Product struct {
	gorm.Model
	Name  string `json:"name"`
	Price int    `json:"price"`
}

var db *gorm.DB

func init() {
	var err error
	db, err = gorm.Open("postgres", "host=myhost port=myport user=myuser dbname=websays password=XPc8DC3FucbPBiR#")
	if err != nil {
		panic("failed to connect database")
	}

	db.AutoMigrate(&Product{})
}

func main() {
	router := gin.Default()
    // Load the categories from a file on startup
	data, err := ioutil.ReadFile("categories.txt")
	if err != nil {
		fmt.Println("Failed to load categories from file:", err)
	} else {
		json.Unmarshal(data, &categories)
	}

	// Initialize some articles for demonstration purposes
	articles = append(articles, Article{ID: 1, Title: "Article 1", Content: "Content of article 1"})
	articles = append(articles, Article{ID: 2, Title: "Article 2", Content: "Content of article 2"})
	articles = append(articles, Article{ID: 3, Title: "Article 3", Content: "Content of article 3"})

	// Create the routes for handling CRUD operations on the articles
	router.GET("/articles", getArticles)
	router.GET("/articles/:id", getArticle)
	router.POST("/articles", createArticle)
	router.PUT("/articles/:id", updateArticle)
	router.DELETE("/articles/:id", deleteArticle)


    // Create the routes for handling CRUD operations on the categories
	router.GET("/categories", getCategories)
	router.GET("/categories/:id", getCategory)
	router.POST("/categories", createCategory)
	router.PUT("/categories/:id", updateCategory)
	router.DELETE("/categories/:id", deleteCategory)


    router.GET("/products", getAllProducts)
	router.GET("/products/:id", getProduct)
	router.POST("/products", createProduct)
	router.PUT("/products/:id", updateProduct)
	router.DELETE("/products/:id", deleteProduct)

	// Start the server
	fmt.Println("Starting the server on port 8000...")
	router.Run(":8000")
}
func getArticles(c *gin.Context) {
	c.Header("Content-Type", "application/json")
    c.JSON(http.StatusOK, articles)
}

func getArticle(c *gin.Context) {
	c.Header("Content-Type", "application/json")
	id, _ := strconv.Atoi(c.Param("id"))

	for _, article := range articles {
		if article.ID == id {
			c.JSON(http.StatusOK, article)
			return
		}
	}
	c.JSON(http.StatusNotFound, gin.H{"message": "Article not found"})
}
func createArticle(c *gin.Context) {
	c.Header("Content-Type", "application/json")
	var article Article
	c.BindJSON(&article)
	article.ID = len(articles) + 1
	articles = append(articles, article)
	c.JSON(http.StatusCreated, article)
}

func updateArticle(c *gin.Context) {
	c.Header("Content-Type", "application/json")
	id, _ := strconv.Atoi(c.Param("id"))
	var article Article
	c.BindJSON(&article)

	for i, a := range articles {
		if a.ID == id {
			article.ID = id
			articles[i] = article
			c.JSON(http.StatusOK, article)
			return
		}
	}
	c.JSON(http.StatusNotFound, gin.H{"message": "Article not found"})
}

func deleteArticle(c *gin.Context) {
	c.Header("Content-Type", "application/json")
	id, _ := strconv.Atoi(c.Param("id"))

	for i, article := range articles {
		if article.ID == id {
			articles = append(articles[:i], articles[i+1:]...)
			c.JSON(http.StatusOK, gin.H{"message": "Article deleted"})
			return
		}
	}
	c.JSON(http.StatusNotFound, gin.H{"message": "Article not found"})
}

func getCategories(c *gin.Context) {
	c.Header("Content-Type", "application/json")
	c.JSON(http.StatusOK, categories)
}
func getCategory(c *gin.Context) {
	c.Header("Content-Type", "application/json")
	id, _ := strconv.Atoi(c.Param("id"))

	for _, category := range categories {
		if category.ID == id {
			c.JSON(http.StatusOK, category)
			return
		}
	}
	c.JSON(http.StatusNotFound, gin.H{"message": "Category not found"})
}

func createCategory(c *gin.Context) {
	c.Header("Content-Type", "application/json")
	var category Category
	c.BindJSON(&category)
	category.ID = len(categories) + 1
	categories = append(categories, category)
	data, _ := json.Marshal(categories)
	ioutil.WriteFile("categories.txt", data, 0644)
	c.JSON(http.StatusCreated, category)
}

func updateCategory(c *gin.Context) {
	c.Header("Content-Type", "application/json")
	id, _ := strconv.Atoi(c.Param("id"))
	var category Category
	c.BindJSON(&category)

	for i, a := range categories {
		if a.ID == id {
			category.ID = id
			categories[i] = category
			data, _ := json.Marshal(categories)
			ioutil.WriteFile("categories.txt", data, 0644)
			c.JSON(http.StatusOK, category)
			return
		}
	}
	c.JSON(http.StatusNotFound, gin.H{"message": "Category not found"})
}

func deleteCategory(c *gin.Context) {
	c.Header("Content-Type", "application/json")
	id, _ := strconv.Atoi(c.Param("id"))

	for i, a := range categories {
		if a.ID == id {
			categories = append(categories[:i], categories[i+1:]...)
			data, _ := json.Marshal(categories)
			ioutil.WriteFile("categories.txt", data, 0644)
			c.JSON(http.StatusOK, gin.H{"message": "Category deleted successfully"})
			return
		}
	}
	c.JSON(http.StatusNotFound, gin.H{"message": "Category not found"})
}

func createProduct(c *gin.Context) {
	c.Header("Content-Type", "application/json")
	var product Product
	c.BindJSON(&product)
	db.Create(&product)
	c.JSON(http.StatusOK, product)
}

func getProduct(c *gin.Context) {
	c.Header("Content-Type", "application/json")
	id := c.Param("id")
	var product Product
	db.First(&product, id)
	if product.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{"message": "Product not found"})
		return
	}
	c.JSON(http.StatusOK, product)
}

func getAllProducts(c *gin.Context) {
	c.Header("Content-Type", "application/json")
	var products []Product
	db.Find(&products)
	c.JSON(http.StatusOK, products)
}

func updateProduct(c *gin.Context) {
	c.Header("Content-Type", "application/json")
	var product Product
	id := c.Param("id")
	db.First(&product, id)
	if product.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{"message": "Product not found"})
		return
	}
	c.BindJSON(&product)
	db.Save(&product)
	c.JSON(http.StatusOK, product)
}

func deleteProduct(c *gin.Context) {
	c.Header("Content-Type", "application/json")
	var product Product
	id := c.Param("id")
	db.First(&product, id)
	if product.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{"message": "Product not found"})
		return
	}
	db.Delete(&product)
	c.JSON(http.StatusOK, gin.H{"message": "Product deleted successfully"})
}



