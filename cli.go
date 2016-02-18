package main

import (
	"fmt"
	// "log"
	"os"
	"runtime"

	// "bitbucket.org/grengojbo/ads-core/config"
	// "github.com/nu7hatch/gouuid"
	"github.com/codegangsta/cli"
	"github.com/grengojbo/gotools"
	"github.com/qor/qor-example/app/models"
	"github.com/qor/qor-example/db"
)

var (
	Version   = "0.1.0"
	BuildTime = "2015-09-20 UTC"
	GitHash   = "c00"
)

func ConfigRuntime() {
	nuCPU := runtime.NumCPU()
	runtime.GOMAXPROCS(nuCPU)
	fmt.Printf("Running with %d CPUs\n", nuCPU)
}

var Commands = []cli.Command{
	cmdMigrate,
	cmdUser,
}

var cmdMigrate = cli.Command{
	Name:   "migrate",
	Usage:  "Migration DB",
	Action: runMigrate,
}

var cmdUser = cli.Command{
	Name:  "user",
	Usage: "Manage User",
	// Action: runUser,
	Subcommands: []cli.Command{
		{
			Name:   "add",
			Usage:  "Create new User",
			Action: runUserAdd,
		},
		{
			Name:   "set",
			Usage:  "Set pasword or email for User",
			Action: runUserSet,
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:  "password, p",
					Usage: "Set User password",
				},
			},
		},
		{
			Name:   "show",
			Usage:  "Show User information",
			Action: runUserShow,
		},
	},
}

func runMigrate(c *cli.Context) {
	fmt.Println("Start Migration ...")
	fmt.Println("End migration :)")
}

func runUserAdd(c *cli.Context) {
	if len(c.Args().First()) > 1 {
		fmt.Println("TODO ...", c.Args().First())
		// user := models.User{}
		// user.Email = "admin@example.com"
		// user.Name = "admin"
		// user.Password = "$2a$10$SXinmKBnwhRcB4EJLlTO2.OebRd0Tv8TzvFMLJ6XNiJeB0//SolS."
		// user.Gender = "Male"
		// user.CreatedAt = time.Now()
		// user.Role = "admin"
		// if err := db.DB.Create(&user).Error; err != nil {
		// 	log.Fatalf("create user (%v) failure, got err %v", user, err)
		// }
	} else {
		fmt.Println("Is not set User :)")
	}
}

func runUserSet(c *cli.Context) {
	if len(c.Args().First()) > 1 {
		var user models.User
		if !db.DB.Where("name = ? OR email = ?", c.Args().First(), c.Args().First()).First(&user).RecordNotFound() {
			if c.IsSet("password") {
				user.Password = gotools.PasswordBcrypt(c.String("password"))
				db.DB.Save(&user)
				fmt.Printf("Set %s password: %s\n", user.Name, c.String("password"))
				// fmt.Printf("Set %s password: %s\n", user.Name, user.Password)
				// fmt.Printf("%v\n", gotools.VerifyPassword(passwd, "admin"))
			}
		} else {
			fmt.Println("Is not exits User:", c.Args().First())
		}
	} else {
		fmt.Println("Is not set User :)")
	}
}

func runUserShow(c *cli.Context) {
	if len(c.Args().First()) > 1 {
		var user models.User
		if !db.DB.Where("name = ? OR email = ?", c.Args().First(), c.Args().First()).First(&user).RecordNotFound() {
			fmt.Printf("User: %s %s [%s]\n", user.LastName, user.FirstName, user.Name)
			fmt.Printf("E-Mail: %s\n", user.Email)
			fmt.Println("Gender:", user.Gender)
			fmt.Println("Role:", user.Role)
			// fmt.Println("Adress:", user.Addresses)
		} else {
			fmt.Println("Is not exits User:", c.Args().First())
		}
	} else {
		fmt.Println("Is not set User :)")
	}
}

// func run(c *cli.Context) {
// 	fmt.Println("...")
// }

func main() {
	app := cli.NewApp()
	app.Name = "qor-cli"
	app.Version = Version
	app.Usage = "Run QOR command"
	app.Author = "Oleg Dolya"
	app.Email = "oleg.dolya@gmail.com"
	app.EnableBashCompletion = true
	app.Commands = Commands
	app.Flags = []cli.Flag{
		cli.BoolFlag{
			Name: "verbose",
			// Value: "false",
			Usage: "Verbose mode",
		},
		cli.BoolFlag{
			Name: "debug",
			// Value: "false",
			Usage: "Debug mode",
		},
		// cli.StringFlag{
		//   Name:   "passwd, p",
		//   Usage:  "SSH pasword",
		//   EnvVar: "CISCO_PASSWD",
		// },
		// cli.StringFlag{
		//   Name:   "user, u",
		//   Usage:  "SSH user name",
		//   EnvVar: "CISCO_USER",
		// },
		// cli.StringFlag{
		//   Name:  "host",
		//   Usage: "connection to only host",
		// },
		// cli.BoolTFlag{
		//   Name:  "teltet, t",
		//   Usage: "Connect to telnet mode",
		// },
		cli.StringFlag{
			Name:   "config, c",
			Value:  "config/config.yml",
			Usage:  "config file to use (config/config.yml)",
			EnvVar: "APP_CONFIG",
		},
	}
	// app.Before = func(c *cli.Context) error {
	// log.Println("Load config:", c.GlobalString("config"))
	// cfg, err := getConfig(c)
	//    if err != nil {
	//      log.Fatal(err)
	//      return
	//    }
	//    log.Println(cfg)
	// setting.CustomConf = c.GlobalString("config")
	// setting.NewConfigContext()
	// setting.NewServices()
	// return nil
	// }
	app.Run(os.Args)
}