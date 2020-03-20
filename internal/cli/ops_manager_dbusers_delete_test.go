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

package cli

import (
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/mongodb/mongocli/internal/fixtures"
	"github.com/mongodb/mongocli/internal/mocks"
)

func TestOpsManagerDBUserDelete_Run(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockStore := mocks.NewMockAutomationPatcher(ctrl)

	defer ctrl.Finish()

	expected := fixtures.AutomationConfig()

	createOpts := &opsManagerDBUsersDeleteOpts{
		globalOpts: newGlobalOpts(),
		deleteOpts: &deleteOpts{
			confirm:        true,
			entry:          "test",
			successMessage: "DB user '%s' deleted\n",
		},
		authDB: "admin",
		store:  mockStore,
	}

	mockStore.
		EXPECT().
		GetAutomationConfig(createOpts.projectID).
		Return(expected, nil).
		Times(1)

	mockStore.
		EXPECT().
		UpdateAutomationConfig(createOpts.projectID, expected).
		Return(nil).
		Times(1)

	err := createOpts.Run()
	if err != nil {
		t.Fatalf("Run() unexpected error: %v", err)
	}
}