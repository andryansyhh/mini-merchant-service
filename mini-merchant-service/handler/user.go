package handler

import (
	"mini-merchant-service/auth"
	"mini-merchant-service/entity"
	"mini-merchant-service/service"
	"mini-merchant-service/helper"
	"net/http"

	"github.com/gin-gonic/gin"
)

type userHandler struct {
	userService service.UserService
	authService auth.Service
}

func NewUserHandler(userService service.UserService, authService auth.Service) *userHandler {
	return &userHandler{userService, authService}
}

// create user / register handler
func (h *userHandler) RegisterUserHandler(c *gin.Context) {
	var inputUser entity.UserInputs

	if err := c.ShouldBindJSON(&inputUser); err != nil {

		splitError := helper.SplitErrorInformation(err)
		responseError := helper.APIResponse("input data required", 400, "bad request", gin.H{"errors": splitError})

		c.JSON(400, responseError)
		return
	}

	if err := helper.ValidatePassword(inputUser.Password); err != nil {

		responseError := helper.APIResponse("input data required", 400, "bad request", gin.H{"error": "error validate password length < 6"})

		c.JSON(400, responseError)
		return
	}

	newUser, err := h.userService.RegisterUser(inputUser)
	if err != nil {
		responseError := helper.APIResponse("internal server error", 500, "error", gin.H{"errors": err.Error()})

		c.JSON(500, responseError)
		return
	}
	response := helper.APIResponse("success create new user", 201, "Status Created", newUser)
	c.JSON(201, response)
}

// login user handler
func (h *userHandler) LoginUserHandler(c *gin.Context) {
	var inputLoginUser entity.LoginUserInputs

	if err := c.ShouldBindJSON(&inputLoginUser); err != nil {
		responseError := helper.APIResponse("input data required", 400, "bad request", gin.H{"errors": err.Error()})

		c.JSON(400, responseError)
		return
	}

	userData, err := h.userService.LoginUser(inputLoginUser)

	if err != nil {
		responseError := helper.APIResponse("input data required", 400, "bad request", gin.H{"errors": err.Error()})

		c.JSON(400, responseError)
		return
	}

	token, err := h.authService.GenerateTokenUser(userData.ID)

	if err != nil {
		responseError := helper.APIResponse("input data required", 500, "bad request", gin.H{"errors": err.Error()})

		c.JSON(500, responseError)
		return
	}
	response := helper.APIResponse("success login user", 200, "success", gin.H{"token": token, "id": userData.ID})
	c.JSON(200, response)
}

// show all user handler
func (h *userHandler) ShowAllUserHandler(c *gin.Context) {
	user, err := h.userService.ShowAllUser()

	if err != nil {
		responseError := helper.APIResponse("internal server error", 500, "error", gin.H{"errors": err.Error()})

		c.JSON(500, responseError)
		return
	}

	response := helper.APIResponse("success get all user", 200, "status OK", user)
	c.JSON(200, response)
}

// find user by id handler
func (h *userHandler) GetUserByIDHandler(c *gin.Context) {
	id := c.Params.ByName("user_id")

	userUser, err := h.userService.FindUserByID(id)
	if err != nil {
		responseError := helper.APIResponse("input params error", 400, "bad request", gin.H{"errors": err.Error()})

		c.JSON(400, responseError)
		return
	}

	response := helper.APIResponse("success get user by ID", 200, "success", userUser)
	c.JSON(200, response)
}

// update user handler
func (h *userHandler) UpdateUserByIDHandler(c *gin.Context) {
	id := c.Params.ByName("user_id")

	var updateUserInput entity.UpdateUserInputs

	if err := c.ShouldBindJSON(&updateUserInput); err != nil {
		splitError := helper.SplitErrorInformation(err)
		responseError := helper.APIResponse("input data required", 400, "bad request", gin.H{"errors": splitError})

		c.JSON(400, responseError)
		return
	}

	userData := c.MustGet("currentUser").(gin.H)
	userID := userData["user_id"]

	if id != userID {
		responseError := helper.APIResponse("Unauthorize", 401, "error", gin.H{"error": "user ID not authorize"})

		c.JSON(401, responseError)
		return
	}

	userUser, err := h.userService.UpdateUserByID(id, updateUserInput)
	if err != nil {
		responseError := helper.APIResponse("internal server error", 500, "error", gin.H{"error": err.Error()})

		c.JSON(500, responseError)
		return
	}

	response := helper.APIResponse("success update user by ID", http.StatusOK, "success", userUser)
	c.JSON(http.StatusOK, response)
}

// delete user handler
func (h *userHandler) DeleteUserByIDHandler(c *gin.Context) {
	id := c.Params.ByName("user_id")

	userData := c.MustGet("currentUser").(gin.H)
	userID := userData["user_id"]

	if id != userID {
		responseError := helper.APIResponse("Unauthorize", 401, "error", gin.H{"error": "user ID not authorize"})

		c.JSON(401, responseError)
		return
	}

	userUser, err := h.userService.DeleteUserByID(id)

	if err != nil {
		responseError := helper.APIResponse("error bad request delete user", 400, "error", gin.H{"error": err.Error()})

		c.JSON(400, responseError)
		return
	}

	response := helper.APIResponse("success delete user by ID", 200, "success", userUser)
	c.JSON(200, response)
}

// create new outlet hand;er
func (h *userHandler) CreateOutletUserHandler(c *gin.Context) {

	// file, err := c.FormFile("Picture")

	userData := c.MustGet("currentUser").(gin.H)
	userID := userData["user_id"].(string)

	if len(userID) == 0 {
		responseError := helper.APIResponse("Unauthorize", 401, "error", gin.H{"error": "user not authorize / not login"})

		c.JSON(401, responseError)
		return
	}

	var inputOutlet entity.OutletInput

	if err := c.ShouldBindJSON(&inputOutlet); err != nil {

		splitError := helper.SplitErrorInformation(err)
		responseError := helper.APIResponse("input data required", 400, "bad request", gin.H{"errors": splitError})

		c.JSON(400, responseError)
		return
	}

	newOutlet, err := h.userService.CreateOutletUser(inputOutlet, userID)
	if err != nil {
		responseError := helper.APIResponse("internal server error", 500, "error", gin.H{"errors": err.Error()})

		c.JSON(500, responseError)
		return
	}
	response := helper.APIResponse("success create new Outlet", 201, "Status OK", newOutlet)
	c.JSON(201, response)
}

// show outlet hand;er
func (h *userHandler) ShowAllOutletUserHandler(c *gin.Context) {
	outlet, err := h.userService.ShowAllOutletUser()

	if err != nil {
		responseError := helper.APIResponse("internal server error", 500, "error", gin.H{"errors": err.Error()})

		c.JSON(500, responseError)
		return
	}

	response := helper.APIResponse("success get all Outlet", 200, "status OK", outlet)
	c.JSON(200, response)
}
