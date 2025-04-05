package activation_code

import (
	"app/internal/entity"
	"context"
	"strconv"
	"strings"
	"testing"

	"github.com/alicebob/miniredis/v2"
	"github.com/gosuit/e"
	"github.com/gosuit/lec"
	"github.com/gosuit/rs"
	"github.com/gosuit/sl"
	"github.com/stretchr/testify/assert"
)

func TestGet(t *testing.T) {
	redis, code := getCode()
	ctx := lec.New(sl.Default())

	tests := []struct {
		Title string
		Code  string
		Id    uint64
		IsErr bool
		Err   e.Error
	}{
		{
			Title: "Successful get",
			Code:  "111111",
			Id:    1,
			IsErr: false,
		},
		{
			Title: "Not found",
			Code:  "000000",
			Id:    0,
			IsErr: true,
			Err:   e.New("Not found", e.NotFound),
		},

		// TODO: Add more test cases
	}

	for _, tc := range tests {
		t.Run(tc.Title, func(t *testing.T) {
			if !tc.IsErr {
				toSet := &entity.ActivationCode{
					Code:   tc.Code,
					UserId: tc.Id,
				}

				setErr := redis.Set(ctx, redisKey(tc.Id), toSet, redisExpires).Err()
				if setErr != nil {
					t.Errorf("unexpected error: %v", setErr.Error())
				}

				res, err := code.Get(ctx, tc.Id)
				if err != nil {
					t.Errorf("unexpected error: %v", err.Error())
				}

				assert.Equal(t, tc.Code, res.Code)
				assert.Equal(t, tc.Id, res.UserId)
			} else {
				if tc.Err.GetCode() == e.NotFound {
					res, err := code.Get(ctx, tc.Id)

					assert.Equal(t, (*entity.ActivationCode)(nil), res)
					assert.Equal(t, e.NotFound, err.GetCode())
				}
			}
		})
	}
}

func TestSet(t *testing.T) {
	redis, code := getCode()
	ctx := lec.New(sl.Default())

	tests := []struct {
		Title string
		Code  string
		Id    uint64
		IsErr bool
		Err   e.Error
	}{
		{
			Title: "Successful set",
			Code:  "111111",
			Id:    1,
			IsErr: false,
		},

		// TODO: Add more test cases
	}

	for _, tc := range tests {
		t.Run(tc.Title, func(t *testing.T) {
			if !tc.IsErr {
				toSet := &entity.ActivationCode{
					Code:   tc.Code,
					UserId: tc.Id,
				}

				err := code.Set(ctx, toSet)
				if err != nil {
					t.Errorf("unexpected error: %v", err.Error())
				}

				var setted entity.ActivationCode

				getErr := redis.Get(ctx, redisKey(tc.Id)).Scan(&setted)
				if getErr != nil {
					t.Errorf("unexpected error: %v", getErr.Error())
				}

				assert.Equal(t, tc.Code, setted.Code)
				assert.Equal(t, tc.Id, setted.UserId)
			}
		})
	}
}

func TestDel(t *testing.T) {
	redis, code := getCode()
	ctx := lec.New(sl.Default())

	tests := []struct {
		Title string
		Code  string
		Id    uint64
		IsErr bool
		Err   e.Error
	}{
		{
			Title: "Successful del",
			Code:  "111111",
			Id:    1,
			IsErr: false,
		},

		// TODO: Add more test cases
	}

	for _, tc := range tests {
		t.Run(tc.Title, func(t *testing.T) {
			if !tc.IsErr {
				toSet := &entity.ActivationCode{
					Code:   tc.Code,
					UserId: tc.Id,
				}

				setErr := redis.Set(ctx, redisKey(tc.Id), toSet, redisExpires).Err()
				if setErr != nil {
					t.Errorf("unexpected error: %v", setErr.Error())
				}

				err := code.Del(ctx, tc.Id)
				if err != nil {
					t.Errorf("unexpected error: %v", err.Error())
				}

				var deleted entity.ActivationCode

				redis.Get(ctx, redisKey(tc.Id)).Scan(&deleted)

				assert.Equal(t, uint64(0), deleted.UserId)
			}
		})
	}
}

func TestRedisKey(t *testing.T) {
	id := uint64(1)
	expectedKey := "activation_codes:1"

	key := redisKey(id)

	assert.Equal(t, expectedKey, key)
}

func getCode() (*rs.Client, *Code) {
	server, err := miniredis.Run()
	if err != nil {
		panic(err)
	}

	parts := strings.Split(server.Addr(), ":")
	port, _ := strconv.ParseInt(parts[1], 10, 32)

	cfg := &rs.Config{
		Host: parts[0],
		Port: int(port),
	}

	client, err := rs.New(context.Background(), cfg)
	if err != nil {
		panic(err)
	}

	return &client, New(client)
}
