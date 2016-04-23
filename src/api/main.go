package main

import (
	"github.com/gin-gonic/gin"
	"flag"
	"io/ioutil"
	"gopkg.in/yaml.v2"
	"database/sql"
	"github.com/go-gorp/gorp"
	"github.com/streadway/amqp"
	_ "github.com/lib/pq"
	"gopkg.in/telegram-bot-api.v4"
	"log"
)

type ConfigInterface interface {
	Load(path string) error
}

type Config struct {
	PgSqlConnectionString    string `yaml:"PgSqlConnectionString"`
	RabbitMQConnectionString string `yaml:"RabbitConnectionString"`
	Port                     string `yaml:"Port"`
	DebugMode                bool 	`yaml:"DebugMode"`
	TelegramToken		 string `yaml:"TelegramToken"`
}

func (c *Config) Load(path string) (err error) {
	data, err := ioutil.ReadFile(path)
	if err != nil {
		return
	}
	err = yaml.Unmarshal(data, &c)
	return
}

func main() {
	var configFile string
	flag.StringVar(&configFile, "config", "", "config filename")
	flag.StringVar(&configFile, "c", "", "config filename (shorthand)")
	flag.Parse()
	configurator := new(Config)
	err := configurator.Load(configFile);
	if err != nil {
		log.Panic(err)
	}
	log.Println("Congfiguration loaded correctly...")
	bot, err := tgbotapi.NewBotAPI(configurator.TelegramToken)
	bot.Debug = true
	log.Printf("Authorized on account %s", bot.Self.UserName)
	if err != nil {
		log.Panic(err)
	}
	log.Println("Bot connection completed...")
	strMode := "Debug"
	if !configurator.DebugMode {
		gin.SetMode(gin.ReleaseMode)
		strMode = "Release"
	}
	log.Printf("API daemon started in the %s mode", strMode)
	r := gin.Default()
	handler := new(Handler)
	v1 := r.Group("v1")
	{
		v1.GET("/ping", func(c *gin.Context) { handler.Ping(c)})
	}
	db, err := sql.Open("postgres", configurator.PgSqlConnectionString)
	if err != nil {
		log.Panic(err)
	}
	postgres := &gorp.DbMap{Db: db, Dialect: gorp.PostgresDialect{}}
	_, err = postgres.Exec("SET SCHEMA 'public'")
	if err != nil {
		log.Panic(err)
	}
	log.Println("Postgres connected...")
	_, err = amqp.Dial(configurator.RabbitMQConnectionString)
	if err != nil {
		log.Panic(err)
	}
	log.Println("RabbitMQ connected...")
	go botRouter(bot)
	r.Run(configurator.Port)
}