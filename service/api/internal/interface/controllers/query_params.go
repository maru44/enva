package controllers

import "errors"

const (
	QueryParamsOrgID     = "orgId"
	QueryParamsProjectID = "projectId"
	QueryParamsSlug      = "slug"
)

var (
	ErrorNoOrgIdParams = errors.New("No org id")
)
