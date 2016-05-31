package main

// These are your imports / libraries / frameworks
import (
	// this is Go's built-in sql library
	"database/sql"
	"log"
	"net/http"
	"os"
	"strconv"

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
		rows, err := db.Query("SELECT first_name, last_name, email, phone_number FROM Customer WHERE customer_id = 1;")
		if err != nil {
			// careful about returning errors to the user!
			c.AbortWithError(http.StatusInternalServerError, err)
		}

		var para string
		var first string
		var last string
		var email string
		var phone string

		for rows.Next() {
			// assign each of them, in order, to the parameters of rows.Scan.
			// preface each variable with &
			rows.Scan(&first, &last, &email, &phone)
			// can't combine ints and strings in Go. Use strconv.Itoa(int) instead
			para += "<p>Name: " + first + " " + last + "</p>"
			para += "<p>Email: " + email + "</p>"
			para += "<p>Phone Number: " + phone + "</p>"
		}
		c.Data(http.StatusOK, "text/html", []byte(para))
	})

	router.GET("/QuserAddr", func(c *gin.Context) {
		rows, err := db.Query("SELECT address, city_name, state_name, zip_code FROM customer_address JOIN city ON city.city_id = customer_address.city_id JOIN state ON state.state_id = customer_address.state_id JOIN zip ON zip.zip_code_id = customer_address.zip_code_id WHERE customer_id = 1;")
		if err != nil {
			// careful about returning errors to the user!
			c.AbortWithError(http.StatusInternalServerError, err)
		}

		cols, _ := rows.Columns()
		if len(cols) == 0 {
			c.AbortWithStatus(http.StatusNoContent)
		}

		var para string
		var address string
		var state_name string
		var city_name string
		var zip_code int

		for rows.Next() {
			// assign each of them, in order, to the parameters of rows.Scan.
			// preface each variable with &
			rows.Scan(&address, &city_name, &state_name, &zip_code)
			// can't combine ints and strings in Go. Use strconv.Itoa(int) instead
			para += "<p>Address: " + address + " " + city_name + ", " + state_name + " " + strconv.Itoa(zip_code) + "</p>"
		}
		c.Data(http.StatusOK, "text/html", []byte(para))
	})

	router.GET("/QuserBooking", func(c *gin.Context) {
		rows, err := db.Query("SELECT first_name, last_name, email, phone_number FROM Customer WHERE customer_id = 1;")
		if err != nil {
			// careful about returning errors to the user!
			c.AbortWithError(http.StatusInternalServerError, err)
		}

		cols, _ := rows.Columns()
		if len(cols) == 0 {
			c.AbortWithStatus(http.StatusNoContent)
		}

		var para string
		var first string
		var last string
		var email string
		var phone string

		for rows.Next() {
			// assign each of them, in order, to the parameters of rows.Scan.
			// preface each variable with &
			rows.Scan(&first, &last, &email, &phone)
			// can't combine ints and strings in Go. Use strconv.Itoa(int) instead
			para += "<p>Name: " + first + " " + last + "</p>"
			para += "<p>Email: " + email + "</p>"
			para += "<p>Phone Number: " + phone + "</p>"
		}
		c.Data(http.StatusOK, "text/html", []byte(para))
	})

	router.GET("/QavailableRooms", func(c *gin.Context) {
		table := "<table class='table'><thead><tr>"
		// put your query here
		rows, err := db.Query("Select hotel_name, room_number, room_type From Hotel Join room on hotel.hotel_id = room.hotel_id join room_type on room.room_type_id = room_type.room_type_id where room.booking_available = 'true';")
		if err != nil {
			// careful about returning errors to the user!
			c.AbortWithError(http.StatusInternalServerError, err)
		}
		// foreach loop over rows.Columns, using value
		cols, _ := rows.Columns()
		if len(cols) == 0 {
			c.AbortWithStatus(http.StatusNoContent)
		}
		table += "<th class='text-center'>Hotel Name</th>"
		table += "<th class='text-center'>Room Number</th>"
		table += "<th class='text-center'>Room Type</th>"
		// once you've added all the columns in, close the header
		table += "</thead><tbody>"
		// declare all your RETURNED columns here
		var hotel string
		var roomNum string
		var roomType string
		for rows.Next() {
			// assign each of them, in order, to the parameters of rows.Scan.
			// preface each variable with &
			rows.Scan(&hotel, &roomNum, &roomType)
			// can't combine ints and strings in Go. Use strconv.Itoa(int) instead
			table += "<tr><td>" + hotel + "</td><td>" + roomNum + "</td><td>" + roomType + "</td></tr>"
		}
		// finally, close out the body and table
		table += "</tbody></table>"
		c.Data(http.StatusOK, "text/html", []byte(table))
	})


	router.GET("/QuserBooking", func(c *gin.Context) {
		table := "<table class='table'><thead><tr>"
		// put your query here
		rows, err := db.Query("SELECT hotel_name, room_number, party_size, begin_date, end_date From Booking JOIN Hotel ON hotel.hotel_id = booking.hotel_id WHERE customer_id = 1 AND (begin_date >= current_date OR end_date >= current_date);")
		if err != nil {
			// careful about returning errors to the user!
			c.AbortWithError(http.StatusInternalServerError, err)
		}
		// foreach loop over rows.Columns, using value
		cols, _ := rows.Columns()
		if len(cols) == 0 {
			c.AbortWithStatus(http.StatusNoContent)
		}
		table += "<th class='text-center'>Hotel Name</th>"
		table += "<th class='text-center'>Room Number</th>"
		table += "<th class='text-center'>Party Size</th>"
		table += "<th class='text-center'>Begin Date</th>"
		table += "<th class='text-center'>End Date</th>"
		// once you've added all the columns in, close the header
		table += "</thead><tbody>"
		// declare all your RETURNED columns here
		var hotel string
		var roomNum string
		var party_size int
		var begin string
		var end string
		for rows.Next() {
			// assign each of them, in order, to the parameters of rows.Scan.
			// preface each variable with &
			rows.Scan(&hotel, &roomNum, &party_size, &begin, &end)
			// can't combine ints and strings in Go. Use strconv.Itoa(int) instead
			table += "<tr><td>" + hotel + "</td><td>" + roomNum + "</td><td>" + strconv.Itoa(party_size) + "</td><td>" + begin + "</td><td>" + end + "</td></tr>"
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
		c.Data(http.StatusOK, "text/html", []byte("New Account Created"))
	})

	// NO code should go after this line. it won't ever reach that point
	router.Run(":" + port)
}
