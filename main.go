package main

import (
	"fmt"
	"kollectionmanager/m/db"
	"kollectionmanager/m/deployment/migrations"
	"kollectionmanager/m/routes"
	"kollectionmanager/m/utils"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
)

func LoadEnvFromPath(envFilePath string) {
	err := godotenv.Load(envFilePath)
	if err != nil {
		err = utils.LoadEnvErr(envFilePath, err)
		log.Fatal(err)
	}
}

func main() {
	//carrega as envs em modo debug
	//LoadEnvFromPath(utils.LoadEnvFromPath)
	LoadEnvFromPath("/vault/secrets/.env")
	app := fiber.New()

	fmt.Printf("🚀 Iniciando servidor Fiber na porta :%s\n", utils.Port)

	newDB, err := db.Connect()
	if err != nil {
		fmt.Println(err)
	}

	fmt.Printf("🚀 Banco conectado")

	migrations.MigrateIfExists(newDB)
	routes.Router(app, newDB)

	fmt.Println("Rotas iniciadas")

	fmt.Println(utils.ServerStatus)
	if err := app.Listen(":" + utils.Port); err != nil {
		panic(err)
	}
}
