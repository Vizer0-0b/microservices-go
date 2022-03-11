package models

import (
  "encoding/json"
  "errors"
  "github.com/jinzhu/gorm"
  "github.com/plasticube/microservices-inspect/config"
  modelErrors "github.com/plasticube/microservices-inspect/models/errors"
  "github.com/plasticube/microservices-inspect/utils"
)

type Inspection struct {
  Id              int                 `json:"id" example:"123" gorm:"primaryKey"`
  No              string              `json:"no" example:"1231231" gorm:"unique"`
  InspectorId     int                 `json:"inspector_id" example:"00336600" gorm:"column:inspector_id"`
  Address         string              `json:"address" example:"Beijing Tiananmen" gorm:"unique"`
  Target          string              `json:"target" example:"34342543634"`
  Result          string              `json:"result,omitempty" example:"inspection result"`
  BelongNetwork   string              `json:"belong_network,omitempty" example:"net work name;column:belong_network"`
  Status          int                 `json:"status" example:"0"`
  CreateTime      utils.LocalTime     `json:"create_time" example:"2021-02-24 20:19:39" gorm:"type:datetime;autoCreateTime;column:created;column:create_time"`
  InspectTime     utils.LocalTime     `json:"inspect_time" example:"2021-02-24 20:19:39" gorm:"type:datetime;column:inspect_time"`
  ExamineTime     utils.LocalTime     `json:"examine_time,omitempty" example:"2021-02-24 20:19:39" gorm:"type:datetime;column:examine_time"`
  Remark          string              `json:"remark" example:"remarks here."`
}

func (b *Inspection) TableName() string {
  return "inspection"
}

// GetAllInspection GetAllInspections Fetch all inspection data
func GetAllInspection(inspection *[]Inspection) (err error) {
  err = config.DB.Find(inspection).Error
  if err != nil {
    return err
  }
  return nil
}

// GetAllInspectionByPage GetAllInspections Fetch all inspection data by page
func GetAllInspectionByPage(inspection *[]Inspection, PageCount *int64, pageSize int, pageIndex int) (err error) {
  err = config.DB.Offset((pageIndex - 1) * pageSize).Limit(pageSize).Find(inspection).Error
  config.DB.Model(&Inspection{}).Count(PageCount)

  if err != nil {
    return err
  }
  return nil
}

// CreateInspection ... Insert New data
func CreateInspection(inspection *Inspection) (err error) {
  err = config.DB.Create(inspection).Error

  if err != nil {
    byteErr, _ := json.Marshal(err)
    var newError modelErrors.GormErr
    err = json.Unmarshal(byteErr, &newError)
    if err != nil {
      return err
    }
    switch newError.Number {
    case 1062:
      err = modelErrors.NewAppErrorWithType(modelErrors.ResourceAlreadyExists)
      return

    default:
      err = modelErrors.NewAppErrorWithType(modelErrors.UnknownError)
    }
  }

  return
}

// GetInspectionById ... Fetch only one inspection by Id
func GetInspectionById(inspection *Inspection, id int) (err error) {
  err = config.DB.Where("id = ?", id).First(inspection).Error

  if err != nil {
    switch err.Error() {
    case gorm.ErrRecordNotFound.Error():
      err = modelErrors.NewAppErrorWithType(modelErrors.NotFound)
    default:
      err = modelErrors.NewAppErrorWithType(modelErrors.UnknownError)
    }
  }

  return
}

// UpdateInspection ... Update inspection
func UpdateInspection(id int, inspectionMap map[string]interface{}) (inspection Inspection, err error) {
  inspection.Id = id
  err = config.DB.Model(&inspection).
    Select("inspector_id","address","target","result","belong_network","status","create_time","inspect_time","examine_time","remark").
    Updates(inspectionMap).Error

  // err = config.DB.Save(inspection).Error
  if err != nil {
    byteErr, _ := json.Marshal(err)
    var newError modelErrors.GormErr
    err = json.Unmarshal(byteErr, &newError)
    err = errors.New(newError.Message)
    return inspection, err
  }

  err = config.DB.Where("id = ?", id).First(&inspection).Error

  return
}

// DeleteInspection ... Delete inspection
func DeleteInspection(id int) (err error) {
  tx := config.DB.Delete(&Inspection{}, id)
  if tx.Error != nil {
    err = modelErrors.NewAppErrorWithType(modelErrors.UnknownError)
    return
  }

  if tx.RowsAffected ==  0 {
    err = modelErrors.NewAppErrorWithType(modelErrors.NotFound)
  }

  return
}
