package services

import "ewallet-framework/internal/interfaces"

type HealthCheck struct {
	HealthCheckRepository interfaces.IHealthCheckRepository
}

func (c *HealthCheck) HealtCheckServices() (string, error) {
	return "service healthy", nil
}
