package util

type DotEnvConfig map[string]string

func (d DotEnvConfig) StringOrDefaultKey(key string, defKey string) string {
	val, ok := d[key]
	if !ok || val == "" {
		return d[defKey]
	}

	return d[key]
}
