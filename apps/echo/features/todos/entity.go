package todos

import "time"

type Core struct {
    ID              int       `json:"id"`
    ActivityGroupID int       `json:"activity_group_id"`
    Title           string    `json:"title"`
    IsActive        bool      `json:"is_active"`
    Priority        string    `json:"priority"`
    CreatedAt       time.Time `json:"created_at"`
    UpdatedAt       time.Time `json:"updated_at"`
}

type Business interface {
    GetAllData(param string) (data []Core, row int, err error)
    GetData(id int) (data Core, row int, err error)
    InsertData(insert Core) (data Core, row int, err error)
    UpdateData(id int, insert Core) (data Core, row int, err error)
    DeleteData(id int) (row int, err error)
}

type Data interface {
    GetAllData(param string) (data []Core, row int, err error)
    GetData(id int) (data Core, row int, err error)
    InsertData(insert Core) (data Core, row int, err error)
    UpdateData(id int, insert Core) (data Core, row int, err error)
    DeleteData(id int) (row int, err error)
}
