/*
 * 这是默认设置,请设置`customMade`, 打开koroFileHeader查看配置 进行设置: https://github.com/OBKoro1/koro1FileHeader/wiki/%E9%85%8D%E7%BD%AE
 */
package jwt

import (
	"fmt"
	"testing"
)

func TestNewJWT(t *testing.T) {
	JWT := NewJWT([]byte("test"))
	fmt.Println(JWT)
}

func TestCreateToken(t *testing.T) {
	JWT := NewJWT([]byte("test"))
	fmt.Println(JWT)
	token, err := JWT.CreateToken(CustomClaims{
		Id: int64(10010),
	})
	fmt.Println(token, err)
}

func TestParseToken(t *testing.T) {
	JWT := NewJWT([]byte("test"))
	fmt.Println(JWT)
	token, err := JWT.CreateToken(CustomClaims{
		Id: int64(10010),
	})
	fmt.Println(token, err)
	claim, err := JWT.ParseToken(token)
	fmt.Println(claim, err)
}
