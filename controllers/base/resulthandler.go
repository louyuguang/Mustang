package base

import (
	"Mustang/utils/hack"
	"Mustang/utils/logs"
	"encoding/json"
	"fmt"
	"github.com/astaxie/beego"
	"github.com/go-sql-driver/mysql"
	"net/http"
)

var _ error = &ErrorResult{}

func (e *ErrorResult) Error() string {
	return fmt.Sprintf("code:%d,subCode:%d,msg:%s", e.Code, e.SubCode, e.Msg)
}

type ErrorResult struct {
	// http code
	Code int `json:"code"`
	// The custom code
	SubCode int    `json:"subCode"`
	Msg     string `json:"msg"`
}

type ResultHandlerController struct {
	beego.Controller
}

func (c *ResultHandlerController) Success(data interface{}) {
	c.Ctx.Output.SetStatus(http.StatusOK)
	c.Data["json"] = map[string]interface{}{"status": 0, "msg": data}
	c.ServeJSON()
}

func (c *ResultHandlerController) Fail(data interface{}) {
	errorResult := &ErrorResult{Code: http.StatusOK}
	switch e := data.(type) {
	case *mysql.MySQLError:
		errorResult.SubCode = -1
		if e.Number == 1062 {
			errorResult.Msg = "Resources already exist! "
		} else {
			errorResult.Msg = e.Message
		}
	case string:
		errorResult.SubCode = 1
		errorResult.Msg = e
	case error:
		errorResult.SubCode = -1
		errorResult.Msg = e.Error()
	default:
		errorResult.SubCode = -1
		errorResult.Msg = "Internal server error."
	}
	c.Ctx.Output.SetStatus(errorResult.Code)
	c.Data["json"] = map[string]interface{}{"status": errorResult.SubCode, "msg": errorResult.Msg}
	c.ServeJSON()
}

func (c *ResultHandlerController) errorResult(code int, msg string) []byte {
	errorResult := ErrorResult{
		Code: code,
		Msg:  msg,
	}
	body, err := json.Marshal(errorResult)
	if err != nil {
		logs.Error("Json Marshal error. %v", err)
		c.CustomAbort(http.StatusInternalServerError, http.StatusText(http.StatusInternalServerError))
	}
	return body
}

// Abort stops controller handler and show the error data， e.g. Prepare
func (c *ResultHandlerController) AbortForbidden(msg string) {
	logs.Info("Abort Forbidden error. %s", msg)
	c.CustomAbort(http.StatusForbidden, hack.String(c.errorResult(http.StatusForbidden, msg)))
}

func (c *ResultHandlerController) AbortInternalServerError(msg string) {
	logs.Error("Abort InternalServerError error. %s", msg)
	c.CustomAbort(http.StatusInternalServerError, hack.String(c.errorResult(http.StatusInternalServerError, msg)))
}

func (c *ResultHandlerController) AbortBadRequest(msg string) {
	logs.Info("Abort BadRequest error. %s", msg)
	c.CustomAbort(http.StatusBadRequest, hack.String(c.errorResult(http.StatusBadRequest, msg)))
}

// format BadRequest with param name.
func (c *ResultHandlerController) AbortBadRequestFormat(paramName string) {
	msg := fmt.Sprintf("Invalid param %s !", paramName)
	c.AbortBadRequest(msg)
}

func (c *ResultHandlerController) AbortUnauthorized(msg string) {
	logs.Info("Abort Unauthorized error. %s", msg)
	c.CustomAbort(http.StatusUnauthorized, hack.String(c.errorResult(http.StatusUnauthorized, msg)))
}
