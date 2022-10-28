package activities

import "time"

type Core struct {
    ID        int       `json:"id"`
    Email     string    `json:"email"`
    Title     string    `json:"title"`
    CreatedAt time.Time `json:"created_at"`
    UpdatedAt time.Time `json:"updated_at"`
}

type Business interface {
    GetAllData() (data []Core, err error)
    GetData(id int) (data Core, row int, err error)
    InsertData(insert Core) (data Core, row int, err error)
    UpdateData(id int, insert Core) (data Core, row int, err error)
    DeleteData(id int) (row int, err error)
}

type Data interface {
    GetAllData() (data []Core, err error)
    GetData(id int) (data Core, row int, err error)
    InsertData(insert Core) (data Core, row int, err error)
    UpdateData(id int, insert Core) (data Core, row int, err error)
    DeleteData(id int) (row int, err error)
    UniqueData(insert Core) (row int, err error)
}
