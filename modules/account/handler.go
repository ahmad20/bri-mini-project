package account

import (
	"errors"
	"net/http"

	"github.com/ahmad20/bri-mini-project/entities"
	"github.com/ahmad20/bri-mini-project/modules/auth"
	"github.com/ahmad20/bri-mini-project/modules/customer"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type handler struct {
	accountUseCase  UseCaseInterface
	customerUseCase customer.UseCaseInterface
	auth            auth.AuthInterface
}

type HandlerInterface interface {
	Login(c *gin.Context)
	AdminRegister(c *gin.Context)
	ReadById(c *gin.Context)
	GetAllAdmin(c *gin.Context)
	GetAdminApproval(c *gin.Context)
	UpdateAdminApproval(c *gin.Context)
	UpdateAdminStatus(c *gin.Context)
	AdminDelete(c *gin.Context)

	CustomerRegister(c *gin.Context)
	CustomerDelete(c *gin.Context)
	GetAllCustomer(c *gin.Context)
}

var (
	ErrAccountMissing = errors.New("account does not exist")
	ErrInternalServer = errors.New("internal server error")
	ErrEmptyList      = errors.New("list empty")
	ErrBadRequest     = errors.New("bad request")
)

func NewHandler(accountUseCase UseCaseInterface, customerUseCase customer.UseCaseInterface, auth auth.AuthInterface) HandlerInterface {
	return &handler{
		accountUseCase:  accountUseCase,
		customerUseCase: customerUseCase,
		auth:            auth,
	}
}

func (h *handler) Login(c *gin.Context) {
	var loginRequest LoginRequest

	if err := c.ShouldBindJSON(&loginRequest); err != nil {
		c.JSON(http.StatusBadRequest, ErrBadRequest.Error())
		return
	}
	account, err := h.accountUseCase.SearchByUsername(loginRequest.Username)

	if err != nil {
		if errors.Is(gorm.ErrRecordNotFound, err) {
			c.JSON(http.StatusNotFound, Response{
				Code:    http.StatusNotFound,
				Message: ErrAccountMissing.Error(),
				Data:    nil,
			})
			return
		}
		c.JSON(http.StatusInternalServerError, Response{
			Code:    http.StatusInternalServerError,
			Message: ErrInternalServer.Error(),
			Data:    nil,
		})
		return
	}
	//Password or role matching
	if err != bcrypt.CompareHashAndPassword([]byte(account.Password), []byte(loginRequest.Password)) || account.Role != loginRequest.Role {
		c.JSON(http.StatusUnauthorized, Response{
			Code:    http.StatusUnauthorized,
			Message: "unauthorized user",
			Data:    nil,
		})
		return
	}
	token, err := h.auth.GenerateToken(loginRequest.Username, loginRequest.Role)

	if err != nil {
		c.JSON(http.StatusUnauthorized, Response{
			Code:    http.StatusUnauthorized,
			Message: "unauthorized",
			Data:    nil,
		})
		return
	}
	c.JSON(http.StatusOK, Response{
		Code:    http.StatusOK,
		Message: "login success",
		Data:    token,
	})

}

func (h *handler) AdminRegister(c *gin.Context) {
	var request RegisterRequest

	// Validation Request and Transform
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": ErrBadRequest.Error()})
		return
	}
	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(request.Password), bcrypt.DefaultCost)

	request.Password = string(hashedPassword)

	newAccount := entities.Account{
		Username: request.Username,
		Password: request.Password,
		Role:     request.Role,
	}
	if err := h.accountUseCase.Register(&newAccount); err != nil {
		c.JSON(http.StatusInternalServerError, Response{
			Code:    http.StatusInternalServerError,
			Message: ErrInternalServer.Error(),
			Data:    nil,
		})
		return
	}
	c.JSON(http.StatusOK, Response{
		Code:    http.StatusOK,
		Message: "success",
		Data:    nil,
	})
}

func (h *handler) ReadById(c *gin.Context) {
	userID := c.Param("id")

	var account *entities.Account
	account, err := h.accountUseCase.GetById(userID)
	if err != nil {
		c.JSON(http.StatusNotFound, Response{
			Code:    http.StatusNotFound,
			Message: ErrAccountMissing.Error(),
			Data:    nil,
		})
		return
	}

	c.JSON(http.StatusOK, Response{
		Code:    http.StatusOK,
		Message: "success",
		Data:    account,
	})
}

func (h *handler) GetAllAdmin(c *gin.Context) {
	if c.Request.URL.RawQuery != "" {
		keyword := c.Query("search")
		page := c.Query("page")
		limit := c.Query("limit")

		admins, err := h.accountUseCase.GetAdminsWithConditions(keyword, page, limit)
		if err != nil {
			c.JSON(http.StatusInternalServerError, Response{
				Code:    http.StatusInternalServerError,
				Message: ErrInternalServer.Error(),
				Data:    nil,
			})
			return
		}

		c.JSON(http.StatusOK, Response{
			Code:    http.StatusOK,
			Message: "success",
			Data:    admins,
		})
		return
	} else {
		admins, err := h.accountUseCase.GetAll()
		if err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": ErrEmptyList.Error()})
			return
		}
		c.JSON(http.StatusOK, Response{
			Code:    http.StatusOK,
			Message: "success",
			Data:    admins,
		})
		return
	}
}

func (h *handler) GetAdminApproval(c *gin.Context) {
	accounts, err := h.accountUseCase.GetWaitingApproval()
	if err != nil {
		c.JSON(http.StatusNotFound, Response{
			Code:    http.StatusNotFound,
			Message: ErrEmptyList.Error(),
			Data:    nil,
		})
		return
	}
	c.JSON(http.StatusOK, Response{
		Code:    http.StatusOK,
		Message: "success",
		Data:    accounts,
	})
}

func (h *handler) UpdateAdminApproval(c *gin.Context) {
	userID := c.Param("id")
	var statusRequest StatusRequest

	if err := c.ShouldBindJSON(&statusRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": ErrBadRequest.Error()})
		return
	}
	account, err := h.accountUseCase.GetById(userID)

	if err != nil {
		if errors.Is(gorm.ErrRecordNotFound, err) {
			c.JSON(http.StatusNotFound, gin.H{"error": ErrAccountMissing.Error()})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": ErrInternalServer.Error()})
		return
	}

	if err := h.accountUseCase.UpdateApproval(statusRequest.Status, account); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": ErrInternalServer.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "success"})
}

func (h *handler) UpdateAdminStatus(c *gin.Context) {
	userID := c.Param("id")
	var statusRequest StatusRequest

	if err := c.ShouldBindJSON(&statusRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": ErrBadRequest.Error()})
		return
	}
	account, err := h.accountUseCase.GetById(userID)

	if err != nil {
		if errors.Is(gorm.ErrRecordNotFound, err) {
			c.JSON(http.StatusNotFound, gin.H{"error": ErrAccountMissing.Error()})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": ErrInternalServer.Error()})
		return
	}

	if err := h.accountUseCase.UpdateStatus(statusRequest.Status, account); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": ErrInternalServer.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "success"})
}

func (h *handler) AdminDelete(c *gin.Context) {
	userID := c.Param("id")
	account, err := h.accountUseCase.GetById(userID)
	if err != nil {
		if errors.Is(gorm.ErrRecordNotFound, err) {
			c.JSON(http.StatusNotFound, gin.H{"error": ErrAccountMissing.Error()})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": ErrInternalServer.Error()})
		return
	}
	if err := h.accountUseCase.Delete(account); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": ErrInternalServer.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "success"})
}

func (h *handler) CustomerRegister(c *gin.Context) {
	var customer entities.Customer

	if err := c.ShouldBindJSON(&customer); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": ErrBadRequest})
		return
	}
	if err := h.customerUseCase.Register(&customer); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"err": ErrInternalServer})
		return
	}
	c.JSON(http.StatusOK, Response{
		Code:    http.StatusOK,
		Message: "register success",
		Data:    nil,
	})

}

func (h *handler) CustomerDelete(c *gin.Context) {
	userID := c.Param("id")
	customer, err := h.customerUseCase.GetById(userID)
	if err != nil {
		if errors.Is(gorm.ErrRecordNotFound, err) {
			c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": ErrInternalServer})
	}
	if err := h.customerUseCase.Delete(customer); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": ErrInternalServer})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "success"})
}

func (h *handler) GetAllCustomer(c *gin.Context) {
	if c.Request.URL.RawQuery != "" {
		handleSearchRequest(h, c)
		return
	}

	count, err := h.customerUseCase.CountList(&entities.Customer{})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": ErrInternalServer.Error()})
		return
	}

	if count < 1 {
		data := FetchData()
		for _, item := range data {
			if err := h.customerUseCase.Register(&item); err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to save"})
				return
			}
		}
	}

	customers, err := h.customerUseCase.GetAll()
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"customer": customers})
}

func handleSearchRequest(h *handler, c *gin.Context) {
	keyword := c.Query("search")
	page := c.Query("page")
	limit := c.Query("limit")

	customers, err := h.customerUseCase.GetCustomersWithConditions(keyword, page, limit)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"customer": customers})
}
