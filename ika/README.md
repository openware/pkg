# Ika: light config for go apps

## Introduction

Openware config for golang supporting json, yaml and ENV configs 12factor compliant

## Usage

`go get github.com/openware/pkg/ika`

### Example config and usage

```go
type YourAppConfig struct {
	Number    int64  `yaml:"number" env:"TEST_NUMBER" env-default:"1"`
	String    string `yaml:"string" env:"TEST_STRING" env-default:"default"`
	NoDefault string `yaml:"no-default" env:"TEST_NO_DEFAULT"`
	NoEnv     string `yaml:"no-env" env-default:"default"`
	Required  string `yaml:"required" env-required:"true"`
}

func LoadYourAppConfig(cfgFilePath string) (*YourAppConfig, error) {
	cfg := &YourAppConfig{}
	if err := ika.ReadConfig(cfgFilePath, cfg); err != nil {
		return nil, err
	}
	return cfg, nil
}

```

### Credits

Forked from: https://github.com/ilyakaznacheev/cleanenv
