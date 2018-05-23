package setting

import (
	"os"
	"log"
	"encoding/json"
)

type Configuration struct {
	ServicePort int

	DB string

	CdnDomain string
	CdnHttps  bool

	S3BucketName      string
	AwsKeyId          string
	AwsSecretKey      string
	UploadImageFolder string

	OriginDomainLocal string
	OriginDomainRun   []string
}

func (configuration Configuration) String() string {
	temp, err := json.Marshal(configuration)
	if err != nil {
		return ""
	}
	return string(temp)
}

var configuration *Configuration = nil

func LoadConfig() (*Configuration, error) {
	// Load configuration
	file, err := os.Open("setting/conf.json")
	if err != nil {
		log.Println("error:", err)
	}
	decoder := json.NewDecoder(file)
	configuration := Configuration{}
	errJson := decoder.Decode(&configuration)
	if errJson != nil {
		log.Println("error:", errJson)
		return nil, err
	}
	log.Println(configuration)
	// End load config
	return &configuration, nil
}

func CurrentConfig() *Configuration {
	if configuration == nil {
		configurationVal, err := LoadConfig()
		if err != nil {
			log.Print(err)
		}
		configuration = configurationVal
	}
	return configuration
}
