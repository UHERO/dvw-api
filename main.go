package main

import (
	"database/sql"
	"fmt"
	"github.com/UHERO/rest-api/common"
	"github.com/UHERO/rest-api/routers"
	"log"
	"net"
	"net/http"
	"net/url"
	"os"
	"time"

	"github.com/garyburd/redigo/redis"
	"github.com/go-sql-driver/mysql"
	"github.com/urfave/negroni"
)

func main() {
	common.StartUp()

	// Set up MySQL
	dbPort, ok := os.LookupEnv("DB_PORT")
	if !ok {
		dbPort = "3306"
	}
	dbName, ok := os.LookupEnv("DB_DBNAME")
	if !ok {
		dbName = "dbedt_visitor_dw"
	}
	mysqlConfig := mysql.Config{
		User:      os.Getenv("DB_USER"),
		Passwd:    os.Getenv("DB_PASSWORD"),
		Net:       "tcp",
		Addr:      net.JoinHostPort(os.Getenv("DB_HOST"), dbPort),
		Loc:       time.Local,
		ParseTime: true,
		DBName:    dbName,
	}
	connectionString := mysqlConfig.FormatDSN()
	db, err := sql.Open("mysql", connectionString)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		log.Fatal("Cannot login to MySQL server - check all DB_* environment variables")
	}

	// Set up Redis
	var redis_server, authpw string
	if redis_url, ok := os.LookupEnv("REDIS_URL"); ok {
		if u, err := url.Parse(redis_url); err == nil {
			redis_server = u.Host // includes port where specified
			authpw, _ = u.User.Password()
		}
	}
	if redis_server == "" {
		log.Print("Valid REDIS_URL var not found; using redis @ localhost:6379")
		redis_server = "localhost:6379"
	}
	pool := &redis.Pool{
		MaxIdle:     10,
		MaxActive:   50,
		IdleTimeout: 240 * time.Second,
		Dial: func() (redis.Conn, error) {
			c, err := redis.Dial("tcp", redis_server)
			if err != nil {
				log.Printf("*** Cannot contact redis server at %s. No caching!", redis_server)
				return nil, err
			}
			if authpw != "" {
				if _, err = c.Do("AUTH", authpw); err != nil {
					c.Close()
					log.Print("*** Redis authentication failure. No caching!")
					return nil, err
				}
			}
			log.Printf("Redis connection to %s established", redis_server)
			return c, nil
		},
		TestOnBorrow: func(c redis.Conn, t time.Time) error {
			if time.Since(t) < time.Minute {
				return nil
			}
			_, err := c.Do("PING")
			return err
		},
	}

	cache := &data.Cache{Pool: pool, TTL: 60 * 10} //TTL in seconds

	// Get the mux router object
	router := routers.InitRoutes(cache)
	// Create a negroni instance
	n := negroni.Classic()
	n.UseHandler(router)

	port := os.Getenv("GO_REST_PORT")
	if port == "" {
		port = "8080"
	}

	server := &http.Server{
		Addr:    fmt.Sprintf(":%s", port),
		Handler: n,
	}
	log.Printf("Listening on %s...", server.Addr)
	server.ListenAndServe()
}
