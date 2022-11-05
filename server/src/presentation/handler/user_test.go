package handler

import (
	"context"

	"github.com/kuma0328/web3hackathon/domain/entity"
	"github.com/kuma0328/web3hackathon/domain/repository"
)

// func Test_CreateUser(t *testing.T) {
// 	tests := []struct {
// 		name      string
// 		reqBody   string
// 		reqMethod string
// 		wantCode  int
// 		wantBody  string
// 	}{
// 		{
// 			name:      "正常に動作した場合",
// 			reqBody:   `{"name":"hoge","mail":"hoge@hoge.com",password:"hoge"}`,
// 			wantCode:  http.StatusOK,
// 			reqMethod: http.MethodPost,
// 			wantBody:  `{"name":"hoge","id":1,"mail":"hoge@hoge.com"}`,
// 		},
// 		{
// 			name:      "request bodyが空だった場合、400エラーになる",
// 			reqBody:   ``,
// 			wantCode:  200,
// 			reqMethod: http.MethodPost,
// 			wantBody:  `{"Status":400,"Result":"name empty"}`,
// 		},
// 		{
// 			name:      "POSTリクエスト以外は 404 ",
// 			reqBody:   `{"name":"hoge","mail":"hoge@hoge.com"}`,
// 			wantCode:  404,
// 			reqMethod: http.MethodGet,
// 			wantBody:  `{"Status":404,"Result":"method not allowed"}`,
// 		},
// 	}

// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			gin.SetMode(gin.TestMode)
// 			router := gin.Default()

// 			userHandler := NewUserHandler(&UserUsecaseMock{})
// 			router.POST("/user/signup", userHandler.Sigunup)

// 			body := bytes.NewBufferString(tt.reqBody)
// 			w := httptest.NewRecorder()
// 			req, _ := http.NewRequest(http.MethodGet, "/user/signup", body)
// 			router.ServeHTTP(w, req)

// 			if tt.wantCode != w.Code {
// 				t.Errorf("TestHandler_CreateTask code Error : want %d but got %d", tt.wantCode, w.Code)
// 			}

// 			wantbody := tt.wantBody + "\n"

// 			if tt.wantBody != w.Body.String() && wantbody != w.Body.String() {
// 				t.Errorf("TestHandler_CreateTask body Error : want %s but got %s", tt.wantBody, w.Body.String())
// 			}

// 		})
// 	}
// }

var (
	testCreateUser *entity.User = &entity.User{
		Id:       1,
		Name:     "hoge",
		Mail:     "hoge@hoge.com",
		Password: "hogehoge",
	}
	testUpdateUser *entity.User = &entity.User{
		Id:       1,
		Name:     "hoge",
		Mail:     "hoge2@hoge.com",
		Password: "hogehoge2",
	}
)

var _ IUserUsecaseMock = &UserUsecaseMock{}

type UserUsecaseMock struct {
	repo repository.IUserRepository
}

type IUserUsecaseMock interface {
	CreateUser(context.Context, *entity.User) (*entity.User, error)
	UpdateUser(context.Context, *entity.User) (*entity.User, error)
	GetUserByID(ctx context.Context, id int) (*entity.User, error)
	GetUserByMail(ctx context.Context, mail string) (*entity.User, error)
	DeleteUser(ctx context.Context, id int) error
	LoginUser(ctx context.Context, user *entity.User) error
}

func (h *UserUsecaseMock) CreateUser(ctx context.Context, user *entity.User) (*entity.User, error) {
	return testCreateUser, nil
}

func (ur UserUsecaseMock) UpdateUser(ctx context.Context, user *entity.User) (*entity.User, error) {
	return testUpdateUser, nil
}

func (ur UserUsecaseMock) GetUserByID(ctx context.Context, id int) (*entity.User, error) {
	return testCreateUser, nil
}

func (ur UserUsecaseMock) DeleteUser(ctx context.Context, id int) error {
	return nil
}

func (ur UserUsecaseMock) GetUserByMail(ctx context.Context, mail string) (*entity.User, error) {
	return testCreateUser, nil
}

func (ur UserUsecaseMock) LoginUser(ctx context.Context, user *entity.User) error {
	return nil
}
