package controllers

import (
	"errors"
	"github.com/golang/mock/gomock"
	"github.com/labstack/echo/v4"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"go-unit-test/internal/usecases"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestAuthRepository(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Controller Suite")
}

var _ = Describe("ControllerTest", func() {
	var (
		ctrl        *gomock.Controller
		authUseCase *usecases.MockAuthUseCase
		controller  ControllerInterface
		e           *echo.Echo
		rec         *httptest.ResponseRecorder
	)

	BeforeEach(func() {
		ctrl = gomock.NewController(GinkgoT())
		authUseCase = usecases.NewMockAuthUseCase(ctrl)
		controller = NewController(authUseCase)
		e = echo.New()
		rec = httptest.NewRecorder()
	})

	AfterEach(func() {
		ctrl.Finish()
	})

	Describe("SignUp", func() {
		Context("when sign up is successful", func() {
			It("should return 200 OK", func() {
				reqBody := `{"username":"john_doe","password":"password123","email":"john@example.com"}`
				req := httptest.NewRequest(http.MethodPost, "/signup", strings.NewReader(reqBody))
				req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
				c := e.NewContext(req, rec)

				authUseCase.EXPECT().SignUp("john_doe", "password123", "john@example.com").Return(nil)

				err := controller.SignUp(c)
				Expect(err).ToNot(HaveOccurred())
				Expect(rec.Code).To(Equal(http.StatusOK))

			})
		})

		Context("when sign up fails", func() {
			It("should return 404 OK", func() {
				reqBody := `{"username":"john_doe","password":"password123","email":"john@example.com"}`
				req := httptest.NewRequest(http.MethodPost, "/signup", strings.NewReader(reqBody))
				req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
				c := e.NewContext(req, rec)

				authUseCase.EXPECT().SignUp("john_doe", "password123", "john@example.com").Return(errors.New("fails"))

				err := controller.SignUp(c)
				Expect(err).ToNot(HaveOccurred())
				Expect(rec.Code).To(Equal(http.StatusNotFound))
				Expect(rec.Body.String()).To(ContainSubstring("fails"))

			})
		})

	})

	Describe("SignIn", func() {
		Context("when sign in is successful", func() {
			It("should return 200 OK", func() {
				reqBody := `{"username":"john_doe","password":"password123"}`
				req := httptest.NewRequest(http.MethodPost, "/signin", strings.NewReader(reqBody))
				req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
				c := e.NewContext(req, rec)

				authUseCase.EXPECT().SignIn("john_doe", "password123").Return("token", nil)

				err := controller.SignIn(c)
				Expect(err).ToNot(HaveOccurred())
				Expect(rec.Code).To(Equal(http.StatusOK))

			})
		})

		Context("when sign in fails", func() {
			It("should return 404 OK", func() {
				reqBody := `{"username":"john_doe","password":"password123"}`
				req := httptest.NewRequest(http.MethodPost, "/signin", strings.NewReader(reqBody))
				req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
				c := e.NewContext(req, rec)

				authUseCase.EXPECT().SignIn("john_doe", "password123").Return("", errors.New("fails"))

				err := controller.SignIn(c)
				Expect(err).ToNot(HaveOccurred())
				Expect(rec.Code).To(Equal(http.StatusNotFound))
				Expect(rec.Body.String()).To(ContainSubstring("fails"))

			})
		})

	})

})
