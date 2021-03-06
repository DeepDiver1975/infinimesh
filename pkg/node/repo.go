//--------------------------------------------------------------------------
// Copyright 2018 Infinite Devices GmbH
// www.infinimesh.io
//
//   Licensed under the Apache License, Version 2.0 (the "License");
//   you may not use this file except in compliance with the License.
//   You may obtain a copy of the License at
//
//       http://www.apache.org/licenses/LICENSE-2.0
//
//   Unless required by applicable law or agreed to in writing, software
//   distributed under the License is distributed on an "AS IS" BASIS,
//   WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//   See the License for the specific language governing permissions and
//   limitations under the License.
//--------------------------------------------------------------------------

package node

import (
	"context"

	"github.com/infinimesh/infinimesh/pkg/node/nodepb"
)

const (
	KindAsset  = "asset"
	KindDevice = "device"
)

type Repo interface {
	CreateUserAccount(ctx context.Context, username, password string, isRoot, enabled bool) (uid string, err error)
	ListAccounts(ctx context.Context) (accounts []*nodepb.Account, err error)
	UpdateAccount(ctx context.Context, account *nodepb.UpdateAccountRequest) (err error)
	GetAccount(ctx context.Context, accountID string) (account *nodepb.Account, err error)
	SetPassword(ctx context.Context, account, password string) error

	IsAuthorized(ctx context.Context, target, who, action string) (decision bool, err error)
	IsAuthorizedNamespace(ctx context.Context, namespace, account string, action nodepb.Action) (decision bool, err error)
	Authorize(ctx context.Context, account, node, action string, inherit bool) (err error)
	AuthorizeNamespace(ctx context.Context, account, namespace string, action nodepb.Action) (err error)
	Authenticate(ctx context.Context, username, password string) (success bool, uid string, defaultNamespace string, err error)

	CreateObject(ctx context.Context, name, parentID, kind, namespaceID string) (id string, err error)
	DeleteObject(ctx context.Context, uid string) (err error)
	ListForAccount(ctx context.Context, account string, namespace string, recurse bool) (inheritedObjects []*nodepb.Object, err error)

	CreateNamespace(ctx context.Context, name string) (id string, err error)
	GetNamespace(ctx context.Context, uid string) (namespace *nodepb.Namespace, err error)
	ListNamespaces(ctx context.Context) (namespaces []*nodepb.Namespace, err error)
	ListNamespacesForAccount(ctx context.Context, accountID string) (namespaces []*nodepb.Namespace, err error)
	ListPermissionsInNamespace(ctx context.Context, namespace string) (permissions []*nodepb.Permission, err error)
	DeletePermissionInNamespace(ctx context.Context, namespace, accountID string) (err error)
}
