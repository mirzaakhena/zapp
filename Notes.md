# Zapp

## Goal
1. As simple as possible
2. Avoid automagically
3. Implement clean architecture
4. As lazy as possible

## Feature
1. Simple RBAC
2. Logging
3. Config
4. User Management (Login, Register, )
5. RESTfull
6. Validator

## Library used
1. github.com/gin-gonic/gin
2. github.com/gin-gonic/autotls
2. github.com/gin-contrib/cors
2. github.com/go-playground/validator
2. github.com/jinzhu/gorm
4. github.com/satori/go.uuid
5. github.com/segmentio/ksuid
6. github.com/spf13/viper
7. github.com/sirupsen/logrus
9. github.com/lestrrat-go/file-rotatelogs
9. github.com/stretchr/testify/assert
9. github.com/rifflock/lfshook
8. github.com/dgrijalva/jwt-go
9. github.com/nsqio/go-nsq
3. golang.org/x/crypto/bcrypt


## Layer
1. Controller
2. Service
3. Repository

## Rest CRUD
1. **Create**

      REQUEST with data body

        POST /person
        {
          ...
        }

      RESPONSE JSON

        {
          "message": "",
          "code": 1,
          "data": {
            ...
          }
        }

2. **GetAll**

      REQUEST with query

        GET /person?
        page=1 &
        size=20 &
        sortBy=someFieldName &
        sortDesc=false &
        f_someFieldName=valueOfField &
        f_anotherField=anotherValue &
        ...
        
      RESPONSE JSON

        {
          "message": "",
          "code": 1,
          "data": {
            "totalItems": 20,
            "items": [
              { ... },
              ...
            ]
          }
        }

3. **GetOne**

      REQUEST with params ID

        GET /person/:personID

      RESPONSE JSON

        {
          "message": "",
          "code": 1,
          "data": {
            ...
          }
        }

4. **Update**

      REQUEST with params ID and data body

        PUT /person/:personID
        {
          ...
        }

      RESPONSE JSON

        {
          "message": "",
          "code": 1,
          "data": {
            ...
          }
        }

5. **Delete**

      REQUEST with params ID

        DELETE /person/:personID
        {
          ...
        }

      RESPONSE JSON

        {
          "message": "",
          "code": 1,
          "data": {
            ...
          }
        }


## Controller
1. Store the userID

## Model
1. Database table structure
2. Will always "extend" BaseModel
3. Request structure
4. Response structure

## Service
1. Have basic CRUD interface with structure

        // IObjService is
        type IObjService interface {
          Create(ctx *Context, obj model.CreateObjRequest) (*model.Obj, error)
          GetOne(ctx *Context, ID string) *model.Obj
          GetAll(ctx *Context, req model.GetAllCommonRequest) *model.GetAllCommonResponse
          Delete(ctx *Context, ID string) (*model.Obj, error)
          Update(ctx *Context, ID string, obj model.UpdateObjRequest) (*model.Obj, error)
        }

2. transaction will used in this layer
3. All method with reserve the first argument for context

## Repository
1. Have basic CRUD interface

        type IOrderJasaRepository interface {
          Create(ctx map[string]interface{}, obj *model.OrderJasa) error
          GetOne(ctx map[string]interface{}, ID string) *model.OrderJasa
          GetAll(ctx map[string]interface{}, page, size int, sortBy string, sortDesc bool, filters map[string]string) ([]model.OrderJasa, uint)
          Delete(ctx map[string]interface{}, ID string) error
          Update(ctx map[string]interface{}, obj *model.OrderJasa) error
        }

