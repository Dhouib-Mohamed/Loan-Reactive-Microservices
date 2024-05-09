package main

import "loan-management/utils"

func main() {
	utils.RequestSubscribe()
	go utils.ValidateRequests()
}
