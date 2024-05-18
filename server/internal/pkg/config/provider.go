package config

type Provider = map[string]ServiceProvider

func CombineProvider(providers ...Provider) Provider {
	combined := Provider{}

	for _, provider := range providers {
		if provider == nil {
			continue
		}
		for key, value := range provider {
			combined[key] = value
		}
	}

	return combined
}

type ServiceProvider interface {
	Boot()
}
