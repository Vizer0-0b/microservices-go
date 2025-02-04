package errors

import (
  "github.com/plasticube/microservices-inspect/controllers/svc"
  "net/http"

  "github.com/gin-gonic/gin"
  models "github.com/plasticube/microservices-inspect/models/errors"
)

type MessagesResponse struct {
  Message string `json:"message"`
}

type ServerResponse struct {
  Code int `json:"code"`
  Message string `json:"message"`
}

// Handler is Gin middleware to handle errors.
func Handler(c *gin.Context) {
  // Execute request handlers and then handle any errors
  c.Next()

  errs := c.Errors

  if len(errs) > 0 {
    err, ok := errs[0].Err.(*models.AppError)
    if ok {
      resp := MessagesResponse{Message: err.Error()}
      switch err.Type {
      case models.NotFound:
        c.JSON(http.StatusNotFound, resp)
        return
      case models.ValidationError:
        c.JSON(http.StatusBadRequest, resp)
        return
      case models.ResourceAlreadyExists:
        c.JSON(http.StatusConflict, resp)
        return
      case models.NotAuthenticated:
        c.JSON(http.StatusUnauthorized, resp)
        return
      case models.NotAuthorized:
        c.JSON(http.StatusForbidden, resp)
        return
      case models.RepositoryError:
        c.JSON(http.StatusInternalServerError, MessagesResponse{Message: "We are working to improve the flow of this request."})
        return
      default:
        c.JSON(http.StatusInternalServerError, MessagesResponse{Message: "We are working to improve the flow of this request."})
        return
      }
    }

    c.JSON(http.StatusInternalServerError, ServerResponse{Code:svc.ERROR, Message: errs.String()})
    return
  }
}
