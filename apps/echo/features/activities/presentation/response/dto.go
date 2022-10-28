package response

import (
    "time"

    "devcode-golang.arief.id/apps/echo/features/activities"
)

type Activities struct {
    ID        int       `json:"id"`
    Email     string    `json:"email"`
    Title     string    `json:"title"`
    CreatedAt time.Time `json:"created_at"`
    UpdatedAt time.Time `json:"updated_at"`
}

func FromCoreList(data []activities.Core) []Activities {
    result := []Activities{}
    for key := range data {
        result = append(result, FromCore(data[key]))
    }
    return result
}

func FromCore(data activities.Core) Activities {
    return Activities{
        ID:        data.ID,
        Email:     data.Email,
        Title:     data.Title,
        CreatedAt: data.CreatedAt,
        UpdatedAt: data.UpdatedAt,
    }
}
