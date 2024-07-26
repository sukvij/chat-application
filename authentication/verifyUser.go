package authentication

import (
	"time"

	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/middleware/jwt"
)

var (
	secret = []byte("signature_hmac_secret_shared_key")
)

type fooClaims struct {
	Email string `json:"email"`
}

func GenerateToken(ctx iris.Context) {
	signer := jwt.NewSigner(jwt.HS256, secret, 100*time.Minute)
	var claims fooClaims
	claims.Email = ctx.Params().Get("userEmail")
	ctx.ReadJSON(&claims)

	token, err := signer.Sign(claims)
	if err != nil {
		ctx.StopWithStatus(iris.StatusInternalServerError)
		return
	}

	ctx.SetCookieKV("token", string(token))
	ctx.Next()
}

func VerifyMiddleware(ctx iris.Context) {
	token := []byte(ctx.GetCookie("token"))
	_, err := jwt.Verify(jwt.HS256, secret, token)
	if err != nil {
		ctx.StopWithJSON(401, "authentication failed")
		return
	}
	ctx.Next()
}

func Logout(ctx iris.Context) {
	cookie := iris.Cookie{
		Name:     "token",
		Value:    "",
		Expires:  time.Now(),
		HttpOnly: true,
	}
	ctx.SetCookie(&cookie)
	ctx.JSON("successfullt logout")
}
