package controllers

import "errors"

const (
	QueryParamsOrgID     = "orgId"
	QueryParamsProjectID = "projectId"
	QueryParamsSlug      = "slug"

	QueryParamsKvID  = "kvId"
	QueryParamsKvKey = "kvKey"
)

var (
	ErrorNoOrgIdParams = errors.New("No org id")
)
