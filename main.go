package main

import (
	"fmt"
	"github.com/UHERO/dvw-api/controllers"
	"github.com/UHERO/dvw-api/data"
	"github.com/UHERO/dvw-api/routers"
	"github.com/garyburd/redigo/redis"
	"github.com/go-sql-driver/mysql"
	"github.com/urfave/negroni"
	"log"
	"net"
	"net/http"
	"net/url"
	"os"
	"strconv"
	"time"
)

const apiName = "dvw"

func main() {
	//common.StartUp()  ///// THIS IS FOR API AUTH FOR EXTERNAL USERS?

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
	db, err := data.CreateDatabase(connectionString)
	if err != nil {
		log.Fatal("Cannot reach MySQL server - check all DB_* environment variables")
	}
	defer db.Close()

	// Set up Redis
	var redisServer, authPw string
	if redisUrl, ok := os.LookupEnv("REDIS_URL"); ok {
		if u, err := url.Parse(redisUrl); err == nil {
			redisServer = u.Host // includes port where specified
			authPw, _ = u.User.Password()
		}
	}
	if redisServer == "" {
		log.Print("Valid REDIS_URL var not found; using redis @ localhost:6379")
		redisServer = "localhost:6379"
	}
	pool := &redis.Pool{
		MaxIdle: 10,
		MaxActive: 50,
		IdleTimeout: 240 * time.Second,
		Dial: func() (redis.Conn, error) {
			c, err := redis.Dial("tcp", redisServer)
			if err != nil {
				log.Printf("*** Cannot contact redis server at %s. No caching!", redisServer)
				return nil, err
			}
			if authPw != "" {
				if _, err = c.Do("AUTH", authPw); err != nil {
					c.Close()
					log.Print("*** Redis authentication failure. No caching!")
					return nil, err
				}
			}
			log.Printf("Redis connection to %s established", redisServer)
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
	cacheTTLStr, ok := os.LookupEnv("API_CACHE_TTL")
	if !ok {
		cacheTTLStr = "10"
	}
	cacheTTLMin, _ := strconv.Atoi(cacheTTLStr)

	controllers.CreateCache(apiName, pool, cacheTTLMin)
	router := routers.CreateRouter(apiName)
	n := negroni.Classic()
	n.UseHandler(router)

	port, ok := os.LookupEnv("GO_REST_PORT")
	if !ok {
		port = "8080"
	}
	server := &http.Server{
		Addr:    fmt.Sprintf(":%s", port),
		Handler: n,
	}

	log.Printf("Listening on %s...", server.Addr)
	err = server.ListenAndServe()
	if err != nil {
		// OUTPUT ERROR
	}
}
