/*
 * Simple Inventory API
 *
 * This is a simple API
 *
 * API version: 1.0.0
 * Contact: you@your-company.com
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package models

import (
	"time"
)

type AllStates struct {
	StateCode string `json:"state_code"`

	TotalCases int64 `json:"total_cases"`

	TotalRecovered int64 `json:"total_recovered"`

	TotalDeath int64 `json:"total_death"`

	TotalVaccinated1 int64 `json:"total_vaccinated1"`

	TotalVaccinated2 int64 `json:"total_vaccinated2"`

	TotalTested int64 `json:"total_tested"`

	LastUpdated time.Time `json:"last_updated"`
}
