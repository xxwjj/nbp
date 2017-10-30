// Copyright 2017 The OpenSDS Authors.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package api

import (
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"

	"github.com/astaxie/beego"
	"github.com/opensds/opensds/pkg/db"
	dbtest "github.com/opensds/opensds/pkg/db/testing"
	"github.com/opensds/opensds/pkg/model"
)

func init() {
	var dockPortal DockPortal
	beego.Router("/v1alpha/docks", &dockPortal, "get:ListDocks")
	beego.Router("/v1alpha/docks/:dockId", &dockPortal, "get:GetDock")
}

func TestListDocks(t *testing.T) {

	var fakeDocks = []*model.DockSpec{
		&model.DockSpec{
			BaseModel: &model.BaseModel{
				Id:        "b7602e18-771e-11e7-8f38-dbd6d291f4e0",
				CreatedAt: "2017-10-11T11:28:58",
			},
			Name:        "cinder",
			Description: "cinder backend service",
			StorageType: "block",
			Endpoint:    "localhost:50050",
			Status:      "available",
			DriverName:  "cinder",
		},
	}

	mockClient := new(dbtest.MockClient)
	mockClient.On("ListDocks").Return(fakeDocks, nil)
	db.C = mockClient

	r, _ := http.NewRequest("GET", "/v1alpha/docks", nil)
	w := httptest.NewRecorder()
	beego.BeeApp.Handlers.ServeHTTP(w, r)

	var output []model.DockSpec
	json.Unmarshal(w.Body.Bytes(), &output)

	expectedJson := `[
		{
			"id": "b7602e18-771e-11e7-8f38-dbd6d291f4e0",
			"name": "cinder",
			"description": "cinder backend service",
			"storageType": "block",
			"status": "available",
			"driverName": "cinder",
			"endpoint": "localhost:50050",
			"createdAt": "2017-10-11T11:28:58",
			"updatedAt": ""		
		}		
	]`

	var expected []model.DockSpec
	json.Unmarshal([]byte(expectedJson), &expected)

	if w.Code != 200 {
		t.Errorf("Expected 200, actual %v", w.Code)
	}

	if !reflect.DeepEqual(expected, output) {
		t.Errorf("Expected %v, actual %v", expected, output)
	}
}

func TestListDocksWithInternalError(t *testing.T) {

	mockClient := new(dbtest.MockClient)
	mockClient.On("ListDocks").Return(nil, errors.New("db error"))
	db.C = mockClient

	r, _ := http.NewRequest("GET", "/v1alpha/docks", nil)
	w := httptest.NewRecorder()
	beego.BeeApp.Handlers.ServeHTTP(w, r)

	if w.Code != 500 {
		t.Errorf("Expected 500, actual %v", w.Code)
	}
}

func TestGetDock(t *testing.T) {

	var fakeDock = &model.DockSpec{
		BaseModel: &model.BaseModel{
			Id:        "b7602e18-771e-11e7-8f38-dbd6d291f4e0",
			CreatedAt: "2017-10-11T11:28:58",
		},
		Name:        "cinder",
		Description: "cinder backend service",
		StorageType: "block",
		Endpoint:    "localhost:50050",
		Status:      "available",
		DriverName:  "cinder",
	}

	mockClient := new(dbtest.MockClient)
	mockClient.On("GetDock", "b7602e18-771e-11e7-8f38-dbd6d291f4e0").Return(fakeDock, nil)
	db.C = mockClient

	r, _ := http.NewRequest("GET",
		"/v1alpha/docks/b7602e18-771e-11e7-8f38-dbd6d291f4e0", nil)
	w := httptest.NewRecorder()
	beego.BeeApp.Handlers.ServeHTTP(w, r)

	var output model.DockSpec
	json.Unmarshal(w.Body.Bytes(), &output)

	expectedJson := `
		{
			"id": "b7602e18-771e-11e7-8f38-dbd6d291f4e0",
			"name": "cinder",
			"description": "cinder backend service",
			"storageType": "block",
			"status": "available",
			"driverName": "cinder",
			"endpoint": "localhost:50050",
			"createdAt": "2017-10-11T11:28:58",
			"updatedAt": ""		
		}`

	var expected model.DockSpec
	json.Unmarshal([]byte(expectedJson), &expected)

	if w.Code != 200 {
		t.Errorf("Expected 200, actual %v", w.Code)
	}

	if !reflect.DeepEqual(expected, output) {
		t.Errorf("Expected %v, actual %v", expected, output)
	}
}

func TestGetDockWithBadRequestError(t *testing.T) {

	mockClient := new(dbtest.MockClient)
	mockClient.On("GetDock", "b7602e18-771e-11e7-8f38-dbd6d291f4e0").Return(
		nil, errors.New("db error"))
	db.C = mockClient

	r, _ := http.NewRequest("GET",
		"/v1alpha/docks/b7602e18-771e-11e7-8f38-dbd6d291f4e0", nil)
	w := httptest.NewRecorder()
	beego.BeeApp.Handlers.ServeHTTP(w, r)

	if w.Code != 400 {
		t.Errorf("Expected 400, actual %v", w.Code)
	}
}
