package inspection

import (
  "errors"
  "fmt"
  "github.com/go-playground/validator/v10"
  errorModels "github.com/plasticube/microservices-inspect/models/errors"
  "strings"
)

func updateValidation(request map[string]interface{}) (err error) {
  var errorsValidation []string

  for k, v := range request {
    if v == "" {
      errorsValidation = append(errorsValidation, fmt.Sprintf("%s cannot be empty", k))
    }
  }

  validationMap := map[string]string{
  }

  validate := validator.New()
  err = validate.RegisterValidation("update_validation", func(fl validator.FieldLevel) bool {
    m, ok := fl.Field().Interface().(map[string]interface{})
    if !ok {
      return false
    }

    for k, v := range validationMap {
      errValidate := validate.Var(m[k], v)
      if errValidate != nil {
        validatorErr := errValidate.(validator.ValidationErrors)
        errorsValidation = append(errorsValidation, fmt.Sprintf("%s do not satisfy condition %v=%v", k, validatorErr[0].Tag(), validatorErr[0].Param()))
      }
    }

    return true
  })

  if err != nil {
    err = errorModels.NewAppError(err, errorModels.UnknownError)
    return
  }

  err = validate.Var(request, "update_validation")
  if err != nil {
    err = errorModels.NewAppError(err, errorModels.UnknownError)
    return
  }
  if errorsValidation != nil {
    err = errorModels.NewAppError(errors.New(strings.Join(errorsValidation, ", ")), errorModels.ValidationError)
  }
  return
}
