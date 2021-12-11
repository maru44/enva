package controllers

import "errors"

const (
	QueryParamsOrgID     = "orgId"
	QueryParamsProjectID = "projectId"
)

var (
	ErrorNoOrgIdParams = errors.New("No org id")
)
