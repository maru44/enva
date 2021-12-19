package controllers

import "errors"

const (
	QueryParamsOrgID = "orgId"

	QueryParamsProjectID   = "projectId"
	QueryParamsProjectSlug = "projectSlug"

	QueryParamsSlug = "slug"

	QueryParamsKvID  = "kvId"
	QueryParamsKvKey = "key"
)

var (
	ErrorNoOrgIdParams = errors.New("No org id")
)
