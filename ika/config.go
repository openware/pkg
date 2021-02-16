package ika

// Setter is an interface for a custom value setter.
//
// To implement a custom value setter you need to add a SetValue function to your type that will receive a string raw value:
//
// 	type MyField string
//
// 	func (f *MyField) SetValue(s string) error {
// 		if s == "" {
// 			return fmt.Errorf("field value can't be empty")
// 		}
// 		*f = MyField("my field is: " + s)
// 		return nil
// 	}
type Setter interface {
	SetValue(string) error
}

// Updater gives an ability to implement custom update function for a field or a whole structure
type Updater interface {
	Update() error
}

// ReadConfig reads configuration file and parses it depending on tags in structure provided.
// Then it reads and parses
//
// Example:
//
//	 type ConfigDatabase struct {
//	 	Port     string `yaml:"port" env:"PORT" env-default:"5432"`
//	 	Host     string `yaml:"host" env:"HOST" env-default:"localhost"`
//	 	Name     string `yaml:"name" env:"NAME" env-default:"postgres"`
//	 	User     string `yaml:"user" env:"USER" env-default:"user"`
//	 	Password string `yaml:"password" env:"PASSWORD"`
//	 }
//
//	 var cfg ConfigDatabase
//
//	 err := cleanenv.ReadConfig("config.yml", &cfg)
//	 if err != nil {
//	     ...
//	 }

type Source interface {
	Load(dst interface{}) error
}

func ReadConfig(cfg interface{}, sources ...Source) error {
	for _, source := range sources {
		err := source.Load(cfg)
		if err != nil {
			return err
		}
	}

	return nil
}
