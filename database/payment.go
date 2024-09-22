package database

import (
	structs "rto/struct"
)
func IsUserPaymentInfoPresent(userId string)bool{
	var present int
	query := `select 1 from payments.user_payment_info
	where userId = $1;`
	err := DB.QueryRow(query, userId).Scan(&present)
	if err != nil {
		return false
	}
	if present == 1 {
		return true
	}
	return false
}

func AddUserPaymentInformation(upf structs.UserPaymentsInfo)(string, error){
	var inputArgs []interface{}
	var userPaymentInfoId string

	query := `insert into payments.user_payment_info(userId, UPI_Id, username) values ($1, $2, $3) returning user_payment_Id;`

	inputArgs = append(inputArgs, upf.UserId)
	inputArgs = append(inputArgs, upf.UPIID)
	inputArgs = append(inputArgs, upf.UserName)
	err := DB.QueryRow(query, inputArgs...).Scan(&userPaymentInfoId)
	if err != nil {
		return "", err
	}

	return userPaymentInfoId, nil
}