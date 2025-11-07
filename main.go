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
	LoadEnvFromPath("/vault/secrets/dotenv")
	app := fiber.New()

	newDB, err := db.Connect()
	if err != nil {
		fmt.Println(err)
	}

	migrations.MigrateIfExists(newDB)
	routes.Router(app, newDB)

	fmt.Println(utils.ServerStatus)
	app.Listen(utils.Port)
}
