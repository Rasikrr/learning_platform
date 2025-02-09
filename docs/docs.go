// Package docs Code generated by swaggo/swag. DO NOT EDIT
package docs

import "github.com/swaggo/swag"

const docTemplate = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{escape .Description}}",
        "title": "{{.Title}}",
        "termsOfService": "http://swagger.io/terms/",
        "contact": {
            "name": "Pick me team",
            "url": "https://github.com/Rasikrr/learning_platform",
            "email": "leaning_platform@gmail.com"
        },
        "license": {
            "name": "Apache 2.0",
            "url": "http://www.apache.org/licenses/LICENSE-2.0.html"
        },
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/api/v1/auth/confirm": {
            "post": {
                "description": "Confirm user registration using confirmation code",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "auth"
                ],
                "summary": "Confirm register",
                "parameters": [
                    {
                        "description": "user email and confirmation code",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/auth.ConfirmRegisterRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Success",
                        "schema": {
                            "$ref": "#/definitions/auth.Auth"
                        }
                    }
                }
            }
        },
        "/api/v1/auth/confirm_reset_password": {
            "post": {
                "description": "Confirm user password reset using confirmation code",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "auth"
                ],
                "summary": "Confirm password reset",
                "parameters": [
                    {
                        "description": "user email and confirmation code",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/auth.ConfirmResetPasswordRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Success",
                        "schema": {
                            "$ref": "#/definitions/api.EmptySuccessResponse"
                        }
                    }
                }
            }
        },
        "/api/v1/auth/login": {
            "post": {
                "description": "Login user with email and password",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "auth"
                ],
                "summary": "Login user",
                "parameters": [
                    {
                        "description": "user credentials",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/auth.LoginRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Success",
                        "schema": {
                            "$ref": "#/definitions/auth.Auth"
                        }
                    }
                }
            }
        },
        "/api/v1/auth/refresh": {
            "post": {
                "description": "Refresh user tokens",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "auth"
                ],
                "summary": "Refresh token",
                "parameters": [
                    {
                        "description": "user's refresh token",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/auth.RefreshRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Success",
                        "schema": {
                            "$ref": "#/definitions/auth.Auth"
                        }
                    }
                }
            }
        },
        "/api/v1/auth/register": {
            "post": {
                "description": "Register user with email and password and confirm password. Then send confirmation code to user email",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "auth"
                ],
                "summary": "Register user",
                "parameters": [
                    {
                        "description": "user credentials",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/auth.RegisterRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Success",
                        "schema": {
                            "$ref": "#/definitions/api.EmptySuccessResponse"
                        }
                    }
                }
            }
        },
        "/api/v1/auth/reset_password": {
            "post": {
                "description": "Reset user password with email and password and confirm password. Then send confirmation code to user email",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "auth"
                ],
                "summary": "Reset password",
                "parameters": [
                    {
                        "description": "user credentials",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/auth.ResetPasswordRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Success",
                        "schema": {
                            "$ref": "#/definitions/api.EmptySuccessResponse"
                        }
                    }
                }
            }
        },
        "/api/v1/courses": {
            "post": {
                "description": "Get courses for catalog",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "courses"
                ],
                "summary": "Get courses",
                "parameters": [
                    {
                        "description": "request",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/queries.getCoursesListRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Success",
                        "schema": {
                            "$ref": "#/definitions/queries.getCoursesListResponse"
                        }
                    }
                }
            }
        },
        "/api/v1/courses/categories": {
            "get": {
                "description": "Get courses categories",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "courses"
                ],
                "summary": "Get courses categories",
                "responses": {
                    "200": {
                        "description": "Success",
                        "schema": {
                            "$ref": "#/definitions/queries.getCategoriesListResponse"
                        }
                    }
                }
            }
        },
        "/api/v1/courses/enroll": {
            "post": {
                "security": [
                    {
                        "BearerAuth": []
                    }
                ],
                "description": "enroll to course by course id",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "courses"
                ],
                "summary": "enroll to course",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Authorization token",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
                    },
                    {
                        "description": "request",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/enrollments.enrollToCourseRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Success",
                        "schema": {
                            "$ref": "#/definitions/api.EmptySuccessResponse"
                        }
                    }
                }
            }
        },
        "/api/v1/courses/{course_id}/topic/{topic_id}/content": {
            "get": {
                "security": [
                    {
                        "BearerAuth": []
                    }
                ],
                "description": "Get course topic content by topic id",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "courses"
                ],
                "summary": "Get course topic content",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Authorization token",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "course id",
                        "name": "course_id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "topic id",
                        "name": "topic_id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Success",
                        "schema": {
                            "$ref": "#/definitions/entity.TopicContent"
                        }
                    }
                }
            }
        },
        "/api/v1/courses/{course_id}/topic/{topic_id}/quiz/reset": {
            "post": {
                "security": [
                    {
                        "BearerAuth": []
                    }
                ],
                "description": "submit quiz",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "courses"
                ],
                "summary": "submit quiz",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Authorization token",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "course id",
                        "name": "course_id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "topic id",
                        "name": "topic_id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Success",
                        "schema": {
                            "$ref": "#/definitions/api.EmptySuccessResponse"
                        }
                    }
                }
            }
        },
        "/api/v1/courses/{course_id}/topic/{topic_id}/quiz/submit": {
            "post": {
                "security": [
                    {
                        "BearerAuth": []
                    }
                ],
                "description": "submit quiz",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "courses"
                ],
                "summary": "submit quiz",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Authorization token",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "course id",
                        "name": "course_id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "topic id",
                        "name": "topic_id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "request",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/commands.submitQuizRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Success",
                        "schema": {
                            "$ref": "#/definitions/api.EmptySuccessResponse"
                        }
                    }
                }
            }
        },
        "/api/v1/courses/{course_id}/topic/{topic_id}/quizzes": {
            "get": {
                "security": [
                    {
                        "BearerAuth": []
                    }
                ],
                "description": "get quizzes list by topic id",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "courses"
                ],
                "summary": "get quizzes by topic id",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Authorization token",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "course id",
                        "name": "course_id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "topic id",
                        "name": "topic_id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Success",
                        "schema": {
                            "$ref": "#/definitions/queries.getTopicQuizzesResponse"
                        }
                    }
                }
            }
        },
        "/api/v1/courses/{course_id}/topic/{topic_id}/tasks/{order}": {
            "get": {
                "security": [
                    {
                        "BearerAuth": []
                    }
                ],
                "description": "Get course topic tasks by topic id and. Also you must pass task order",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "courses"
                ],
                "summary": "Get course topic tasks",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Authorization token",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "course id",
                        "name": "course_id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "topic id",
                        "name": "topic_id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "number of task",
                        "name": "order",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Success",
                        "schema": {
                            "$ref": "#/definitions/entity.PracticalTask"
                        }
                    }
                }
            }
        },
        "/api/v1/courses/{id}": {
            "get": {
                "description": "Get detailed course info with topics by id",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "courses"
                ],
                "summary": "Get course by id",
                "parameters": [
                    {
                        "type": "string",
                        "description": "course id",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Success",
                        "schema": {
                            "$ref": "#/definitions/queries.getCourseDetailedResponse"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "api.EmptySuccessResponse": {
            "type": "object",
            "properties": {
                "status": {
                    "type": "string"
                }
            }
        },
        "auth.Auth": {
            "type": "object",
            "properties": {
                "access_token": {
                    "type": "string"
                },
                "refresh_token": {
                    "type": "string"
                }
            }
        },
        "auth.ConfirmRegisterRequest": {
            "type": "object",
            "properties": {
                "code": {
                    "type": "string"
                },
                "email": {
                    "type": "string"
                }
            }
        },
        "auth.ConfirmResetPasswordRequest": {
            "type": "object",
            "properties": {
                "code": {
                    "type": "string"
                },
                "email": {
                    "type": "string"
                }
            }
        },
        "auth.LoginRequest": {
            "type": "object",
            "properties": {
                "email": {
                    "type": "string"
                },
                "password": {
                    "type": "string"
                }
            }
        },
        "auth.RefreshRequest": {
            "type": "object",
            "properties": {
                "refresh_token": {
                    "type": "string"
                }
            }
        },
        "auth.RegisterRequest": {
            "type": "object",
            "properties": {
                "email": {
                    "type": "string"
                },
                "password": {
                    "type": "string"
                },
                "password_confirm": {
                    "type": "string"
                }
            }
        },
        "auth.ResetPasswordRequest": {
            "type": "object",
            "properties": {
                "email": {
                    "type": "string"
                },
                "password": {
                    "type": "string"
                },
                "password_confirm": {
                    "type": "string"
                }
            }
        },
        "commands.answerQuiz": {
            "type": "object",
            "properties": {
                "answer": {
                    "type": "array",
                    "items": {
                        "type": "boolean"
                    }
                },
                "question_id": {
                    "type": "string"
                }
            }
        },
        "commands.submitQuizRequest": {
            "type": "object",
            "properties": {
                "answers": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/commands.answerQuiz"
                    }
                }
            }
        },
        "enrollments.enrollToCourseRequest": {
            "type": "object",
            "properties": {
                "course_id": {
                    "type": "string"
                }
            }
        },
        "entity.PracticalTask": {
            "type": "object",
            "properties": {
                "created_at": {
                    "type": "string"
                },
                "description": {
                    "type": "string"
                },
                "difficulty_level": {
                    "$ref": "#/definitions/enum.Difficulty"
                },
                "expected_output": {
                    "type": "string"
                },
                "id": {
                    "type": "string"
                },
                "order_number": {
                    "type": "integer"
                },
                "starter_code": {
                    "type": "string"
                },
                "topic_id": {
                    "type": "string"
                },
                "updated_at": {
                    "type": "string"
                }
            }
        },
        "entity.TopicContent": {
            "type": "object",
            "properties": {
                "additional_resources": {
                    "type": "array",
                    "items": {
                        "type": "string"
                    }
                },
                "content": {
                    "type": "string"
                },
                "created_at": {
                    "type": "string"
                },
                "id": {
                    "type": "string"
                },
                "image_urls": {
                    "type": "array",
                    "items": {
                        "type": "string"
                    }
                },
                "topic_id": {
                    "type": "string"
                },
                "updated_at": {
                    "type": "string"
                },
                "video_urls": {
                    "type": "array",
                    "items": {
                        "type": "string"
                    }
                }
            }
        },
        "enum.Difficulty": {
            "type": "integer",
            "enum": [
                0,
                1,
                2
            ],
            "x-enum-varnames": [
                "DifficultyEasy",
                "DifficultyMedium",
                "DifficultyHard"
            ]
        },
        "queries.category": {
            "type": "object",
            "properties": {
                "created_at": {
                    "type": "string"
                },
                "created_by": {
                    "type": "string"
                },
                "id": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                },
                "updated_at": {
                    "type": "string"
                }
            }
        },
        "queries.getCategoriesListResponse": {
            "type": "object",
            "properties": {
                "categories": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/queries.category"
                    }
                }
            }
        },
        "queries.getCourseDetailedResponse": {
            "type": "object",
            "properties": {
                "course": {
                    "$ref": "#/definitions/queries.getCourseResponse"
                },
                "topics": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/queries.topic"
                    }
                }
            }
        },
        "queries.getCourseResponse": {
            "type": "object",
            "properties": {
                "category": {
                    "$ref": "#/definitions/queries.category"
                },
                "created_at": {
                    "type": "string"
                },
                "created_by": {
                    "type": "string"
                },
                "description": {
                    "type": "string"
                },
                "id": {
                    "type": "string"
                },
                "image_url": {
                    "type": "string"
                },
                "title": {
                    "type": "string"
                },
                "updated_at": {
                    "type": "string"
                }
            }
        },
        "queries.getCoursesListRequest": {
            "type": "object",
            "properties": {
                "categories_ids": {
                    "type": "array",
                    "items": {
                        "type": "string"
                    }
                },
                "limit": {
                    "type": "integer"
                },
                "offset": {
                    "type": "integer"
                }
            }
        },
        "queries.getCoursesListResponse": {
            "type": "object",
            "properties": {
                "courses": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/queries.getCourseResponse"
                    }
                }
            }
        },
        "queries.getTopicQuizzesResponse": {
            "type": "object",
            "properties": {
                "quizzes": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/queries.quizz"
                    }
                }
            }
        },
        "queries.quizz": {
            "type": "object",
            "properties": {
                "created_at": {
                    "type": "string"
                },
                "id": {
                    "type": "string"
                },
                "multiple_choice": {
                    "type": "boolean"
                },
                "options": {
                    "type": "array",
                    "items": {
                        "type": "string"
                    }
                },
                "question": {
                    "type": "string"
                },
                "topic_id": {
                    "type": "string"
                },
                "updated_at": {
                    "type": "string"
                }
            }
        },
        "queries.topic": {
            "type": "object",
            "properties": {
                "course_id": {
                    "type": "string"
                },
                "created_at": {
                    "type": "string"
                },
                "description": {
                    "type": "string"
                },
                "id": {
                    "type": "string"
                },
                "order_number": {
                    "type": "integer"
                },
                "title": {
                    "type": "string"
                },
                "updated_at": {
                    "type": "string"
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "localhost:8081",
	BasePath:         "/",
	Schemes:          []string{},
	Title:            "Learning Platform API",
	Description:      "This is docs for Learning Platform API",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
