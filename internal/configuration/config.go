package configuration

import (
	"encoding/json"
	"fmt"

	configuration "github.com/AlpacaLabs/go-config"
	"github.com/rs/xid"

	flag "github.com/spf13/pflag"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

const (
	flagForGrpcPort = "grpc_port"
	flagForHTTPPort = "http_port"

	flagForAccountGrpcAddress = "account_service_address"
	flagForAccountGrpcHost    = "account_service_host"
	flagForAccountGrpcPort    = "account_service_port_grpc"

	flagForMFAGrpcAddress = "mfa_service_address"
	flagForMFAGrpcHost    = "mfa_service_host"
	flagForMFAGrpcPort    = "mfa_service_port_grpc"

	flagForPasswordGrpcAddress = "password_service_address"
	flagForPasswordGrpcHost    = "password_service_host"
	flagForPasswordGrpcPort    = "password_service_port_grpc"
)

type Config struct {
	// AppName is a low cardinality identifier for this service.
	AppName string

	// AppID is a unique identifier for the instance (pod) running this app.
	AppID string

	// KafkaConfig provides configuration for connecting to Apache Kafka.
	KafkaConfig configuration.KafkaConfig

	// SQLConfig provides configuration for connecting to a SQL database.
	SQLConfig configuration.SQLConfig

	// GrpcPort controls what port our gRPC server runs on.
	GrpcPort int

	// HTTPPort controls what port our HTTP server runs on.
	HTTPPort int

	// AccountGRPCAddress is the gRPC address of the Account service.
	AccountGRPCAddress string

	// MFAGRPCAddress is the gRPC address of the MFA service.
	MFAGRPCAddress string

	// PasswordGRPCAddress is the gRPC address of the Password service.
	PasswordGRPCAddress string
}

func (c Config) String() string {
	b, err := json.Marshal(c)
	if err != nil {
		log.Fatalf("Could not marshal config to string: %v", err)
	}
	return string(b)
}

func LoadConfig() Config {
	c := Config{
		AppName:  "api-auth",
		AppID:    xid.New().String(),
		GrpcPort: 8081,
		HTTPPort: 8083,
	}

	c.KafkaConfig = configuration.LoadKafkaConfig()
	c.SQLConfig = configuration.LoadSQLConfig()

	flag.Int(flagForGrpcPort, c.GrpcPort, "gRPC port")
	flag.Int(flagForHTTPPort, c.HTTPPort, "HTTP port")

	flag.String(flagForAccountGrpcAddress, "", "Address of Account gRPC service")
	flag.String(flagForAccountGrpcHost, "", "Host of Account gRPC service")
	flag.String(flagForAccountGrpcPort, "", "Port of Account gRPC service")

	flag.String(flagForMFAGrpcAddress, "", "Address of MFA gRPC service")
	flag.String(flagForMFAGrpcHost, "", "Host of MFA gRPC service")
	flag.String(flagForMFAGrpcPort, "", "Port of MFA gRPC service")

	flag.String(flagForPasswordGrpcAddress, "", "Address of Password gRPC service")
	flag.String(flagForPasswordGrpcHost, "", "Host of Password gRPC service")
	flag.String(flagForPasswordGrpcPort, "", "Port of Password gRPC service")

	flag.Parse()

	viper.BindPFlag(flagForGrpcPort, flag.Lookup(flagForGrpcPort))
	viper.BindPFlag(flagForHTTPPort, flag.Lookup(flagForHTTPPort))

	viper.BindPFlag(flagForAccountGrpcAddress, flag.Lookup(flagForAccountGrpcAddress))
	viper.BindPFlag(flagForAccountGrpcHost, flag.Lookup(flagForAccountGrpcHost))
	viper.BindPFlag(flagForAccountGrpcPort, flag.Lookup(flagForAccountGrpcPort))

	viper.BindPFlag(flagForMFAGrpcAddress, flag.Lookup(flagForMFAGrpcAddress))
	viper.BindPFlag(flagForMFAGrpcHost, flag.Lookup(flagForMFAGrpcHost))
	viper.BindPFlag(flagForMFAGrpcPort, flag.Lookup(flagForMFAGrpcPort))

	viper.BindPFlag(flagForPasswordGrpcAddress, flag.Lookup(flagForPasswordGrpcAddress))
	viper.BindPFlag(flagForPasswordGrpcHost, flag.Lookup(flagForPasswordGrpcHost))
	viper.BindPFlag(flagForPasswordGrpcPort, flag.Lookup(flagForPasswordGrpcPort))

	viper.AutomaticEnv()

	c.GrpcPort = viper.GetInt(flagForGrpcPort)
	c.HTTPPort = viper.GetInt(flagForHTTPPort)

	c.AccountGRPCAddress = getGrpcAddress(flagForAccountGrpcAddress, flagForAccountGrpcHost, flagForAccountGrpcPort)
	c.MFAGRPCAddress = getGrpcAddress(flagForMFAGrpcAddress, flagForMFAGrpcHost, flagForMFAGrpcPort)
	c.PasswordGRPCAddress = getGrpcAddress(flagForPasswordGrpcAddress, flagForPasswordGrpcHost, flagForPasswordGrpcPort)

	return c
}

func getGrpcAddress(addrFlag, hostFlag, portFlag string) string {
	addr := viper.GetString(addrFlag)
	host := viper.GetString(hostFlag)
	port := viper.GetInt(portFlag)

	if port != 0 {
		return fmt.Sprintf("%s:%d", host, port)
	}

	return addr
}
