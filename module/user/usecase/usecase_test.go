package usecase

//
//import (
//	"context"
//	"github.com/pkg/errors"
//	userDomain "myapp/module/user/domain"
//	"testing"
//)
//
//type mockHasher struct {
//}
//
//func (mockHasher) RandomStr(length int) (string, error) {
//	return "abcd", nil
//}
//
//func (mockHasher) HashPassword(salt, password string) (string, error) {
//	return "abcdefgfasd", nil
//}
//
//func (mockHasher) CompareHashPassword(hashedPassword, salt, password string) bool {
//	return true
//}
//
//type mockUserRepo struct {
//}
//
//func (mockUserRepo) FindByEmail(ctx context.Context, email string) (*userDomain.User, error) {
//	if email == "existed@gmail.com" {
//		return &userDomain.User{}, nil
//	}
//
//	if email == "error@gmail.com" {
//		return nil, errors.New("cannot get record")
//	}
//
//	return &userDomain.User{}, nil
//}
//
//func (mockUserRepo) Create(ctx context.Context, data *userDomain.User) error {
//	return nil
//}
//
//type mockTokenProvider struct {
//}
//
//func (mockTokenProvider) IssueToken(ctx context.Context, id, sub string) (token string, err error) {
//	return "", nil
//}
//func (mockTokenProvider) TokenExpireInSeconds() int {
//	return 0
//}
//func (mockTokenProvider) RefreshExpireInSeconds() int {
//	return 0
//}
//
//type mockSessionRepo struct {
//}
//
//func (mockSessionRepo) Create(ctx context.Context, data *userDomain.Session) error {
//	return nil
//}
//
//func TestUseCase_Register(t *testing.T) {
//	uc := NewUserUseCase(mockUserRepo{}, mockHasher{}, mockTokenProvider{}, mockSessionRepo{})
//
//	type testData struct {
//		Input    EmailPasswordRegistrationDTO
//		Expected error
//	}
//
//	table := []testData{
//		{
//			Input: EmailPasswordRegistrationDTO{
//				FirstName: "Viet",
//				LastName:  "Tran",
//				Email:     "existed@gmail.com",
//				Password:  "123456",
//			},
//			Expected: userDomain.ErrEmailHasExisted,
//		},
//		{
//			Input: EmailPasswordRegistrationDTO{
//				FirstName: "Viet",
//				LastName:  "Tran",
//				Email:     "error@gmail.com",
//				Password:  "123456",
//			},
//			Expected: errors.New("cannot get record"),
//		},
//	}
//
//	for i := range table {
//		actualError := uc.Register(context.Background(), table[i].Input)
//
//		if actualError.Error() != table[i].Expected.Error() {
//			t.Errorf("Register failed. Expected %s, but actual is %s", table[i].Expected.Error(), actualError.Error())
//		}
//	}
//}
