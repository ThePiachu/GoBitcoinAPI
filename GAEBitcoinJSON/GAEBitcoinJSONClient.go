package GAEBitcoinJSON

// Copyright 2011-2014 ThePiachu. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

import (
	//"gaehttpjsonrpc"
	//"os"
	"appengine"
	"appengine/urlfetch"
	"encoding/json"
	"github.com/ThePiachu/Go/Log"
	"io/ioutil"
	"net/http"
	"strings"
	"time"
)

var Address string
var Username string
var Password string

func SetAddress(newAddress string, newUsername string, newPassword string) {
	Address = newAddress
	Username = newUsername
	Password = newPassword
}

var TimeoutDuration time.Duration

func init() {
	TimeoutDuration, _ = time.ParseDuration("60s")
}

func CallWithBasicAuth(c appengine.Context, address string, username string, password string, allowInvalidServerCertificate bool, method string, id interface{}, params []interface{}) (map[string]interface{}, error) {
	Log.Debugf(c, "GAEBitcoinJSON CallWithBasicAuth - %v", method)
	data, err := json.Marshal(map[string]interface{}{
		"method": method,
		"id":     id,
		"params": params,
	})
	if err != nil {
		Log.Infof(c, "Marshal: %v", err)
		return nil, err
	}

	req, err := http.NewRequest("POST", address, strings.NewReader(string(data)))
	if err != nil {
		Log.Infof(c, "Request: %v", err)
		return nil, err
	}
	req.SetBasicAuth(username, password)

	tr := &urlfetch.Transport{Context: c, Deadline: TimeoutDuration, AllowInvalidServerCertificate: allowInvalidServerCertificate}

	resp, err := tr.RoundTrip(req)
	if err != nil {
		Log.Infof(c, "Post: %v", err)
		return nil, err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		Log.Infof(c, "ReadAll: %v", err)
		return nil, err
	}
	result := make(map[string]interface{})
	err = json.Unmarshal(body, &result)
	if err != nil {
		Log.Infof(c, "Unmarshal: %v", err)
		Log.Infof(c, "%s", body)
		return nil, err
	}
	return result, nil
}

func CallWithBasicAuthSingleParam(c appengine.Context, address string, username string, password string, allowInvalidServerCertificate bool, method string, id interface{}, params interface{}) (map[string]interface{}, error) {
	Log.Debugf(c, "GAEBitcoinJSON CallWithBasicAuthSingleParam - %v", method)
	data, err := json.Marshal(map[string]interface{}{
		"method": method,
		"id":     id,
		"params": params,
	})
	if err != nil {
		Log.Infof(c, "Marshal: %v", err)
		return nil, err
	}

	req, err := http.NewRequest("POST", address, strings.NewReader(string(data)))
	if err != nil {
		Log.Infof(c, "Request: %v", err)
		return nil, err
	}
	req.SetBasicAuth(username, password)

	tr := &urlfetch.Transport{Context: c, Deadline: TimeoutDuration, AllowInvalidServerCertificate: allowInvalidServerCertificate}

	resp, err := tr.RoundTrip(req)
	if err != nil {
		Log.Infof(c, "Post: %v", err)
		return nil, err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		Log.Infof(c, "ReadAll: %v", err)
		return nil, err
	}
	result := make(map[string]interface{})
	err = json.Unmarshal(body, &result)
	if err != nil {
		Log.Infof(c, "Unmarshal: %v", err)
		Log.Infof(c, "%s", body)
		return nil, err
	}
	return result, nil
}

/*
//https://en.bitcoin.it/wiki/Api#Full_list\
*/

func BackupWallet(c appengine.Context, id interface{}, destination []interface{}) (map[string]interface{}, error) {
	//Safely copies wallet.dat to destination, which can be a directory or a path with filename.
	resp, err := CallWithBasicAuth(c, Address, Username, Password, true, "backupwallet", id, destination)
	if err != nil {
		c.Infof(err.Error())
		return resp, err
	}
	//result:=resp["result"]
	//c.Infof(result)

	return resp, err
}

func CreateRawTransaction(c appengine.Context, id interface{}, data []interface{}) (map[string]interface{}, error) {
	//version 0.7 Creates a raw transaction spending given inputs.
	resp, err := CallWithBasicAuth(c, Address, Username, Password, true, "createrawtransaction", id, data)
	if err != nil {
		c.Infof(err.Error())
		return resp, err
	}
	//result:=resp["result"]
	//c.Infof(result)

	return resp, err
}

func EncryptWallet(c appengine.Context, id interface{}, passphrase []interface{}) (map[string]interface{}, error) {
	//Encrypts the wallet with <passphrase>.
	resp, err := CallWithBasicAuth(c, Address, Username, Password, true, "encryptwallet", id, passphrase)
	if err != nil {
		c.Infof(err.Error())
		return resp, err
	}
	//result:=resp["result"]
	//c.Infof(result)

	return resp, err
}

func GetAccount(c appengine.Context, id interface{}, bitcoinaddress []interface{}) (map[string]interface{}, error) {
	//Returns the account associated with the given address.
	resp, err := CallWithBasicAuth(c, Address, Username, Password, true, "getaccount", id, bitcoinaddress)
	if err != nil {
		c.Infof(err.Error())
		return resp, err
	}
	//result:=resp["result"]
	//c.Infof(result)

	return resp, err
}

func GetAccountAddress(c appengine.Context, id interface{}, account []interface{}) (map[string]interface{}, error) {
	//Returns the current bitcoin address for receiving payments to this account.
	resp, err := CallWithBasicAuth(c, Address, Username, Password, true, "getaccountaddress", id, account)
	if err != nil {
		c.Infof(err.Error())
		return resp, err
	}
	//result:=resp["result"]
	//c.Infof(result)

	return resp, err
}
func GetAddressesByAccount(c appengine.Context, id interface{}, account []interface{}) (map[string]interface{}, error) {
	//Returns the list of addresses for the given account.
	resp, err := CallWithBasicAuth(c, Address, Username, Password, true, "getaddressesbyaccount", id, account)
	if err != nil {
		c.Infof(err.Error())
		return resp, err
	}
	//result:=resp["result"]
	//c.Infof(result)

	return resp, err
}

func GetBalance(c appengine.Context, id interface{}, data []interface{}) (map[string]interface{}, error) {
	//If [account] is not specified, returns the server's total available balance.
	//If [account] is specified, returns the balance in the account.
	resp, err := CallWithBasicAuth(c, Address, Username, Password, true, "getbalance", id, data)
	if err != nil {
		c.Infof(err.Error())
		return resp, err
	}
	//result:=resp["result"]
	//c.Infof(result)

	return resp, err
}

/*func GetBlockByCount(c appengine.Context, id interface{}, height []interface{})(map[string]interface{}, error){
	//Dumps the block existing at specified height. Note: this is not available in the official release
	resp, err:=CallWithBasicAuth(c, Address, Username, Password, true, "getblockbycount", id, height)
	if err!=nil{
		c.Infof(err.Error())
		return resp, err
	}
	//result:=resp["result"]
	//c.Infof(result)

	return resp, err
}*/

func GetBlockCount(c appengine.Context, id interface{}) (map[string]interface{}, error) {
	//Returns the number of blocks in the longest block chain.
	resp, err := CallWithBasicAuth(c, Address, Username, Password, true, "getblockcount", id, nil)
	if err != nil {
		c.Infof(err.Error())
		return resp, err
	}
	//result:=resp["result"]
	//c.Infof(result)

	return resp, err
}

func GetBlockHash(c appengine.Context, id interface{}, index []interface{}) (map[string]interface{}, error) {
	//Returns hash of block in best-block-chain at <index>; index 0 is the genesis block
	resp, err := CallWithBasicAuth(c, Address, Username, Password, true, "getblockhash", id, index)
	if err != nil {
		c.Infof(err.Error())
		return resp, err
	}
	return resp, err
}

func GetBlockNumber(c appengine.Context, id interface{}) (map[string]interface{}, error) {
	//Returns the block number of the latest block in the longest block chain.
	resp, err := CallWithBasicAuth(c, Address, Username, Password, true, "getblocknumber", id, nil)
	if err != nil {
		c.Infof(err.Error())
		return resp, err
	}
	//result:=resp["result"]
	//c.Infof(result)

	return resp, err
}
func GetConnectionCount(c appengine.Context, id interface{}) (map[string]interface{}, error) {
	//Returns the number of connections to other nodes.
	resp, err := CallWithBasicAuth(c, Address, Username, Password, true, "getconnectioncount", id, nil)
	if err != nil {
		c.Infof(err.Error())
		return resp, err
	}
	//result:=resp["result"]
	//c.Infof(result)

	return resp, err
}
func GetDifficulty(c appengine.Context, id interface{}) (map[string]interface{}, error) {
	//Returns the proof-of-work difficulty as a multiple of the minimum difficulty.
	resp, err := CallWithBasicAuth(c, Address, Username, Password, true, "getdifficulty", id, nil)
	if err != nil {
		c.Infof(err.Error())
		return resp, err
	}
	//result:=resp["result"]
	//c.Infof(result)

	return resp, err
}
func GetGenerate(c appengine.Context, id interface{}) (map[string]interface{}, error) {
	//Returns true or false whether bitcoind is currently generating hashes
	resp, err := CallWithBasicAuth(c, Address, Username, Password, true, "getgenerate", id, nil)
	if err != nil {
		c.Infof(err.Error())
		return resp, err
	}
	//result:=resp["result"]
	//c.Infof(result)

	return resp, err
}
func GetHashesPerSec(c appengine.Context, id interface{}) (map[string]interface{}, error) {
	//Returns a recent hashes per second performance measurement while generating.
	resp, err := CallWithBasicAuth(c, Address, Username, Password, true, "gethashespersec", id, nil)
	if err != nil {
		c.Infof(err.Error())
		return resp, err
	}
	//result:=resp["result"]
	//c.Infof(result)

	return resp, err
}

func GetMemoryPool(c appengine.Context, id interface{}, data []interface{}) (map[string]interface{}, error) {
	//If [data] is not specified, returns data needed to construct a block to work on:
	//"version" : block version
	//"previousblockhash" : hash of current highest block
	//"transactions" : contents of non-coinbase transactions that should be included in the next block
	//"coinbasevalue" : maximum allowable input to coinbase transaction, including the generation award and transaction fees
	//"time" : timestamp appropriate for next block
	//"bits" : compressed target of next block
	//If [data] is specified, tries to solve the block and returns true if it was successful.
	resp, err := CallWithBasicAuth(c, Address, Username, Password, true, "getmemorypool", id, data)
	if err != nil {
		c.Infof(err.Error())
		return resp, err
	}
	//result:=resp["result"]
	//c.Infof(result)

	return resp, err
}

func GetInfo(c appengine.Context, id interface{}) (map[string]interface{}, error) {
	//Returns an object containing various state info.
	resp, err := CallWithBasicAuth(c, Address, Username, Password, true, "getinfo", id, nil)
	if err != nil {
		c.Infof(err.Error())
		return resp, err
	}
	//result:=resp["result"]
	//c.Infof(result)

	return resp, err
}

func GetNewAddress(c appengine.Context, id interface{}, account []interface{}) (map[string]interface{}, error) {
	//Returns a new bitcoin address for receiving payments. If [account] is specified (recommended), it is added to the address book so payments received with the address will be credited to [account].
	resp, err := CallWithBasicAuth(c, Address, Username, Password, true, "getnewaddress", id, account)
	if err != nil {
		c.Infof(err.Error())
		return resp, err
	}
	//result:=resp["result"]
	//c.Infof(result)

	return resp, err
}

func GetReceivedByAccount(c appengine.Context, id interface{}, data []interface{}) (map[string]interface{}, error) {
	//Returns the total amount received by addresses with [account] in transactions with at least [minconf] confirmations. If [account] not provided return will include all transactions to all accounts. (version 0.3.24-beta)
	resp, err := CallWithBasicAuth(c, Address, Username, Password, true, "getreceivedbyaccount", id, data)
	if err != nil {
		c.Infof(err.Error())
		return resp, err
	}
	//result:=resp["result"]
	//c.Infof(result)

	return resp, err
}

func GetReceivedByAddress(c appengine.Context, id interface{}, data []interface{}) (map[string]interface{}, error) {
	//Returns the total amount received by <bitcoinaddress> in transactions with at least [minconf] confirmations. While some might consider this obvious, value reported by this only considers *receiving* transactions. It does not check payments that have been made *from* this address. In other words, this is not "getaddressbalance".
	resp, err := CallWithBasicAuth(c, Address, Username, Password, true, "getreceivedbyaddress", id, data)
	if err != nil {
		c.Infof(err.Error())
		return resp, err
	}
	//result:=resp["result"]
	//c.Infof(result)

	return resp, err
}

func GetBlock(c appengine.Context, id interface{}, txid []interface{}) (map[string]interface{}, error) {
	//Returns information about the block with the given hash.
	resp, err := CallWithBasicAuth(c, Address, Username, Password, true, "getblock", id, txid)
	if err != nil {
		c.Infof(err.Error())
		return resp, err
	}
	return resp, err
}

func GetTransaction(c appengine.Context, id interface{}, txid []interface{}) (map[string]interface{}, error) {
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
	resp, err := CallWithBasicAuth(c, Address, Username, Password, true, "gettransaction", id, txid)
	if err != nil {
		c.Infof(err.Error())
		return resp, err
	}
	//result:=resp["result"]
	//c.Infof(result)

	return resp, err
}

func GetRawTransaction(c appengine.Context, id interface{}, txid []interface{}) (map[string]interface{}, error) {
	resp, err := CallWithBasicAuth(c, Address, Username, Password, true, "getrawtransaction", id, txid)
	if err != nil {
		c.Infof(err.Error())
		return resp, err
	}
	return resp, err
}

func GetWork(c appengine.Context, id interface{}, data []interface{}) (map[string]interface{}, error) {
	//If [data] is not specified, returns formatted hash data to work on:
	//"midstate" : precomputed hash state after hashing the first half of the data
	//"data" : block data
	//"hash1" : formatted hash buffer for second hash
	//"target" : little endian hash target
	//If [data] is specified, tries to solve the block and returns true if it was successful.
	resp, err := CallWithBasicAuth(c, Address, Username, Password, true, "getwork", id, data)
	if err != nil {
		c.Infof(err.Error())
		return resp, err
	}
	////result:=resp["result"]
	////c.Infof(result)

	return resp, err
}

func Help(c appengine.Context, id interface{}, command string) (map[string]interface{}, error) {
	//List commands, or get help for a command.
	resp, err := CallWithBasicAuth(c, Address, Username, Password, true, "help", id, []interface{}{command})
	if err != nil {
		c.Infof(err.Error())
		return resp, err
	}
	//result:=resp["result"]
	//c.Infof(result)

	return resp, err
}

func KeyPoolRefill(c appengine.Context, id interface{}) (map[string]interface{}, error) {
	//Fills the keypool, requires wallet passphrase to be set.
	resp, err := CallWithBasicAuth(c, Address, Username, Password, true, "keypoolrefill", id, nil)
	if err != nil {
		c.Infof(err.Error())
		return resp, err
	}
	//result:=resp["result"]
	//c.Infof(result)

	return resp, err
}

func ListAccounts(c appengine.Context, id interface{}, minconf interface{}) (map[string]interface{}, error) {
	//Returns Object that has account names as keys, account balances as values.
	resp, err := CallWithBasicAuthSingleParam(c, Address, Username, Password, true, "listaccounts", id, minconf)
	if err != nil {
		c.Infof(err.Error())
		return resp, err
	}
	//result:=resp["result"]
	//c.Infof(result)
	Log.Infof(c, "%v", resp)

	return resp, err
}

func ListSinceBlock(c appengine.Context, id interface{}, blockid, targetconfirmations interface{}) (map[string]interface{}, error) {
	//Get all transactions in blocks since block [blockid], or all transactions if omitted.
	resp, err := CallWithBasicAuth(c, Address, Username, Password, true, "listsinceblock", id, []interface{}{blockid, targetconfirmations})
	if err != nil {
		c.Infof(err.Error())
		return resp, err
	}
	//result:=resp["result"]
	//c.Infof(result)

	return resp, err
}

func ListReceivedByAccount(c appengine.Context, id interface{}, data []interface{}) (map[string]interface{}, error) {
	//Returns an array of objects containing:
	//"account" : the account of the receiving addresses
	//"amount" : total amount received by addresses with this account
	//"confirmations" : number of confirmations of the most recent transaction included
	resp, err := CallWithBasicAuth(c, Address, Username, Password, true, "listreceivedbyaccount", id, data)
	if err != nil {
		c.Infof(err.Error())
		return resp, err
	}
	//result:=resp["result"]
	//c.Infof(result)

	return resp, err
}

func ListReceivedByAddress(c appengine.Context, id interface{}, data []interface{}) (map[string]interface{}, error) {
	//Returns an array of objects containing:
	//"address" : receiving address
	//"account" : the account of the receiving address
	//"amount" : total amount received by the address
	//"confirmations" : number of confirmations of the most recent transaction included
	//To get a list of accounts on the system, execute bitcoind listreceivedbyaddress 0 true
	resp, err := CallWithBasicAuth(c, Address, Username, Password, true, "listreceivedbyaddress", id, data)
	if err != nil {
		c.Infof(err.Error())
		return resp, err
	}
	//result:=resp["result"]
	//c.Infof(result)

	return resp, err
}

func ListTransactions(c appengine.Context, id interface{}, data []interface{}) (map[string]interface{}, error) {
	//Returns up to [count] most recent transactions skipping the first [from] transactions for account [account]. If [account] not provided will return recent transaction from all accounts.
	resp, err := CallWithBasicAuth(c, Address, Username, Password, true, "listtransactions", id, data)
	if err != nil {
		c.Infof(err.Error())
		return resp, err
	}
	//result:=resp["result"]
	//c.Infof(result)

	return resp, err
}

func ListUnspent(c appengine.Context, id interface{}, data []interface{}) (map[string]interface{}, error) {
	//Returns up to [count] most recent transactions skipping the first [from] transactions for account [account]. If [account] not provided will return recent transaction from all accounts.
	resp, err := CallWithBasicAuth(c, Address, Username, Password, true, "listunspent", id, data)
	if err != nil {
		c.Infof(err.Error())
		return resp, err
	}
	//result:=resp["result"]
	//c.Infof(result)

	return resp, err
}

func Move(c appengine.Context, id interface{}, data []interface{}) (map[string]interface{}, error) {
	//Move from one account in your wallet to another
	resp, err := CallWithBasicAuth(c, Address, Username, Password, true, "move", id, data)
	if err != nil {
		c.Infof(err.Error())
		return resp, err
	}
	//result:=resp["result"]
	//c.Infof(result)

	return resp, err
}

func SendFrom(c appengine.Context, id interface{}, data []interface{}) (map[string]interface{}, error) {
	//<amount> is a real and is rounded to 8 decimal places. Will send the given amount to the given address, ensuring the account has a valid balance using [minconf] confirmations. Returns the transaction ID if successful (not in JSON object).
	resp, err := CallWithBasicAuth(c, Address, Username, Password, true, "sendfrom", id, data)
	if err != nil {
		c.Infof(err.Error())
		return resp, err
	}
	//result:=resp["result"]
	//c.Infof(result)

	return resp, err
}

func SendMany(c appengine.Context, id interface{}, data []interface{}) (map[string]interface{}, error) {
	//amounts are double-precision floating point numbers
	resp, err := CallWithBasicAuth(c, Address, Username, Password, true, "sendmany", id, data)
	if err != nil {
		c.Infof(err.Error())
		return resp, err
	}
	//result:=resp["result"]
	//c.Infof(result)

	return resp, err
}

func SendToAddress(c appengine.Context, id interface{}, data []interface{}) (map[string]interface{}, error) {
	//<amount> is a real and is rounded to 8 decimal places. Returns the transaction ID <txid> if successful.
	resp, err := CallWithBasicAuth(c, Address, Username, Password, true, "sendtoaddress", id, data)
	if err != nil {
		c.Infof(err.Error())
		return resp, err
	}
	//result:=resp["result"]
	//c.Infof(result)

	return resp, err
}

func SetAccount(c appengine.Context, id interface{}, data []interface{}) (map[string]interface{}, error) {
	//Sets the account associated with the given address. Assigning address that is already assigned to the same account will create a new address associated with that account.
	resp, err := CallWithBasicAuth(c, Address, Username, Password, true, "setaccount", id, data)
	if err != nil {
		c.Infof(err.Error())
		return resp, err
	}
	//result:=resp["result"]
	//c.Infof(result)

	return resp, err
}

func SetGenerate(c appengine.Context, id interface{}, data []interface{}) (map[string]interface{}, error) {
	//<generate> is true or false to turn generation on or off.
	//Generation is limited to [genproclimit] processors, -1 is unlimited.
	resp, err := CallWithBasicAuth(c, Address, Username, Password, true, "setgenerate", id, data)
	if err != nil {
		c.Infof(err.Error())
		return resp, err
	}
	//result:=resp["result"]
	//c.Infof(result)

	return resp, err
}

func SetTxFee(c appengine.Context, id interface{}, amount []interface{}) (map[string]interface{}, error) {
	//<amount> is a real and is rounded to the nearest 0.00000001
	resp, err := CallWithBasicAuth(c, Address, Username, Password, true, "settxfee", id, amount)
	if err != nil {
		c.Infof(err.Error())
		return resp, err
	}
	//result:=resp["result"]
	//c.Infof(result)

	return resp, err
}

func SignMessage(c appengine.Context, id interface{}, bitcoinaddress, message interface{}) (map[string]interface{}, error) {
	//Sign a message with the private key of an address.
	resp, err := CallWithBasicAuth(c, Address, Username, Password, true, "signmessage", id, []interface{}{bitcoinaddress, message})
	if err != nil {
		c.Infof(err.Error())
		return resp, err
	}
	//result:=resp["result"]
	//c.Infof(result)

	return resp, err
}
func Stop(c appengine.Context, id interface{}) (map[string]interface{}, error) {
	//Stop bitcoin server.
	resp, err := CallWithBasicAuth(c, Address, Username, Password, true, "stop", id, nil)
	if err != nil {
		c.Infof(err.Error())
		return resp, err
	}
	//result:=resp["result"]
	//c.Infof(result)

	return resp, err
}

func ValidateAddress(c appengine.Context, id interface{}, bitcoinaddress interface{}) (map[string]interface{}, error) {
	//Return information about <bitcoinaddress>.
	resp, err := CallWithBasicAuth(c, Address, Username, Password, true, "validateaddress", id, []interface{}{bitcoinaddress})
	if err != nil {
		c.Infof(err.Error())
		return resp, err
	}
	//result:=resp["result"]
	//c.Infof(result)

	return resp, err
}

func VerifyMessage(c appengine.Context, id interface{}, bitcoinaddress, signature, message interface{}) (map[string]interface{}, error) {
	//Verify a signed message.
	resp, err := CallWithBasicAuth(c, Address, Username, Password, true, "verifymessage", id, []interface{}{bitcoinaddress, signature, message})
	if err != nil {
		c.Infof(err.Error())
		return resp, err
	}
	//result:=resp["result"]
	//c.Infof(result)

	return resp, err
}

func WalletLock(c appengine.Context, id interface{}) (map[string]interface{}, error) {
	//Removes the wallet encryption key from memory, locking the wallet. After calling this method, you will need to call walletpassphrase again before being able to call any methods which require the wallet to be unlocked.
	resp, err := CallWithBasicAuth(c, Address, Username, Password, true, "walletlock", id, nil)
	if err != nil {
		c.Infof(err.Error())
		return resp, err
	}
	//result:=resp["result"]
	//c.Infof(result)

	return resp, err
}

func WalletPassPhrase(c appengine.Context, id interface{}, passphrase, timeout interface{}) (map[string]interface{}, error) {
	//Stores the wallet decryption key in memory for <timeout> seconds.
	resp, err := CallWithBasicAuth(c, Address, Username, Password, true, "walletpassphrase", id, []interface{}{passphrase, timeout})
	if err != nil {
		c.Infof(err.Error())
		return resp, err
	}
	//result:=resp["result"]
	//c.Infof(result)

	return resp, err
}

func WalletPassPhraseChange(c appengine.Context, id interface{}, data []interface{}) (map[string]interface{}, error) {
	//Changes the wallet passphrase from <oldpassphrase> to <newpassphrase>.
	resp, err := CallWithBasicAuth(c, Address, Username, Password, true, "walletpassphrasechange", id, data)
	if err != nil {
		c.Infof(err.Error())
		return resp, err
	}
	//result:=resp["result"]
	//c.Infof(result)

	return resp, err
}
