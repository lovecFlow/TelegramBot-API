package config

import (
	"log"
	"os"
	"time"

	"github.com/ilyakaznacheev/cleanenv"
)

//? Структурные теги в Go представляют собой аннотации, которые отображаются после типа в декларации структуры Go. Каждый тег состоит из коротких строк, которым назначены определенные значения. Структурный тег выделяется символами апострофа ````` и выглядит следующим образом:

type Config struct { //* Все параметры конфига будут содержать параметры yaml файла
	Env         string `yaml:"env" env-default:"local"` //* Тег yaml нужен для определения имя соотвествующего параметра в yaml файле
	StoragePath string `yaml:"storage_path" env-required:"true"`
	HTTPServer  `yaml:"http_server"`
}

type HTTPServer struct {
	Address     string        `yaml:"address" env-default:"localhost:8080"`
	Timeot      time.Duration `yaml:"timeout" env-default:"4s"`
	IdleTimeout time.Duration `yaml:"idle_timeout" env-default:"60s"`
}

// * Функция, которая читает файл с конфига и заполнит объект конфига, который написан выше
func Load() *Config {
	//* Сначала определяемся откуда считывать файл с конфигом.
	configPath := os.Getenv("CONFIG_PATH") //* Получаем из переменной окружения
	if configPath == "" {                  //* Если не находим, то роняем приложение с фаталом
		log.Fatal("CONFIG_PATH is not set")
	}

	//* Проверяем существование файла
	if _, err := os.Stat(configPath); os.IsNotExist(err) { //* Проверка происходит с определенной ошибкой
		log.Fatalf("Config file does not exist: %s", configPath) //* Если не находим, то роняем с фаталом
	}

	var cfg Config //* Объявляеи объект конфига

	//*Считываем файл по пути, который указан у нас
	if err := cleanenv.ReadConfig(configPath, &cfg); err != nil {
		log.Fatalf("Cannot read config: %s", err)
	}

	return &cfg
}
