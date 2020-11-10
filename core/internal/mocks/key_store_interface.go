// Code generated by mockery v2.4.0-beta. DO NOT EDIT.

package mocks

import (
	big "math/big"

	accounts "github.com/ethereum/go-ethereum/accounts"

	common "github.com/ethereum/go-ethereum/common"

	mock "github.com/stretchr/testify/mock"

	models "github.com/smartcontractkit/chainlink/core/store/models"

	types "github.com/ethereum/go-ethereum/core/types"
)

// KeyStoreInterface is an autogenerated mock type for the KeyStoreInterface type
type KeyStoreInterface struct {
	mock.Mock
}

// Accounts provides a mock function with given fields:
func (_m *KeyStoreInterface) Accounts() []accounts.Account {
	ret := _m.Called()

	var r0 []accounts.Account
	if rf, ok := ret.Get(0).(func() []accounts.Account); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]accounts.Account)
		}
	}

	return r0
}

// Delete provides a mock function with given fields: address
func (_m *KeyStoreInterface) Delete(address common.Address) error {
	ret := _m.Called(address)

	var r0 error
	if rf, ok := ret.Get(0).(func(common.Address) error); ok {
		r0 = rf(address)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Export provides a mock function with given fields: address, newPassword
func (_m *KeyStoreInterface) Export(address common.Address, newPassword string) ([]byte, error) {
	ret := _m.Called(address, newPassword)

	var r0 []byte
	if rf, ok := ret.Get(0).(func(common.Address, string) []byte); ok {
		r0 = rf(address, newPassword)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]byte)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(common.Address, string) error); ok {
		r1 = rf(address, newPassword)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetAccountByAddress provides a mock function with given fields: _a0
func (_m *KeyStoreInterface) GetAccountByAddress(_a0 common.Address) (accounts.Account, error) {
	ret := _m.Called(_a0)

	var r0 accounts.Account
	if rf, ok := ret.Get(0).(func(common.Address) accounts.Account); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Get(0).(accounts.Account)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(common.Address) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetAccounts provides a mock function with given fields:
func (_m *KeyStoreInterface) GetAccounts() []accounts.Account {
	ret := _m.Called()

	var r0 []accounts.Account
	if rf, ok := ret.Get(0).(func() []accounts.Account); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]accounts.Account)
		}
	}

	return r0
}

// GetFirstAccount provides a mock function with given fields:
func (_m *KeyStoreInterface) GetFirstAccount() (accounts.Account, error) {
	ret := _m.Called()

	var r0 accounts.Account
	if rf, ok := ret.Get(0).(func() accounts.Account); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(accounts.Account)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// HasAccounts provides a mock function with given fields:
func (_m *KeyStoreInterface) HasAccounts() bool {
	ret := _m.Called()

	var r0 bool
	if rf, ok := ret.Get(0).(func() bool); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(bool)
	}

	return r0
}

// Import provides a mock function with given fields: keyJSON, oldPassword
func (_m *KeyStoreInterface) Import(keyJSON []byte, oldPassword string) (accounts.Account, error) {
	ret := _m.Called(keyJSON, oldPassword)

	var r0 accounts.Account
	if rf, ok := ret.Get(0).(func([]byte, string) accounts.Account); ok {
		r0 = rf(keyJSON, oldPassword)
	} else {
		r0 = ret.Get(0).(accounts.Account)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func([]byte, string) error); ok {
		r1 = rf(keyJSON, oldPassword)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// NewAccount provides a mock function with given fields:
func (_m *KeyStoreInterface) NewAccount() (accounts.Account, error) {
	ret := _m.Called()

	var r0 accounts.Account
	if rf, ok := ret.Get(0).(func() accounts.Account); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(accounts.Account)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// SignHash provides a mock function with given fields: hash
func (_m *KeyStoreInterface) SignHash(hash common.Hash) (models.Signature, error) {
	ret := _m.Called(hash)

	var r0 models.Signature
	if rf, ok := ret.Get(0).(func(common.Hash) models.Signature); ok {
		r0 = rf(hash)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(models.Signature)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(common.Hash) error); ok {
		r1 = rf(hash)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// SignTx provides a mock function with given fields: account, tx, chainID
func (_m *KeyStoreInterface) SignTx(account accounts.Account, tx *types.Transaction, chainID *big.Int) (*types.Transaction, error) {
	ret := _m.Called(account, tx, chainID)

	var r0 *types.Transaction
	if rf, ok := ret.Get(0).(func(accounts.Account, *types.Transaction, *big.Int) *types.Transaction); ok {
		r0 = rf(account, tx, chainID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*types.Transaction)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(accounts.Account, *types.Transaction, *big.Int) error); ok {
		r1 = rf(account, tx, chainID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Unlock provides a mock function with given fields: password
func (_m *KeyStoreInterface) Unlock(password string) error {
	ret := _m.Called(password)

	var r0 error
	if rf, ok := ret.Get(0).(func(string) error); ok {
		r0 = rf(password)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Wallets provides a mock function with given fields:
func (_m *KeyStoreInterface) Wallets() []accounts.Wallet {
	ret := _m.Called()

	var r0 []accounts.Wallet
	if rf, ok := ret.Get(0).(func() []accounts.Wallet); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]accounts.Wallet)
		}
	}

	return r0
}
