package data

import (
    "devcode-golang.arief.id/apps/echo/features/activities"

    "gorm.io/gorm"
)

type Activities struct {
    gorm.Model
    Email string `gorm:"unique" json:"email" form:"email"`
    Title string `json:"title" form:"title"`
}

func toCoreList(data []Activities) []activities.Core {
    result := []activities.Core{}
    for key := range data {
        result = append(result, data[key].toCore())
    }
    return result
}

func (data *Activities) toCore() activities.Core {
    return activities.Core{
        ID:        int(data.ID),
        Email:     data.Email,
        Title:     data.Title,
        CreatedAt: data.CreatedAt,
        UpdatedAt: data.UpdatedAt,
    }
}

func fromCore(core activities.Core) Activities {
    return Activities{
        Email: core.Email,
        Title: core.Title,
    }
}
