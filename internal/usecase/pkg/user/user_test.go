package user

import (
	"app/internal/entity"
	"app/internal/entity/types"
	mock_user "app/internal/usecase/pkg/user/mocks/user"
	"testing"

	"github.com/gosuit/e"
	"github.com/gosuit/lec"
	"github.com/gosuit/sl"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
)

func TestGet(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	ctx := lec.New(sl.Default())

	tests := []struct {
		TestName  string
		User      *entity.User
		IsErr     bool
		Err       e.Error
		SetupMock func(mock *mock_user.MockUserStorage)
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
			IsErr: false,
			SetupMock: func(mock *mock_user.MockUserStorage) {
				mock.EXPECT().GetById(ctx, uint64(1)).Return(&entity.User{
					Id:       1,
					Email:    "test",
					Name:     "test",
					Password: "test",
					Age:      10,
					Role:     types.USER,
					Verified: true,
				}, nil)
			},
		},
		{
			TestName: "User not found",
			User: &entity.User{
				Id: 1,
			},
			IsErr: true,
			Err:   e.New("Not found", e.NotFound),
			SetupMock: func(mock *mock_user.MockUserStorage) {
				mock.EXPECT().GetById(ctx, uint64(1)).Return(nil, e.New("Not found", e.NotFound))
			},
		},

		// TODO: add more test cases
	}

	for _, tc := range tests {
		t.Run(tc.TestName, func(t *testing.T) {
			userStorage := mock_user.NewMockUserStorage(ctrl)
			tc.SetupMock(userStorage)

			user := User{
				user: userStorage,
			}

			res, err := user.Get(ctx, tc.User.Id)
			if err != nil {
				if !tc.IsErr {
					t.Errorf("Test failed. Error: %v", err.Error())
				} else {
					assert.Equal(t, tc.Err.GetCode(), err.GetCode())
				}
			} else {
				if !tc.IsErr {
					assert.Equal(t, res, tc.User)
				} else {
					t.Error("Error was expected")
				}
			}
		})
	}
}

// TODO: Test other funcs
