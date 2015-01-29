package BitcoinAPI

// Copyright 2014 ThePiachu. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

import (
	"appengine"
	"errors"
	"github.com/ThePiachu/GoBitcoinAPI/GAEBitcoinJSON"
	"github.com/ThePiachu/Go/Log"
)

func SetAddress(newAddress string, newUsername string, newPassword string) {
	GAEBitcoinJSON.SetAddress(newAddress, newUsername, newPassword)
}

func CreateRawTransaction(c appengine.Context, transactions []map[string]interface{}, outputs map[string]float64) (string, error) {
	//version 0.7 Creates a raw transaction spending given inputs.
	resp, err := GAEBitcoinJSON.CreateRawTransaction(c, 1, []interface{}{transactions, outputs})

	if err != nil {
		return "", err
	}

	result := ""
	if resp["result"] != nil {
		result = resp["result"].(string)
	}
	if resp["error"] != nil {
		err = errors.New(resp["error"].(map[string]interface{})["message"].(string))
	}

	return result, err
}

func GetAccountAddress(c appengine.Context, account string) (string, error) {
	//Returns the current bitcoin address for receiving payments to this account.
	resp, err := GAEBitcoinJSON.GetAccountAddress(c, 1, []interface{}{account})

	if err != nil {
		return "", err
	}

	result := ""
	if resp["result"] != nil {
		result = resp["result"].(string)
	}
	if resp["error"] != nil {
		err = errors.New(resp["error"].(map[string]interface{})["message"].(string))
	}

	return result, err
}

func GetBalanceSimple(c appengine.Context) (float64, error) {
	//If [account] is not specified, returns the server's total available balance.
	//If [account] is specified, returns the balance in the account.
	resp, err := GAEBitcoinJSON.GetBalance(c, 1, nil)

	if err != nil {
		return 0.0, err
	}

	var result float64
	if resp["result"] != nil {
		result = resp["result"].(float64)
	}
	if resp["error"] != nil {
		err = errors.New(resp["error"].(map[string]interface{})["message"].(string))
	}

	return result, err
}

func GetBalance(c appengine.Context, account string, minconf float64) (float64, error) {
	//If [account] is not specified, returns the server's total available balance.
	//If [account] is specified, returns the balance in the account.
	resp, err := GAEBitcoinJSON.GetBalance(c, 1, []interface{}{account, minconf})

	if err != nil {
		return 0.0, err
	}

	var result float64
	if resp["result"] != nil {
		result = resp["result"].(float64)
	}
	if resp["error"] != nil {
		err = errors.New(resp["error"].(map[string]interface{})["message"].(string))
	}

	return result, err
}

func GetBlockHash(c appengine.Context, index int64) (string, error) {
	resp, err := GAEBitcoinJSON.GetBlockHash(c, 1, []interface{}{index})

	if err != nil {
		return "", err
	}

	var result string
	if resp["result"] != nil {
		result = resp["result"].(string)
	}
	if resp["error"] != nil {
		err = errors.New(resp["error"].(map[string]interface{})["message"].(string))
	}

	return result, err
}

func GetBlock(c appengine.Context, hash string) (map[string]interface{}, error) {
	resp, err := GAEBitcoinJSON.GetBlock(c, 1, []interface{}{hash})

	if err != nil {
		return nil, err
	}

	var result map[string]interface{}
	if resp["result"] != nil {
		result = resp["result"].(map[string]interface{})
	}
	if resp["error"] != nil {
		err = errors.New(resp["error"].(map[string]interface{})["message"].(string))
	}

	return result, err
}

func GetNewAddress(c appengine.Context, account string) (string, error) {
	//Returns the current bitcoin address for receiving payments to this account.
	resp, err := GAEBitcoinJSON.GetNewAddress(c, 1, []interface{}{account})

	if err != nil {
		return "", err
	}

	result := ""
	if resp["result"] != nil {
		result = resp["result"].(string)
	}
	if resp["error"] != nil {
		err = errors.New(resp["error"].(map[string]interface{})["message"].(string))
	}

	return result, err
}

func GetReceivedByAccount(c appengine.Context, account string, minconf float64) (string, error) {
	//Returns the current bitcoin address for receiving payments to this account.
	resp, err := GAEBitcoinJSON.GetReceivedByAccount(c, 1, []interface{}{account, minconf})

	if err != nil {
		return "", err
	}

	result := ""
	if resp["result"] != nil {
		result = resp["result"].(string)
	}
	if resp["error"] != nil {
		err = errors.New(resp["error"].(map[string]interface{})["message"].(string))
	}

	return result, err
}

func GetReceivedByAddress(c appengine.Context, address string, minconf float64) (string, error) {
	//Returns the current bitcoin address for receiving payments to this account.
	resp, err := GAEBitcoinJSON.GetReceivedByAddress(c, 1, []interface{}{address, minconf})

	if err != nil {
		return "", err
	}

	result := ""
	if resp["result"] != nil {
		result = resp["result"].(string)
	}
	if resp["error"] != nil {
		err = errors.New(resp["error"].(map[string]interface{})["message"].(string))
	}

	return result, err
}

func GetTransaction(c appengine.Context, txid string) (string, error) {
	//Returns the current bitcoin address for receiving payments to this account.
	resp, err := GAEBitcoinJSON.GetTransaction(c, 1, []interface{}{txid})

	if err != nil {
		return "", err
	}

	result := ""
	if resp["result"] != nil {
		result = resp["result"].(string)
	}
	if resp["error"] != nil {
		err = errors.New(resp["error"].(map[string]interface{})["message"].(string))
	}

	return result, err
}

func GetRawTransaction(c appengine.Context, txid string) (string, error) {
	//Returns the current bitcoin address for receiving payments to this account.
	resp, err := GAEBitcoinJSON.GetRawTransaction(c, 1, []interface{}{txid})

	if err != nil {
		return "", err
	}

	result := ""
	if resp["result"] != nil {
		result = resp["result"].(string)
	}
	if resp["error"] != nil {
		err = errors.New(resp["error"].(map[string]interface{})["message"].(string))
	}

	return result, err
}

func GetRawTransactionVerbose(c appengine.Context, txid string) (map[string]interface{}, error) {
	//Returns the current bitcoin address for receiving payments to this account.
	resp, err := GAEBitcoinJSON.GetRawTransaction(c, 1, []interface{}{txid, 1})

	if err != nil {
		return nil, err
	}
	var result map[string]interface{}
	if resp["result"] != nil {
		result = resp["result"].(map[string]interface{})
	}
	if resp["error"] != nil {
		err = errors.New(resp["error"].(map[string]interface{})["message"].(string))
	}
	return result, err
}

func ListAccounts(c appengine.Context, minconf float64) (map[string]float64, error) {
	//Returns Object that has account names as keys, account balances as values.
	resp, err := GAEBitcoinJSON.ListAccounts(c, 1, []interface{}{minconf})
	if err != nil {
		return nil, err
	}

	var result map[string]interface{}
	if resp["result"] != nil {
		result = resp["result"].(map[string]interface{})
	}
	if resp["error"] != nil {
		err = errors.New(resp["error"].(map[string]interface{})["message"].(string))
	}

	answer := map[string]float64{}
	for k, v := range result {
		answer[k] = v.(float64)
	}

	return answer, err
}

func ListReceivedByAccount(c appengine.Context, minconf float64, includeempty bool) ([]map[string]interface{}, error) {
	//
	resp, err := GAEBitcoinJSON.ListReceivedByAccount(c, 1, []interface{}{minconf, includeempty})
	if err != nil {
		return nil, err
	}

	var result []map[string]interface{}
	if resp["result"] != nil {
		result = resp["result"].([]map[string]interface{})
	}
	if resp["error"] != nil {
		err = errors.New(resp["error"].(map[string]interface{})["message"].(string))
	}

	return result, err
}

func ListReceivedByAddress(c appengine.Context, minconf float64, includeempty bool) ([]map[string]interface{}, error) {
	//
	resp, err := GAEBitcoinJSON.ListReceivedByAddress(c, 1, []interface{}{minconf, includeempty})
	if err != nil {
		return nil, err
	}

	var result []map[string]interface{}
	if resp["result"] != nil {
		result = resp["result"].([]map[string]interface{})
	}
	if resp["error"] != nil {
		err = errors.New(resp["error"].(map[string]interface{})["message"].(string))
	}

	return result, err
}

func ListSinceBlock(c appengine.Context, blockhash string, targetConfirmations float64) (map[string]interface{}, error) {
	//Get all transactions in blocks since block [blockhash], or all transactions if omitted.
	//[target-confirmations] intentionally does not affect the list of returned transactions, but only affects the returned "lastblock" value.
	resp, err := GAEBitcoinJSON.ListSinceBlock(c, 1, blockhash, targetConfirmations)
	if err != nil {
		return nil, err
	}

	var result map[string]interface{}
	if resp["result"] != nil {
		result = resp["result"].(map[string]interface{})
	}
	if resp["error"] != nil {
		err = errors.New(resp["error"].(map[string]interface{})["message"].(string))
	}
	return result, err
}

func ListTransactions(c appengine.Context, account string, count float64, from float64) ([]interface{}, error) {
	//
	resp, err := GAEBitcoinJSON.ListTransactions(c, 1, []interface{}{account, count, from})
	if err != nil {
		return nil, err
	}

	var result []interface{}
	if resp["result"] != nil {
		result = resp["result"].([]interface{})
	}
	if resp["error"] != nil {
		err = errors.New(resp["error"].(map[string]interface{})["message"].(string))
	}

	return result, err
}

func ListUnspent(c appengine.Context, minconf int, maxconf int) ([]interface{}, error) {
	//version 0.7 Returns array of unspent transaction inputs in the wallet.
	resp, err := GAEBitcoinJSON.ListUnspent(c, 1, []interface{}{minconf, maxconf})
	if err != nil {
		return nil, err
	}

	var result []interface{}
	if resp["result"] != nil {
		result = resp["result"].([]interface{})
	}
	if resp["error"] != nil {
		err = errors.New(resp["error"].(map[string]interface{})["message"].(string))
	}

	return result, err
}

func Move(c appengine.Context, fromAccount string, toAccount string, amount float64, minconf float64, comment string) (bool, error) {
	//Move from one account in your wallet to another
	resp, err := GAEBitcoinJSON.Move(c, 1, []interface{}{fromAccount, toAccount, amount, minconf, comment})

	if err != nil {
		return false, err
	}
	var result bool
	if resp["result"] != nil {
		result = resp["result"].(bool)
	}
	if resp["error"] != nil {
		err = errors.New(resp["error"].(map[string]interface{})["message"].(string))
	}
	return result, err
}

func SendFrom(c appengine.Context, fromaccount, tobitcoinaddress string, amount float64, minconf float64, comment string, commentTo string) (string, error) {
	//<amount> is a real and is rounded to 8 decimal places. Will send the given amount to the given address, ensuring the account has a valid balance using [minconf] confirmations. Returns the transaction ID if successful (not in JSON object).
	resp, err := GAEBitcoinJSON.SendFrom(c, 1, []interface{}{fromaccount, tobitcoinaddress, amount, minconf, comment, commentTo})
	if err != nil {
		return "", err
	}

	result := ""
	if resp["result"] != nil {
		result = resp["result"].(string)
	}
	if resp["error"] != nil {
		err = errors.New(resp["error"].(map[string]interface{})["message"].(string))
	}

	return result, err
}

func SendFromSimple(c appengine.Context, fromaccount, tobitcoinaddress string, amount float64) (string, error) {
	//<amount> is a real and is rounded to 8 decimal places. Will send the given amount to the given address, ensuring the account has a valid balance using [minconf] confirmations. Returns the transaction ID if successful (not in JSON object).

	Log.Debugf(c, "SendFromSimple - %v, %v, %v", fromaccount, tobitcoinaddress, amount)

	resp, err := GAEBitcoinJSON.SendFrom(c, 1, []interface{}{fromaccount, tobitcoinaddress, amount})

	Log.Debugf(c, "SendFromSimple 2 - %v, %v", resp, err)
	if err != nil {
		return "", err
	}

	result := ""
	if resp["result"] != nil {
		result = resp["result"].(string)
	}
	if resp["error"] != nil {
		err = errors.New(resp["error"].(map[string]interface{})["message"].(string))
	}

	return result, err
}

func SendMany(c appengine.Context, fromAccount string, addressamount map[string]float64, minconf float64, comment string) (string, error) {
	//amounts are double-precision floating point numbers

	resp, err := GAEBitcoinJSON.SendMany(c, 1, []interface{}{fromAccount, addressamount, minconf, comment})

	if err != nil {
		return "", err
	}

	result := ""
	if resp["result"] != nil {
		result = resp["result"].(string)
	}
	if resp["error"] != nil {
		err = errors.New(resp["error"].(map[string]interface{})["message"].(string))
	}

	return result, err
}

func SetAccount(c appengine.Context, bitcoinAddress, account string) error {
	//Sets the account associated with the given address. Assigning address that is already assigned to the same account will create a new address associated with that account.
	resp, err := GAEBitcoinJSON.SetAccount(c, 1, []interface{}{bitcoinAddress, account})

	if err != nil {
		return err
	}
	if resp["error"] != nil {
		err = errors.New(resp["error"].(map[string]interface{})["message"].(string))
	}
	return err
}

func WalletPassPhrase(c appengine.Context, passphrase string, timeout float64) error {
	resp, err := GAEBitcoinJSON.WalletPassPhrase(c, 1, passphrase, timeout)

	if err != nil {
		return err
	}
	if resp["error"] != nil {
		err = errors.New(resp["error"].(map[string]interface{})["message"].(string))
	}
	return err
}

func GetAddressesByAccount(c appengine.Context, account string) ([]string, error) {
	//Returns the list of addresses for the given account.
	resp, err := GAEBitcoinJSON.GetAddressesByAccount(c, 1, []interface{}{account})

	if err != nil {
		return nil, err
	}

	var result []interface{}
	if resp["result"] != nil {
		result = resp["result"].([]interface{})
	}
	if resp["error"] != nil {
		err = errors.New(resp["error"].(map[string]interface{})["message"].(string))
	}

	answer := make([]string, len(result))
	for i := 0; i < len(result); i++ {
		answer[i] = result[i].(string)
	}

	return answer, err
}

func WalletLock(c appengine.Context) error {
	//Removes the wallet encryption key from memory, locking the wallet. After calling this method, you will need to call walletpassphrase again before being able to call any methods which require the wallet to be unlocked.
	resp, err := GAEBitcoinJSON.WalletLock(c, 1)

	if err != nil {
		return err
	}

	if resp["error"] != nil {
		err = errors.New(resp["error"].(map[string]interface{})["message"].(string))
	}

	return err
}

func ValidateAddress(c appengine.Context, bitcoinaddress string) (map[string]interface{}, error) {
	//Return information about <bitcoinaddress>.
	resp, err := GAEBitcoinJSON.ValidateAddress(c, 1, bitcoinaddress)
	if err != nil {
		return nil, err
	}

	var result map[string]interface{}
	if resp["result"] != nil {
		result = resp["result"].(map[string]interface{})
	}
	if resp["error"] != nil {
		err = errors.New(resp["error"].(map[string]interface{})["message"].(string))
	}
	return result, err
}

func VerifyMessage(c appengine.Context, bitcoinaddress string, signature string, message string) (bool, error) {
	//Return information about <bitcoinaddress>.
	resp, err := GAEBitcoinJSON.VerifyMessage(c, 1, bitcoinaddress, signature, message)
	if err != nil {
		return false, err
	}

	result := false
	if resp["result"] != nil {
		result = resp["result"].(bool)
	}
	if resp["error"] != nil {
		err = errors.New(resp["error"].(map[string]interface{})["message"].(string))
	}
	return result, err
}

/*
func BackupWallet(c appengine.Context, id interface{}, destination []interface{})(map[string]interface{}, error){
	//Safely copies wallet.dat to destination, which can be a directory or a path with filename.
	resp, err:=gaehttpjsonrpc.CallWithBasicAuth(c, Address, Username, Password, true, "backupwallet", id, destination)
	if err!=nil{
		//log.Println(err)
		return resp, err
	}
	result:=resp["result"]
	//log.Println(result)

	return resp, err
}
func EncryptWallet(c appengine.Context, id interface{}, passphrase []interface{})(map[string]interface{}, error){
	//Encrypts the wallet with <passphrase>.
	resp, err:=gaehttpjsonrpc.CallWithBasicAuth(c, Address, Username, Password, true, "encryptwallet", id, passphrase)
	if err!=nil{
		//log.Println(err)
		return resp, err
	}
	result:=resp["result"]
	//log.Println(result)

	return resp, err
}

func GetAccount(c appengine.Context, id interface{}, bitcoinaddress []interface{})(map[string]interface{}, error){
	//Returns the account associated with the given address.
	resp, err:=gaehttpjsonrpc.CallWithBasicAuth(c, Address, Username, Password, true, "getaccount", id, bitcoinaddress)
	if err!=nil{
		//log.Println(err)
		return resp, err
	}
	result:=resp["result"]
	//log.Println(result)

	return resp, err
}




func GetBlockCount(c appengine.Context, id interface{})(map[string]interface{}, error){
	//Returns the number of blocks in the longest block chain.
	resp, err:=gaehttpjsonrpc.CallWithBasicAuth(c, Address, Username, Password, true, "getblockcount", id, nil)
	if err!=nil{
		//log.Println(err)
		return resp, err
	}
	result:=resp["result"]
	//log.Println(result)

	return resp, err
}








func GetBlockNumber(c appengine.Context, id interface{})(map[string]interface{}, error){
	//Returns the block number of the latest block in the longest block chain.
	resp, err:=gaehttpjsonrpc.CallWithBasicAuth(c, Address, Username, Password, true, "getblocknumber", id, nil)
	if err!=nil{
		//log.Println(err)
		return resp, err
	}
	result:=resp["result"]
	//log.Println(result)

	return resp, err
}
func GetConnectionCount(c appengine.Context, id interface{})(map[string]interface{}, error){
	//Returns the number of connections to other nodes.
	resp, err:=gaehttpjsonrpc.CallWithBasicAuth(c, Address, Username, Password, true, "getconnectioncount", id, nil)
	if err!=nil{
		//log.Println(err)
		return resp, err
	}
	result:=resp["result"]
	//log.Println(result)

	return resp, err
}
func GetDifficulty(c appengine.Context, id interface{})(map[string]interface{}, error){
	//Returns the proof-of-work difficulty as a multiple of the minimum difficulty.
	resp, err:=gaehttpjsonrpc.CallWithBasicAuth(c, Address, Username, Password, true, "getdifficulty", id, nil)
	if err!=nil{
		//log.Println(err)
		return resp, err
	}
	result:=resp["result"]
	//log.Println(result)

	return resp, err
}
func GetGenerate(c appengine.Context, id interface{})(map[string]interface{}, error){
	//Returns true or false whether bitcoind is currently generating hashes
	resp, err:=gaehttpjsonrpc.CallWithBasicAuth(c, Address, Username, Password, true, "getgenerate", id, nil)
	if err!=nil{
		//log.Println(err)
		return resp, err
	}
	result:=resp["result"]
	//log.Println(result)

	return resp, err
}
func GetHashesPerSec(c appengine.Context, id interface{})(map[string]interface{}, error){
	//Returns a recent hashes per second performance measurement while generating.
	resp, err:=gaehttpjsonrpc.CallWithBasicAuth(c, Address, Username, Password, true, "gethashespersec", id, nil)
	if err!=nil{
		//log.Println(err)
		return resp, err
	}
	result:=resp["result"]
	//log.Println(result)

	return resp, err
}

func GetMemoryPool(c appengine.Context, id interface{}, data []interface{})(map[string]interface{}, error){
	//If [data] is not specified, returns data needed to construct a block to work on:
  	//"version" : block version
  	//"previousblockhash" : hash of current highest block
  	//"transactions" : contents of non-coinbase transactions that should be included in the next block
  	//"coinbasevalue" : maximum allowable input to coinbase transaction, including the generation award and transaction fees
  	//"time" : timestamp appropriate for next block
  	//"bits" : compressed target of next block
	//If [data] is specified, tries to solve the block and returns true if it was successful.
	resp, err:=gaehttpjsonrpc.CallWithBasicAuth(c, Address, Username, Password, true, "getmemorypool", id, data)
	if err!=nil{
		//log.Println(err)
		return resp, err
	}
	result:=resp["result"]
	//log.Println(result)

	return resp, err
}

func GetInfo(c appengine.Context, id interface{})(map[string]interface{}, error){
	//Returns an object containing various state info.
	resp, err:=gaehttpjsonrpc.CallWithBasicAuth(c, Address, Username, Password, true, "getinfo", id, nil)
	if err!=nil{
		//log.Println(err)
		return resp, err
	}
	result:=resp["result"]
	//log.Println(result)

	return resp, err
}

func GetNewAddress(c appengine.Context, id interface{}, account []interface{})(map[string]interface{}, error){
	//Returns a new bitcoin address for receiving payments. If [account] is specified (recommended), it is added to the address book so payments received with the address will be credited to [account].
	resp, err:=gaehttpjsonrpc.CallWithBasicAuth(c, Address, Username, Password, true, "getnewaddress", id, account)
	if err!=nil{
		//log.Println(err)
		return resp, err
	}
	result:=resp["result"]
	//log.Println(result)

	return resp, err
}

func GetReceivedByAccount(c appengine.Context, id interface{}, data []interface{})(map[string]interface{}, error){
	//Returns the total amount received by addresses with [account] in transactions with at least [minconf] confirmations. If [account] not provided return will include all transactions to all accounts. (version 0.3.24-beta)
	resp, err:=gaehttpjsonrpc.CallWithBasicAuth(c, Address, Username, Password, true, "getreceivedbyaccount", id, data)
	if err!=nil{
		//log.Println(err)
		return resp, err
	}
	result:=resp["result"]
	//log.Println(result)

	return resp, err
}

func GetReceivedByAddress(c appengine.Context, id interface{}, data []interface{})(map[string]interface{}, error){
	//Returns the total amount received by <bitcoinaddress> in transactions with at least [minconf] confirmations. While some might consider this obvious, value reported by this only considers *receiving* transactions. It does not check payments that have been made *from* this address. In other words, this is not "getaddressbalance".
	resp, err:=gaehttpjsonrpc.CallWithBasicAuth(c, Address, Username, Password, true, "getreceivedbyaddress", id, data)
	if err!=nil{
		//log.Println(err)
		return resp, err
	}
	result:=resp["result"]
	//log.Println(result)

	return resp, err
}

func GetTransaction(c appengine.Context, id interface{}, txid []interface{})(map[string]interface{}, error){
	//Returns an object about the given transaction containing:
	//"amount" : total amount of the transaction
	//"confirmations" : number of confirmations of the transaction
	//"txid" : the transaction ID
	//"time" : time the transaction occurred
	//"details" - An array of objects containing:
	//"account"
	//"address"
	//"category"
	//"amount"
	resp, err:=gaehttpjsonrpc.CallWithBasicAuth(c, Address, Username, Password, true, "gettransaction", id, txid)
	if err!=nil{
		//log.Println(err)
		return resp, err
	}
	result:=resp["result"]
	//log.Println(result)

	return resp, err
}

func GetWork(c appengine.Context, id interface{}, data []interface{})(map[string]interface{}, error){
	//If [data] is not specified, returns formatted hash data to work on:
	//"midstate" : precomputed hash state after hashing the first half of the data
	//"data" : block data
	//"hash1" : formatted hash buffer for second hash
	//"target" : little endian hash target
	//If [data] is specified, tries to solve the block and returns true if it was successful.
	resp, err:=gaehttpjsonrpc.CallWithBasicAuth(c, Address, Username, Password, true, "getwork", id, data)
	if err!=nil{
		//log.Println(err)
		return resp, err
	}
	//result:=resp["result"]
	//log.Println(result)

	return resp, err
}

func Help(c appengine.Context, id interface{}, command string)(map[string]interface{}, error){
	//List commands, or get help for a command.
	resp, err:=gaehttpjsonrpc.CallWithBasicAuth(c, Address, Username, Password, true, "help", id, []interface{}{command})
	if err!=nil{
		//log.Println(err)
		return resp, err
	}
	result:=resp["result"]
	//log.Println(result)

	return resp, err
}

func KeyPoolRefill(c appengine.Context, id interface{})(map[string]interface{}, error){
	//Fills the keypool, requires wallet passphrase to be set.
	resp, err:=gaehttpjsonrpc.CallWithBasicAuth(c, Address, Username, Password, true, "keypoolrefill", id, nil)
	if err!=nil{
		//log.Println(err)
		return resp, err
	}
	result:=resp["result"]
	//log.Println(result)

	return resp, err
}


func ListReceivedByAccount(c appengine.Context, id interface{}, data []interface{})(map[string]interface{}, error){
	//Returns an array of objects containing:
	//"account" : the account of the receiving addresses
	//"amount" : total amount received by addresses with this account
	//"confirmations" : number of confirmations of the most recent transaction included
	resp, err:=gaehttpjsonrpc.CallWithBasicAuth(c, Address, Username, Password, true, "listreceivedbyaccount", id, data)
	if err!=nil{
		//log.Println(err)
		return resp, err
	}
	result:=resp["result"]
	//log.Println(result)

	return resp, err
}

func ListReceivedByAddress(c appengine.Context, id interface{}, data []interface{})(map[string]interface{}, error){
	//Returns an array of objects containing:
	//"address" : receiving address
	//"account" : the account of the receiving address
	//"amount" : total amount received by the address
	//"confirmations" : number of confirmations of the most recent transaction included
	//To get a list of accounts on the system, execute bitcoind listreceivedbyaddress 0 true
	resp, err:=gaehttpjsonrpc.CallWithBasicAuth(c, Address, Username, Password, true, "listreceivedbyaddress", id, data)
	if err!=nil{
		//log.Println(err)
		return resp, err
	}
	result:=resp["result"]
	//log.Println(result)

	return resp, err
}

func ListTransactions(c appengine.Context, id interface{}, account, count, from interface{})(map[string]interface{}, error){
	//Returns up to [count] most recent transactions skipping the first [from] transactions for account [account]. If [account] not provided will return recent transaction from all accounts.
	resp, err:=gaehttpjsonrpc.CallWithBasicAuth(c, Address, Username, Password, true, "listtransactions", id, []interface{}{account, count, from})
	if err!=nil{
		//log.Println(err)
		return resp, err
	}
	result:=resp["result"]
	//log.Println(result)

	return resp, err
}

func SendToAddress(c appengine.Context, id interface{}, data []interface{})(map[string]interface{}, error){
	//<amount> is a real and is rounded to 8 decimal places. Returns the transaction ID <txid> if successful.
	resp, err:=gaehttpjsonrpc.CallWithBasicAuth(c, Address, Username, Password, true, "sendtoaddress", id, data)
	if err!=nil{
		//log.Println(err)
		return resp, err
	}
	result:=resp["result"]
	//log.Println(result)

	return resp, err
}


func SetGenerate(c appengine.Context, id interface{}, data []interface{})(map[string]interface{}, error){
	//<generate> is true or false to turn generation on or off.
	//Generation is limited to [genproclimit] processors, -1 is unlimited.
	resp, err:=gaehttpjsonrpc.CallWithBasicAuth(c, Address, Username, Password, true, "setgenerate", id, data)
	if err!=nil{
		//log.Println(err)
		return resp, err
	}
	result:=resp["result"]
	//log.Println(result)

	return resp, err
}

func SetTxFee(c appengine.Context, id interface{}, amount []interface{})(map[string]interface{}, error){
	//<amount> is a real and is rounded to the nearest 0.00000001
	resp, err:=gaehttpjsonrpc.CallWithBasicAuth(c, Address, Username, Password, true, "settxfee", id, amount)
	if err!=nil{
		//log.Println(err)
		return resp, err
	}
	result:=resp["result"]
	//log.Println(result)

	return resp, err
}

func SignMessage(c appengine.Context, id interface{}, bitcoinaddress, message interface{})(map[string]interface{}, error){
	//Sign a message with the private key of an address.
	resp, err:=gaehttpjsonrpc.CallWithBasicAuth(c, Address, Username, Password, true, "signmessage", id, []interface{}{bitcoinaddress, message})
	if err!=nil{
		//log.Println(err)
		return resp, err
	}
	result:=resp["result"]
	//log.Println(result)

	return resp, err
}
func Stop(c appengine.Context, id interface{})(map[string]interface{}, error){
	//Stop bitcoin server.
	resp, err:=gaehttpjsonrpc.CallWithBasicAuth(c, Address, Username, Password, true, "stop", id, nil)
	if err!=nil{
		//log.Println(err)
		return resp, err
	}
	result:=resp["result"]
	//log.Println(result)

	return resp, err
}


func VerifyMessage(c appengine.Context, id interface{}, signature, message interface{})(map[string]interface{}, error){
	//Verify a signed message.
	resp, err:=gaehttpjsonrpc.CallWithBasicAuth(c, Address, Username, Password, true, "verifymessage", id, []interface{}{signature, message})
	if err!=nil{
		//log.Println(err)
		return resp, err
	}
	result:=resp["result"]
	//log.Println(result)

	return resp, err
}

func WalletPassPhraseChange(c appengine.Context, id interface{}, data []interface{})(map[string]interface{}, error){
	//Changes the wallet passphrase from <oldpassphrase> to <newpassphrase>.
	resp, err:=gaehttpjsonrpc.CallWithBasicAuth(c, Address, Username, Password, true, "walletpassphrasechange", id, data)
	if err!=nil{
		//log.Println(err)
		return resp, err
	}
	result:=resp["result"]
	//log.Println(result)

	return resp, err
}
*/
