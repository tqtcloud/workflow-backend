package conf

import (
	"context"
	"fmt"
	"sync"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	mgoclient *mongo.Client
)

func newConfig() *Config {
	return &Config{
		App: newDefaultAPP(),
		Log: newDefaultLog(),

		Mongo: newDefaultMongoDB(),
	}
}

// Config 应用配置
type Config struct {
	App     *app     `toml:"app"`
	Log     *log     `toml:"log"`
	Jenkins *jenkins `toml:"jenkins"`
	Mongo   *mongodb `toml:"mongodb"`
	SshExec *sshexec `toml:"sshexec"`
}

type app struct {
	Name       string `toml:"name" env:"APP_NAME"`
	EncryptKey string `toml:"encrypt_key" env:"APP_ENCRYPT_KEY"`
	HTTP       *http  `toml:"http"`
	GRPC       *grpc  `toml:"grpc"`
}

func newDefaultAPP() *app {
	return &app{
		Name:       "cmdb",
		EncryptKey: "defualt app encrypt key",
		HTTP:       newDefaultHTTP(),
		GRPC:       newDefaultGRPC(),
	}
}

type http struct {
	Host      string `toml:"host" env:"HTTP_HOST"`
	Port      string `toml:"port" env:"HTTP_PORT"`
	EnableSSL bool   `toml:"enable_ssl" env:"HTTP_ENABLE_SSL"`
	CertFile  string `toml:"cert_file" env:"HTTP_CERT_FILE"`
	KeyFile   string `toml:"key_file" env:"HTTP_KEY_FILE"`
}

func (a *http) Addr() string {
	return a.Host + ":" + a.Port
}

func newDefaultHTTP() *http {
	return &http{
		Host: "127.0.0.1",
		Port: "8050",
	}
}

type grpc struct {
	Host      string `toml:"host" env:"GRPC_HOST"`
	Port      string `toml:"port" env:"GRPC_PORT"`
	EnableSSL bool   `toml:"enable_ssl" env:"GRPC_ENABLE_SSL"`
	CertFile  string `toml:"cert_file" env:"GRPC_CERT_FILE"`
	KeyFile   string `toml:"key_file" env:"GRPC_KEY_FILE"`
}

func (a *grpc) Addr() string {
	return a.Host + ":" + a.Port
}

func newDefaultGRPC() *grpc {
	return &grpc{
		Host: "127.0.0.1",
		Port: "18050",
	}
}

type log struct {
	Level   string    `toml:"level" env:"LOG_LEVEL"`
	PathDir string    `toml:"path_dir" env:"LOG_PATH_DIR"`
	Format  LogFormat `toml:"format" env:"LOG_FORMAT"`
	To      LogTo     `toml:"to" env:"LOG_TO"`
}

type jenkins struct {
	DevEndpoints  string `toml:"devendpoints" env:"Dev_ENDPOINTS" `
	TestEndpoints string `toml:"testendpoints" env:"Test_ENDPOINTS" `
	UatEndpoints  string `toml:"uatendpoints" env:"Uat_ENDPOINTS" `
	LptEndpoints  string `toml:"lptendpoints" env:"Lpt_ENDPOINTS" `
	ProdEndpoints string `toml:"prodendpoints" env:"Prod_ENDPOINTS" `
	User          string `toml:"user" env:"MONGO_USERNAME"`
	Password      string `toml:"password" env:"MONGO_PASSWORD"`
}

type sshexec struct {
	DevNode        string `toml:"devnode" env:"Dev_Node" `
	TestNode       string `toml:"testnode" env:"Test_Node" `
	UatNode        string `toml:"uatnode" env:"Uat_Node" `
	ProdNode       string `toml:"prodnode" env:"Prod_Node" `
	User           string `toml:"user" env:"SSH_EXEC_USER"`
	DevPort        string `toml:"devport" env:"Dev_Port"`
	TestPort       string `toml:"testport" env:"Test_Port"`
	UatPort        string `toml:"uatport" env:"Uat_Port"`
	ProdPort       string `toml:"prodport" env:"Prod_Port"`
	DevSshKeyPath  string `toml:"devsshKeyPath" env:"Dev_SSH_Key_Path"`
	ProdSshKeyPath string `toml:"prodsshKeyPath" env:"Prod_SSH_Key_Path"`
	TestSshKeyPath string `toml:"testsshKeyPath" env:"Test_SSH_Key_Path"`
	UatSshKeyPath  string `toml:"uatsshKeyPath" env:"Uat_SSH_Key_Path"`
	DevSshShell    string `toml:"DevSshShell" env:"Dev_SSH_Shell"`
	TestSshShell   string `toml:"TestSshShell" env:"Test_SSH_Shell"`
	UatSshShell    string `toml:"UatSshShell" env:"Uat_SSH_Shell"`
	ProdSshShell   string `toml:"ProdSshShell" env:"Prod_SSH_Shell"`
}

func newDefaultLog() *log {
	return &log{
		Level:   "debug",
		PathDir: "logs",
		Format:  "text",
		To:      "stdout",
	}
}

func newDefaultMongoDB() *mongodb {
	return &mongodb{
		Database:  "",
		Endpoints: []string{"127.0.0.1:27017"},
	}
}

type mongodb struct {
	Endpoints []string `toml:"endpoints" env:"MONGO_ENDPOINTS" envSeparator:","`
	UserName  string   `toml:"username" env:"MONGO_USERNAME"`
	Password  string   `toml:"password" env:"MONGO_PASSWORD"`
	Database  string   `toml:"database" env:"MONGO_DATABASE"`
	lock      sync.Mutex
}

// Client 获取一个全局的mongodb客户端连接
func (m *mongodb) Client() (*mongo.Client, error) {
	// 加载全局数据量单例
	m.lock.Lock()
	defer m.lock.Unlock()
	if mgoclient == nil {
		conn, err := m.getClient()
		if err != nil {
			return nil, err
		}
		mgoclient = conn
	}

	return mgoclient, nil
}

func (m *mongodb) GetDB() (*mongo.Database, error) {
	conn, err := m.Client()
	if err != nil {
		return nil, err
	}
	return conn.Database(m.Database), nil
}

func (m *mongodb) getClient() (*mongo.Client, error) {
	opts := options.Client()

	cred := options.Credential{
		AuthSource: m.Database,
	}

	if m.UserName != "" && m.Password != "" {
		cred.Username = m.UserName
		cred.Password = m.Password
		cred.PasswordSet = true
		opts.SetAuth(cred)
	}
	opts.SetHosts(m.Endpoints)
	opts.SetConnectTimeout(5 * time.Second)

	// Connect to MongoDB
	client, err := mongo.Connect(context.TODO(), opts)
	if err != nil {
		return nil, fmt.Errorf("new mongodb client error, %s", err)
	}

	if err = client.Ping(context.TODO(), nil); err != nil {
		return nil, fmt.Errorf("ping mongodb server(%s) error, %s", m.Endpoints, err)
	}

	return client, nil
}
