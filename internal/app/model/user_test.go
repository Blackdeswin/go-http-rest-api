package model_test

import (
	"http-rest-api/internal/app/model"
	"testing"

	// "github.com/godoctor/godoctor/analysis/names"
	"github.com/stretchr/testify/assert"
)


func TestUser_Validation(t *testing.T){
	// u := model.TestUser(t)
	// assert.NoError(t, u.Validate())

	testCases := []struct{
		name string
		u func() *model.User
		isVaild bool
	}{
		{
			name: "valid", 
			u: func() *model.User {
				return model.TestUser(t)
			},
			isVaild: true,
		},
		{
			name: "with encrypt password", 
			u: func() *model.User {
				u := model.TestUser(t)
				u.Password = ""
				u.EncryptedPassword = "encryptedpassword"

				return u
			},
			isVaild: true,
		},
		{
			name: "empty email", 
			u: func() *model.User {
				u := model.TestUser(t)
				u.Email = ""

				return u
			},
			isVaild: false,
		},
		{
			name: "invalid email", 
			u: func() *model.User {
				u := model.TestUser(t)
				u.Email = "invalid"

				return u
			},
			isVaild: false,
		},
		{
			name: "empty password", 
			u: func() *model.User {
				u := model.TestUser(t)
				u.Password = ""

				return u
			},
			isVaild: false,
		},
		{
			name: "short password", 
			u: func() *model.User {
				u := model.TestUser(t)
				u.Password = "123"

				return u
			},
			isVaild: false,
		},
	
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			if tc.isVaild{
				assert.NoError(t, tc.u().Validate())
			} else {
				assert.Error(t, tc.u().Validate())
			}
		})
	}
}

func TestUser_BeforeCreate(t *testing.T){
	u := model.TestUser(t)
	assert.NoError(t, u.BeforeCreate())
	assert.NotEmpty(t, u.EncryptedPassword)
}