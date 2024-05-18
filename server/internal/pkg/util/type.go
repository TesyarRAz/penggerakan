package util

type DotEnvConfig map[string]string

func (d DotEnvConfig) StringOrDefaultKey(key string, defKey string) string {
	val, ok := d[key]
	if !ok || val == "" {
		return d[defKey]
	}

	return d[key]
}

func (d DotEnvConfig) Modify(data map[string]string) DotEnvConfig {
	for k, v := range data {
		d[k] = v
	}

	return d
}
