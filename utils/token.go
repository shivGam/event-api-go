package utils

import(
	"time"
	"os"
	"fmt"
	"github.com/golang-jwt/jwt/v5"
)

func GenerateToken(email string,user_id int64) (string,error){
	token,err:=jwt.NewWithClaims(jwt.SigningMethodHS256,jwt.MapClaims{
		"email":email,
		"user_id":user_id,
		"exp":time.Now().Add(time.Hour*2).Unix(),
	}).SignedString([]byte(os.Getenv("JWT_SECRET")))
	if err!=nil{
		return "",err
	}
	return token,nil
}

func VerifyToken(token string) error {
	parsedToken,err:=jwt.Parse(token,func(token *jwt.Token)(interface{},error){
		_,ok:=token.Method.(*jwt.SigningMethodHMAC)
		if !ok{
			return nil,fmt.Errorf("unexpected signing method: %v",token.Header["alg"])
		}
		return []byte(os.Getenv("JWT_SECRET")),nil
	})
	if err!=nil{
		return err
	}
	if !parsedToken.Valid{
		return fmt.Errorf("invalid token")
	}
	return nil
}