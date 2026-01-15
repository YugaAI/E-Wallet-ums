package services

import "ewallet-ums/internal/interfaces"

type HealthCheck struct {
	HealthCheckRepository interfaces.IHealthCheckRepository
}

func (c *HealthCheck) HealtCheckServices() (string, error) {
	return "service healthy", nil
}
