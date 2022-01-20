// Copyright 2022 Google LLC. All Rights Reserved.
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
package google

import (
	_ "github.com/GoogleCloudPlatform/declarative-resource-client-library/unstructured/google/apigee"
	_ "github.com/GoogleCloudPlatform/declarative-resource-client-library/unstructured/google/apigee/alpha"
	_ "github.com/GoogleCloudPlatform/declarative-resource-client-library/unstructured/google/apigee/beta"
	_ "github.com/GoogleCloudPlatform/declarative-resource-client-library/unstructured/google/apikeys"
	_ "github.com/GoogleCloudPlatform/declarative-resource-client-library/unstructured/google/apikeys/alpha"
	_ "github.com/GoogleCloudPlatform/declarative-resource-client-library/unstructured/google/apikeys/beta"
	_ "github.com/GoogleCloudPlatform/declarative-resource-client-library/unstructured/google/assuredworkloads"
	_ "github.com/GoogleCloudPlatform/declarative-resource-client-library/unstructured/google/assuredworkloads/alpha"
	_ "github.com/GoogleCloudPlatform/declarative-resource-client-library/unstructured/google/assuredworkloads/beta"
	_ "github.com/GoogleCloudPlatform/declarative-resource-client-library/unstructured/google/bigquery"
	_ "github.com/GoogleCloudPlatform/declarative-resource-client-library/unstructured/google/bigquery/alpha"
	_ "github.com/GoogleCloudPlatform/declarative-resource-client-library/unstructured/google/bigquery/beta"
	_ "github.com/GoogleCloudPlatform/declarative-resource-client-library/unstructured/google/bigqueryreservation"
	_ "github.com/GoogleCloudPlatform/declarative-resource-client-library/unstructured/google/bigqueryreservation/alpha"
	_ "github.com/GoogleCloudPlatform/declarative-resource-client-library/unstructured/google/bigqueryreservation/beta"
	_ "github.com/GoogleCloudPlatform/declarative-resource-client-library/unstructured/google/billingbudgets"
	_ "github.com/GoogleCloudPlatform/declarative-resource-client-library/unstructured/google/billingbudgets/alpha"
	_ "github.com/GoogleCloudPlatform/declarative-resource-client-library/unstructured/google/billingbudgets/beta"
	_ "github.com/GoogleCloudPlatform/declarative-resource-client-library/unstructured/google/binaryauthorization"
	_ "github.com/GoogleCloudPlatform/declarative-resource-client-library/unstructured/google/binaryauthorization/alpha"
	_ "github.com/GoogleCloudPlatform/declarative-resource-client-library/unstructured/google/binaryauthorization/beta"
	_ "github.com/GoogleCloudPlatform/declarative-resource-client-library/unstructured/google/cloudbuild"
	_ "github.com/GoogleCloudPlatform/declarative-resource-client-library/unstructured/google/cloudbuild/alpha"
	_ "github.com/GoogleCloudPlatform/declarative-resource-client-library/unstructured/google/cloudbuild/beta"
	_ "github.com/GoogleCloudPlatform/declarative-resource-client-library/unstructured/google/cloudfunctions"
	_ "github.com/GoogleCloudPlatform/declarative-resource-client-library/unstructured/google/cloudfunctions/alpha"
	_ "github.com/GoogleCloudPlatform/declarative-resource-client-library/unstructured/google/cloudfunctions/beta"
	_ "github.com/GoogleCloudPlatform/declarative-resource-client-library/unstructured/google/cloudidentity"
	_ "github.com/GoogleCloudPlatform/declarative-resource-client-library/unstructured/google/cloudidentity/alpha"
	_ "github.com/GoogleCloudPlatform/declarative-resource-client-library/unstructured/google/cloudidentity/beta"
	_ "github.com/GoogleCloudPlatform/declarative-resource-client-library/unstructured/google/cloudkms"
	_ "github.com/GoogleCloudPlatform/declarative-resource-client-library/unstructured/google/cloudkms/alpha"
	_ "github.com/GoogleCloudPlatform/declarative-resource-client-library/unstructured/google/cloudkms/beta"
	_ "github.com/GoogleCloudPlatform/declarative-resource-client-library/unstructured/google/cloudresourcemanager"
	_ "github.com/GoogleCloudPlatform/declarative-resource-client-library/unstructured/google/cloudresourcemanager/alpha"
	_ "github.com/GoogleCloudPlatform/declarative-resource-client-library/unstructured/google/cloudresourcemanager/beta"
	_ "github.com/GoogleCloudPlatform/declarative-resource-client-library/unstructured/google/cloudscheduler"
	_ "github.com/GoogleCloudPlatform/declarative-resource-client-library/unstructured/google/cloudscheduler/alpha"
	_ "github.com/GoogleCloudPlatform/declarative-resource-client-library/unstructured/google/cloudscheduler/beta"
	_ "github.com/GoogleCloudPlatform/declarative-resource-client-library/unstructured/google/compute"
	_ "github.com/GoogleCloudPlatform/declarative-resource-client-library/unstructured/google/compute/alpha"
	_ "github.com/GoogleCloudPlatform/declarative-resource-client-library/unstructured/google/compute/beta"
	_ "github.com/GoogleCloudPlatform/declarative-resource-client-library/unstructured/google/configcontroller/alpha"
	_ "github.com/GoogleCloudPlatform/declarative-resource-client-library/unstructured/google/containeranalysis"
	_ "github.com/GoogleCloudPlatform/declarative-resource-client-library/unstructured/google/containeranalysis/alpha"
	_ "github.com/GoogleCloudPlatform/declarative-resource-client-library/unstructured/google/containeranalysis/beta"
	_ "github.com/GoogleCloudPlatform/declarative-resource-client-library/unstructured/google/containeraws"
	_ "github.com/GoogleCloudPlatform/declarative-resource-client-library/unstructured/google/containeraws/alpha"
	_ "github.com/GoogleCloudPlatform/declarative-resource-client-library/unstructured/google/containeraws/beta"
	_ "github.com/GoogleCloudPlatform/declarative-resource-client-library/unstructured/google/containerazure"
	_ "github.com/GoogleCloudPlatform/declarative-resource-client-library/unstructured/google/containerazure/alpha"
	_ "github.com/GoogleCloudPlatform/declarative-resource-client-library/unstructured/google/containerazure/beta"
	_ "github.com/GoogleCloudPlatform/declarative-resource-client-library/unstructured/google/datafusion/alpha"
	_ "github.com/GoogleCloudPlatform/declarative-resource-client-library/unstructured/google/datafusion/beta"
	_ "github.com/GoogleCloudPlatform/declarative-resource-client-library/unstructured/google/dataproc"
	_ "github.com/GoogleCloudPlatform/declarative-resource-client-library/unstructured/google/dataproc/alpha"
	_ "github.com/GoogleCloudPlatform/declarative-resource-client-library/unstructured/google/dataproc/beta"
	_ "github.com/GoogleCloudPlatform/declarative-resource-client-library/unstructured/google/eventarc"
	_ "github.com/GoogleCloudPlatform/declarative-resource-client-library/unstructured/google/eventarc/alpha"
	_ "github.com/GoogleCloudPlatform/declarative-resource-client-library/unstructured/google/eventarc/beta"
	_ "github.com/GoogleCloudPlatform/declarative-resource-client-library/unstructured/google/filestore"
	_ "github.com/GoogleCloudPlatform/declarative-resource-client-library/unstructured/google/filestore/alpha"
	_ "github.com/GoogleCloudPlatform/declarative-resource-client-library/unstructured/google/filestore/beta"
	_ "github.com/GoogleCloudPlatform/declarative-resource-client-library/unstructured/google/firebaserules"
	_ "github.com/GoogleCloudPlatform/declarative-resource-client-library/unstructured/google/firebaserules/alpha"
	_ "github.com/GoogleCloudPlatform/declarative-resource-client-library/unstructured/google/firebaserules/beta"
	_ "github.com/GoogleCloudPlatform/declarative-resource-client-library/unstructured/google/gameservices"
	_ "github.com/GoogleCloudPlatform/declarative-resource-client-library/unstructured/google/gameservices/alpha"
	_ "github.com/GoogleCloudPlatform/declarative-resource-client-library/unstructured/google/gameservices/beta"
	_ "github.com/GoogleCloudPlatform/declarative-resource-client-library/unstructured/google/gkehub/alpha"
	_ "github.com/GoogleCloudPlatform/declarative-resource-client-library/unstructured/google/gkehub/beta"
	_ "github.com/GoogleCloudPlatform/declarative-resource-client-library/unstructured/google/iam"
	_ "github.com/GoogleCloudPlatform/declarative-resource-client-library/unstructured/google/iam/alpha"
	_ "github.com/GoogleCloudPlatform/declarative-resource-client-library/unstructured/google/iam/beta"
	_ "github.com/GoogleCloudPlatform/declarative-resource-client-library/unstructured/google/iap"
	_ "github.com/GoogleCloudPlatform/declarative-resource-client-library/unstructured/google/iap/alpha"
	_ "github.com/GoogleCloudPlatform/declarative-resource-client-library/unstructured/google/iap/beta"
	_ "github.com/GoogleCloudPlatform/declarative-resource-client-library/unstructured/google/identitytoolkit"
	_ "github.com/GoogleCloudPlatform/declarative-resource-client-library/unstructured/google/identitytoolkit/alpha"
	_ "github.com/GoogleCloudPlatform/declarative-resource-client-library/unstructured/google/identitytoolkit/beta"
	_ "github.com/GoogleCloudPlatform/declarative-resource-client-library/unstructured/google/logging"
	_ "github.com/GoogleCloudPlatform/declarative-resource-client-library/unstructured/google/logging/alpha"
	_ "github.com/GoogleCloudPlatform/declarative-resource-client-library/unstructured/google/logging/beta"
	_ "github.com/GoogleCloudPlatform/declarative-resource-client-library/unstructured/google/monitoring"
	_ "github.com/GoogleCloudPlatform/declarative-resource-client-library/unstructured/google/monitoring/alpha"
	_ "github.com/GoogleCloudPlatform/declarative-resource-client-library/unstructured/google/monitoring/beta"
	_ "github.com/GoogleCloudPlatform/declarative-resource-client-library/unstructured/google/networkconnectivity"
	_ "github.com/GoogleCloudPlatform/declarative-resource-client-library/unstructured/google/networkconnectivity/alpha"
	_ "github.com/GoogleCloudPlatform/declarative-resource-client-library/unstructured/google/networkconnectivity/beta"
	_ "github.com/GoogleCloudPlatform/declarative-resource-client-library/unstructured/google/networksecurity/alpha"
	_ "github.com/GoogleCloudPlatform/declarative-resource-client-library/unstructured/google/networksecurity/beta"
	_ "github.com/GoogleCloudPlatform/declarative-resource-client-library/unstructured/google/networkservices/alpha"
	_ "github.com/GoogleCloudPlatform/declarative-resource-client-library/unstructured/google/orgpolicy"
	_ "github.com/GoogleCloudPlatform/declarative-resource-client-library/unstructured/google/orgpolicy/alpha"
	_ "github.com/GoogleCloudPlatform/declarative-resource-client-library/unstructured/google/orgpolicy/beta"
	_ "github.com/GoogleCloudPlatform/declarative-resource-client-library/unstructured/google/osconfig"
	_ "github.com/GoogleCloudPlatform/declarative-resource-client-library/unstructured/google/osconfig/alpha"
	_ "github.com/GoogleCloudPlatform/declarative-resource-client-library/unstructured/google/osconfig/beta"
	_ "github.com/GoogleCloudPlatform/declarative-resource-client-library/unstructured/google/privateca"
	_ "github.com/GoogleCloudPlatform/declarative-resource-client-library/unstructured/google/privateca/alpha"
	_ "github.com/GoogleCloudPlatform/declarative-resource-client-library/unstructured/google/privateca/beta"
	_ "github.com/GoogleCloudPlatform/declarative-resource-client-library/unstructured/google/pubsub"
	_ "github.com/GoogleCloudPlatform/declarative-resource-client-library/unstructured/google/pubsub/alpha"
	_ "github.com/GoogleCloudPlatform/declarative-resource-client-library/unstructured/google/pubsub/beta"
	_ "github.com/GoogleCloudPlatform/declarative-resource-client-library/unstructured/google/recaptchaenterprise"
	_ "github.com/GoogleCloudPlatform/declarative-resource-client-library/unstructured/google/recaptchaenterprise/alpha"
	_ "github.com/GoogleCloudPlatform/declarative-resource-client-library/unstructured/google/recaptchaenterprise/beta"
	_ "github.com/GoogleCloudPlatform/declarative-resource-client-library/unstructured/google/run/alpha"
	_ "github.com/GoogleCloudPlatform/declarative-resource-client-library/unstructured/google/servicemanagement"
	_ "github.com/GoogleCloudPlatform/declarative-resource-client-library/unstructured/google/servicemanagement/alpha"
	_ "github.com/GoogleCloudPlatform/declarative-resource-client-library/unstructured/google/servicemanagement/beta"
	_ "github.com/GoogleCloudPlatform/declarative-resource-client-library/unstructured/google/storage"
	_ "github.com/GoogleCloudPlatform/declarative-resource-client-library/unstructured/google/storage/alpha"
	_ "github.com/GoogleCloudPlatform/declarative-resource-client-library/unstructured/google/storage/beta"
	_ "github.com/GoogleCloudPlatform/declarative-resource-client-library/unstructured/google/vmware/alpha"
	_ "github.com/GoogleCloudPlatform/declarative-resource-client-library/unstructured/google/vpcaccess"
	_ "github.com/GoogleCloudPlatform/declarative-resource-client-library/unstructured/google/vpcaccess/alpha"
	_ "github.com/GoogleCloudPlatform/declarative-resource-client-library/unstructured/google/vpcaccess/beta"
)
