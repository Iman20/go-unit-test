package usecases

import (
	"errors"
	"github.com/golang/mock/gomock"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"go-unit-test/internal/models"
	"go-unit-test/internal/repositories"
	"go-unit-test/internal/utils/commons/bcrypt"
	"go-unit-test/internal/utils/commons/jwt"
	bcrypt2 "golang.org/x/crypto/bcrypt"
	"testing"
)

func TestAuthUseCase(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "AuthUseCase Suite")
}

var _ = Describe("AuthUseCase", func() {
	var (
		ctrl               *gomock.Controller
		authRepo           *repositories.MockAuthRepository
		authUseCase        AuthUseCase
		mockPasswordHasher *bcrypt.MockBcryptBuilder
		jwtBuilder         *jwt.MockJwtBuilder
	)

	BeforeEach(func() {
		ctrl = gomock.NewController(GinkgoT())
		authRepo = repositories.NewMockAuthRepository(ctrl)
		mockPasswordHasher = bcrypt.NewMockBcryptBuilder(ctrl)
		jwtBuilder = jwt.NewMockJwtBuilder(ctrl)
		authUseCase = NewAuthUseCase(authRepo, mockPasswordHasher, jwtBuilder)
	})

	AfterEach(func() {
		ctrl.Finish()
	})

	Describe("SignUp", func() {
		Context("when signing up a user successfully", func() {
			It("should sign up a user successfully", func() {
				username := "iman_20"
				password := "password123"
				email := "iman@gmail.com"

				hashPassword, err := bcrypt2.GenerateFromPassword([]byte(password), bcrypt2.DefaultCost)
				Expect(err).ToNot(HaveOccurred())

				mockPasswordHasher.EXPECT().GenerateFromPassword(gomock.Any(), gomock.Any()).Return(hashPassword, nil)

				var captureUser *models.User
				authRepo.EXPECT().CreateUser(gomock.Any()).Return(nil).Do(func(user *models.User) {
					captureUser = user
				})

				err = authUseCase.SignUp(username, password, email)
				Expect(err).ToNot(HaveOccurred())

				Expect(captureUser.Username).To(Equal(username))
				Expect(captureUser.Email).To(Equal(email))

				err = bcrypt2.CompareHashAndPassword([]byte(captureUser.Password), []byte(password))
				Expect(err).ToNot(HaveOccurred())

			})
		})

		Context("when invalid password", func() {
			It("should return an error invalid password", func() {
				username := "iman_20"
				password := ""
				email := "iman@gmail.com"

				err := authUseCase.SignUp(username, password, email)
				Expect(err).To(HaveOccurred())
				Expect(err.Error()).To(Equal("invalid password"))
			})
		})

		Context("when invalid email", func() {
			It("should return an error invalid email", func() {
				username := "iman_20"
				password := "password123"
				email := "iman.gmail.com"

				err := authUseCase.SignUp(username, password, email)
				Expect(err).To(HaveOccurred())
				Expect(err.Error()).To(Equal("invalid email"))
			})
		})

		Context("when bcrypt generate hash password fails", func() {
			It("should return an bcrypt error", func() {
				username := "iman_20"
				password := "password123"
				email := "iman@gmail.com"

				mockPasswordHasher.EXPECT().GenerateFromPassword(gomock.Any(), gomock.Any()).Return(nil, errors.New("bcrypt error"))

				err := authUseCase.SignUp(username, password, email)
				Expect(err).To(HaveOccurred())
				Expect(err.Error()).To(Equal("bcrypt error"))
			})
		})
	})

	Describe("SignIn", func() {
		Context("when signing in a user successfully", func() {
			It("should sign in a user successfully", func() {
				username := "iman_20"
				password := "password123"

				hashPassword, err := bcrypt2.GenerateFromPassword([]byte(password), bcrypt2.DefaultCost)
				Expect(err).ToNot(HaveOccurred())

				mockUser := models.User{
					Username: username,
					Password: string(hashPassword),
				}
				authRepo.EXPECT().FindUserByUsername(gomock.Any()).Return(&mockUser, nil)
				mockPasswordHasher.EXPECT().CompareHashAndPassword([]byte(mockUser.Password), []byte(password)).Return(nil)
				jwtBuilder.EXPECT().GenerateJwt(gomock.Any()).Return("kjskjdnfwesdf", nil)

				_, err = authUseCase.SignIn(username, password)
				Expect(err).ToNot(HaveOccurred())

			})
		})

		Context("when invalid username", func() {
			It("should return an error invalid username", func() {
				username := ""
				password := "password123"

				_, err := authUseCase.SignIn(username, password)
				Expect(err).To(HaveOccurred())
				Expect(err.Error()).To(Equal("invalid username"))
			})
		})

		Context("when password not match", func() {
			It("should return an error invalid credentials", func() {
				username := "iman_20"
				password := "password123"

				hashPassword, err := bcrypt2.GenerateFromPassword([]byte(password), bcrypt2.DefaultCost)
				Expect(err).ToNot(HaveOccurred())

				mockUser := models.User{
					Username: username,
					Password: string(hashPassword),
				}

				authRepo.EXPECT().FindUserByUsername(gomock.Any()).Return(&mockUser, nil)
				mockPasswordHasher.EXPECT().CompareHashAndPassword(gomock.Any(), gomock.Any()).Return(errors.New("invalid credentials"))

				_, err = authUseCase.SignIn(username, password)
				Expect(err).To(HaveOccurred())
				Expect(err.Error()).To(Equal("invalid credentials"))
			})
		})

		Context("when token fails", func() {
			It("should return errors", func() {
				username := "iman_20"
				password := "password123"

				hashPassword, err := bcrypt2.GenerateFromPassword([]byte(password), bcrypt2.DefaultCost)
				Expect(err).ToNot(HaveOccurred())

				mockUser := models.User{
					Username: username,
					Password: string(hashPassword),
				}
				authRepo.EXPECT().FindUserByUsername(gomock.Any()).Return(&mockUser, nil)
				mockPasswordHasher.EXPECT().CompareHashAndPassword([]byte(mockUser.Password), []byte(password)).Return(nil)
				jwtBuilder.EXPECT().GenerateJwt(gomock.Any()).Return("", errors.New("token fails"))

				_, err = authUseCase.SignIn(username, password)
				Expect(err).To(HaveOccurred())
				Expect(err.Error()).To(Equal("token fails"))

			})
		})

	})
})
