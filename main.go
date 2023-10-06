package main

import (
	"context"
	"flag"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/redis/go-redis/v9"
	"github.com/spf13/viper"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/wwfyde/demo-gin/api"
	"github.com/wwfyde/demo-gin/api/v1"
	_ "github.com/wwfyde/demo-gin/docs"
	"github.com/wwfyde/demo-gin/handlers"
	"github.com/wwfyde/demo-gin/middleware/logger"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.uber.org/zap"
	"log"
	"net/http"
	"os"
	"time"
)

var (
	port = flag.Int("port", 8000, "Server listening port")
)

func init() {
	// TODO init

	//gin.SetMode(gin.TestMode)

	// get env
	if err := godotenv.Load(); err != nil {
		panic("Error loading .env file")
	}
	fmt.Println(os.Getenv("GIN_MODE"))
	gin.SetMode(gin.ReleaseMode)

}

// @title				Demo Gin API
// @version			3.1.0
// @host				localhost:8009
// @BasePath			/api/v1
// @externalDocs.url	https://swagger.io/resources/open-api/
func main() {
	// get ENV
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	MongoURI := os.Getenv("MONGO_URI")
	if MongoURI == "" {
		log.Fatal("MongoURI not configured!")
	}
	// 强制日志颜色化
	gin.ForceConsoleColor()
	//// 禁止日志的颜色
	//gin.DisableConsoleColor()

	// Parse Command line flags
	flag.Parse()
	//docs.SwaggerInfo.BasePath = "/api/v1"

	// get cfg
	cfg := readConfig()

	// get ctx
	ctx, cancel := context.WithTimeout(context.Background(), cfg.GetDuration("app.timeout")*time.Second)
	log.Println("timeout set:", cfg.GetDuration("app.timeout")*time.Second)
	defer cancel()

	// get mdb
	mdb, err := mongo.Connect(ctx, options.Client().ApplyURI(MongoURI))
	defer func() {
		if err = mdb.Disconnect(ctx); err != nil {
			panic(err)
		}
	}()

	// get rdb
	rdb := redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
		DB:   0,
	})
	PodcastHandler := handlers.NewPodcastHandler(context.TODO(), cfg, mdb, rdb)
	r := gin.Default()
	r.Use(logger.Logger())
	// Recovery
	r.Use(gin.Recovery())
	apiv1 := r.Group("/api/v1")
	{
		//apiv1.GET("", api.v12.HelloWorld)

		//apiv1.GET("/podcast")
	}
	apiv1.GET("/ping", api.Ping)
	apiv1.GET("/hello", handlers.HelloWorld)
	apiv1.GET("/hello2", v1.HelloWorld)
	apiv1.GET("", func(c *gin.Context) { c.String(http.StatusOK, "api/v1") })
	r.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	r.GET("/podcast/:id", PodcastHandler.GetPodcast)
	r.Use()
	{
		logger, _ := zap.NewProduction()
		defer logger.Sync()
		logger.Info("My First Structured log with zap!",
			zap.String("hello", "世界!"),
			zap.Int("attempt", 3),
			zap.Duration("backoff", time.Second),
		)
		if err := r.Run(fmt.Sprintf(":%d", *port)); err != nil {
			log.Fatalf("Run server failed: %v", err)
		}
	}

	// serve

}

func readConfig() *viper.Viper {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("config")
	if err := viper.ReadInConfig(); err != nil {
		panic(err)
	}
	return viper.GetViper()
}
