package inspection

import (
  "errors"
  "github.com/gin-gonic/gin"
  "github.com/plasticube/microservices-inspect/controllers"
  "github.com/plasticube/microservices-inspect/controllers/svc"
  "github.com/plasticube/microservices-inspect/tools"
  "net/http"
  "strconv"

  _ "github.com/plasticube/microservices-inspect/controllers/errors"
  "github.com/plasticube/microservices-inspect/models"
  errorModels "github.com/plasticube/microservices-inspect/models/errors"
)

// NewInspection godoc
// @Tags inspection
// @Summary Create New Inspection
// @Description Create new inspection on the system
// @Accept  json
// @Produce  json
// @Param data body NewInspectionRequest true "body data"
// @Success 200 {object} models.Inspection
// @Failure 400 {object} svc.Resp
// @Failure 500 {object} svc.Resp
// @Router /inspection [post]
func NewInspection(c *gin.Context) {
  var request NewInspectionRequest

  if err := controllers.BindJSON(c, &request); err != nil {
    appError := errorModels.NewAppError(err, errorModels.ValidationError)
    _ = c.Error(appError)
    return
  }
  inspection := models.Inspection{
    InspectorId   : request.InspectorId  ,
    Address       : request.Address      ,
    Target        : request.Target       ,
    Result        : request.Result       ,
    BelongNetwork : request.BelongNetwork,
    Status        : request.Status       ,
    CreateTime    : request.CreateTime   ,
    InspectTime   : request.InspectTime  ,
    ExamineTime   : request.ExamineTime  ,
  }

  err := models.CreateInspection(&inspection)
  if err != nil {
    _ = c.Error(err)
    return
  }

  c.JSON(http.StatusOK, svc.Resp{svc.OK, inspection})
}

// GetAllInspection godoc
// @Tags inspection
// @Summary Get all Inspections
// @Description Get all Inspections on the system
// @Success 200 {object} svc.Resp(svc.OK, []models.Inspection)
// @Failure 400 {object} svc.Resp(svc.ERROR, string)
// @Failure 500 {object} svc.Resp(svc.ERROR, string)
// @Router /inspection [get]
func GetAllInspection(c *gin.Context) {
  var inspections []models.Inspection
  err := models.GetAllInspection(&inspections)
  if err != nil {
    appError := errorModels.NewAppErrorWithType(errorModels.UnknownError)
    _ = c.Error(appError)
    return
  }

  c.JSON(http.StatusOK, svc.Resp{svc.OK, inspections})
}

// GetAllInspectionByPage godoc
// @Tags inspection
// @Summary Get all Inspections by page
// @Param page_size query int false "Page size"
// @Param page_index  query int false "Page index"
// @Description Get all Inspections by page on the system
// @Success 200 {object} svc.PageResp{svc.OK,[]models.Inspection}
// @Failure 400 {object} svc.Resp
// @Failure 500 {object} svc.Resp
// @Router /inspection/page [get]
func GetAllInspectionByPage(c *gin.Context) {
  var (
    pageIndex = 1
    pageSize  = 10
    pageCount int64
    err       error
    inspections      []models.Inspection
  )

  size := c.Request.FormValue("page_size")
  index := c.Request.FormValue("page_index")

  if size != "" {
    pageSize = tools.StrToInt(err, size)
  }
  if index != "" {
    pageIndex = tools.StrToInt(err, index)
  }
  err = models.GetAllInspectionByPage(&inspections, &pageCount, pageSize, pageIndex)
  if err != nil {
    appError := errorModels.NewAppErrorWithType(errorModels.UnknownError)
    _ = c.Error(appError)
    return
  }

  c.JSON(http.StatusOK, svc.PageResp{svc.OK, inspections, pageIndex,pageSize,int(pageCount)})
}

// GetInspectionById godoc
// @Tags inspection
// @Summary Get inspections by Id
// @Description Get Inspections by Id on the system
// @Param inspection_id path int true "id of inspection"
// @Success 200 {object} models.Inspection
// @Failure 400 {object} svc.Resp
// @Failure 500 {object} svc.Resp
// @Router /inspection/{inspection_id} [get]
func GetInspectionById(c *gin.Context) {
  var inspection models.Inspection
  inspectionId, err := strconv.Atoi(c.Param("id"))
  if err != nil {
    appError := errorModels.NewAppError(errors.New("inspection id is invalid"), errorModels.ValidationError)
    _ = c.Error(appError)
    return
  }

  err = models.GetInspectionById(&inspection, inspectionId)
  if err != nil {
    appError := errorModels.NewAppError(err, errorModels.ValidationError)
    _ = c.Error(appError)
    return
  }

  c.JSON(http.StatusOK, svc.Resp{svc.OK, inspection})
}

func UpdateInspection(c *gin.Context) {
  inspectionId, err := strconv.Atoi(c.Param("id"))
  if err != nil {
    appError := errorModels.NewAppError(errors.New("param id is necessary in the url"), errorModels.ValidationError)
    _ = c.Error(appError)
    return
  }
  var requestMap map[string]interface{}

  err = controllers.BindJSONMap(c, &requestMap)
  if err != nil {
    appError := errorModels.NewAppError(err, errorModels.ValidationError)
    _ = c.Error(appError)
    return
  }

  err = updateValidation(requestMap)
  if err != nil {
    _ = c.Error(err)
    return
  }

  inspection, err := models.UpdateInspection(inspectionId, requestMap)
  if err != nil {
    _ = c.Error(err)
    return
  }

  c.JSON(http.StatusOK, svc.Resp{svc.OK, inspection})
}

func DeleteInspection(c *gin.Context) {
  inspectionId, err := strconv.Atoi(c.Param("id"))
  if err != nil {
    appError := errorModels.NewAppError(errors.New("param id is necessary in the url"), errorModels.ValidationError)
    _ = c.Error(appError)
    return
  }

  err = models.DeleteInspection(inspectionId)
  if err != nil {
    _ = c.Error(err)
    return
  }

  c.JSON(http.StatusOK, svc.Resp{svc.OK, string("resource deleted successfully")})
}

// DoInspection godoc
// @Tags inspection
// @Summary Post do Inspection
// @Description do Inspection
// @Param inspection_id path int true "id of inspection"
// @Success 200 {object} models.Inspection
// @Failure 400 {object} svc.Resp
// @Failure 500 {object} svc.Resp
// @Router /inspection/{inspection_id} [get]
func DoInspection(c *gin.Context){
  var request NewInspectionRequest

  if err := controllers.BindJSON(c, &request); err != nil {
    appError := errorModels.NewAppError(err, errorModels.ValidationError)
    _ = c.Error(appError)
    return
  }
  inspection := models.Inspection{
    InspectorId   : request.InspectorId  ,
    Target        : request.Target       ,
    BelongNetwork : request.BelongNetwork,
  }

  err := models.CreateInspection(&inspection)
  if err != nil {
    _ = c.Error(err)
    return
  }

  c.JSON(http.StatusOK, svc.Resp{svc.OK, inspection})
}
