package request

import "devcode-golang.arief.id/apps/echo/features/activities"

type Activities struct {
    Email string `json:"email" form:"email"`
    Title string `json:"title" form:"title"`
}

func ToCore(req Activities) activities.Core {
    return activities.Core{
        Email: req.Email,
        Title: req.Title,
    }
}
