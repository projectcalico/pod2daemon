// Copyright (c) 2021 Tigera, Inc. All rights reserved.
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

package binder

type BinderOption func(*binder)

type funcCredentialFilter func(Credentials) error

// FilterCredential is a function that receives the connection Credential info
// return a non nil erro to deny the credential.
func CredentialFilter(f func(Credentials) error) BinderOption {
	return func(b *binder) {
		b.workloads.credentialFilter = f
	}
}
