package factory

import (
    _activityBusiness "devcode-golang.arief.id/apps/echo/features/activities/business"
    _activityData "devcode-golang.arief.id/apps/echo/features/activities/data"
    _activityPresentation "devcode-golang.arief.id/apps/echo/features/activities/presentation"

    _todoBusiness "devcode-golang.arief.id/apps/echo/features/todos/business"
    _todoData "devcode-golang.arief.id/apps/echo/features/todos/data"
    _todoPresentation "devcode-golang.arief.id/apps/echo/features/todos/presentation"

    "gorm.io/gorm"
)

type Presenter struct {
    ActivityPresenter *_activityPresentation.ActivityHandler
    TodoPresenter     *_todoPresentation.TodoHandler
}

func InitFactory(dbConn *gorm.DB) Presenter {
    activityData := _activityData.NewActivityRepository(dbConn)
    activityBusiness := _activityBusiness.NewActivityBusiness(activityData)
    activityPresentation := _activityPresentation.NewActivityHandler(activityBusiness)

    todoData := _todoData.NewTodoRepository(dbConn)
    todoBusiness := _todoBusiness.NewTodoBusiness(todoData)
    todoPresentation := _todoPresentation.NewTodoHandler(todoBusiness)

    return Presenter{
        ActivityPresenter: activityPresentation,
        TodoPresenter:     todoPresentation,
    }
}
