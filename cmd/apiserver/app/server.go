package app

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
	"k8s.io/klog/v2"

	"github.com/huangjiasingle/suyi/cmd/apiserver/app/options"
	"github.com/huangjiasingle/suyi/pkg/apis/order"
	"github.com/huangjiasingle/suyi/pkg/tools/http/server"
	"github.com/huangjiasingle/suyi/pkg/tools/storage/db"
)

// NewAPIServerCommand creates a *cobra.Command object with default parameters
func NewAPIServerCommand() *cobra.Command {
	s := options.NewServerRunOptions()
	// 1.完善flag参数
	completedOptions := Complete(s)
	cmd := &cobra.Command{
		Use:  "apiserver",
		Long: ``,
		RunE: func(cmd *cobra.Command, args []string) error {
			klog.InfoS("Loaded ipaas configuration", "Configuration", completedOptions)
			// 1.验证参数
			// 2.初始化各种数据库的链接
			// 3.初始化缓存
			// 4.生成http 处理api
			// 5.启动http server服务
			if errs := Validate(completedOptions); len(errs) != 0 {
				return errs[0]
			}

			if err := InitDB(completedOptions); err != nil {
				return err
			}

			server := CreateServerChain(completedOptions)
			klog.Infof("Listening and serving HTTP on :%d", completedOptions.System.Port)
			if err := server.ServerHTTP(fmt.Sprintf("%v:%d", completedOptions.System.Address, completedOptions.System.Port)); err != nil {
				return err
			}
			return nil
		},
	}
	return cmd
}

// Complet 完成参数的初始化
func Complete(s *options.ServerRunOptions) *options.ServerRunOptions {
	pflag.StringVar(&s.Mysql.IP, "mysql-ip", "", "--mysql-ip")
	pflag.IntVar(&s.Mysql.Port, "mysql-port", 0, "--mysql-port")
	pflag.StringVar(&s.Mysql.DB, "mysql-db", "", "--mysql-db")
	pflag.StringVar(&s.Mysql.User, "mysql-user", "", "--mysql-user")
	pflag.StringVar(&s.Mysql.Password, "mysql-password", "", "--mysql-password")

	pflag.StringVar(&s.Config, "config", s.Config, "The server config.")
	pflag.Parse()
	return s
}

// Validate 验证参数是否有误
func Validate(s *options.ServerRunOptions) (errors []error) {
	if s.Config != "" {
		if err := s.LoadConfig(); err != nil {
			errors = append(errors, err)
		}
	}

	if s.Mysql.IP == "" {
		errors = append(errors, fmt.Errorf("mysql ip is null, it must be define"))
	}
	if s.Mysql.Port == 0 {
		errors = append(errors, fmt.Errorf("mysql port is null, it must be define"))
	}
	if s.Mysql.DB == "" {
		errors = append(errors, fmt.Errorf("mysql database is null, it must be define"))
	}
	if s.Mysql.User == "" {
		errors = append(errors, fmt.Errorf("mysql user is null, it must be define"))
	}
	if s.Mysql.Password == "" {
		errors = append(errors, fmt.Errorf("mysql password is null, it must be define"))
	}

	return errors
}

// InitDB 初始化各种数据库的链接
func InitDB(s *options.ServerRunOptions) error {
	dsn := fmt.Sprintf("%v:%v@(%v:%v)/%v?timeout=30s&loc=Local&parseTime=true", s.Mysql.User, s.Mysql.Password, s.Mysql.IP, s.Mysql.Port, s.Mysql.DB)
	if err := db.Init(dsn, s.System.Debug); err != nil {
		return fmt.Errorf("init mysql client failed: %v", err)
	}
	return nil
}

// CreateServerChain ....
func CreateServerChain(s *options.ServerRunOptions) *server.Server {
	sv := server.New(s.System.Prefix)
	order.RegistryOrderAPI(sv)
	return sv
}
