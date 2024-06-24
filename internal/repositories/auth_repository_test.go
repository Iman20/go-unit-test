package repositories

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"go-unit-test/config"
	"go-unit-test/internal/models"
	"gorm.io/gorm"
	"testing"
)

func TestAuthRepository(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "AuthRepository Suite")
}

var _ = Describe("AuthRepository", func() {
	var (
		db       *gorm.DB
		authRepo AuthRepository
	)

	BeforeEach(func() {
		db = config.InitMockDB()
		authRepo = NewAuthRepository(db)
	})

	Describe("CreateUser", func() {
		Context("with valid data", func() {
			It("should create new user successfully", func() {
				err := authRepo.CreateUser(&models.User{Username: "iman", Password: "password123", Email: "iman@gmail.com"})
				Expect(err).NotTo(HaveOccurred())
			})
		})

		Context("with invalid data", func() {
			It("should create new user error", func() {
				err := authRepo.CreateUser(&models.User{Username: "iman", Password: "password123", Email: "iman@gmail.com"})
				Expect(err).To(Equal(gorm.ErrInvalidData))
			})
		})
	})

	Describe("FindUserByUsername", func() {
		Context("with valid data", func() {
			It("should return valid user", func() {
				user, err := authRepo.FindUserByUsername("iman")
				Expect(err).To(BeNil())
				Expect(user.Username).Should(Equal("iman"))
			})
		})

		Context("with invalid data", func() {
			It("should return error not found", func() {
				user, err := authRepo.FindUserByUsername("iman lain")
				Expect(err).Should(Equal(gorm.ErrRecordNotFound))
				Expect(user).To(BeNil())
			})
		})
	})

})
