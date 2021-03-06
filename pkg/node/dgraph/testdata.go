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

package dgraph

import (
	"context"
	"fmt"

	"github.com/dgraph-io/dgo"
	"github.com/dgraph-io/dgo/protos/api"

	"github.com/infinimesh/infinimesh/pkg/node"
	"github.com/infinimesh/infinimesh/pkg/node/nodepb"
)

func ImportSchema(dg *dgo.Dgraph, drop bool) error {
	if drop {
		err := dg.Alter(context.Background(), &api.Operation{DropAll: drop})
		if err != nil {
			return err
		}
	}
	schema := `
  tags: [string] .
  name: string @index(exact) .
  username: string @index(exact) .
  action: string @index(term) .
  type: string @index(exact) .
  access.to: uid @reverse .
  children: uid @reverse .
  owns: uid @reverse .
  kind: string @index(exact) .
  has.credentials: uid @reverse .
  access.to.namespace: uid @reverse .
  fingerprint: string @index(exact) .
  certificates: uid @reverse .
  password: password .`
	fmt.Println("Apply schema ", schema)
	return dg.Alter(context.Background(), &api.Operation{
		Schema: schema,
	})

}

func ImportStandardSet(repo node.Repo) (userID string, adminID string, err error) {
	// careful,  currently when referencing a namespace, the name of it has to be used, not the id (0x...)
	sharedNs := "shared-project"
	_, err = repo.CreateNamespace(context.Background(), sharedNs)
	if err != nil {
		return "", "", err
	}

	ns := "joe"
	joe, err := repo.CreateUserAccount(context.Background(), "joe", "test123", false, true)
	if err != nil {
		return "", "", err
	}
	fmt.Println("User joe: ", joe)

	hanswurst, err := repo.CreateUserAccount(context.Background(), "hanswurst", "hanswurst", false, true)
	if err != nil {
		return "", "", err
	}

	fmt.Println("User hanswurst: ", hanswurst)

	// Authorize both users on a shared project
	{
		err = repo.AuthorizeNamespace(context.Background(), joe, sharedNs, nodepb.Action_WRITE)
		if err != nil {
			return "", "", err
		}

		err = repo.AuthorizeNamespace(context.Background(), hanswurst, sharedNs, nodepb.Action_WRITE)
		if err != nil {
			return "", "", err
		}
	}

	admin, err := repo.CreateUserAccount(context.Background(), "admin", "admin123", true, true)
	if err != nil {
		return "", "", err
	}
	fmt.Println("Admin: ", admin)

	building, err := repo.CreateObject(context.Background(), "Angerstr 14", "", node.KindAsset, ns)
	if err != nil {
		return "", "", err
	}

	first, err := repo.CreateObject(context.Background(), "First Floor", building, node.KindAsset, ns)
	if err != nil {
		return "", "", err
	}

	_, err = repo.CreateObject(context.Background(), "Second Floor", building, node.KindAsset, ns)
	if err != nil {
		return "", "", err
	}

	apartment1Right, err := repo.CreateObject(context.Background(), "Apartment right side", first, node.KindAsset, ns)
	if err != nil {
		return "", "", err
	}

	_, err = repo.CreateObject(context.Background(), "Entrance", apartment1Right, node.KindAsset, ns)
	if err != nil {
		return "", "", err
	}

	_, err = repo.CreateObject(context.Background(), "Bathroom", apartment1Right, node.KindAsset, ns)
	if err != nil {
		return "", "", err
	}

	_, err = repo.CreateObject(context.Background(), "Kitchen", apartment1Right, node.KindAsset, ns)
	if err != nil {
		return "", "", err
	}

	_, err = repo.CreateObject(context.Background(), "Bedroom", apartment1Right, node.KindAsset, ns)
	if err != nil {
		return "", "", err
	}

	_, err = repo.CreateObject(context.Background(), "Kinderzimmer", apartment1Right, node.KindAsset, ns)
	if err != nil {
		return "", "", err
	}

	_, err = repo.CreateObject(context.Background(), "Walk-through room", apartment1Right, node.KindAsset, ns)
	if err != nil {
		return "", "", err
	}

	livingRoom, err := repo.CreateObject(context.Background(), "Living room", apartment1Right, node.KindAsset, ns)
	if err != nil {
		return "", "", err
	}

	_, err = repo.CreateObject(context.Background(), "Test-device", livingRoom, node.KindDevice, ns)
	if err != nil {
		return "", "", err
	}

	_, err = repo.CreateObject(context.Background(), "Test-device-no-parent", "", node.KindDevice, ns)
	if err != nil {
		return "", "", err
	}

	fmt.Println("User: ", joe)

	// result := repo.Authorize(context.Background(), joe, apartment1Right, "WRITE", true)

	return joe, admin, err
}
