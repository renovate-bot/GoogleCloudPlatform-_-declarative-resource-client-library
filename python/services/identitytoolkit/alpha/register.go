// Copyright 2025 Google LLC. All Rights Reserved.
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
package server

import (
	"google.golang.org/grpc"

	sdkgrpc "github.com/GoogleCloudPlatform/declarative-resource-client-library/python/proto/identitytoolkit/alpha/identitytoolkit_alpha_go_proto"
)

// RegisterServers registers each resource with the gRPC server.
func RegisterServers(s *grpc.Server) {
	sdkgrpc.RegisterIdentitytoolkitAlphaConfigServiceServer(s, &ConfigServer{})
	sdkgrpc.RegisterIdentitytoolkitAlphaOAuthIdpConfigServiceServer(s, &OAuthIdpConfigServer{})
	sdkgrpc.RegisterIdentitytoolkitAlphaTenantServiceServer(s, &TenantServer{})
	sdkgrpc.RegisterIdentitytoolkitAlphaTenantOAuthIdpConfigServiceServer(s, &TenantOAuthIdpConfigServer{})
}
