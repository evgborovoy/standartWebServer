package main

import (
	"flag"
	"log"

	"github.com/BurntSushi/toml"
	"github.com/evgborovoy/StandardWebServer/internal/app/api"
)

var (
	configPath string
)

func init() {
	// приложение при запуске будет получать путь до конфиг файла
	flag.StringVar(&configPath, "path", "configs/api.toml", "path to config file in .toml format")
}

func main() {
	// В этот момент происходит инициализация configPath значением из командной строки
	flag.Parse()
	log.Println("start")
	// server instance initializations
	config := api.NewConfig()
	_, err := toml.DecodeFile(configPath, config) // Десериализация toml файла
	if err != nil {
		log.Println("can not find config file, use default values", err)
	}
	// Если файл прочитан, то данные будут браться из него, а если нет, то дефолтные
	server := api.New(config)

	//api server start
	if err := server.Start(); err != nil {
		log.Fatal(err)
	}
}
