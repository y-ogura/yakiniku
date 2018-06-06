package usecase

import model "github.com/y-ogura/yakiniku/account"

func (a *accountUsecase) List() ([]*model.Client, error) {
	accounts, err := a.accountRepository.List()
	if err != nil {
		return nil, err
	}
	res := formatClientList(accounts)
	return res, nil
}
