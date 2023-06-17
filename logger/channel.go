package logger

type Channel struct {
	Path     string `mapstructure:"path" json:"path" yaml:"path"`
	Level    string `mapstructure:"level" json:"level" yaml:"level"`
	Days     int    `mapstructure:"days" json:"days" yaml:"days"`
	Console  bool   `mapstructure:"console" json:"console" yaml:"console"`
	Format   string `mapstructure:"format" json:"format" yaml:"format"`
	MaxSize  int    `mapstructure:"maxSize" json:"maxSize" yaml:"maxSize"`
	Compress bool   `mapstructure:"compress" json:"compress" yaml:"compress"`
}
