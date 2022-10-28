package migrations

import (
    _modelActivities "devcode-golang.arief.id/apps/echo/features/activities/data"
    _modelTodos "devcode-golang.arief.id/apps/echo/features/todos/data"

    "gorm.io/gorm"
)

func InitMigrate(db *gorm.DB) {
    db.AutoMigrate(&_modelActivities.Activities{})
    db.AutoMigrate(&_modelTodos.Todos{})
}
