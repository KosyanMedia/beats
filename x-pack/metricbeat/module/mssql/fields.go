// Copyright Elasticsearch B.V. and/or licensed to Elasticsearch B.V. under one
// or more contributor license agreements. Licensed under the Elastic License;
// you may not use this file except in compliance with the Elastic License.

// Code generated by beats/dev-tools/cmd/asset/asset.go - DO NOT EDIT.

package mssql

import (
	"github.com/elastic/beats/v7/libbeat/asset"
)

func init() {
	if err := asset.SetFields("metricbeat", "mssql", asset.ModuleFieldsPri, AssetMssql); err != nil {
		panic(err)
	}
}

// AssetMssql returns asset data.
// This is the base64 encoded gzipped contents of module/mssql.
func AssetMssql() string {
	return "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"
}
