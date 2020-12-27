# Ika: light config for go apps

## Introductions

Openware config for golang supporting json, yaml and ENV configs 12factor compliant

## Usage

`go get github.com/openware/config`

### Example config

```
type config struct {
	Number    int64  `yaml:"number" env:"TEST_NUMBER" env-default:"1"`
	String    string `yaml:"string" env:"TEST_STRING" env-default:"default"`
	NoDefault string `yaml:"no-default" env:"TEST_NO_DEFAULT"`
	NoEnv     string `yaml:"no-env" env-default:"default"`
}
```

### Credits

Forked from: https://github.com/ilyakaznacheev/cleanenv
