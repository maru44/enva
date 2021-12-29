package controllers

import "errors"

const (
	QueryParamsOrgID   = "orgId"
	QueryParamsOrgSlug = "orgSlug"

	QueryParamsProjectID   = "projectId"
	QueryParamsProjectSlug = "projectSlug"

	QueryParamsID   = "id"
	QueryParamsSlug = "slug"

	QueryParamsKvID  = "kvId"
	QueryParamsKvKey = "key"
)

var (
	ErrorNoOrgIdParams = errors.New("No org id")
)
