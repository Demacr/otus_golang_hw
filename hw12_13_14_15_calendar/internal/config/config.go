package config

import (
	"context"
	"fmt"
	"io/ioutil"
	"log"
	"net"
	"net/http"
	"os"
	"strconv"
	"sync"
	"time"

	"google.golang.org/grpc"

	"github.com/Demacr/otus_golang_hw/hw12_13_14_15_calendar/internal/grpcserver"
	"github.com/Demacr/otus_golang_hw/hw12_13_14_15_calendar/internal/httpserver"
	"github.com/Demacr/otus_golang_hw/hw12_13_14_15_calendar/internal/logger"
	"github.com/Demacr/otus_golang_hw/hw12_13_14_15_calendar/internal/storage"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq" // Register postgres driver
	"github.com/pkg/errors"
	"github.com/spf13/pflag"
	"gopkg.in/yaml.v2"
)

type Config struct {
	Host string `yaml:"host"`
	Port struct {
		HTTP int `yaml:"http"`
		GRPC int `yaml:"grpc"`
	} `yaml:"port"`
	Database   string `yaml:"database"`
	PostgreSQL struct {
		Host     string `yaml:"host"`
		Port     int    `yaml:"port"`
		Login    string `yaml:"login"`
		Password string `yaml:"password"`
		Database string `yaml:"database"`
	} `yaml:"postgresql"`
	Log struct {
		File  string `yaml:"file"`
		Level string `yaml:"level"`
	} `yaml:"log"`

	strg storage.Storager
}

func Configure() *Config {
	var configPath string
	var config Config
	pflag.StringVar(&configPath, "config", "", "Path to config file")
	pflag.Parse()

	if configPath == "" {
		log.Fatal("missing config file")
	}

	configFile, err := os.Open(configPath)
	if err != nil {
		log.Fatal(err)
	}

	configBytes, err := ioutil.ReadAll(configFile)
	if err != nil {
		log.Fatal(err)
	}

	err = yaml.Unmarshal(configBytes, &config)
	if err != nil {
		log.Fatal(err)
	}

	return &config
}

func (cfg *Config) ConfigureLoggerByConfig() {
	fd, err := os.OpenFile(cfg.Log.File, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal(err)
	}
	level := logger.LevelFromString(cfg.Log.Level)
	logger.Log = logger.NewLogger(level, fd, log.New(fd, "", log.LstdFlags))

	logger.Debug = logger.GenerateLoggerFunc(logger.DEBUG, level)
	logger.Info = logger.GenerateLoggerFunc(logger.INFORMATIONAL, level)
	logger.Warning = logger.GenerateLoggerFunc(logger.WARNING, level)
	logger.Error = logger.GenerateLoggerFunc(logger.ERROR, level)
	logger.Fatal = logger.GenerateLoggerFunc(logger.FATAL, level)
}

func (cfg *Config) RunHTTPServer(quitCh chan interface{}, wg *sync.WaitGroup) {
	router := httpserver.NewRouter(cfg.strg)
	server := &http.Server{
		Addr:         cfg.Host + ":" + strconv.Itoa(cfg.Port.HTTP),
		Handler:      httpserver.MiddlewareLogger(router.RootHandler()),
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	wg.Add(1)
	go func() {
		defer wg.Done()

		err := server.ListenAndServe()
		if err != nil {
			return
		}
	}()

	go func() {
		<-quitCh
		if err := server.Shutdown(context.Background()); err != nil {
			logger.Error("error during shutdown HTTP server:", err)
		}
	}()
}

func (cfg *Config) RunGRPCServer(quitCh chan interface{}, wg *sync.WaitGroup) error {
	lis, err := net.Listen("tcp", cfg.Host+":"+strconv.Itoa(cfg.Port.GRPC))
	if err != nil {
		logger.Error("failed to listen: %v", err)
		return errors.Wrap(err, "GRPC failed to listen")
	}
	grpcServer := grpc.NewServer(
		grpc.UnaryInterceptor(
			grpcserver.LogInterceptor(),
		))
	grpcserver.RegisterCalendarServer(grpcServer, &grpcserver.CalendarService{})

	wg.Add(1)
	go func() {
		defer wg.Done()

		err := grpcServer.Serve(lis)
		if err != nil {
			logger.Error(err)
		}
	}()

	go func() {
		<-quitCh
		grpcServer.GracefulStop()
	}()

	return nil
}

func (cfg *Config) RunStorage() {
	switch cfg.Database {
	case "inmemory":
		cfg.strg = storage.NewInMemoryStorage()

	case "postgresql":
		cfg.strg = cfg.NewPgSQLStorage()
	}
}

func (cfg *Config) NewPgSQLStorage() *storage.PgSQLStorage {
	dsn := fmt.Sprintf(
		"port=%d host=%s user=%s password=%s dbname=%s sslmode=disable",
		cfg.PostgreSQL.Port,
		cfg.PostgreSQL.Host,
		cfg.PostgreSQL.Login,
		cfg.PostgreSQL.Password,
		cfg.PostgreSQL.Database,
	)

	db, err := sqlx.Connect("postgres", dsn)
	if err != nil {
		logger.Fatal(err)
		return nil
	}
	return storage.NewPgSQLStorageStruct(db)
}
