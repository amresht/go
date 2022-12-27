package main
/*
@author				amresht
@description		this is the main module for the Loading JSON text as defined in the conf.json
*/
import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"os"
)

type MySQLDBConn struct {
	DbUsername  string `json:"DB_USERNAME"`
	DbPassword  string `json:"DB_PASSWORD"`
	DbTimeout   string `json:"DB_TIMEOUT"`
	DbFqdn      string `json:"DB_SERVER_FQDN"`
	DbIpAddress string `json:"DB_SERVER_IP"`
}

type CloudProvider struct {
	IsValid       string        `json:"VALID"`
	Name          string        `json:"NAME"`
	StartTime     int           `json:"START_TIME"`
	EndTime       int           `json:"END_TIME"`
	ResourceGroup string        `json:"RESOURCE_GROUP"`
	Machine       MachineConfig `json:"MACHINE_CONFIG"`
}

type MachineConfig struct {
	CloudAPIBin  string `json:"CLOUD_API_BIN"`
	InstanceTyep string `json:"INSTANCE_TYPE"`
	Country      string `json:"COUNTRY"`
	Region       string `json:"REGION"`
}

type MyConfig struct {
	AppLogPath     string          `json:"APP_LOGPATH"`
	ErrLogPath     string          `json:"ERR_LOGPATH"`
	DebugMode      string          `json:"DEBUG_MODE"`
	EnvTag         string          `json:"ENV_TAG"`
	DBConn         MySQLDBConn     `json:"DATABASE"`
	CloudProviders []CloudProvider `json:"CLOUD_PROVIDERS"`
	Consumers      []string        `json:"CONSUMERS"`
}

var (
	InfoLogger    *log.Logger
	WarningLogger *log.Logger
	DebugLogger   *log.Logger
	ErrorLogger   *log.Logger
	FatalLogger   *log.Logger
)

func init() {
	file, err := os.OpenFile("C:\\AMRESH\\go\\src\\goApp\\logs\\App.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		log.Fatal(err)
	}

	InfoLogger = log.New(file, "INFO:", log.Ldate|log.Ltime|log.Lshortfile)
	DebugLogger = log.New(file, "DEBUG:", log.Ldate|log.Ltime|log.Lshortfile)
	WarningLogger = log.New(file, "WARN:", log.Ldate|log.Ltime|log.Lshortfile)
	ErrorLogger = log.New(file, "ERROR:", log.Ldate|log.Ltime|log.Lshortfile)
	FatalLogger = log.New(file, "FATAL:", log.Ldate|log.Ltime|log.Lshortfile)
}

func MyLoadJSON() MyConfig {
	InfoLogger.Println("Starting to load JSON.....")

	jsonFile, err := os.Open("conf.json")
	if err != nil {
		fmt.Println(err)
		FatalLogger.Println("Could not open json file")
	} else {
		InfoLogger.Println("JSON file opened successfully.....")
	}

	byteValue, err := io.ReadAll(jsonFile)

	if err != nil {
		fmt.Println("Could not read JSON, file possibly corrupt")
		FatalLogger.Println("Could not read JSON, file possibly corrupt")

	}

	jsonFile.Close()

	var VarConfig MyConfig
	err = json.Unmarshal(byteValue, &VarConfig)
	if err != nil {
		fmt.Println(err)
		FatalLogger.Println(err)
	} else {
		InfoLogger.Println("JSON config loaded successfully.....")
	}
	return VarConfig

}
