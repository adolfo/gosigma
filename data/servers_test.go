// Copyright 2014 ALTOROS
// Licensed under the AGPLv3, see LICENSE file for details.

package data

import (
	"encoding/json"
	"testing"
)

const jsonServers = `
{
    "meta": {
        "limit": 0,
        "offset": 0,
        "total_count": 5
    },
    "objects": [
        {
            "name": "test_server_4",
            "owner": {
                "resource_uri": "/api/2.0/user/80cb30fb-0ea3-43db-b27b-a125752cc0bf/",
                "uuid": "80cb30fb-0ea3-43db-b27b-a125752cc0bf"
            },
            "resource_uri": "/api/2.0/servers/43b1110a-31c5-41cc-a3e7-0b806076a913/",
            "runtime": null,
            "status": "stopped",
            "uuid": "43b1110a-31c5-41cc-a3e7-0b806076a913"
        },
        {
            "name": "test_server_2",
            "owner": {
                "resource_uri": "/api/2.0/user/80cb30fb-0ea3-43db-b27b-a125752cc0bf/",
                "uuid": "80cb30fb-0ea3-43db-b27b-a125752cc0bf"
            },
            "resource_uri": "/api/2.0/servers/3be1ebc6-1d03-4c4b-88ff-02557b940d19/",
            "runtime": null,
            "status": "stopped",
            "uuid": "3be1ebc6-1d03-4c4b-88ff-02557b940d19"
        },
        {
            "name": "test_server_0",
            "owner": {
                "resource_uri": "/api/2.0/user/80cb30fb-0ea3-43db-b27b-a125752cc0bf/",
                "uuid": "80cb30fb-0ea3-43db-b27b-a125752cc0bf"
            },
            "resource_uri": "/api/2.0/servers/b1defe23-e725-474d-acba-e46baa232611/",
            "runtime": null,
            "status": "stopped",
            "uuid": "b1defe23-e725-474d-acba-e46baa232611"
        },
        {
            "name": "test_server_3",
            "owner": {
                "resource_uri": "/api/2.0/user/80cb30fb-0ea3-43db-b27b-a125752cc0bf/",
                "uuid": "80cb30fb-0ea3-43db-b27b-a125752cc0bf"
            },
            "resource_uri": "/api/2.0/servers/cff0f338-2b84-4846-a028-3ec9e1b86184/",
            "runtime": null,
            "status": "stopped",
            "uuid": "cff0f338-2b84-4846-a028-3ec9e1b86184"
        },
        {
            "name": "test_server_1",
            "owner": {
                "resource_uri": "/api/2.0/user/80cb30fb-0ea3-43db-b27b-a125752cc0bf/",
                "uuid": "80cb30fb-0ea3-43db-b27b-a125752cc0bf"
            },
            "resource_uri": "/api/2.0/servers/93a04cd5-84cb-41fc-af17-683e3868ee95/",
            "runtime": null,
            "status": "stopped",
            "uuid": "93a04cd5-84cb-41fc-af17-683e3868ee95"
        }
    ]
}
`

func TestJsonUnmarshal(t *testing.T) {
	var ii Servers
	ii.Meta.Limit = 12345
	ii.Meta.Offset = 12345
	ii.Meta.TotalCount = 12345
	err := json.Unmarshal([]byte(jsonServers), &ii)
	if err != nil {
		t.Error(err)
	}

	if ii.Meta.Limit != 0 {
		t.Errorf("Meta.Limit = %d, wants 0", ii.Meta.Limit)
	}
	if ii.Meta.Offset != 0 {
		t.Errorf("Meta.Offset = %d, wants 0", ii.Meta.Offset)
	}
	if ii.Meta.TotalCount != 5 {
		t.Errorf("Meta.TotalCount = %d, wants 5", ii.Meta.TotalCount)
	}
	if len(ii.Objects) != 5 {
		t.Errorf("Meta.Objects.len = %d, wants 5", len(ii.Objects))
	}

	verify := func(i int, name, uri, status, uuid string) {
		obj := ii.Objects[i]
		if obj.Name != name {
			t.Errorf("Object %d, Name = '%s', wants '%s'", i, obj.Name, name)
		}
		if obj.URI != uri {
			t.Errorf("Object %d, URI = '%s', wants '%s'", i, obj.URI, uri)
		}
		if obj.Status != status {
			t.Errorf("Object %d, Status = '%s', wants '%s'", i, obj.Status, status)
		}
		if obj.UUID != uuid {
			t.Errorf("Object %d, UUID = '%s', wants '%s'", i, obj.UUID, uuid)
		}
	}

	verify(0, "test_server_4", "/api/2.0/servers/43b1110a-31c5-41cc-a3e7-0b806076a913/",
		"stopped", "43b1110a-31c5-41cc-a3e7-0b806076a913")
	verify(1, "test_server_2", "/api/2.0/servers/3be1ebc6-1d03-4c4b-88ff-02557b940d19/",
		"stopped", "3be1ebc6-1d03-4c4b-88ff-02557b940d19")
	verify(2, "test_server_0", "/api/2.0/servers/b1defe23-e725-474d-acba-e46baa232611/",
		"stopped", "b1defe23-e725-474d-acba-e46baa232611")
	verify(3, "test_server_3", "/api/2.0/servers/cff0f338-2b84-4846-a028-3ec9e1b86184/",
		"stopped", "cff0f338-2b84-4846-a028-3ec9e1b86184")
	verify(4, "test_server_1", "/api/2.0/servers/93a04cd5-84cb-41fc-af17-683e3868ee95/",
		"stopped", "93a04cd5-84cb-41fc-af17-683e3868ee95")
}
