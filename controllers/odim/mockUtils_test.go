//(C) Copyright [2023] Hewlett Packard Enterprise Development LP
//
//Licensed under the Apache License, Version 2.0 (the "License"); you may
//not use this file except in compliance with the License. You may obtain
//a copy of the License at
//
//    http://www.apache.org/licenses/LICENSE-2.0
//
//Unless required by applicable law or agreed to in writing, software
//distributed under the License is distributed on an "AS IS" BASIS, WITHOUT
//WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied. See the
//License for the specific language governing permissions and limitations
// under the License.

package controllers

import (
	"context"
	"errors"
	"net/http"

	infraiov1 "github.com/ODIM-Project/BMCOperator/api/v1"
	"k8s.io/apimachinery/pkg/api/meta"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/types"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

type mockCommonRec struct {
	variable string
}
type mockRestClient struct{ url string }

type mockCommonUtil struct{}

// Mock Common Rec Methods
// get objects
func (m *mockCommonRec) GetBmcObject(ctx context.Context, field, value, ns string) *infraiov1.Bmc {
	return nil
}

func (m *mockCommonRec) GetAllBmcObject(ctx context.Context, ns string) *[]infraiov1.Bmc {
	return nil
}

func (m *mockCommonRec) GetBiosSchemaObject(ctx context.Context, field, value, ns string) *infraiov1.BiosSchemaRegistry {
	return nil
}
func (m *mockCommonRec) GetBiosObject(ctx context.Context, field, value, ns string) *infraiov1.BiosSetting {
	return nil
}

func (m *mockCommonRec) GetBootObject(ctx context.Context, field, value, ns string) *infraiov1.BootOrderSetting {
	return nil
}

func (m *mockCommonRec) GetOdimObject(ctx context.Context, field, value, ns string) *infraiov1.Odim {
	return nil
}

func (m *mockCommonRec) GetVolumeObject(ctx context.Context, bmcIP, ns string) *infraiov1.Volume {
	return nil
}

func (m *mockCommonRec) GetFirmwareObject(ctx context.Context, field, value, ns string) *infraiov1.Firmware {
	return nil
}

// create objects
func (m *mockCommonRec) CreateBiosSettingObject(ctx context.Context, biosAttributes map[string]string, bmcObj *infraiov1.Bmc) bool {
	return true
}

func (m *mockCommonRec) CreateBootOrderSettingObject(ctx context.Context, bootAttributes *infraiov1.BootSetting, bmcObj *infraiov1.Bmc) bool {
	return true
}

func (m *mockCommonRec) CheckAndCreateBiosSchemaObject(ctx context.Context, attributeResp map[string]interface{}, bmcObj *infraiov1.Bmc) bool {
	return true
}

func (m *mockCommonRec) CreateEventSubscriptionObject(ctx context.Context, subscriptionDetails map[string]interface{}, ns string, originResources []string) bool {
	return true
}

func (m *mockCommonRec) GetAllEventSubscriptionObjects(ctx context.Context, ns string) *[]infraiov1.Eventsubscription {
	return nil
}

func (m *mockCommonRec) GetEventsubscriptionObject(ctx context.Context, field, value, ns string) *infraiov1.Eventsubscription {
	return nil
}

func (m *mockCommonRec) GetEventMessageRegistryObject(ctx context.Context, registryPrefix, registryVersion, ns string) *infraiov1.EventsMessageRegistry {
	return nil
}

func (m *mockCommonRec) CheckAndCreateEventMessageObject(ctx context.Context, messageRegistryResp map[string]interface{}, bmcObj *infraiov1.Bmc) bool {
	return true
}

// update objects
func (m *mockCommonRec) UpdateBmcStatus(ctx context.Context, bmcObj *infraiov1.Bmc) {
}

func (m *mockCommonRec) UpdateOdimStatus(ctx context.Context, status string, odimObj *infraiov1.Odim) {
}

func (m *mockCommonRec) UpdateBmcObjectOnReset(ctx context.Context, bmcObject *infraiov1.Bmc, status string) {
}

func (m *mockCommonRec) UpdateVolumeStatus(ctx context.Context, volObject *infraiov1.Volume, volumeID, volumeName, capBytes, durableName, durableNameFormat string) {
}

// get updated objects
func (m *mockCommonRec) GetUpdatedBmcObject(ctx context.Context, ns types.NamespacedName, bmcObj *infraiov1.Bmc) {
}

func (m *mockCommonRec) GetUpdatedOdimObject(ctx context.Context, ns types.NamespacedName, odimObj *infraiov1.Odim) {
}

func (m *mockCommonRec) GetUpdatedVolumeObject(ctx context.Context, ns types.NamespacedName, volObj *infraiov1.Volume) {
}

func (m *mockCommonRec) GetUpdatedFirmwareObject(ctx context.Context, ns types.NamespacedName, firmObj *infraiov1.Firmware) {
}

func (m *mockCommonRec) UpdateBiosSettingObject(ctx context.Context, biosAttributes map[string]string, biosObj *infraiov1.BiosSetting) bool {
	return true
}

func (m *mockCommonRec) GetAllVolumeObjectIds(ctx context.Context, bmc *infraiov1.Bmc, ns string) map[string][]string {
	return nil
}

func (m *mockCommonRec) GetAllBiosSchemaRegistryObjects(ctx context.Context, ns string) *[]infraiov1.BiosSchemaRegistry {
	return nil
}

func (m *mockCommonRec) DeleteVolumeObject(ctx context.Context, volObj *infraiov1.Volume) {}

func (m *mockCommonRec) DeleteBmcObject(ctx context.Context, bmcObj *infraiov1.Bmc) {}

func (m *mockCommonRec) UpdateEventsubscriptionStatus(ctx context.Context, eventsubObj *infraiov1.Eventsubscription, eventsubscriptionDetails map[string]interface{}, originResouces []string) {
}

func (m *mockCommonRec) GetUpdatedEventsubscriptionObjects(ctx context.Context, ns types.NamespacedName, eventSubObj *infraiov1.Eventsubscription) {
}

func (m *mockCommonRec) GetAllVolumeObjects(ctx context.Context, bmcIP, ns string) []*infraiov1.Volume {
	return nil
}

func (m *mockCommonRec) GetVolumeObjectByVolumeID(ctx context.Context, volumeName, ns string) *infraiov1.Volume {
	return nil
}

func (m *mockCommonRec) GetRootCAFromSecret(ctx context.Context) []byte {
	return []byte{}
}

// common reconciler funcs
func (m *mockCommonRec) GetCommonReconcilerClient() client.Client {
	return mockCommonRec{}
}

func (m *mockCommonRec) GetCommonReconcilerScheme() *runtime.Scheme {
	return nil
}

// Client mocks
func (m mockCommonRec) Get(ctx context.Context, key client.ObjectKey, obj client.Object) error {
	return nil
}

func (m mockCommonRec) List(ctx context.Context, list client.ObjectList, opts ...client.ListOption) error {
	return nil
}

func (m mockCommonRec) Create(ctx context.Context, obj client.Object, opts ...client.CreateOption) error {
	return nil
}

func (m mockCommonRec) Delete(ctx context.Context, obj client.Object, opts ...client.DeleteOption) error {
	return nil
}

func (m mockCommonRec) Update(ctx context.Context, obj client.Object, opts ...client.UpdateOption) error {
	return nil
}

func (m mockCommonRec) Patch(ctx context.Context, obj client.Object, patch client.Patch, opts ...client.PatchOption) error {
	return nil
}

func (m mockCommonRec) DeleteAllOf(ctx context.Context, obj client.Object, opts ...client.DeleteAllOfOption) error {
	return nil
}

func (m mockCommonRec) Status() client.StatusWriter {
	return mockCommonRec{}
}

func (m mockCommonRec) RESTMapper() meta.RESTMapper {
	return nil
}

func (m mockCommonRec) Scheme() *runtime.Scheme {
	return nil
}

// mock Rest calls

func (m *mockRestClient) Post(url string, reason string, body interface{}) (*http.Response, error) {
	if url == "/redfish/v1/EventService/Subscriptions" {
		return &http.Response{StatusCode: 202}, nil
	}
	return nil, nil
}
func (m *mockRestClient) Get(url string, reason string) (map[string]interface{}, int, error) {
	//Test_odimUtils_checkOdimConnection mock1,2,3
	if m.url == "connected_case" {
		return map[string]interface{}{"@odata.id": "connected"}, 200, nil
	} else if m.url == "not_connected_case" {
		return map[string]interface{}{"@odata.id": nil}, 200, nil
	} else if m.url == "error_case" {
		return nil, 200, errors.New("fakeErr")
		//Test_odimUtils_updateConnectionMethodVariants mock1
	} else if url == "/redfish/v1/AggregationService/ConnectionMethods" {
		return map[string]interface{}{
			"Members": []interface{}{
				map[string]interface{}{
					"@odata.id": "/redfish/v1/ConnectionMethods/variant1",
				},
			},
		}, 200, nil
		//Test_odimUtils_updateConnectionMethodVariants mock2
	} else if url == "/redfish/v1/ConnectionMethods/variant1" {
		return map[string]interface{}{
			"ConnectionMethodVariant": "ILO2.0",
			"@odata.id":               "/redfish/v1/ConnectionMethods/id",
		}, 200, nil
		//
	} else if url == "/redfish/v1/EventService/Subscriptions" {
		return map[string]interface{}{
			"Members": []interface{}{
				map[string]interface{}{
					"@odata.id": "/redfish/v1/EventService/Subscriptions/1",
				},
			},
		}, 200, nil
	} else if url == "/redfish/v1/EventService/Subscriptions/1" {
		return map[string]interface{}{
			"Name": "BmcOperatorSubscription",
		}, 200, nil
	}
	return nil, 0, nil
}
func (m *mockRestClient) Patch(string, string, interface{}) (*http.Response, error) {
	return nil, nil
}
func (m *mockRestClient) Delete(url string, reason string) (*http.Response, error) {
	if url == "/redfish/v1/EventService/Subscriptions/1" {
		return &http.Response{StatusCode: 200}, nil
	}
	return nil, nil
}
func (m *mockRestClient) Put(string, string, interface{}) (*http.Response, error) {
	return nil, nil
}
func (m *mockRestClient) RecallWithNewToken(url, reason, method string, body interface{}) (*http.Response, error) {
	return nil, nil
}

// --------------------------COMMON UTIL MOCKS------------------------
func (m mockCommonUtil) GetBmcSystemDetails(context.Context, *infraiov1.Bmc) map[string]interface{} {
	return nil
}
func (m mockCommonUtil) MoniteringTaskmon(ctx context.Context, headerInfo http.Header, operation, resourceName string) (bool, map[string]interface{}) {
	// taskmon success
	return true, nil
}
func (m mockCommonUtil) BmcAddition(ctx context.Context, bmcObject *infraiov1.Bmc, body []byte) (bool, map[string]interface{}) {
	return false, nil
}
func (m mockCommonUtil) BmcDeleteOperation(ctx context.Context, aggregationURL string, resource string) bool {
	return false
}