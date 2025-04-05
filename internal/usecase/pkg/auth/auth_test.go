package auth

import (
	"app/internal/entity"
	"app/internal/entity/types"
	mock_auth "app/internal/usecase/pkg/auth/mocks/auth"
	mock_coder "app/internal/usecase/pkg/auth/mocks/coder"
	"testing"
	"time"

	"github.com/gosuit/e"
	"github.com/gosuit/lec"
	"github.com/gosuit/sl"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
)

func TestLogin(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	ctx := lec.New(sl.Default())

	tests := []struct {
		TestName  string
		User      *entity.User
		Tokens    *entity.Tokens
		IsErr     bool
		Err       e.Error
		SetupMock func(user *mock_auth.MockUserStorage, coder *mock_coder.MockCoder, jwt *mock_auth.MockJwtUseCase)
	}{
		{
			TestName: "Successful",
			User: &entity.User{
				Id:       1,
				Email:    "test",
				Name:     "test",
				Password: "test",
				Age:      10,
				Role:     types.USER,
				Verified: true,
			},
			Tokens: &entity.Tokens{
				Access:  "test",
				Refresh: "test",
			},
			IsErr: false,
			SetupMock: func(user *mock_auth.MockUserStorage, coder *mock_coder.MockCoder, jwt *mock_auth.MockJwtUseCase) {
				user.EXPECT().GetByEmail(ctx, "test").Return(&entity.User{
					Id:       1,
					Email:    "test",
					Name:     "test",
					Password: "test",
					Age:      10,
					Role:     types.USER,
					Verified: true,
				}, nil)

				coder.EXPECT().CompareHash("test", "test").Return(nil)

				jwt.EXPECT().GenerateToken(&entity.User{
					Id:       1,
					Email:    "test",
					Name:     "test",
					Password: "test",
					Age:      10,
					Role:     types.USER,
					Verified: true,
				}, 1*time.Hour, false).Return("test", nil)

				jwt.EXPECT().GenerateToken(&entity.User{
					Id:       1,
					Email:    "test",
					Name:     "test",
					Password: "test",
					Age:      10,
					Role:     types.USER,
					Verified: true,
				}, 72*time.Hour, true).Return("test", nil)
			},
		},
		{
			TestName: "Not found",
			User: &entity.User{
				Email: "test",
			},
			IsErr: true,
			Err:   e.New("Not found", e.NotFound),
			SetupMock: func(user *mock_auth.MockUserStorage, coder *mock_coder.MockCoder, jwt *mock_auth.MockJwtUseCase) {
				user.EXPECT().GetByEmail(ctx, "test").Return(nil, e.New("Not found", e.NotFound))
			},
		},

		// TODO: add more test cases
	}

	for _, tc := range tests {
		t.Run(tc.TestName, func(t *testing.T) {
			userStorage := mock_auth.NewMockUserStorage(ctrl)
			coder := mock_coder.NewMockCoder(ctrl)
			jwt := mock_auth.NewMockJwtUseCase(ctrl)
			tc.SetupMock(userStorage, coder, jwt)

			auth := Auth{
				user:  userStorage,
				coder: coder,
				jwt:   jwt,
			}

			res, err := auth.Login(ctx, tc.User)
			if err != nil {
				if !tc.IsErr {
					t.Errorf("Test failed. Error: %v", err.Error())
				} else {
					assert.Equal(t, tc.Err.GetCode(), err.GetCode())
				}
			} else {
				if !tc.IsErr {
					assert.Equal(t, res, tc.Tokens)
				} else {
					t.Error("Error was expected")
				}
			}
		})
	}
}

// TODO: Test other funcs
