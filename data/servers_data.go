// Copyright 2014 ALTOROS
// Licensed under the AGPLv3, see LICENSE file for details.

package data

var serversData = []Server{
	Server{
		ServerRecord: ServerRecord{
			Resource{
				"/api/2.0/servers/43b1110a-31c5-41cc-a3e7-0b806076a913/",
				"43b1110a-31c5-41cc-a3e7-0b806076a913"},
			"test_server_4",
			"stopped",
		},
	},
	Server{
		ServerRecord: ServerRecord{
			Resource{
				"/api/2.0/servers/3be1ebc6-1d03-4c4b-88ff-02557b940d19/",
				"3be1ebc6-1d03-4c4b-88ff-02557b940d19"},
			"test_server_2",
			"stopped",
		},
	},
}

const jsonServersData = `{
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

var serversDetailData = []Server{
	Server{
		ServerRecord: ServerRecord{
			Resource{
				"/api/2.0/servers/43b1110a-31c5-41cc-a3e7-0b806076a913/",
				"43b1110a-31c5-41cc-a3e7-0b806076a913"},
			"test_server_4",
			"stopped",
		},
		CPU: 1000,
		Mem: 536870912,
		NICs: []NIC{
			NIC{
				IPv4: &IPv4{
					Conf: "static",
					IP:   Resource{URI: "/api/2.0/ips/31.171.246.37/", UUID: "31.171.246.37"},
				},
				Model: "virtio",
				MAC:   "22:40:85:4f:d3:ce",
			},
			NIC{
				Model: "virtio",
				MAC:   "22:aa:fe:07:48:3b",
				VLAN: &Resource{
					URI:  "/api/2.0/vlans/5bc05e7e-6555-4f40-add8-3b8e91447702/",
					UUID: "5bc05e7e-6555-4f40-add8-3b8e91447702",
				},
			},
		},
		Drives:      []ServerDrive{},
		VNCPassword: "testserver",
	},
}

const jsonServersDetailData = `{
    "meta": {
        "limit": 0,
        "offset": 0,
        "total_count": 5
    },
    "objects": [
        {
            "context": true,
            "cpu": 1000,
            "cpu_model": null,
            "cpus_instead_of_cores": false,
            "drives": [],
            "enable_numa": false,
            "hv_relaxed": false,
            "hv_tsc": false,
            "jobs": [],
            "mem": 536870912,
            "meta": {},
            "name": "test_server_4",
            "nics": [
                {
                    "boot_order": null,
                    "firewall_policy": null,
                    "ip_v4_conf": {
                        "conf": "static",
                        "ip": {
                            "resource_uri": "/api/2.0/ips/31.171.246.37/",
                            "uuid": "31.171.246.37"
                        }
                    },
                    "ip_v6_conf": null,
                    "mac": "22:40:85:4f:d3:ce",
                    "model": "virtio",
                    "runtime": null,
                    "vlan": null
                },
                {
                    "boot_order": null,
                    "firewall_policy": null,
                    "ip_v4_conf": null,
                    "ip_v6_conf": null,
                    "mac": "22:aa:fe:07:48:3b",
                    "model": "virtio",
                    "runtime": null,
                    "vlan": {
                        "resource_uri": "/api/2.0/vlans/5bc05e7e-6555-4f40-add8-3b8e91447702/",
                        "uuid": "5bc05e7e-6555-4f40-add8-3b8e91447702"
                    }
                }
            ],
            "owner": {
                "resource_uri": "/api/2.0/user/80cb30fb-0ea3-43db-b27b-a125752cc0bf/",
                "uuid": "80cb30fb-0ea3-43db-b27b-a125752cc0bf"
            },
            "requirements": [],
            "resource_uri": "/api/2.0/servers/43b1110a-31c5-41cc-a3e7-0b806076a913/",
            "runtime": null,
            "smp": 1,
            "status": "stopped",
            "tags": [],
            "uuid": "43b1110a-31c5-41cc-a3e7-0b806076a913",
            "vnc_password": "testserver"
        },
        {
            "context": true,
            "cpu": 1000,
            "cpu_model": null,
            "cpus_instead_of_cores": false,
            "drives": [],
            "enable_numa": false,
            "hv_relaxed": false,
            "hv_tsc": false,
            "jobs": [],
            "mem": 536870912,
            "meta": {},
            "name": "test_server_2",
            "nics": [],
            "owner": {
                "resource_uri": "/api/2.0/user/80cb30fb-0ea3-43db-b27b-a125752cc0bf/",
                "uuid": "80cb30fb-0ea3-43db-b27b-a125752cc0bf"
            },
            "requirements": [],
            "resource_uri": "/api/2.0/servers/3be1ebc6-1d03-4c4b-88ff-02557b940d19/",
            "runtime": null,
            "smp": 1,
            "status": "stopped",
            "tags": [],
            "uuid": "3be1ebc6-1d03-4c4b-88ff-02557b940d19",
            "vnc_password": "testserver"
        },
        {
            "context": true,
            "cpu": 1000,
            "cpu_model": null,
            "cpus_instead_of_cores": false,
            "drives": [],
            "enable_numa": false,
            "hv_relaxed": false,
            "hv_tsc": false,
            "jobs": [],
            "mem": 536870912,
            "meta": {},
            "name": "test_server_0",
            "nics": [],
            "owner": {
                "resource_uri": "/api/2.0/user/80cb30fb-0ea3-43db-b27b-a125752cc0bf/",
                "uuid": "80cb30fb-0ea3-43db-b27b-a125752cc0bf"
            },
            "requirements": [],
            "resource_uri": "/api/2.0/servers/b1defe23-e725-474d-acba-e46baa232611/",
            "runtime": null,
            "smp": 1,
            "status": "stopped",
            "tags": [],
            "uuid": "b1defe23-e725-474d-acba-e46baa232611",
            "vnc_password": "testserver"
        },
        {
            "context": true,
            "cpu": 1000,
            "cpu_model": null,
            "cpus_instead_of_cores": false,
            "drives": [],
            "enable_numa": false,
            "hv_relaxed": false,
            "hv_tsc": false,
            "jobs": [],
            "mem": 536870912,
            "meta": {},
            "name": "test_server_3",
            "nics": [],
            "owner": {
                "resource_uri": "/api/2.0/user/80cb30fb-0ea3-43db-b27b-a125752cc0bf/",
                "uuid": "80cb30fb-0ea3-43db-b27b-a125752cc0bf"
            },
            "requirements": [],
            "resource_uri": "/api/2.0/servers/cff0f338-2b84-4846-a028-3ec9e1b86184/",
            "runtime": null,
            "smp": 1,
            "status": "stopped",
            "tags": [],
            "uuid": "cff0f338-2b84-4846-a028-3ec9e1b86184",
            "vnc_password": "testserver"
        },
        {
            "context": true,
            "cpu": 1000,
            "cpu_model": null,
            "cpus_instead_of_cores": false,
            "drives": [],
            "enable_numa": false,
            "hv_relaxed": false,
            "hv_tsc": false,
            "jobs": [],
            "mem": 536870912,
            "meta": {},
            "name": "test_server_1",
            "nics": [],
            "owner": {
                "resource_uri": "/api/2.0/user/80cb30fb-0ea3-43db-b27b-a125752cc0bf/",
                "uuid": "80cb30fb-0ea3-43db-b27b-a125752cc0bf"
            },
            "requirements": [],
            "resource_uri": "/api/2.0/servers/93a04cd5-84cb-41fc-af17-683e3868ee95/",
            "runtime": null,
            "smp": 1,
            "status": "stopped",
            "tags": [],
            "uuid": "93a04cd5-84cb-41fc-af17-683e3868ee95",
            "vnc_password": "testserver"
        }
    ]
}
`

var serverData = Server{
	ServerRecord: ServerRecord{
		Resource{
			"/api/2.0/servers/472835d5-2bbb-4d87-9d08-7364bc373691/",
			"472835d5-2bbb-4d87-9d08-7364bc373691"},
		"trusty-server-cloudimg-amd64",
		"starting",
	},
	CPU: 2000,
	Mem: 2147483648,
	Meta: map[string]string{"description": "trusty-server-cloudimg-amd64",
		"ssh_public_key": "ssh-rsa AAAAB3NzaC1yc2EAAAADAQABAAABAQDDiwTGBsmFKBYHcKaVy5IgsYBR4XVYLS6KP/NKClE7gONlIGURE3+/45BX8TfHJHM5WTN8NBqJejKDHqwfyueR1f2VGoPkJxODGt/X/ZDNftLZLYwPd2DfDBs27ahOadZCk4Cl5l7mU0aoE74UnIcQoNPl6w7axkIFTIXr8+0HMk8DFB0iviBSJK118p1RGwhsoA1Hudn1CsgqARGPmNn6mxwvmQfQY7hZxZoOH9WMcvkNZ7rAFrwS/BuvEpEXkoC95K/JDPvmQVVJk7we+WeHfTYSmApkDFcSaypyjL2HOV8pvE+VntcIIhZccHiOubyjsBAx5aoTI+ueCsoz5AL1 maxim.perenesenko@altoros.com"},
	NICs: []NIC{
		NIC{
			IPv4: &IPv4{
				Conf: "static",
				IP:   Resource{"/api/2.0/ips/31.171.246.37/", "31.171.246.37"},
			},
			Model: "virtio",
			MAC:   "22:40:85:4f:d3:ce",
		},
		NIC{
			Model: "virtio",
			MAC:   "22:aa:fe:07:48:3b",
			VLAN: &Resource{
				"/api/2.0/vlans/5bc05e7e-6555-4f40-add8-3b8e91447702/",
				"5bc05e7e-6555-4f40-add8-3b8e91447702",
			},
		},
	},
	Drives: []ServerDrive{
		ServerDrive{
			BootOrder: 1,
			Channel:   "0:0",
			Device:    "virtio",
			Drive: Resource{
				URI:  "/api/2.0/drives/ddce5beb-6cfe-4a80-81bd-3ae5f71e0c00/",
				UUID: "ddce5beb-6cfe-4a80-81bd-3ae5f71e0c00",
			},
		},
	},
	VNCPassword: "Pim3UkEc",
}

const jsonServerData = `{
    "context": true,
    "cpu": 2000,
    "cpu_model": null,
    "cpus_instead_of_cores": false,
    "drives": [
        {
            "boot_order": 1,
            "dev_channel": "0:0",
            "device": "virtio",
            "drive": {
                "resource_uri": "/api/2.0/drives/ddce5beb-6cfe-4a80-81bd-3ae5f71e0c00/",
                "uuid": "ddce5beb-6cfe-4a80-81bd-3ae5f71e0c00"
            },
            "runtime": null
        }
    ],
    "enable_numa": false,
    "hv_relaxed": false,
    "hv_tsc": false,
    "jobs": [],
    "mem": 2147483648,
    "meta": {
        "description": "trusty-server-cloudimg-amd64",
        "ssh_public_key": "ssh-rsa AAAAB3NzaC1yc2EAAAADAQABAAABAQDDiwTGBsmFKBYHcKaVy5IgsYBR4XVYLS6KP/NKClE7gONlIGURE3+/45BX8TfHJHM5WTN8NBqJejKDHqwfyueR1f2VGoPkJxODGt/X/ZDNftLZLYwPd2DfDBs27ahOadZCk4Cl5l7mU0aoE74UnIcQoNPl6w7axkIFTIXr8+0HMk8DFB0iviBSJK118p1RGwhsoA1Hudn1CsgqARGPmNn6mxwvmQfQY7hZxZoOH9WMcvkNZ7rAFrwS/BuvEpEXkoC95K/JDPvmQVVJk7we+WeHfTYSmApkDFcSaypyjL2HOV8pvE+VntcIIhZccHiOubyjsBAx5aoTI+ueCsoz5AL1 maxim.perenesenko@altoros.com"
    },
    "name": "trusty-server-cloudimg-amd64",
    "nics": [
        {
            "boot_order": null,
            "firewall_policy": null,
            "ip_v4_conf": {
                "conf": "static",
                "ip": {
                    "resource_uri": "/api/2.0/ips/31.171.246.37/",
                    "uuid": "31.171.246.37"
                }
            },
            "ip_v6_conf": null,
            "mac": "22:40:85:4f:d3:ce",
            "model": "virtio",
            "runtime": null,
            "vlan": null
        },
        {
            "boot_order": null,
            "firewall_policy": null,
            "ip_v4_conf": null,
            "ip_v6_conf": null,
            "mac": "22:aa:fe:07:48:3b",
            "model": "virtio",
            "runtime": null,
            "vlan": {
                "resource_uri": "/api/2.0/vlans/5bc05e7e-6555-4f40-add8-3b8e91447702/",
                "uuid": "5bc05e7e-6555-4f40-add8-3b8e91447702"
            }
        }
    ],
    "owner": {
        "resource_uri": "/api/2.0/user/c25eb0ed-161f-44f4-ac1d-d584ce3a5312/",
        "uuid": "c25eb0ed-161f-44f4-ac1d-d584ce3a5312"
    },
    "requirements": [],
    "resource_uri": "/api/2.0/servers/472835d5-2bbb-4d87-9d08-7364bc373691/",
    "runtime": null,
    "smp": 1,
    "status": "starting",
    "tags": [],
    "uuid": "472835d5-2bbb-4d87-9d08-7364bc373691",
    "vnc_password": "Pim3UkEc"
}
`

/*
var jsonTestServerCreateData = Server{
	ServerRecord: ServerRecord{
		Resource{
			"/api/2.0/servers/472835d5-2bbb-4d87-9d08-7364bc373691/",
			"472835d5-2bbb-4d87-9d08-7364bc373691"},
		"test",
		"starting",
	},
	CPU: 2000,
	Mem: 2147483648,
	Meta: map[string]string{"description": "trusty-server-cloudimg-amd64",
		"ssh_public_key": "ssh-rsa AAAAB3NzaC1yc2EAAAADAQABAAABAQDDiwTGBsmFKBYHcKaVy5IgsYBR4XVYLS6KP/NKClE7gONlIGURE3+/45BX8TfHJHM5WTN8NBqJejKDHqwfyueR1f2VGoPkJxODGt/X/ZDNftLZLYwPd2DfDBs27ahOadZCk4Cl5l7mU0aoE74UnIcQoNPl6w7axkIFTIXr8+0HMk8DFB0iviBSJK118p1RGwhsoA1Hudn1CsgqARGPmNn6mxwvmQfQY7hZxZoOH9WMcvkNZ7rAFrwS/BuvEpEXkoC95K/JDPvmQVVJk7we+WeHfTYSmApkDFcSaypyjL2HOV8pvE+VntcIIhZccHiOubyjsBAx5aoTI+ueCsoz5AL1 maxim.perenesenko@altoros.com"},
	NICs: []NIC{
		NIC{
			IPv4:  &IPv4{Conf: "dhcp"},
			Model: "virtio",
		},
		NIC{
			IPv4:  &IPv4{Conf: "manual"},
			Model: "virtio",
		},
		NIC{
			IPv4: &IPv4{
				Conf: "static",
				IP:   Resource{"/api/2.0/ips/31.171.246.37/", "31.171.246.37"},
			},
			Model: "virtio",
			MAC:   "22:40:85:4f:d3:ce",
		},
		NIC{
			Model: "virtio",
			MAC:   "22:aa:fe:07:48:3b",
			VLAN: &Resource{
				"/api/2.0/vlans/5bc05e7e-6555-4f40-add8-3b8e91447702/",
				"5bc05e7e-6555-4f40-add8-3b8e91447702",
			},
		},
	},
	Drives: []ServerDrive{
		ServerDrive{
			BootOrder: 1,
			Channel:   "0:0",
			Device:    "virtio",
			Drive: Resource{
				URI:  "/api/2.0/drives/ddce5beb-6cfe-4a80-81bd-3ae5f71e0c00/",
				UUID: "ddce5beb-6cfe-4a80-81bd-3ae5f71e0c00",
			},
		},
	},
	VNCPassword: "Pim3UkEc",
}

const jsonTestServerCreateRequestj = `{"cpu":2000,"mem":2147483648,"name":"test","nics":[{"ip_v4_conf":{"conf":"dhcp"},"model":"virtio"},{"ip_v4_conf":{"conf":"manual"},"model":"virtio"},{"ip_v4_conf":{"conf":"static","ip":"ipaddr"},"model":"virtio"},{"model":"virtio","vlan":"vlanid"}],"vnc_password":"testserver"}`
*/
