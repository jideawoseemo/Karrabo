package controller

import (
	"bytes"
	"encoding/json"
	"github.com/Tharolloo/go-playground/config"
	"github.com/sirupsen/logrus"
	"net/http"
	"os"
	"sync"

	"github.com/Tharolloo/go-playground/vin"
	"github.com/gin-gonic/gin"
)

var username = os.Getenv("USERNAME")

var password = os.Getenv("PASSWORD")

type Virtual struct {
	App          *config.AppTools
	accessToken  string
	AuthKeyMutex sync.RWMutex
	//Client config.HTTPClient
}

func (p *Virtual) ConfirmVirtualPayment() gin.HandlerFunc {
	return func(c *gin.Context) {

	}

}

// CreateDynamicVirtualAccount implements vin.Kmainstore.
//func (p *Virtual) CreateDynamicVirtualAccount() gin.HandlerFunc {
//
//
//
//}

// CreateStaticVirtualAccount implements vin.Kmainstore.
func (p *Virtual) CreateStaticVirtualAccount() gin.HandlerFunc {
	return func(c *gin.Context) {
		var staticAccount vin.NinePSBCreateStaticVirtualRequest
		if err := c.ShouldBindJSON(&staticAccount); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		requestBody, err := json.Marshal(staticAccount)
		if err != nil {
			logrus.WithError(err).Error("Error marshaling request body")
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to marshal request body"})
			return
		}

		url := "https://baastest.9psb.com.ng/iva-api/v1/merchant/virtualaccount/create"
		req, err := http.NewRequest("POST", url, bytes.NewBuffer(requestBody))
		if err != nil {
			logrus.WithError(err).Error("Error creating request")
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create request"})
			return
		}
		req.Header.Set("Content-Type", "application/json")
		req.Header.Set("Authorization", "Bearer "+p.accessToken)

		logrus.WithFields(logrus.Fields{
			"url":   url,
			"token": p.accessToken,
		}).Info("Sending request to URL")
		client := &http.Client{}
		resp, err := client.Do(req)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to send request"})
			return
		}
		defer resp.Body.Close()

		if resp.StatusCode != http.StatusOK {
			logrus.WithField("statusCode", resp.StatusCode).Error("Received non-200 response")
			c.JSON(resp.StatusCode, gin.H{"error": "Received non-200 response from CoralPay"})
			return
		}

		var staticResponse vin.NinePSBStaticVirtualAccountResponse
		if err := json.NewDecoder(resp.Body).Decode(&staticResponse); err != nil {
			logrus.WithError(err).Error("fail to decode response")
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to decode response"})
			return
		}
		logrus.WithField("response", staticResponse).Info("Received response")

		c.JSON(http.StatusOK, staticResponse)
	}
}

// CreateDynamicVirtualAccount implements vin.Kmainstore.
func (p *Virtual) CreateDynamicVirtualAccount() gin.HandlerFunc {
	return func(c *gin.Context) {
		var staticAccount vin.NinePSBCreateStaticVirtualRequest
		if err := c.ShouldBindJSON(&staticAccount); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		requestBody, err := json.Marshal(staticAccount)
		if err != nil {
			logrus.WithError(err).Error("Error marshaling request body")
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to marshal request body"})
			return
		}

		url := "https://baastest.9psb.com.ng/iva-api/v1/merchant/virtualaccount/create"
		req, err := http.NewRequest("POST", url, bytes.NewBuffer(requestBody))
		if err != nil {
			logrus.WithError(err).Error("Error creating request")
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create request"})
			return
		}
		req.Header.Set("Content-Type", "application/json")
		req.Header.Set("Authorization", "Bearer "+p.accessToken)

		logrus.WithFields(logrus.Fields{
			"url":   url,
			"token": p.accessToken,
		}).Info("Sending request to URL")
		client := &http.Client{}
		resp, err := client.Do(req)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to send request"})
			return
		}
		defer resp.Body.Close()

		if resp.StatusCode != http.StatusOK {
			logrus.WithField("statusCode", resp.StatusCode).Error("Received non-200 response")
			c.JSON(resp.StatusCode, gin.H{"error": "Received non-200 response from CoralPay"})
			return
		}

		var staticResponse vin.NinePSBStaticVirtualAccountResponse
		if err := json.NewDecoder(resp.Body).Decode(&staticResponse); err != nil {
			logrus.WithError(err).Error("fail to decode response")
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to decode response"})
			return
		}
		logrus.WithField("response", staticResponse).Info("Received response")

		c.JSON(http.StatusOK, staticResponse)
	}
}

// DeactivateVirtualAccount implements vin.Kmainstore.
func (p *Virtual) DeactivateVirtualAccount() gin.HandlerFunc {
	return func(c *gin.Context) {
		var deactivateAccount vin.NinePSBDeactivateVirtualAccountRequest
		if err := c.ShouldBindJSON(&deactivateAccount); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		requestBody, err := json.Marshal(deactivateAccount)
		if err != nil {
			logrus.WithError(err).Error("Error marshaling request body")
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to marshal request body"})
			return
		}

		url := "https://baastest.9psb.com.ng/iva-api/v1/merchant/virtualaccount/deactivate"
		req, err := http.NewRequest("POST", url, bytes.NewBuffer(requestBody))
		if err != nil {
			logrus.WithError(err).Error("Error creating request")
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create request"})
			return
		}
		req.Header.Set("Content-Type", "application/json")
		req.Header.Set("Authorization", "Bearer "+p.accessToken)

		logrus.WithFields(logrus.Fields{
			"url":   url,
			"token": p.accessToken,
		}).Info("Sending request to URL")
		client := &http.Client{}
		resp, err := client.Do(req)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to send request"})
			return
		}
		defer resp.Body.Close()

		if resp.StatusCode != http.StatusOK {
			logrus.WithField("statusCode", resp.StatusCode).Error("Received non-200 response")
			c.JSON(resp.StatusCode, gin.H{"error": "Received non-200 response from CoralPay"})
			return
		}

		var deactivateResponse vin.NinePSBDeactivateVirtualAccountResponse
		if err := json.NewDecoder(resp.Body).Decode(&deactivateResponse); err != nil {
			logrus.WithError(err).Error("fail to decode response")
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to decode response"})
			return
		}
		logrus.WithField("response", deactivateResponse).Info("Received response")

		c.JSON(http.StatusOK, deactivateResponse)
	}
}

// DeactivateVirtualAccount implements vin.Kmainstore.
func (p *Virtual) ReallocateVirtualAccount() gin.HandlerFunc {
	return func(c *gin.Context) {
		var deactivateAccount vin.ReallocateVirtualAccountRequest
		if err := c.ShouldBindJSON(&deactivateAccount); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		requestBody, err := json.Marshal(deactivateAccount)
		if err != nil {
			logrus.WithError(err).Error("Error marshaling request body")
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to marshal request body"})
			return
		}

		url := "https://baastest.9psb.com.ng/iva-api/v1/merchant/virtualaccount/reallocate"
		req, err := http.NewRequest("POST", url, bytes.NewBuffer(requestBody))
		if err != nil {
			logrus.WithError(err).Error("Error creating request")
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create request"})
			return
		}
		req.Header.Set("Content-Type", "application/json")
		req.Header.Set("Authorization", "Bearer "+p.accessToken)

		logrus.WithFields(logrus.Fields{
			"url":   url,
			"token": p.accessToken,
		}).Info("Sending request to URL")
		client := &http.Client{}
		resp, err := client.Do(req)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to send request"})
			return
		}
		defer resp.Body.Close()

		if resp.StatusCode != http.StatusOK {
			logrus.WithField("statusCode", resp.StatusCode).Error("Received non-200 response")
			c.JSON(resp.StatusCode, gin.H{"error": "Received non-200 response from CoralPay"})
			return
		}

		var reallocateResponse vin.ReallocateVirtualAccountResponse
		if err := json.NewDecoder(resp.Body).Decode(&reallocateResponse); err != nil {
			logrus.WithError(err).Error("fail to decode response")
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to decode response"})
			return
		}
		logrus.WithField("response", reallocateResponse).Info("Received response")

		c.JSON(http.StatusOK, reallocateResponse)
	}
}

// ReactivateVirtualAccount implements vin.Kmainstore.
//func (p *Virtual) ReactivateVirtualAccount() gin.HandlerFunc {
//	return func(c *gin.Context) {
//
//	}
//}

func (p *Virtual) ReactivateVirtualAccount() gin.HandlerFunc {
	return func(c *gin.Context) {
		var reactivateAccount vin.NinePSBReactivateVirtualAccountRequest
		if err := c.ShouldBindJSON(&reactivateAccount); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		requestBody, err := json.Marshal(reactivateAccount)
		if err != nil {
			logrus.WithError(err).Error("Error marshaling request body")
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to marshal request body"})
			return
		}

		url := "https://baastest.9psb.com.ng/iva-api/v1/merchant/virtualaccount/reactivate"
		req, err := http.NewRequest("POST", url, bytes.NewBuffer(requestBody))
		if err != nil {
			logrus.WithError(err).Error("Error creating request")
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create request"})
			return
		}
		req.Header.Set("Content-Type", "application/json")
		req.Header.Set("Authorization", "Bearer "+p.accessToken)

		logrus.WithFields(logrus.Fields{
			"url":   url,
			"token": p.accessToken,
		}).Info("Sending request to URL")
		client := &http.Client{}
		resp, err := client.Do(req)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to send request"})
			return
		}
		defer resp.Body.Close()

		if resp.StatusCode != http.StatusOK {
			logrus.WithField("statusCode", resp.StatusCode).Error("Received non-200 response")
			c.JSON(resp.StatusCode, gin.H{"error": "Received non-200 response from CoralPay"})
			return
		}

		var reactivateResponse vin.NinePSBReactivateVirtualAccountResponse
		if err := json.NewDecoder(resp.Body).Decode(&reactivateResponse); err != nil {
			logrus.WithError(err).Error("fail to decode response")
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to decode response"})
			return
		}
		logrus.WithField("response", reactivateResponse).Info("Received response")

		c.JSON(http.StatusOK, reactivateResponse)
	}
}

func NewVirtual(app *config.AppTools) vin.Kmainstore {
	return &Virtual{
		App: app,
	}
}

func (p *Virtual) GetAuthenticationToken() gin.HandlerFunc {
	return func(c *gin.Context) {
		var authRequest vin.NinePSBVirtualAuthenticationRequest
		if err := c.ShouldBindJSON(&authRequest); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		requestBody, err := json.Marshal(authRequest)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to marshal request body"})
			return
		}
		url := "https://baastest.9psb.com.ng/iva-api/v1/merchant/virtualaccount/authenticate"
		req, err := http.NewRequest("POST", url, bytes.NewBuffer(requestBody))
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create request"})
			return
		}
		req.Header.Set("Content-Type", "application/json")
		client := &http.Client{}
		resp, err := client.Do(req)
		if err != nil {
			logrus.WithField("error sending request", resp.StatusCode).Error("received non 200 response")
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to send request"})
			return
		}
		defer resp.Body.Close()
		var authResponse vin.NinePSBVirtualAuthenticationResponse
		if err := json.NewDecoder(resp.Body).Decode(&authResponse); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to decode response"})
			return
		}

		p.AuthKeyMutex.Lock()
		//p.AuthKey = authResponse.Key
		p.accessToken = authResponse.Access_token
		p.AuthKeyMutex.Unlock()

		c.JSON(http.StatusOK, authResponse)

	}
}
