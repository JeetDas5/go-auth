package helpers

import (
	"errors"

	"github.com/gin-gonic/gin"
)

func CheckUserType(c *gin.Context, role string) (err error) {
	err = nil
	userType := c.GetString("user_type")

	if userType != role {
		err = errors.New("Unauthorized")
		return err
	}

	return err
}

func MatchUserTypeToUid(c *gin.Context, userId string) (err error) {
	userType := c.GetString("user_type")
	uid := c.GetString("uid")

	if userType == "USER" && uid != userId {
		return errors.New("unauthorized to access this resource")
	}

	return nil
}
