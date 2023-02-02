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
    "title": "The Neuro-PetStore",
    "version": "1.0.0"
  },
  "host": "the-neuro-pet-store.com",
  "basePath": "/v1",
  "paths": {
    "/buyPet/{id}": {
      "put": {
        "description": "It returns 200 status code when the pet has been unavailable",
        "tags": [
          "pets"
        ],
        "summary": "Buy pet and make unavailable",
        "operationId": "buyPet",
        "parameters": [
          {
            "type": "integer",
            "description": "pet id",
            "name": "id",
            "in": "path",
            "required": true
          }
        ],
        "responses": {
          "200": {
            "description": "200"
          },
          "400": {
            "description": "Bad Request"
          },
          "404": {
            "description": "Pets Not Found"
          },
          "500": {
            "description": "Server Error",
            "schema": {
              "type": "string"
            }
          }
        }
      }
    },
    "/list": {
      "get": {
        "description": "It returns a list of nested objects which contains all pets and their qualities",
        "tags": [
          "pets"
        ],
        "summary": "Get list of currently available pets",
        "operationId": "PetsList",
        "responses": {
          "200": {
            "description": "List of Pets",
            "schema": {
              "$ref": "#/definitions/Pets"
            }
          },
          "400": {
            "description": "Bad Request"
          },
          "404": {
            "description": "Pets Not Found"
          },
          "500": {
            "description": "Server Error",
            "schema": {
              "type": "string"
            }
          }
        }
      }
    }
  },
  "definitions": {
    "Pet": {
      "type": "object",
      "properties": {
        "Available": {
          "type": "boolean",
          "x-omitempty": false
        },
        "Id": {
          "type": "integer"
        },
        "ImageUrl": {
          "type": "string"
        },
        "Name": {
          "type": "string"
        },
        "Qualities": {
          "type": "array",
          "items": {
            "type": "string"
          }
        },
        "Species": {
          "type": "string"
        }
      }
    },
    "Pets": {
      "type": "array",
      "items": {
        "$ref": "#/definitions/Pet"
      }
    }
  },
  "x-ibm-configuration": {
    "cors": {
      "enabled": false
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
    "title": "The Neuro-PetStore",
    "version": "1.0.0"
  },
  "host": "the-neuro-pet-store.com",
  "basePath": "/v1",
  "paths": {
    "/buyPet/{id}": {
      "put": {
        "description": "It returns 200 status code when the pet has been unavailable",
        "tags": [
          "pets"
        ],
        "summary": "Buy pet and make unavailable",
        "operationId": "buyPet",
        "parameters": [
          {
            "type": "integer",
            "description": "pet id",
            "name": "id",
            "in": "path",
            "required": true
          }
        ],
        "responses": {
          "200": {
            "description": "200"
          },
          "400": {
            "description": "Bad Request"
          },
          "404": {
            "description": "Pets Not Found"
          },
          "500": {
            "description": "Server Error",
            "schema": {
              "type": "string"
            }
          }
        }
      }
    },
    "/list": {
      "get": {
        "description": "It returns a list of nested objects which contains all pets and their qualities",
        "tags": [
          "pets"
        ],
        "summary": "Get list of currently available pets",
        "operationId": "PetsList",
        "responses": {
          "200": {
            "description": "List of Pets",
            "schema": {
              "$ref": "#/definitions/Pets"
            }
          },
          "400": {
            "description": "Bad Request"
          },
          "404": {
            "description": "Pets Not Found"
          },
          "500": {
            "description": "Server Error",
            "schema": {
              "type": "string"
            }
          }
        }
      }
    }
  },
  "definitions": {
    "Pet": {
      "type": "object",
      "properties": {
        "Available": {
          "type": "boolean",
          "x-omitempty": false
        },
        "Id": {
          "type": "integer"
        },
        "ImageUrl": {
          "type": "string"
        },
        "Name": {
          "type": "string"
        },
        "Qualities": {
          "type": "array",
          "items": {
            "type": "string"
          }
        },
        "Species": {
          "type": "string"
        }
      }
    },
    "Pets": {
      "type": "array",
      "items": {
        "$ref": "#/definitions/Pet"
      }
    }
  },
  "x-ibm-configuration": {
    "cors": {
      "enabled": false
    }
  }
}`))
}
