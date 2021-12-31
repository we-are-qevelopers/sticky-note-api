package useCase

import (
	"github.com/we-are-qevelopers/sticky-note-api/domain/model"
	"github.com/we-are-qevelopers/sticky-note-api/domain/repository"
)

/*
useCaseはアプリケーションができることを記述(In/Outは考えない)
ex) ユーザを作成
このアプリがユーザの作成をできるということのみを考える
どういう実行をされ方をするのかや、どこに保存するのか等は考えない
*/
type UserUsecase interface {
	View() (todo []*model.User, err error)
	Add(*model.User) (err error)
}

/*
func (interactor *UserInteractor) Get(id int) (user domain.UsersForGet, resultStatus *ResultStatus) {
	db := interactor.DB.Connect()
	// Users の取得
	foundUser, err := interactor.User.FindByID(db, id)
	if err != nil {
		return domain.UsersForGet{}, NewResultStatus(404, err)
	}
	user = foundUser.BuildForGet()
	return user, NewResultStatus(200, nil)
}
*/

type userUsecase struct {
	userRepo repository.UserRepository
}

func NewUserUsecase(userRepository repository.UserRepository) UserUsecase {
	newUserUsecase := userUsecase{userRepo: userRepository}

	return &newUserUsecase
}

func (usecase *userUsecase) View() (user []*model.User, err error) {
	user, err = usecase.userRepo.FindAll()
	return
}

func (usecase *userUsecase) Add(user *model.User) (err error) {
	_, err = usecase.userRepo.Create(user)
	return
}
