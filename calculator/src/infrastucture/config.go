package infrastucture

type Config struct {
	Port                  int    `json:"port"`
	Password              string `json:"password"`
	Database              string `json:"database"`
	Host                  string `json:"host"`
	User                  string `json:"user"`
	UserCollectionName    string `json:"user_collection_name"`
	ProductCollectionName string `json:"product_collection_name"`
}
