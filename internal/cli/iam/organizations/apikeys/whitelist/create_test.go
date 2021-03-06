// Copyright 2020 MongoDB Inc
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// +build unit

package whitelist

import (
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/mongodb/mongocli/internal/mocks"
	"go.mongodb.org/atlas/mongodbatlas"
)

func TestCreate_Run(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockStore := mocks.NewMockOrganizationAPIKeyWhitelistCreator(ctrl)
	defer ctrl.Finish()

	expected := &mongodbatlas.WhitelistAPIKeys{}

	createOpts := &CreateOpts{
		store:  mockStore,
		apyKey: "1",
		ips:    []string{"77.54.32.11"},
	}

	whitelistReq, err := createOpts.newWhitelistAPIKeysReq()
	if err != nil {
		t.Fatalf("Run() unexpected error: %v", err)
	}

	mockStore.
		EXPECT().
		CreateOrganizationAPIKeyWhite(createOpts.OrgID, createOpts.apyKey, whitelistReq).
		Return(expected, nil).
		Times(1)

	err = createOpts.Run()
	if err != nil {
		t.Fatalf("Run() unexpected error: %v", err)
	}
}
