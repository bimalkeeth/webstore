package tokenvalidator
import (
	"net/http"
	"strings"
	"github.com/dgrijalva/jwt-go"
	"fmt"
	"encoding/json"
	"github.com/gorilla/context"
	 conf "cinamaservice/common/contacts/configuration"
	 com "cinamaservice/common"
)
func ValidateAccess(next http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		cinemaAuthorizationHeader := req.Header.Get("x-access-token")
		if cinemaAuthorizationHeader != "" {
			cinemaToken := strings.Split("Bearer " +cinemaAuthorizationHeader, " ")
			if len(cinemaToken) == 2 {
				token, error := jwt.Parse(cinemaToken[1], func(token *jwt.Token) (interface{}, error) {
					if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
						return nil, fmt.Errorf(com.GetErrorDescription("ERR005_AUTHOERROR"))
					}
					return []byte("secret"), nil
				})
				if error != nil {
					json.NewEncoder(w).Encode(conf.Exception{Message: error.Error()})
					return
				}
				if token.Valid {
					context.Set(req, "decoded", token.Claims)
					next(w, req)
				} else {
					json.NewEncoder(w).Encode(conf.Exception{Message: com.GetErrorDescription("ERR006_INVALIDTOKEN")})
				}
			}else {
				json.NewEncoder(w).Encode(conf.Exception{Message: com.GetErrorDescription("ERR006_INVALIDTOKEN")})
			}

		} else {
			json.NewEncoder(w).Encode(conf.Exception{Message: com.GetErrorDescription("ERR007_AUTHOHEADER_REQUIRED")})
		}
	})
}