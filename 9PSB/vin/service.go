package vin

import "github.com/gin-gonic/gin"

type Kmainstore interface {
	CreateStaticVirtualAccount() gin.HandlerFunc
	GetAuthenticationToken() gin.HandlerFunc
	CreateDynamicVirtualAccount() gin.HandlerFunc
	ReallocateVirtualAccount() gin.HandlerFunc
	DeactivateVirtualAccount() gin.HandlerFunc
	ReactivateVirtualAccount() gin.HandlerFunc
	ConfirmVirtualPayment() gin.HandlerFunc
}
