package e2e

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"Epam_final/pkg/front"
)

func TestUser(t *testing.T) {
	createUserBody := front.UserReq{
		Name:    "jon",
		Surname: "jonson",
		Email:   "jonson2@gmail.com",
		Pass:    "jonson12345",
	}
	// Create user
	req, err := makeRequest(http.MethodPost, basePath+"/users", &createUserBody)
	require.NoError(t, err)

	var createUser front.UserRes

	err = doRequest(req, &createUser)
	require.NoError(t, err)

	require.NotNil(t, createUser)
	assert.Equal(t, createUserBody.Name, createUser.Name)
	assert.Equal(t, createUserBody.Surname, createUser.Surname)
	assert.Equal(t, createUserBody.Email, createUser.Email)
	assert.NotEmpty(t, createUser.ID)
	assert.NotEmpty(t, createUser.Status)
	assert.Equal(t, createUser.Status, "active")
	assert.Equal(t, createUser.IsAdmin, false)

	// Get user
	req, err = makeRequest(http.MethodGet, fmt.Sprintf("%s%s%d", basePath, "/users/", createUser.ID), nil)
	require.NoError(t, err)

	var getUser front.UserRes

	err = doRequest(req, &getUser)
	require.NoError(t, err)

	require.NotNil(t, getUser)
	assert.NotEmpty(t, getUser.ID)
	assert.NotEmpty(t, getUser.Name)
	assert.NotEmpty(t, getUser.Surname)
	assert.NotEmpty(t, getUser.Email)
	assert.NotEmpty(t, getUser.Status)

	//Create account
	req, err = makeRequest(http.MethodPost, fmt.Sprintf("%s%s%d%s", basePath, "/users/", createUser.ID, "/accounts"), nil)
	require.NoError(t, err)

	var createAcc front.AccountRes

	err = doRequest(req, &createAcc)
	require.NoError(t, err)

	require.NotNil(t, createAcc)
	assert.NotEmpty(t, createAcc.Acc)
	assert.NotEmpty(t, createAcc.UserID)

	// Get list accounts
	req, err = makeRequest(http.MethodGet, fmt.Sprintf("%s%s%d%s", basePath, "/users/", createUser.ID, "/accounts"), nil)
	require.NoError(t, err)

	var getListAcc []front.AccountRes

	err = doRequest(req, &getListAcc)
	require.NoError(t, err)

	require.NotNil(t, getListAcc)
	assert.NotEmpty(t, getListAcc)
	assert.Equal(t, len(getListAcc), 1)

	//Create card
	req, err = makeRequest(http.MethodPost, fmt.Sprintf("%s%s%d%s", basePath, "/accounts/", createAcc.ID, "/card"), nil)
	require.NoError(t, err)

	var createCard front.CardRes

	err = doRequest(req, &createCard)
	require.NoError(t, err)

	require.NotNil(t, createCard)
	assert.NotEmpty(t, createCard.CardID)
	assert.NotEmpty(t, createCard.AccID)

	//Get card
	req, err = makeRequest(http.MethodGet, fmt.Sprintf("%s%s%d", basePath, "/card/", createCard.ID), nil)
	require.NoError(t, err)

	var getCard front.CardRes

	err = doRequest(req, &getCard)
	require.NoError(t, err)

	require.NotNil(t, getCard)
	assert.NotEmpty(t, getCard.CardID)
	assert.NotEmpty(t, getCard.AccID)

	//Create payments
	createPaymentBody := front.PaymentReq{
		AccID:  createAcc.ID,
		Amount: 200,
	}

	req, err = makeRequest(http.MethodPost, fmt.Sprintf("%s%s", basePath, "/payments"), &createPaymentBody)
	require.NoError(t, err)

	var createPayment front.PaymentsRes

	err = doRequest(req, &createPayment)
	require.NoError(t, err)

	require.NotNil(t, createPayment)
	assert.NotEmpty(t, createPayment.AccID)
	assert.NotEmpty(t, createPayment.Date)
	assert.NotEmpty(t, createPayment.Amount)
	assert.Equal(t, createPayment.Amount, int64(200))

	//Get list payment
	req, err = makeRequest(http.MethodGet, fmt.Sprintf("%s%s%d%s", basePath, "/accounts/", createAcc.ID, "/payments"), nil)
	require.NoError(t, err)

	var getListPayment []front.PaymentsRes

	err = doRequest(req, &getListPayment)
	require.NoError(t, err)

	require.NotNil(t, getListPayment)
	assert.NotEmpty(t, getListPayment)
	assert.Equal(t, len(getListPayment), 1)

	// Create user request
	createUserRequestBody := front.UserRequest{
		AccID: createAcc.ID,
	}
	req, err = makeRequest(http.MethodPost, fmt.Sprintf("%s%s", basePath, "/userRequest"), &createUserRequestBody)
	require.NoError(t, err)

	var createUserRequest front.UserRequestRes
	err = doRequest(req, &createUserRequest)

	require.NotNil(t, createUserRequest)
	require.NotEmpty(t, createUserRequest.AccID)
	require.NotEmpty(t, createUserRequest.Date)

	//Block user
	req, err = makeRequest(http.MethodPut, fmt.Sprintf("%s%s%d%s", basePath, "/admin/users/", createUser.ID, "/block"), nil)
	require.NoError(t, err)

	var blockUser front.UserRes

	err = doRequest(req, &blockUser)
	require.NoError(t, err)

	require.NotNil(t, blockUser)
	require.NotEmpty(t, blockUser.ID)
	require.NotEmpty(t, blockUser.Name)
	require.NotEmpty(t, blockUser.Surname)
	require.NotEmpty(t, blockUser.Email)
	require.NotEmpty(t, blockUser.Status)

	//UnBlock user
	req, err = makeRequest(http.MethodPut, fmt.Sprintf("%s%s%d%s", basePath, "/admin/users/", createUser.ID, "/unblock"), nil)
	require.NoError(t, err)

	var unBlockUser front.UserRes

	err = doRequest(req, &unBlockUser)
	require.NoError(t, err)

	require.NotNil(t, unBlockUser)
	require.NotEmpty(t, unBlockUser.ID)
	require.NotEmpty(t, unBlockUser.Name)
	require.NotEmpty(t, unBlockUser.Surname)
	require.NotEmpty(t, unBlockUser.Email)
	require.NotEmpty(t, unBlockUser.Status)

	//User request approve
	//req, err = makeRequest(http.MethodPut, fmt.Sprintf("%s%s%d%s", basePath, "/admin/request/", createUserRequest.ID, "/approve"), nil)
	//require.NoError(t, err)
	//
	//var userRequestApprove front.UserRequestRes
	//
	//err = doRequest(req, &userRequestApprove)
	//require.NoError(t, err)
	//
	//require.NotNil(t, userRequestApprove)
	//require.NotEmpty(t, userRequestApprove.AccID)
	//require.NotEmpty(t, userRequestApprove.Status)
	//require.NotEmpty(t, userRequestApprove.Date)

	// Delete user
	req, err = makeRequest(http.MethodDelete, fmt.Sprintf("%s%s%d", basePath, "/users/", createUser.ID), nil)
	require.NoError(t, err)

	err = doRequest(req, nil)
	require.NoError(t, err)
}
