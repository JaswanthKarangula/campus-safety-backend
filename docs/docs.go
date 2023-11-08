// Code generated by swaggo/swag. DO NOT EDIT
package docs

import "github.com/swaggo/swag"

const docTemplate = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{escape .Description}}",
        "title": "{{.Title}}",
        "contact": {},
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/admin/createAdmin": {
            "post": {
                "description": "adds a New admin in Db.",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "admin"
                ],
                "summary": "creates a new admin",
                "parameters": [
                    {
                        "description": "adds a New Drone For Customer  in Db",
                        "name": "device",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/api.CreateNewAdminRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/customer/addNewDrone": {
            "post": {
                "description": "adds a New Drone For Customer  in Db.",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "customer"
                ],
                "summary": "addNewDroneForCustomer",
                "parameters": [
                    {
                        "description": "adds a New Drone For Customer  in Db",
                        "name": "device",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/api.AddNewDroneRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/customer/createCustomer": {
            "post": {
                "description": "adds a New Drone For Customer  in Db.",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "customer"
                ],
                "summary": "addNewDroneForCustomer",
                "parameters": [
                    {
                        "description": "adds a New Drone For Customer  in Db",
                        "name": "device",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/api.AddNewDroneRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/customer/getAllActiveOfficers": {
            "get": {
                "description": "returns all active security officers.",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "customer"
                ],
                "summary": "getActiveOfficers",
                "parameters": [
                    {
                        "description": "returns active security officers For Customer  in Db",
                        "name": "device",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/api.GetActiveOfficersRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/db.User"
                            }
                        }
                    }
                }
            }
        },
        "/customer/getAllDrones": {
            "get": {
                "description": "eturns all drones  of a Customer.",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "customer"
                ],
                "summary": "getAllDrones",
                "parameters": [
                    {
                        "description": "returns all drones  of a Customer  in Db",
                        "name": "device",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/api.GetAllDronesRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/db.Drone"
                            }
                        }
                    }
                }
            }
        },
        "/customer/getAllIssuesByCustomer": {
            "get": {
                "description": "returns all issues raised by Customer.",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "customer"
                ],
                "summary": "getAllIssuesByCustomer",
                "parameters": [
                    {
                        "description": "returns all issues raised by customers  in Db",
                        "name": "device",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/api.GetAllIssuesRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/db.Issue"
                            }
                        }
                    }
                }
            }
        },
        "/customer/getAllOfficerIssues": {
            "get": {
                "description": "returns all issues raised by officers  of a Customer.",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "customer"
                ],
                "summary": "getAllOfficerIssues",
                "parameters": [
                    {
                        "description": "returns all issues raised by officers  of a Customer in Db",
                        "name": "device",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/api.GetAllIssuesRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/db.Issue"
                            }
                        }
                    }
                }
            }
        },
        "/customer/raiseCustomerIssue": {
            "post": {
                "description": "raises a new issue by customer  in Db.",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "customer"
                ],
                "summary": "createNewCustomerIssue",
                "parameters": [
                    {
                        "description": "raises a new issue by customer in Db",
                        "name": "device",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/api.CreateNewCustomerIssueRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/customer/stopDroneStream": {
            "put": {
                "description": "update stream.",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "customer"
                ],
                "summary": "update stream",
                "parameters": [
                    {
                        "description": "stop the stream",
                        "name": "device",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/api.UpdateStreamRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/issue/updateIssue": {
            "put": {
                "description": "updates an issue by officer  in Db.",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "issue"
                ],
                "summary": "updateIssue",
                "parameters": [
                    {
                        "description": "updates an  issue  in Db",
                        "name": "device",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/api.UpdateIssueRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/officer/createOfficer": {
            "post": {
                "description": "adds a New Security officer For Customer  in Db.",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "officer"
                ],
                "summary": "createNewSecurityOfficer",
                "parameters": [
                    {
                        "description": "adds a New officer For Customer  in Db",
                        "name": "device",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/api.CreateNewSecurityOfficerRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/officer/raiseOfficerIssue": {
            "post": {
                "description": "raises a new issue by officer  in Db.",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "customer"
                ],
                "summary": "createNewOfficerIssue",
                "parameters": [
                    {
                        "description": "raises a new issue by officer in Db",
                        "name": "device",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/api.CreateNewOfficerIssueRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "api.AddNewDroneRequest": {
            "type": "object",
            "required": [
                "customer_id",
                "model",
                "status"
            ],
            "properties": {
                "customer_id": {
                    "type": "integer"
                },
                "model": {
                    "type": "string"
                },
                "status": {
                    "type": "string"
                }
            }
        },
        "api.CreateNewAdminRequest": {
            "type": "object",
            "required": [
                "email",
                "password",
                "username"
            ],
            "properties": {
                "email": {
                    "type": "string"
                },
                "password": {
                    "type": "string",
                    "minLength": 6
                },
                "username": {
                    "type": "string"
                }
            }
        },
        "api.CreateNewCustomerIssueRequest": {
            "type": "object",
            "required": [
                "customer_id",
                "description"
            ],
            "properties": {
                "customer_id": {
                    "type": "integer"
                },
                "description": {
                    "type": "string"
                }
            }
        },
        "api.CreateNewOfficerIssueRequest": {
            "type": "object",
            "required": [
                "description",
                "officer_id"
            ],
            "properties": {
                "description": {
                    "type": "string"
                },
                "officer_id": {
                    "type": "integer"
                }
            }
        },
        "api.CreateNewSecurityOfficerRequest": {
            "type": "object",
            "required": [
                "customerid",
                "email",
                "password",
                "username"
            ],
            "properties": {
                "customerid": {
                    "type": "integer"
                },
                "email": {
                    "type": "string",
                    "minLength": 6
                },
                "password": {
                    "type": "string"
                },
                "username": {
                    "type": "string"
                }
            }
        },
        "api.GetActiveOfficersRequest": {
            "type": "object",
            "required": [
                "customerid"
            ],
            "properties": {
                "customerid": {
                    "type": "integer"
                }
            }
        },
        "api.GetAllDronesRequest": {
            "type": "object",
            "required": [
                "customerid"
            ],
            "properties": {
                "customerid": {
                    "type": "integer"
                }
            }
        },
        "api.GetAllIssuesRequest": {
            "type": "object",
            "required": [
                "customerid",
                "limit",
                "offset"
            ],
            "properties": {
                "customerid": {
                    "type": "integer"
                },
                "limit": {
                    "type": "integer"
                },
                "offset": {
                    "type": "integer"
                }
            }
        },
        "api.UpdateIssueRequest": {
            "type": "object",
            "required": [
                "comments",
                "issue_id",
                "status"
            ],
            "properties": {
                "comments": {
                    "type": "string"
                },
                "issue_id": {
                    "type": "integer"
                },
                "status": {
                    "type": "string"
                }
            }
        },
        "api.UpdateStreamRequest": {
            "type": "object",
            "required": [
                "stream_id"
            ],
            "properties": {
                "stream_id": {
                    "type": "integer"
                }
            }
        },
        "db.Drone": {
            "type": "object",
            "properties": {
                "customer_id": {
                    "type": "integer"
                },
                "drone_id": {
                    "type": "integer"
                },
                "last_active": {
                    "$ref": "#/definitions/sql.NullTime"
                },
                "model": {
                    "type": "string"
                },
                "status": {
                    "description": "active,inactive,streaming",
                    "type": "string"
                }
            }
        },
        "db.Issue": {
            "type": "object",
            "properties": {
                "comments": {
                    "type": "string"
                },
                "created_at": {
                    "type": "string"
                },
                "description": {
                    "type": "string"
                },
                "issue_id": {
                    "type": "integer"
                },
                "status": {
                    "type": "string"
                }
            }
        },
        "db.User": {
            "type": "object",
            "properties": {
                "created_at": {
                    "type": "string"
                },
                "email": {
                    "type": "string"
                },
                "hashedpassword": {
                    "type": "string"
                },
                "role": {
                    "description": "Admin,Customer,Security Officer",
                    "type": "string"
                },
                "user_id": {
                    "type": "integer"
                },
                "username": {
                    "type": "string"
                }
            }
        },
        "sql.NullTime": {
            "type": "object",
            "properties": {
                "time": {
                    "type": "string"
                },
                "valid": {
                    "description": "Valid is true if Time is not NULL",
                    "type": "boolean"
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "0.0.0.0:8080",
	BasePath:         "",
	Schemes:          []string{},
	Title:            "Drone Saefty Backend",
	Description:      "Drone Saefty Backend",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
