package infrastructure

type Config struct {
	Port                  int    `mapstructure:"port"`
	Password              string `mapstructure:"password"`
	Database              string `mapstructure:"database"`
	Host                  string `mapstructure:"host"`
	User                  string `mapstructure:"user"`
	UserCollectionName    string `mapstructure:"user_collection_name"`
	ProductCollectionName string `mapstructure:"product_collection_name"`
}
