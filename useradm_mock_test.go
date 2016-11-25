// Copyright 2016 Mender Software AS
//
//    Licensed under the Apache License, Version 2.0 (the "License");
//    you may not use this file except in compliance with the License.
//    You may obtain a copy of the License at
//
//        http://www.apache.org/licenses/LICENSE-2.0
//
//    Unless required by applicable law or agreed to in writing, software
//    distributed under the License is distributed on an "AS IS" BASIS,
//    WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//    See the License for the specific language governing permissions and
//    limitations under the License.
package main

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/stretchr/testify/mock"
)

// MockUserAdmApp is an autogenerated mock type for the UserAdmApp type
type MockUserAdmApp struct {
	mock.Mock
}

// Login provides a mock function with given fields: email, pass
func (_m *MockUserAdmApp) Login(email string, pass string) (*jwt.Token, error) {
	ret := _m.Called(email, pass)

	var r0 *jwt.Token
	if rf, ok := ret.Get(0).(func(string, string) *jwt.Token); ok {
		r0 = rf(email, pass)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*jwt.Token)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string, string) error); ok {
		r1 = rf(email, pass)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}