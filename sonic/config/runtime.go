package config

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// FIXME: it's named Runtime, so this is kind of application running context.
// But what I see in the code that it's used just wrap configs in one struct and pass to 'Setups'
// 1. If usage will not change it can be removed.
// 2. If we change the usage to match it's name and purpose, I believe it should be implement
// Service Locator (Service Discovery) Pattern.

// Runtime configuration of the application
type Runtime struct {
	Conf    Config
	DB      *gorm.DB
	Srv     *gin.Engine
	Version string
}
