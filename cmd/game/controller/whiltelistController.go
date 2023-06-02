package controller

import "context"

func (controller *Controller) GetWhitelist(wallet_address string) (bool, error) {
	_, err := controller.db.Queries.GetWhitelistByWallet(context.Background(), wallet_address)
	if err != nil {
		return false, err
	}
	return true, nil
}
