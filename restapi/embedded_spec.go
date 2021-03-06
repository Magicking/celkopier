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
  "swagger": "2.0",
  "info": {
    "description": "Use your friendly RasberryPi to scan your file\n",
    "title": "Celkopier",
    "version": "0.1.0"
  },
  "paths": {
    "/listscanned": {
      "get": {
        "description": "List scanned pictures",
        "summary": "List scanned pictures",
        "operationId": "listScanned",
        "responses": {
          "200": {
            "description": "List of scanned pictures",
            "schema": {
              "type": "array",
              "items": {
                "type": "string"
              }
            }
          },
          "default": {
            "description": "Internal error",
            "schema": {
              "$ref": "#/definitions/Error"
            }
          }
        }
      }
    },
    "/scan": {
      "post": {
        "description": "Start scanner",
        "operationId": "scan",
        "responses": {
          "200": {},
          "default": {
            "description": "Internal error",
            "schema": {
              "$ref": "#/definitions/Error"
            }
          }
        }
      }
    },
    "/uploadscanned": {
      "post": {
        "description": "Upload a scanned picture",
        "consumes": [
          "multipart/form-data"
        ],
        "summary": "Upload a scanned picture",
        "operationId": "upload",
        "parameters": [
          {
            "type": "file",
            "description": "Scanned image",
            "name": "scanned",
            "in": "formData",
            "required": true
          }
        ],
        "responses": {
          "200": {},
          "default": {
            "description": "Internal error",
            "schema": {
              "$ref": "#/definitions/Error"
            }
          }
        }
      }
    }
  },
  "definitions": {
    "Error": {
      "type": "object",
      "required": [
        "code",
        "message"
      ],
      "properties": {
        "code": {
          "type": "integer",
          "format": "int32"
        },
        "fields": {
          "type": "string"
        },
        "message": {
          "type": "string"
        }
      }
    }
  },
  "responses": {
    "scanned": {
      "description": "Scanned picture",
      "schema": {
        "type": "string",
        "format": "binary"
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
  "swagger": "2.0",
  "info": {
    "description": "Use your friendly RasberryPi to scan your file\n",
    "title": "Celkopier",
    "version": "0.1.0"
  },
  "paths": {
    "/listscanned": {
      "get": {
        "description": "List scanned pictures",
        "summary": "List scanned pictures",
        "operationId": "listScanned",
        "responses": {
          "200": {
            "description": "List of scanned pictures",
            "schema": {
              "type": "array",
              "items": {
                "type": "string"
              }
            }
          },
          "default": {
            "description": "Internal error",
            "schema": {
              "$ref": "#/definitions/Error"
            }
          }
        }
      }
    },
    "/scan": {
      "post": {
        "description": "Start scanner",
        "operationId": "scan",
        "responses": {
          "200": {},
          "default": {
            "description": "Internal error",
            "schema": {
              "$ref": "#/definitions/Error"
            }
          }
        }
      }
    },
    "/uploadscanned": {
      "post": {
        "description": "Upload a scanned picture",
        "consumes": [
          "multipart/form-data"
        ],
        "summary": "Upload a scanned picture",
        "operationId": "upload",
        "parameters": [
          {
            "type": "file",
            "description": "Scanned image",
            "name": "scanned",
            "in": "formData",
            "required": true
          }
        ],
        "responses": {
          "200": {},
          "default": {
            "description": "Internal error",
            "schema": {
              "$ref": "#/definitions/Error"
            }
          }
        }
      }
    }
  },
  "definitions": {
    "Error": {
      "type": "object",
      "required": [
        "code",
        "message"
      ],
      "properties": {
        "code": {
          "type": "integer",
          "format": "int32"
        },
        "fields": {
          "type": "string"
        },
        "message": {
          "type": "string"
        }
      }
    }
  },
  "responses": {
    "scanned": {
      "description": "Scanned picture",
      "schema": {
        "type": "string",
        "format": "binary"
      }
    }
  }
}`))
}
