package business

import (
    "errors"

    "devcode-golang.arief.id/apps/echo/features/activities"
    "github.com/go-playground/validator"
)

type activityAvtCase struct {
    activityData activities.Data
}

func NewActivityBusiness(avtData activities.Data) activities.Business {
    return &activityAvtCase{
        activityData: avtData,
    }
}

func (avtcase *activityAvtCase) GetAllData() (data []activities.Core, err error) {
    data, err = avtcase.activityData.GetAllData()
    return data, err
}

func (avtcase *activityAvtCase) GetData(id int) (data activities.Core, row int, err error) {
    data, row, err = avtcase.activityData.GetData(id)
    return data, row, err
}

func (avtcase *activityAvtCase) InsertData(insert activities.Core) (data activities.Core, row int, err error) {
    v := validator.New()
    errEmail := v.Var(insert.Email, "required,email")
    if errEmail != nil {
        return data, -1, errors.New("invalid format email")
    }
    errTitle := v.Var(insert.Title, "required")
    if errTitle != nil {
        return data, -1, errors.New("title cannot be null")
    }
    rowUnique, _ := avtcase.activityData.UniqueData(insert)
    if rowUnique == 1 {
        return data, -1, errors.New("email already exists")
    }
    data, row, err = avtcase.activityData.InsertData(insert)
    return data, row, err
}

func (avtcase *activityAvtCase) UpdateData(id int, insert activities.Core) (data activities.Core, row int, err error) {
    dataGet, _, _ := avtcase.activityData.GetData(id)
    v := validator.New()
    if insert.Email != "" {
        errEmail := v.Var(insert.Email, "required,email")
        if errEmail != nil {
            return data, -1, errors.New("invalid format email")
        }
        rowUnique, _ := avtcase.activityData.UniqueData(insert)
        if rowUnique == 1 {
            return data, -1, errors.New("email already exists")
        }
    } else if insert.Email == "" {
        insert.Email = dataGet.Email
    }
    if insert.Title == "" {
        insert.Title = dataGet.Title
    }
    data, row, err = avtcase.activityData.UpdateData(id, insert)
    return data, row, err
}

func (avtcase *activityAvtCase) DeleteData(id int) (row int, err error) {
    row, err = avtcase.activityData.DeleteData(id)
    return row, err
}
