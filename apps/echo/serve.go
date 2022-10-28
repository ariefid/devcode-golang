package main

import (
    "os"

    "devcode-golang.arief.id/apps/echo/config"
    "devcode-golang.arief.id/apps/echo/factory"
    "devcode-golang.arief.id/apps/echo/migrations"
    "devcode-golang.arief.id/apps/echo/routes"
)

func main() {
    dbConn := config.InitDB()
    migrations.InitMigrate(dbConn)
    presenter := factory.InitFactory(dbConn)
    e := routes.New(presenter)
    echoDockerPort := os.Getenv("ECHO_DOCKER_PORT")
    e.Logger.Fatal(e.Start(":" + echoDockerPort))
}
