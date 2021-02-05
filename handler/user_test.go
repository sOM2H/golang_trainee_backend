package handler

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
	"testing"

	"github.com/jinzhu/gorm"
	"github.com/labstack/echo/v4"
	"github.com/sOM2H/golang_trainee_backend/controllers/comment"
	"github.com/sOM2H/golang_trainee_backend/controllers/post"
	"github.com/sOM2H/golang_trainee_backend/controllers/user"
	"github.com/sOM2H/golang_trainee_backend/db"
	"github.com/sOM2H/golang_trainee_backend/model"
	"github.com/sOM2H/golang_trainee_backend/router"
	"github.com/sOM2H/golang_trainee_backend/store"
	"github.com/stretchr/testify/assert"
)

var (
	d  *gorm.DB
	us user.Store
	ps post.Store
	cs comment.Store
	h  *Handler
	e  *echo.Echo
)

func TestMain(m *testing.M) {
	setup()
	code := m.Run()
	tearDown()
	os.Exit(code)
}

func TestSignUpCaseSuccess(t *testing.T) {
	tearDown()
	setup()
	var (
		reqJSON = `{"user":{"email":"test1@gmail.com","password":"secret1234"}}`
	)
	req := httptest.NewRequest(echo.POST, "/api/users", strings.NewReader(reqJSON))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	assert.NoError(t, h.SignUp(c))
	fmt.Println(rec.Body)
	if assert.Equal(t, http.StatusCreated, rec.Code) {
		m := responseMap(rec.Body.Bytes(), "user")
		assert.Equal(t, "test1@gmail.com", m["email"])
		assert.NotEmpty(t, m["token"])
	}
}

func TestLoginCaseSuccess(t *testing.T) {
	tearDown()
	setup()
	var (
		reqJSON = `{"user":{"email":"test@gmail.com","password":"secret1234"}}`
	)
	req := httptest.NewRequest(echo.POST, "/api/users/login", strings.NewReader(reqJSON))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	assert.NoError(t, h.Login(c))
	if assert.Equal(t, http.StatusOK, rec.Code) {
		m := responseMap(rec.Body.Bytes(), "user")
		assert.Equal(t, "test@gmail.com", m["email"])
		assert.NotEmpty(t, m["token"])
	}
}

func tearDown() {
	_ = d.Close()
	if err := db.DropTestDB(); err != nil {
		log.Fatal(err)
	}
}

func setup() {
	d = db.TestDB()
	db.AutoMigrate(d)
	us = store.NewUserStore(d)
	h = NewHandler(us, ps, cs)
	e = router.New()
	loadFixtures()
}

func loadFixtures() error {
	u1 := model.User{
		Email: "test@gmail.com",
	}
	u1.Password, _ = u1.HashPassword("secret1234")
	if err := us.Create(&u1); err != nil {
		return err
	}

	return nil
}

func responseMap(b []byte, key string) map[string]interface{} {
	var m map[string]interface{}
	json.Unmarshal(b, &m)
	return m[key].(map[string]interface{})
}
