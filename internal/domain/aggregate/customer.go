package aggregate

import "decard/internal/domain/entity"

type Customer struct {
	profile entity.Profile
}

func (c *Customer) GetProfile() entity.Profile {
	return c.profile
}

func (c *Customer) SetProfile(profile entity.Profile) {
	c.profile = profile
}
