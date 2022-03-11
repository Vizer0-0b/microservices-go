package inspection

import "github.com/plasticube/microservices-inspect/utils"

type NewInspectionRequest struct {
  InspectorId     int                 `json:"inspectorId" example:"00336600" binding:"required"`
  Address         string              `json:"address" example:"Beijing Tienanmen" gorm:"unique"`
  Target          string              `json:"target" example:"QRCode content:34342543634"`
  Result          string              `json:"result,omitempty" example:"inspection result"`
  BelongNetwork   string              `json:"belong_network,omitempty" example:"net work name"`
  Status          int                 `json:"status" example:"0"`
  CreateTime      utils.LocalTime     `json:"create_time" example:"2021-02-24 20:19:39" gorm:"type:datetime;autoCreateTime;column:created;column:create_time"`
  InspectTime     utils.LocalTime     `json:"inspect_time" example:"2021-02-24 20:19:39" gorm:"type:datetime;column:inspect_time"`
  ExamineTime     utils.LocalTime     `json:"examine_time,omitempty" example:"2021-02-24 20:19:39" gorm:"type:datetime;column:examine_time"`
  Remark          string              `json:"remark" example:"remarks here."`
}

type PreInspectionRequest struct {
  InspectorId     int                 `json:"inspectorId" example:"00336600" binding:"required"`
  Target          string              `json:"target" example:"QRCode content:34342543634" binding:"required"`
}

type DoInspectionRequest struct {
  InspectorId     int                 `json:"inspectorId" example:"00336600" binding:"required"`
  Target          string              `json:"target" example:"QRCode content:34342543634" binding:"required"`
  Lon             float64             `json:"lon" example:"116.349939" binding:"required"`
  Lat             float64             `json:"lat" example:"39.947689" binding:"required"`
}

type ExamineInspectionRequest struct {
  Examine         int                `json:"examine" example:"0:OK 1:Deny" binding:"required"`
  Remark          string             `json:"remark" example:"remark text ..."`
}
