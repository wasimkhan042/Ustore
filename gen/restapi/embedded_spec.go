// Code generated by go-swagger; DO NOT EDIT.

package restapi

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
)

var (
	// SwaggerJSON embedded version of the swagger document used at generation time
	SwaggerJSON json.RawMessage
	// FlatSwaggerJSON embedded flattened version of the swagger document used at generation time
	FlatSwaggerJSON json.RawMessage
)

func init() {
	SwaggerJSON = json.RawMessage([]byte(`{
  "consumes": [
    "application/json"
  ],
  "produces": [
    "application/json"
  ],
  "schemes": [
    "http"
  ],
  "swagger": "2.0",
  "info": {
    "title": "ustore",
    "version": "1.0.0"
  },
  "host": "localhost:8080",
  "basePath": "/v1",
  "paths": {
    "/signup": {
      "post": {
        "description": "To register a new user",
        "tags": [
          "signup"
        ],
        "operationId": "signup",
        "parameters": [
          {
            "description": "signup model",
            "name": "signup",
            "in": "body",
            "required": true,
            "schema": {
              "$ref": "#/definitions/SignUp"
            }
          }
        ],
        "responses": {
          "200": {
            "description": "Successful registeration",
            "schema": {
              "$ref": "#/definitions/SignUpResponse"
            }
          },
          "400": {
            "description": "Bad Request"
          },
          "404": {
            "description": "User not found",
            "schema": {
              "type": "string"
            }
          },
          "500": {
            "description": "Server error",
            "schema": {
              "type": "string"
            }
          }
        }
      }
    }
  },
  "definitions": {
    "SignUp": {
      "type": "object",
      "required": [
        "email",
        "first_name",
        "password",
        "username"
      ],
      "properties": {
        "email": {
          "type": "string"
        },
        "first_name": {
          "type": "string"
        },
        "last_name": {
          "type": "string"
        },
        "middle_name": {
          "type": "string"
        },
        "password": {
          "type": "string"
        },
        "profile_image": {
          "type": "string"
        },
        "username": {
          "type": "string"
        }
      }
    },
    "SignUpResponse": {
      "type": "object",
      "properties": {
        "message": {
          "type": "string"
        },
        "success": {
          "type": "boolean"
        }
      }
    }
  }
}`))
	FlatSwaggerJSON = json.RawMessage([]byte(`{
  "consumes": [
    "application/json"
  ],
  "produces": [
    "application/json"
  ],
  "schemes": [
    "http"
  ],
  "swagger": "2.0",
  "info": {
    "title": "ustore",
    "version": "1.0.0"
  },
  "host": "localhost:8080",
  "basePath": "/v1",
  "paths": {
    "/signup": {
      "post": {
        "description": "To register a new user",
        "tags": [
          "signup"
        ],
        "operationId": "signup",
        "parameters": [
          {
            "description": "signup model",
            "name": "signup",
            "in": "body",
            "required": true,
            "schema": {
              "$ref": "#/definitions/SignUp"
            }
          }
        ],
        "responses": {
          "200": {
            "description": "Successful registeration",
            "schema": {
              "$ref": "#/definitions/SignUpResponse"
            }
          },
          "400": {
            "description": "Bad Request"
          },
          "404": {
            "description": "User not found",
            "schema": {
              "type": "string"
            }
          },
          "500": {
            "description": "Server error",
            "schema": {
              "type": "string"
            }
          }
        }
      }
    }
  },
  "definitions": {
    "SignUp": {
      "type": "object",
      "required": [
        "email",
        "first_name",
        "password",
        "username"
      ],
      "properties": {
        "email": {
          "type": "string"
        },
        "first_name": {
          "type": "string"
        },
        "last_name": {
          "type": "string"
        },
        "middle_name": {
          "type": "string"
        },
        "password": {
          "type": "string"
        },
        "profile_image": {
          "type": "string"
        },
        "username": {
          "type": "string"
        }
      }
    },
    "SignUpResponse": {
      "type": "object",
      "properties": {
        "message": {
          "type": "string"
        },
        "success": {
          "type": "boolean"
        }
      }
    }
  }
}`))
}
