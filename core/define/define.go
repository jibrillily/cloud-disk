package define

import (
	"github.com/dgrijalva/jwt-go"
)

type UerClaim struct {
	Id       int
	Identity string
	Name     string
	jwt.StandardClaims
}

var Jwtkey = "cloud-disk-key"

//var MailPassword = os.Getenv("VQLOOVVVKBFHEPLI")
var MailPassword = "VQLOOVVVKBFHEPLI"

//CodeLength 验证码长度
var CodeLength = 6

//CodeExpire 验证码过期时间(s)
var CodeExpire = 300

// TencentSecretKey 腾讯云对象存储
//var TencentSecretKey = os.Getenv("xf6F7j48vSihPpzsTnqc7KHxMULvFH2W")
//var TencentSecretID = os.Getenv("AKIDjJeiW3JJ2UfaDZeRxB8oDBsIRJtCF1Vi")
var TencentSecretKey = "xf6F7j48vSihPpzsTnqc7KHxMULvFH2W"
var TencentSecretID = "AKIDjJeiW3JJ2UfaDZeRxB8oDBsIRJtCF1Vi"

//存储桶
var CosBucket = "https://cloud-disk-1312768688.cos.ap-nanjing.myqcloud.com/"

// PageSize 分页的默认参数
var PageSize = 20

var Datetime = "2006-01-02 15:04:05"

var TokenExpire = 3600
var RefreshTokenExpire = 7200
