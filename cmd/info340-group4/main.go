package main

// These are your imports / libraries / frameworks
import (
	// this is Go's built-in sql library
	"database/sql"
	"log"
	"net/http"
	"os"

	// this allows us to run our web server
	"github.com/gin-gonic/gin"
	// this lets us connect to Postgres DB's
	_ "github.com/lib/pq"
)

var (
	// this is the pointer to the database we will be working with
	// this is a "global" variable (sorta kinda, but you can use it as such)
	db *sql.DB
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		log.Fatal("$PORT must be set")
	}

	var errd error
	// here we want to open a connection to the database using an environemnt variable.
	// This isn't the best technique, but it is the simplest one for heroku
	db, errd = sql.Open("postgres", os.Getenv("DATABASE_URL"))
	if errd != nil {
		log.Fatalf("Error opening database: %q", errd)
	}
	router := gin.New()
	router.Use(gin.Logger())
	router.LoadHTMLGlob("html/*")
	router.Static("/static", "static")

	router.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", nil)
	})

	router.GET("/account.html", func(c *gin.Context) {
		c.HTML(http.StatusOK, "account.html", nil)
	})

	router.GET("/newaccount.html", func(c *gin.Context) {
		c.HTML(http.StatusOK, "newaccount.html", nil)
	})

	router.GET("/QuserInfo", func(c *gin.Context) {
		table := "<table class='table'><thead><tr>"
		// put your query here
		rows, err := db.Query("SELECT first_name, last_name, email, phone_number FROM Customer WHERE customer_id = 1;")
		if err != nil {
			// careful about returning errors to the user!
			c.AbortWithError(http.StatusInternalServerError, err)
		}
		// foreach loop over rows.Columns, using value
		cols, _ := rows.Columns()
		if len(cols) == 0 {
			c.AbortWithStatus(http.StatusNoContent)
		}
		for _, value := range cols {
			table += "<th class='text-center'>" + value + "</th>"
		}
		// once you've added all the columns in, close the header
		table += "</thead><tbody>"
		// declare all your RETURNED columns here
		var first string
		var last string
		var email string
		var phone string 

		for rows.Next() {
			// assign each of them, in order, to the parameters of rows.Scan.
			// preface each variable with &
			rows.Scan(&first, &last, &email, &phone)
			// can't combine ints and strings in Go. Use strconv.Itoa(int) instead
			table += "<tr><td>" + first + "</td><td>" + last + "</td><td>" + email + "</td><td>" + phone + "</td></tr>"
		}
		// finally, close out the body and table
		table += "</tbody></table>"
		c.Data(http.StatusOK, "text/html", []byte(table))
	})

	router.POST("/Qnewaccount", func(c *gin.Context) {
		fname := c.PostForm("fname")
		lname := c.PostForm("lname")
		email := c.PostForm("email")
		phone := c.PostForm("phone")
		password := c.PostForm("password")
		db.Query("SELECT create_customer($5, $1, $2, $3, $4);", fname, lname, email, phone, password)
	})

	// NO code should go after this line. it won't ever reach that point
	router.Run(":" + port)
}