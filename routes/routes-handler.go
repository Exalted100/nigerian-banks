package routes

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
)

type Bank struct {
	Bank string `json:"bank"`
	Code []int `json:"code"`
}

func getBanks(c *gin.Context) {
	type accountNumberDto struct {
		AccountNumber string `json:"accountNumber" binding:"required"`
	}
	var accountNumber accountNumberDto
	if err := c.ShouldBindJSON(&accountNumber); err != nil {
		c.AbortWithStatusJSON(400, gin.H{"error": "Please send a valid account number"})
		return
	}
	if len(accountNumber.AccountNumber) != 10 {
		c.AbortWithStatusJSON(400, gin.H{"error": "Please send a valid account number"})
		return
	}

	accountNumberArray := strings.Split(accountNumber.AccountNumber, "")
	var accountNumberFinal [10]int

	for i, accountNumber := range accountNumberArray {
		number, err := strconv.Atoi(accountNumber)

		if err != nil {
			fmt.Println("Error during conversion")
			c.AbortWithStatusJSON(400, gin.H{"error": "Account number should only include digits"})
		}
		accountNumberFinal[i] = number
	}

	banksList := getBanksList(accountNumberFinal)

	c.JSON(200, gin.H{"data": banksList})
	return
}

func checkDigits(accountNumber [10]int, bankCode []int) bool {
	checkBankCode := 0
	if len(bankCode) == 5 {
		checkBankCode = 3*9 + 7*bankCode[0] + 3*bankCode[1] + 3*bankCode[2] + 7*bankCode[3] + 3*bankCode[4]
	} else if len(bankCode) == 3 {
		checkBankCode = 3*0 + 7*0 + 3*0 + 3*bankCode[0] + 7*bankCode[1] + 3*bankCode[2]
	}
	multiplication := checkBankCode + 3*accountNumber[0] + 7*accountNumber[1] + 3*accountNumber[2] + 3*accountNumber[3] + 7*accountNumber[4] + 3*accountNumber[5] + 3*accountNumber[6] + 7*accountNumber[7] + 3*accountNumber[8]
	modulus := 10 - multiplication%10

	if modulus == accountNumber[9] {
		return true
	}
	return false
}

func getBanksList(accountNumber [10]int) []Bank {
	bankCodes := []Bank{
		{Bank: "ABBEY MORTGAGE BANK", Code: []int{7, 0, 7}},
		{Bank: "ACCESS BANK", Code: []int{0, 4, 4}},
		{Bank: "AG MORTGAGE BANK", Code: []int{7, 0, 4}},
		{Bank: "ASO SAVINGS AND LOANS PLC", Code: []int{4, 0, 1}},
		{Bank: "CITIBANK", Code: []int{0, 2, 3}},
		{Bank: "CORONATION MERCHANT BANK", Code: []int{5, 9, 9}},
		{Bank: "DIAMOND BANK", Code: []int{0, 6, 3}},
		{Bank: "ECOBANK NIGERIA", Code: []int{0, 5, 0}},
		{Bank: "ENTERPRISE BANK", Code: []int{0, 8, 4}},
		{Bank: "EQUITORIAL TRUST BANK", Code: []int{2, 3, 2}},
		{Bank: "FBN MORTGAGE LIMITED", Code: []int{4, 1, 3}},
		{Bank: "FBN QUEST MERCHANT BANK", Code: []int{7, 0, 5}},
		{Bank: "FIDELITY BANK", Code: []int{0, 7, 0}},
		{Bank: "FINBANK", Code: []int{0, 8, 5}},
		{Bank: "FIRST BANK OF NIGERIA", Code: []int{0, 1, 1}},
		{Bank: "FIRST CITY MONUMENT BANK", Code: []int{2, 1, 4}},
		{Bank: "FSDH MERCHANT BANK", Code: []int{6, 0, 1}},
		{Bank: "GATEWAY MORTGAGE BANK", Code: []int{7, 0, 6}},
		{Bank: "GUARANTY TRUST BANK", Code: []int{0, 5, 8}},
		{Bank: "HERITAGE BANK", Code: []int{0, 3, 0}},
		{Bank: "IMPERIAL HOMES MORTGAGE BANK", Code: []int{4, 1, 5}},
		{Bank: "INTERCONTINENTAL BANK", Code: []int{0, 6, 9}},
		{Bank: "JAIZ BANK", Code: []int{3, 0, 1}},
		{Bank: "JUBILEE LIFE MORTGAGE BANK", Code: []int{4, 0, 2}},
		{Bank: "KEYSTONE BANK", Code: []int{0, 8, 2}},
		{Bank: "KUDA BANK", Code: []int{9, 0, 5, 6, 7}},
		{Bank: "LAGOS BUILDING INVESTMENT COMPANY", Code: []int{7, 3, 5}},
		{Bank: "MAINSTREET BANK", Code: []int{0, 1, 4}},
		{Bank: "NIBSS PSEUDO BANK", Code: []int{9, 9, 8}},
		{Bank: "NOVA MERCHANT BANK", Code: []int{7, 4, 2}},
		{Bank: "OCEANIC BANK", Code: []int{0, 5, 6}},
		{Bank: "PARALLEX", Code: []int{5, 0, 2}},
		{Bank: "POLARIS BANK LIMITED", Code: []int{0, 7, 6}},
		{Bank: "PROVIDUS BANK", Code: []int{1, 0, 1}},
		{Bank: "SAFETRUST MORTGAGE BANK", Code: []int{4, 0, 3}},
		{Bank: "SKYE BANK", Code: []int{0, 7, 6}},
		{Bank: "STANBIC IBTC BANK", Code: []int{2, 2, 1}},
		{Bank: "STANDARD CHARTERED BANK", Code: []int{0, 6, 8}},
		{Bank: "STERLING BANK", Code: []int{2, 3, 2}},
		{Bank: "SUNTRUST", Code: []int{1, 0, 0}},
		{Bank: "TRUSTBOND MORTGAGE", Code: []int{5, 2, 3}},
		{Bank: "UNION BANK OF NIGERIA", Code: []int{0, 3, 2}},
		{Bank: "UNITED BANK FOR AFRICA", Code: []int{0, 3, 3}},
		{Bank: "UNITY BANK", Code: []int{2, 1, 5}},
		{Bank: "WEMA BANK", Code: []int{0, 3, 5}},
		{Bank: "ZENITH BANK", Code: []int{0, 5, 7}},
	}
	bankList := []Bank{}
	for _, bank := range bankCodes {
		digitValid := checkDigits(accountNumber, bank.Code)
		if digitValid {
			bankList = append(bankList, Bank{Bank: bank.Bank, Code: bank.Code})
		}
	}
	return bankList
}
