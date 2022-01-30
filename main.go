package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/dlclark/regexp2"
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
	"github.com/joho/godotenv"
	gologger "github.com/mo-taufiq/go-logger"
	"github.com/spf13/cast"
	"github.com/thedevsaddam/govalidator"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

type Server struct {
	Port           string
	ReadTimeOut    int
	WriteTimeOut   int
	MaxHeaderBytes int
}

type Log struct {
	Path                string
	TimeZone            string
	NestedFuncLevel     int
	NestedLocationLevel int
}

type Database struct {
	Host          string
	Port          string
	UserName      string
	Password      string
	DatabaseName  string
	MigrationPath string
}

type Redis struct {
	Host     string
	Port     string
	UserName string
	Password string
}

type Global struct {
	DebugMode bool
	Server    Server
	Log       Log
	Database  Database
	Redis     Redis
}

var GlobalVariable = Global{}

func init() {
	// load environment configuration file
	err := godotenv.Load(os.Getenv("ENV"))
	if err != nil {
		fmt.Println(err)
	}

	GlobalVariable = Global{
		DebugMode: cast.ToBool(os.Getenv("DEBUG_MODE")),
		Server: Server{
			Port:           os.Getenv("PORT"),
			ReadTimeOut:    cast.ToInt(os.Getenv("READ_TIME_OUT")),
			WriteTimeOut:   cast.ToInt(os.Getenv("WRITE_TIME_OUT")),
			MaxHeaderBytes: cast.ToInt(os.Getenv("MAX_HEADER_BYTES")),
		},
		Log: Log{
			Path:                os.Getenv("PATH_LOGS"),
			TimeZone:            os.Getenv("LOG_TIMEZONE"),
			NestedFuncLevel:     cast.ToInt(os.Getenv("LOG_NESTED_FUNC_LEVEL")),
			NestedLocationLevel: cast.ToInt(os.Getenv("LOG_NESTED_LOCATION_LEVEL")),
		},
		Database: Database{
			Host:          os.Getenv("DB_HOST"),
			Port:          os.Getenv("DB_PORT"),
			UserName:      os.Getenv("DB_USER_NAME"),
			Password:      os.Getenv("DB_PASSWORD"),
			DatabaseName:  os.Getenv("DB_NAME"),
			MigrationPath: os.Getenv("MIGRATION_PATH"),
		},
		Redis: Redis{
			Host:     os.Getenv("REDIS_HOST"),
			Port:     os.Getenv("REDIS_PORT"),
			UserName: os.Getenv("REDIS_USER_NAME"),
			Password: os.Getenv("REDIS_PASSWORD"),
		},
	}

	// setting log configuration
	gologger.LogConf.DebugMode = GlobalVariable.DebugMode
	gologger.LogConf.Path = GlobalVariable.Log.Path
	gologger.LogConf.TimeZone = GlobalVariable.Log.TimeZone
	gologger.LogConf.NestedFuncLevel = GlobalVariable.Log.NestedFuncLevel
	gologger.LogConf.NestedLocationLevel = GlobalVariable.Log.NestedLocationLevel
	gologger.LogConf.LogFuncName = true

	// handle migration
	migrationSetting()
	// custom govalidator
	addCustomGoValidator()
}

func main() {
	// handle db
	db := databaseConnection()
	// handle redis
	rc := redisConnection()

	router := gin.New()
	// router := gin.Default() // with log hit history

	Routers(router, db, rc)

	s := &http.Server{
		Addr:           ":" + GlobalVariable.Server.Port,
		Handler:        router,
		ReadTimeout:    time.Duration(GlobalVariable.Server.ReadTimeOut) * time.Second,
		WriteTimeout:   time.Duration(GlobalVariable.Server.WriteTimeOut) * time.Second,
		MaxHeaderBytes: GlobalVariable.Server.MaxHeaderBytes,
	}

	gologger.Info(fmt.Sprintf("Starting running service on port :%s", GlobalVariable.Server.Port))

	s.ListenAndServe()
}

func migrationSetting() {
	gologger.Info("Start migration up")
	DBConnectionURL := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable", GlobalVariable.Database.UserName, GlobalVariable.Database.Password, GlobalVariable.Database.Host, GlobalVariable.Database.Port, GlobalVariable.Database.DatabaseName)
	m, err := migrate.New(
		"file://"+GlobalVariable.Database.MigrationPath,
		DBConnectionURL)
	if err != nil {
		gologger.Error(err.Error())
	}
	if err := m.Up(); err != nil {
		if !strings.EqualFold(err.Error(), "no change") {
			gologger.Error(err.Error())
			panic(err)
		}
		gologger.Info(fmt.Sprintf("migration: %s", err.Error()))
	}
}

func databaseConnection() *gorm.DB {
	gologger.Info("Start connecting to database")
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=%s", GlobalVariable.Database.Host, GlobalVariable.Database.UserName, GlobalVariable.Database.Password, GlobalVariable.Database.DatabaseName, GlobalVariable.Database.Port, GlobalVariable.Log.TimeZone)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		gologger.Error(err.Error())
		panic(err)
	}
	return db
}

func redisConnection() *redis.Client {
	gologger.Info("Start connecting to redis")
	redisClient := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%s", GlobalVariable.Redis.Host, GlobalVariable.Redis.Port),
		Username: GlobalVariable.Redis.UserName,
		Password: GlobalVariable.Redis.Password,
		DB:       0,
	})

	err := redisClient.Set(context.Background(), "ping", "pong", 0).Err()
	if err != nil {
		gologger.Error(err.Error())
		panic(err)
	}

	return redisClient
}

func addCustomGoValidator() {
	govalidator.AddCustomRule("valid_password", func(field string, rule string, message string, value interface{}) error {
		regex := regexp2.MustCompile(`^(?=.*[A-Za-z])(?=.*[\d])[a-zA-Z\d]{6,}$`, regexp2.RE2)
		isMatch, err := regex.MatchString(cast.ToString(value))
		if err != nil {
			gologger.Error(err.Error())
		}

		if !isMatch {
			return fmt.Errorf("%s must be at least 6 characters long and consist of a combination of letters and numbers", field)
		}
		return nil
	})
}
