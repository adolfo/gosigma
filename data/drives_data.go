// Copyright 2014 ALTOROS
// Licensed under the AGPLv3, see LICENSE file for details.

package data

var driveOwner = Resource{
	"/api/2.0/user/80cb30fb-0ea3-43db-b27b-a125752cc0bf/",
	"80cb30fb-0ea3-43db-b27b-a125752cc0bf",
}

var drivesData = []Drive{
	Drive{
		DriveRecord: DriveRecord{
			Resource{"/api/2.0/drives/2ef7b7c7-7ec4-47a7-9b69-087c9417c0ff/",
				"2ef7b7c7-7ec4-47a7-9b69-087c9417c0ff"},
			driveOwner, "unmounted"},
	},
	Drive{
		DriveRecord: DriveRecord{
			Resource{"/api/2.0/drives/3b30c7ef-1fda-416d-91d1-ba616859360c/",
				"3b30c7ef-1fda-416d-91d1-ba616859360c"},
			driveOwner, "unmounted"},
	},
	Drive{
		DriveRecord: DriveRecord{
			Resource{"/api/2.0/drives/464aed14-8604-4277-be3c-9d53151d53b4/",
				"464aed14-8604-4277-be3c-9d53151d53b4"},
			driveOwner, "unmounted"},
	},
}

const jsonDrivesData = `{
    "meta": {
        "limit": 0,
        "offset": 0,
        "total_count": 9
    },
    "objects": [
        {
            "owner": {
                "resource_uri": "/api/2.0/user/80cb30fb-0ea3-43db-b27b-a125752cc0bf/",
                "uuid": "80cb30fb-0ea3-43db-b27b-a125752cc0bf"
            },
            "resource_uri": "/api/2.0/drives/2ef7b7c7-7ec4-47a7-9b69-087c9417c0ff/",
            "status": "unmounted",
            "uuid": "2ef7b7c7-7ec4-47a7-9b69-087c9417c0ff"
        },
        {
            "owner": {
                "resource_uri": "/api/2.0/user/80cb30fb-0ea3-43db-b27b-a125752cc0bf/",
                "uuid": "80cb30fb-0ea3-43db-b27b-a125752cc0bf"
            },
            "resource_uri": "/api/2.0/drives/3b30c7ef-1fda-416d-91d1-ba616859360c/",
            "status": "unmounted",
            "uuid": "3b30c7ef-1fda-416d-91d1-ba616859360c"
        },
        {
            "owner": {
                "resource_uri": "/api/2.0/user/80cb30fb-0ea3-43db-b27b-a125752cc0bf/",
                "uuid": "80cb30fb-0ea3-43db-b27b-a125752cc0bf"
            },
            "resource_uri": "/api/2.0/drives/464aed14-8604-4277-be3c-9d53151d53b4/",
            "status": "unmounted",
            "uuid": "464aed14-8604-4277-be3c-9d53151d53b4"
        },
        {
            "owner": {
                "resource_uri": "/api/2.0/user/80cb30fb-0ea3-43db-b27b-a125752cc0bf/",
                "uuid": "80cb30fb-0ea3-43db-b27b-a125752cc0bf"
            },
            "resource_uri": "/api/2.0/drives/47ec5074-6058-4b0f-9505-78c83bd5a88b/",
            "status": "unmounted",
            "uuid": "47ec5074-6058-4b0f-9505-78c83bd5a88b"
        },
        {
            "owner": {
                "resource_uri": "/api/2.0/user/80cb30fb-0ea3-43db-b27b-a125752cc0bf/",
                "uuid": "80cb30fb-0ea3-43db-b27b-a125752cc0bf"
            },
            "resource_uri": "/api/2.0/drives/7949e52e-c8ba-461b-a84f-3f247221c644/",
            "status": "unmounted",
            "uuid": "7949e52e-c8ba-461b-a84f-3f247221c644"
        },
        {
            "owner": {
                "resource_uri": "/api/2.0/user/80cb30fb-0ea3-43db-b27b-a125752cc0bf/",
                "uuid": "80cb30fb-0ea3-43db-b27b-a125752cc0bf"
            },
            "resource_uri": "/api/2.0/drives/81b020b0-8ea0-4602-b778-e4df4539f0f7/",
            "status": "unmounted",
            "uuid": "81b020b0-8ea0-4602-b778-e4df4539f0f7"
        },
        {
            "owner": {
                "resource_uri": "/api/2.0/user/80cb30fb-0ea3-43db-b27b-a125752cc0bf/",
                "uuid": "80cb30fb-0ea3-43db-b27b-a125752cc0bf"
            },
            "resource_uri": "/api/2.0/drives/baf8fed4-757f-4d9e-a23a-3b3ff81e16c4/",
            "status": "unmounted",
            "uuid": "baf8fed4-757f-4d9e-a23a-3b3ff81e16c4"
        },
        {
            "owner": {
                "resource_uri": "/api/2.0/user/80cb30fb-0ea3-43db-b27b-a125752cc0bf/",
                "uuid": "80cb30fb-0ea3-43db-b27b-a125752cc0bf"
            },
            "resource_uri": "/api/2.0/drives/cae6df75-00a1-490c-a96b-51777b3ec515/",
            "status": "unmounted",
            "uuid": "cae6df75-00a1-490c-a96b-51777b3ec515"
        },
        {
            "owner": {
                "resource_uri": "/api/2.0/user/80cb30fb-0ea3-43db-b27b-a125752cc0bf/",
                "uuid": "80cb30fb-0ea3-43db-b27b-a125752cc0bf"
            },
            "resource_uri": "/api/2.0/drives/e15dd971-3ef8-497c-9f92-90d5ca1722bd/",
            "status": "unmounted",
            "uuid": "e15dd971-3ef8-497c-9f92-90d5ca1722bd"
        }
    ]
}`

var drivesDetailData = []Drive{
	Drive{
		DriveRecord: DriveRecord{
			Resource{"/api/2.0/drives/2ef7b7c7-7ec4-47a7-9b69-087c9417c0ff/",
				"2ef7b7c7-7ec4-47a7-9b69-087c9417c0ff"},
			driveOwner, "unmounted"},
		Media:       "disk",
		Meta:        nil,
		Size:        1073741824,
		StorageType: "dssd",
		Jobs:        nil,
		Name:        "test_drive_2",
	},
	Drive{
		DriveRecord: DriveRecord{
			Resource{"/api/2.0/drives/3b30c7ef-1fda-416d-91d1-ba616859360c/",
				"3b30c7ef-1fda-416d-91d1-ba616859360c"},
			driveOwner, "unmounted"},
		Media:       "disk",
		Meta:        nil,
		Size:        10737418240,
		StorageType: "dssd",
		Jobs: []Resource{
			Resource{URI: "/api/2.0/jobs/fbe05708-fd42-43d5-814c-9cb805edd4cb/", UUID: "fbe05708-fd42-43d5-814c-9cb805edd4cb"},
			Resource{URI: "/api/2.0/jobs/32513930-6815-4cd4-ae8e-2eb89733c206/", UUID: "32513930-6815-4cd4-ae8e-2eb89733c206"},
		},
		Name: "atom",
	},
	Drive{
		DriveRecord: DriveRecord{
			Resource{"/api/2.0/drives/464aed14-8604-4277-be3c-9d53151d53b4/",
				"464aed14-8604-4277-be3c-9d53151d53b4"},
			driveOwner, "unmounted"},
		Media:       "disk",
		Meta:        nil,
		Size:        1073741824,
		StorageType: "dssd",
		Jobs:        nil,
		Name:        "test_drive_1",
	},
}

const jsonDrivesDetailData = `{
    "meta": {
        "limit": 0,
        "offset": 0,
        "total_count": 9
    },
    "objects": [
        {
            "affinities": [],
            "allow_multimount": false,
            "jobs": [],
            "licenses": [],
            "media": "disk",
            "meta": {},
            "mounted_on": [],
            "name": "test_drive_2",
            "owner": {
                "resource_uri": "/api/2.0/user/80cb30fb-0ea3-43db-b27b-a125752cc0bf/",
                "uuid": "80cb30fb-0ea3-43db-b27b-a125752cc0bf"
            },
            "resource_uri": "/api/2.0/drives/2ef7b7c7-7ec4-47a7-9b69-087c9417c0ff/",
            "runtime": {
                "is_snapshotable": true,
                "snapshots_allocated_size": 0,
                "storage_type": "dssd"
            },
            "size": 1073741824,
            "snapshots": [],
            "status": "unmounted",
            "storage_type": "dssd",
            "tags": [],
            "uuid": "2ef7b7c7-7ec4-47a7-9b69-087c9417c0ff"
        },
        {
            "affinities": [],
            "allow_multimount": false,
            "jobs": [
                {
                    "resource_uri": "/api/2.0/jobs/fbe05708-fd42-43d5-814c-9cb805edd4cb/",
                    "uuid": "fbe05708-fd42-43d5-814c-9cb805edd4cb"
                },
                {
                    "resource_uri": "/api/2.0/jobs/32513930-6815-4cd4-ae8e-2eb89733c206/",
                    "uuid": "32513930-6815-4cd4-ae8e-2eb89733c206"
                }
            ],
            "licenses": [],
            "media": "disk",
            "meta": {},
            "mounted_on": [],
            "name": "atom",
            "owner": {
                "resource_uri": "/api/2.0/user/80cb30fb-0ea3-43db-b27b-a125752cc0bf/",
                "uuid": "80cb30fb-0ea3-43db-b27b-a125752cc0bf"
            },
            "resource_uri": "/api/2.0/drives/3b30c7ef-1fda-416d-91d1-ba616859360c/",
            "runtime": {
                "is_snapshotable": true,
                "snapshots_allocated_size": 0,
                "storage_type": "dssd"
            },
            "size": 10737418240,
            "snapshots": [],
            "status": "unmounted",
            "storage_type": "dssd",
            "tags": [],
            "uuid": "3b30c7ef-1fda-416d-91d1-ba616859360c"
        },
        {
            "affinities": [],
            "allow_multimount": false,
            "jobs": [],
            "licenses": [],
            "media": "disk",
            "meta": {},
            "mounted_on": [],
            "name": "test_drive_1",
            "owner": {
                "resource_uri": "/api/2.0/user/80cb30fb-0ea3-43db-b27b-a125752cc0bf/",
                "uuid": "80cb30fb-0ea3-43db-b27b-a125752cc0bf"
            },
            "resource_uri": "/api/2.0/drives/464aed14-8604-4277-be3c-9d53151d53b4/",
            "runtime": {
                "is_snapshotable": true,
                "snapshots_allocated_size": 0,
                "storage_type": "dssd"
            },
            "size": 1073741824,
            "snapshots": [],
            "status": "unmounted",
            "storage_type": "dssd",
            "tags": [],
            "uuid": "464aed14-8604-4277-be3c-9d53151d53b4"
        },
        {
            "affinities": [],
            "allow_multimount": false,
            "jobs": [],
            "licenses": [],
            "media": "disk",
            "meta": {},
            "mounted_on": [],
            "name": "test_drive_4",
            "owner": {
                "resource_uri": "/api/2.0/user/80cb30fb-0ea3-43db-b27b-a125752cc0bf/",
                "uuid": "80cb30fb-0ea3-43db-b27b-a125752cc0bf"
            },
            "resource_uri": "/api/2.0/drives/47ec5074-6058-4b0f-9505-78c83bd5a88b/",
            "runtime": {
                "is_snapshotable": true,
                "snapshots_allocated_size": 0,
                "storage_type": "dssd"
            },
            "size": 1073741824,
            "snapshots": [],
            "status": "unmounted",
            "storage_type": "dssd",
            "tags": [],
            "uuid": "47ec5074-6058-4b0f-9505-78c83bd5a88b"
        },
        {
            "affinities": [],
            "allow_multimount": false,
            "jobs": [],
            "licenses": [],
            "media": "disk",
            "meta": {
                "description": ""
            },
            "mounted_on": [],
            "name": "xxx",
            "owner": {
                "resource_uri": "/api/2.0/user/80cb30fb-0ea3-43db-b27b-a125752cc0bf/",
                "uuid": "80cb30fb-0ea3-43db-b27b-a125752cc0bf"
            },
            "resource_uri": "/api/2.0/drives/7949e52e-c8ba-461b-a84f-3f247221c644/",
            "runtime": {
                "is_snapshotable": true,
                "snapshots_allocated_size": 0,
                "storage_type": "dssd"
            },
            "size": 1073741824,
            "snapshots": [],
            "status": "unmounted",
            "storage_type": "dssd",
            "tags": [],
            "uuid": "7949e52e-c8ba-461b-a84f-3f247221c644"
        },
        {
            "affinities": [],
            "allow_multimount": false,
            "jobs": [],
            "licenses": [],
            "media": "disk",
            "meta": {},
            "mounted_on": [],
            "name": "t1",
            "owner": {
                "resource_uri": "/api/2.0/user/80cb30fb-0ea3-43db-b27b-a125752cc0bf/",
                "uuid": "80cb30fb-0ea3-43db-b27b-a125752cc0bf"
            },
            "resource_uri": "/api/2.0/drives/81b020b0-8ea0-4602-b778-e4df4539f0f7/",
            "runtime": {
                "is_snapshotable": false,
                "snapshots_allocated_size": 0,
                "storage_type": "zadara"
            },
            "size": 3221225472,
            "snapshots": [],
            "status": "unmounted",
            "storage_type": "zadara",
            "tags": [],
            "uuid": "81b020b0-8ea0-4602-b778-e4df4539f0f7"
        },
        {
            "affinities": [],
            "allow_multimount": false,
            "jobs": [],
            "licenses": [],
            "media": "disk",
            "meta": {},
            "mounted_on": [],
            "name": "test_drive_3",
            "owner": {
                "resource_uri": "/api/2.0/user/80cb30fb-0ea3-43db-b27b-a125752cc0bf/",
                "uuid": "80cb30fb-0ea3-43db-b27b-a125752cc0bf"
            },
            "resource_uri": "/api/2.0/drives/baf8fed4-757f-4d9e-a23a-3b3ff81e16c4/",
            "runtime": {
                "is_snapshotable": true,
                "snapshots_allocated_size": 0,
                "storage_type": "dssd"
            },
            "size": 1073741824,
            "snapshots": [],
            "status": "unmounted",
            "storage_type": "dssd",
            "tags": [],
            "uuid": "baf8fed4-757f-4d9e-a23a-3b3ff81e16c4"
        },
        {
            "affinities": [],
            "allow_multimount": false,
            "jobs": [],
            "licenses": [],
            "media": "disk",
            "meta": {},
            "mounted_on": [],
            "name": "test_drive_0",
            "owner": {
                "resource_uri": "/api/2.0/user/80cb30fb-0ea3-43db-b27b-a125752cc0bf/",
                "uuid": "80cb30fb-0ea3-43db-b27b-a125752cc0bf"
            },
            "resource_uri": "/api/2.0/drives/cae6df75-00a1-490c-a96b-51777b3ec515/",
            "runtime": {
                "is_snapshotable": true,
                "snapshots_allocated_size": 0,
                "storage_type": "dssd"
            },
            "size": 1073741824,
            "snapshots": [],
            "status": "unmounted",
            "storage_type": "dssd",
            "tags": [],
            "uuid": "cae6df75-00a1-490c-a96b-51777b3ec515"
        },
        {
            "affinities": [],
            "allow_multimount": false,
            "jobs": [],
            "licenses": [],
            "media": "disk",
            "meta": {
                "arch": "64",
                "category": "general",
                "description": "",
                "favourite": "False",
                "image_type": "preinst",
                "install_notes": "",
                "os": "linux",
                "paid": "False",
                "url": ""
            },
            "mounted_on": [],
            "name": "otom",
            "owner": {
                "resource_uri": "/api/2.0/user/80cb30fb-0ea3-43db-b27b-a125752cc0bf/",
                "uuid": "80cb30fb-0ea3-43db-b27b-a125752cc0bf"
            },
            "resource_uri": "/api/2.0/drives/e15dd971-3ef8-497c-9f92-90d5ca1722bd/",
            "runtime": {
                "is_snapshotable": true,
                "snapshots_allocated_size": 0,
                "storage_type": "dssd"
            },
            "size": 10737418240,
            "snapshots": [],
            "status": "unmounted",
            "storage_type": "dssd",
            "tags": [],
            "uuid": "e15dd971-3ef8-497c-9f92-90d5ca1722bd"
        }
    ]
}`

var driveData = Drive{
	DriveRecord: DriveRecord{
		Resource{"/api/2.0/drives/2ef7b7c7-7ec4-47a7-9b69-087c9417c0ff/",
			"2ef7b7c7-7ec4-47a7-9b69-087c9417c0ff"},
		driveOwner, "unmounted"},
	Media:       "disk",
	Meta:        nil,
	Size:        1073741824,
	StorageType: "dssd",
	Jobs:        nil,
	Name:        "test_drive_2",
}

const jsonDriveData = `{
    "affinities": [],
    "allow_multimount": false,
    "jobs": [],
    "licenses": [],
    "media": "disk",
    "meta": {},
    "mounted_on": [],
    "name": "test_drive_2",
    "owner": {
        "resource_uri": "/api/2.0/user/80cb30fb-0ea3-43db-b27b-a125752cc0bf/",
        "uuid": "80cb30fb-0ea3-43db-b27b-a125752cc0bf"
    },
    "resource_uri": "/api/2.0/drives/2ef7b7c7-7ec4-47a7-9b69-087c9417c0ff/",
    "runtime": {
        "is_snapshotable": true,
        "snapshots_allocated_size": 0,
        "storage_type": "dssd"
    },
    "size": 1073741824,
    "snapshots": [],
    "status": "unmounted",
    "storage_type": "dssd",
    "tags": [],
    "uuid": "2ef7b7c7-7ec4-47a7-9b69-087c9417c0ff"
}`
