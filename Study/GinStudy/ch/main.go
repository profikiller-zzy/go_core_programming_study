package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/locales/en"
	"github.com/go-playground/locales/zh"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	enTranslations "github.com/go-playground/validator/v10/translations/en"
	zhTranslations "github.com/go-playground/validator/v10/translations/zh"
	"net/http"
	"reflect"
	"strings"
)

type Login struct {
	User     string `form:"user" json:"user" xml:"user" binding:"required,min=5,max=20"`
	Password string `form:"password" json:"password" xml:"password" binding:"required"`
}

type SignUpParam struct {
	Age        uint8  `json:"age" binding:"gte=1,lte=130"`
	Name       string `json:"name" binding:"required"`
	Email      string `json:"email" binding:"required,email"`
	Password   string `json:"password" binding:"required"`
	RePassword string `json:"re_password" binding:"required,eqfield=Password"`
}

// removeMapPrefix 去除errs.Translate方法返回的map的key中的前缀
func removeMapPrefix(fields map[string]string) map[string]string {
	rsp := map[string]string{}
	for field, err := range fields {
		rsp[field[strings.Index(field, ".")+1:]] = err
	}
	return rsp
}

// 定义全局翻译器
var trans ut.Translator

// InitTrans 初始化翻译器。locale指定翻译的语言
func InitTrans(locale string) (err error) {
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		// 注册一个获取tag中key为json的方法
		v.RegisterTagNameFunc(func(field reflect.StructField) string {
			name := strings.Split(field.Tag.Get("json"), ",")[0] // 只取第一个json tag
			if name == "-" {
				return ""
			}
			return name
		})

		zhT := zh.New() // 中文翻译器
		enT := en.New() // 英文翻译器

		// 通用翻译器, 可以接受多个翻译器，第一个参数是备用（fallback）的翻译器，后面的参数是应该支持的翻译器
		uni := ut.New(enT, zhT, enT)

		var ok bool
		trans, ok = uni.GetTranslator(locale)
		if !ok {
			return fmt.Errorf("uni.GetTranslator(%s) failed", locale)
		}

		// 将翻译器注册到验证器中
		// 使得验证器在返回错误信息时，能够根据注册的翻译器将错误信息翻译成指定的语言
		switch locale {
		case "en":
			err = enTranslations.RegisterDefaultTranslations(v, trans)
		case "zh":
			err = zhTranslations.RegisterDefaultTranslations(v, trans)
		default: // 默认是英文翻译器
			err = enTranslations.RegisterDefaultTranslations(v, trans)
		}
		return err
	}
	return err
}

func main() {
	// 首先初始化翻译器
	if err := InitTrans("zh"); err != nil {
		fmt.Println("初始化翻译器错误", err)
		return
	}

	router := gin.Default()

	router.POST("/loginJSON", func(c *gin.Context) {
		var login Login
		if err := c.ShouldBindJSON(&login); err != nil {
			errs, ok := err.(validator.ValidationErrors)
			if !ok { // 表示这个报错不是 validator.ValidationErrors 类型
				c.JSON(http.StatusOK, gin.H{
					"error": err.Error(),
				})
				return
			} else {
				c.JSON(http.StatusBadRequest, gin.H{
					"error": removeMapPrefix(errs.Translate(trans)),
				})
				return
			}
		}
		if login.User != "manu" || login.Password != "123" { // 模拟验证账号密码
			c.JSON(http.StatusUnauthorized, gin.H{"status": "unauthorized"})
			return
		}
		c.JSON(http.StatusOK, gin.H{"status": "you are logged in"})
	})

	router.POST("/loginForm", func(c *gin.Context) {
		var form Login
		if err := c.ShouldBind(&form); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		if form.User != "manu" || form.Password != "123" {
			c.JSON(http.StatusUnauthorized, gin.H{"status": "unauthorized"})
			return
		}
		c.JSON(http.StatusOK, gin.H{"status": "you are logged in"})
	})

	router.POST("/signup", func(c *gin.Context) {
		var u SignUpParam
		if err := c.ShouldBind(&u); err != nil {
			c.JSON(http.StatusOK, gin.H{
				"msg": err.Error(),
			})
			return
		}
		c.JSON(http.StatusOK, "success")
	})

	router.Run(":8080")
}
