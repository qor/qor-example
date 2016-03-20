package commands

import (
	// 	"bytes"
	"fmt"
	// 	"log"
	// 	// "net"

	// 	"os"
	// 	// "time"

	// 	// "bitbucket.org/grengojbo/ads-core/config"
	// 	// "bitbucket.org/grengojbo/ads-core/core"
	// 	// "bitbucket.org/grengojbo/ads-core/db"
	// 	// "bitbucket.org/grengojbo/ads-core/services"
	// 	apim "github.com/Netwurx/routeros-api-go"
	"github.com/codegangsta/cli"
	// 	"github.com/grengojbo/sw-cli/config"
	// 	"github.com/grengojbo/sw-cli/mikrotik"
	// 	"github.com/jinzhu/configor"
	// 	"golang.org/x/crypto/ssh"
	// 	// "crypto/ssh"
	// 	// "github.com/gin-gonic/gin"
	// 	// "github.com/jinzhu/configor"
	// 	// "github.com/fatih/color"
	// 	// "github.com/qor/qor-example/config/admin"
	// 	// "github.com/qor/qor-example/db/migrations"
)

// // Show debug message
// func debug(v ...interface{}) {
// 	if os.Getenv("DEBUG") != "" {
// 		log.Println(v...)
// 	}
// }

// // Error assert
// func assert(err error) {
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// }

// type Routerboard struct {
// 	Config  *config.Config
// 	Verbose bool
// }

// // Loading configuration from yaml file
// func getConfig(c *cli.Context) (config.Config, error) {
// 	yamlPath := c.GlobalString("config")
// 	conf := config.Config{}

// 	err := configor.Load(&conf, yamlPath)
// 	return conf, err
// }

var Commands = []cli.Command{
	// 	cmdShow,
	// 	cmdRun,
	// 	cmdPing,
	cmdMigrate,
}

var cmdMigrate = cli.Command{
	Name:   "migrate",
	Usage:  "Migration DB",
	Action: runMigrate,
}

func runMigrate(c *cli.Context) {
	fmt.Println("Start Migration ...")
	// conf, err := getConfig(c)
	// assert(err)
	// m := Routerboard{Config: &conf, Verbose: c.GlobalBool("verbose")}
	// err, _ = m.RouterboardInterface("178.151.111.129", "admin", "1AzRss53")
	// assert(err)
	fmt.Println("End migration :)")

}

// var cmdShow = cli.Command{
// 	Name:  "show",
// 	Usage: "show subcommand",
// 	Subcommands: []cli.Command{
// 		{
// 			Name:   "firewall",
// 			Usage:  "access list",
// 			Action: runShowFireWall,
// 		},
// 		{
// 			Name:   "int",
// 			Usage:  "interface list",
// 			Action: runShowInterface,
// 		},
// 	},
// 	// Description: `Start QOR Admin web server`,
// 	// Action: runShow,
// 	// Flags: []cli.Flag{
// 	// 	cli.IntFlag{
// 	// 		Name:  "port, p",
// 	// 		Usage: "port number to start web server",
// 	// 	},
// 	// 	cli.StringFlag{"host", "", "Host to start web server", ""},
// 	// 	cli.BoolFlag{
// 	// 		Name:   "release",
// 	// 		Usage:  "Release mode in production.",
// 	// 		EnvVar: "GIN_MODE",
// 	// 	},
// 	// },
// }

// var cmdRun = cli.Command{
// 	Name:   "run",
// 	Usage:  "Run command",
// 	Action: runRun,
// }

// var cmdPing = cli.Command{
// 	Name:   "ping",
// 	Usage:  "Metric RouteOS",
// 	Action: runPing,
// }

// func runShowInterface(c *cli.Context) {
// 	conf, err := getConfig(c)
// 	assert(err)
// 	m := Routerboard{Config: &conf, Verbose: c.GlobalBool("verbose")}
// 	err, _ = m.RouterboardInterface("178.151.111.129", "admin", "1AzRss53")
// 	assert(err)
// }

// func runShowFireWall(c *cli.Context) {
// 	conf, err := getConfig(c)
// 	assert(err)
// 	m := Routerboard{Config: &conf, Verbose: c.GlobalBool("verbose")}
// 	err, _ = m.RouterboardInterface("178.151.111.129", "admin", "1AzRss53")
// 	assert(err)
// }

// func runPing(c *cli.Context) {
// 	if c.GlobalBool("verbose") {
// 		fmt.Println("Connetion to mikrotik...")
// 	}
// 	conf, err := getConfig(c)
// 	assert(err)
// 	// fmt.Printf("config: %#v", conf.Secret)
// 	m := Routerboard{Config: &conf, Verbose: c.GlobalBool("verbose")}
// 	// fmt.Printf("m.Config.Hosts: %v\n", m.Config.Hosts)
// 	// debug(conf)
// 	for i, h := range m.Config.Hosts {
// 		err, code, res := m.RouterboardResource(h.Ip, h.Username, h.Passwd)
// 		if err != nil {
// 			// Unexpected result on login
// 			// fmt.Printf("Error: %v\n", err.Error())
// 			if code == 1 {
// 				fmt.Printf("%v) Error no ping: %s [ %s ] %s (# %s)\n", i+1, h.Name, h.Ip, h.Adress, h.Dogovor)
// 			} else if code == 2 {
// 				fmt.Printf("%v) Error no connect: %s [ %s ] %s (# %s)\n", i+1, h.Name, h.Ip, h.Adress, h.Dogovor)
// 			}
// 		}
// 		if m.Verbose && code == 200 {
// 			fmt.Printf("%v) %s [ %s ] %s | cpu load %s\n", i+1, h.Name, h.Ip, h.Adress, res.CpuLoad)
// 		}
// 	}
// 	// h := Host{Ip: '178.151.111.129', name: 'kievhleb032', adress: 'Харківське шосе, 144-б 2254709', ping: true, changefreq: 'weekly', priority: 1.0, username: 'admin', passwd: '1AzRss53' },
// }

// func (self *Routerboard) RouterboardFirewallFilter(ip, user, passwd string) (err error, code uint8) {
// 	c, err := apim.New(fmt.Sprintf("%s:8728", ip))
// 	if err != nil {
// 		return err, 1
// 	}
// 	err = c.Connect(user, passwd)
// 	if err != nil {
// 		return err, 2
// 	}
// 	defer c.Close()
// 	// res, err = mikrotik.GetResource(c)
// 	// if err != nil {
// 	// 	return err, 3
// 	// }
// 	return nil, 200
// }

// func (self *Routerboard) RouterboardInterface(ip, user, passwd string) (err error, code uint8) {
// 	c, err := apim.New(fmt.Sprintf("%s:8728", ip))
// 	if err != nil {
// 		return err, 1
// 	}
// 	err = c.Connect(user, passwd)
// 	if err != nil {
// 		return err, 2
// 	}
// 	defer c.Close()
// 	// res, err = mikrotik.GetInterface(c)
// 	// if err != nil {
// 	// 	return err, 3
// 	// }
// 	return nil, 200

// }

// func (self *Routerboard) RouterboardResource(ip, user, passwd string) (err error, code uint8, res mikrotik.Resource) {
// 	// res := mikrotik.Resource{}
// 	// h := mikrotik.Host{Ip: "178.151.111.129", Name: "kievhleb032", Adress: "Харківське шосе, 144-б 2254709", Ping: true, Username: "admin", Passwd: "1AzRss53"}
// 	// h := mikrotik.Host{Ip: "178.150.65.139", Name: "kievhleb032", Adress: "Харківське шосе, 144-б 2254709", Ping: true, Username: "admin", Passwd: "1AzRss53"}
// 	c, err := apim.New(fmt.Sprintf("%s:8728", ip))
// 	if err != nil {
// 		return err, 1, res
// 	}
// 	err = c.Connect(user, passwd)
// 	if err != nil {
// 		return err, 2, res
// 	}
// 	defer c.Close()
// 	res, err = mikrotik.GetResource(c)
// 	if err != nil {
// 		return err, 3, res
// 	}
// 	// if self.Verbose {
// 	// fmt.Printf("%s [ %s ] cpu load %s\n", h.Name, h.Ip, res.CpuLoad)

// 	// }
// 	// _, err = mikrotik.GetEthernet("ether1-gateway", c)
// 	// if err != nil {
// 	// 	return err
// 	// }
// 	return nil, 200, res

// }

// func runShow(c *cli.Context) {
// 	// config := &ssh.ClientConfig{
// 	// User: c.GlobalString("user"),
// 	// // Auth: []ssh.ClientAuth{makeKeyring()},
// 	// Auth: []ssh.AuthMethod{ssh.Password(c.GlobalString("passwd"))}
// 	// }
// 	fmt.Printf("User: %s passwd: %s to --> %s\n", c.GlobalString("user"), c.GlobalString("passwd"), c.GlobalString("host"))
// 	if !c.GlobalBool("telnet") {
// 		fmt.Printf("Connect telnet: %s\n", c.GlobalString("host"))
// 	}
// 	client, session, err := connectToSsh(c.GlobalString("user"), c.GlobalString("passwd"), c.GlobalString("host"))
// 	if err != nil {
// 		// panic(err)
// 		log.Fatal(err.Error())
// 	}
// 	execLine(session, c.Args().First())
// 	client.Close()
// }

// func runRun(c *cli.Context) {
// 	// runToHost(c.GlobalString("user"), c.GlobalString("passwd"), c.GlobalString("host"), c.Args().First())
// 	runPipe(c.GlobalString("user"), c.GlobalString("passwd"), c.GlobalString("host"), c.Args().First())
// }

// func execLine(session *ssh.Session, cmd string) {
// 	out, err := session.CombinedOutput(cmd)
// 	if err != nil {
// 		// panic(err)
// 		log.Fatal(err.Error())
// 	}
// 	fmt.Println(string(out))
// }
// func connectToTelnet(user, pass, host string) {
// }

// func connectToSsh(user, pass, host string) (*ssh.Client, *ssh.Session, error) {
// 	// var pass string
// 	// fmt.Print("Password: ")
// 	// fmt.Scanf("%s\n", &pass)

// 	sshConfig := &ssh.ClientConfig{
// 		User: user,
// 		Auth: []ssh.AuthMethod{ssh.Password(pass)},
// 	}

// 	client, err := ssh.Dial("tcp", host, sshConfig)
// 	if err != nil {
// 		return nil, nil, err
// 	}

// 	session, err := client.NewSession()
// 	if err != nil {
// 		client.Close()
// 		return nil, nil, err
// 	}

// 	return client, session, nil
// }

// func runToHost(user, pass, host, cmd string) {
// 	sshConfig := &ssh.ClientConfig{
// 		User: user,
// 		Auth: []ssh.AuthMethod{ssh.Password(pass)},
// 	}
// 	sshConfig.SetDefaults()

// 	client, err := ssh.Dial("tcp", host, sshConfig)
// 	if err != nil {
// 		// errors.Wrap(err, err.Error())
// 		fmt.Println(err.Error())
// 	}

// 	var session *ssh.Session

// 	session, err = client.NewSession()
// 	if err != nil {
// 		// errors.Wrap(err, err.Error())
// 		fmt.Println(err.Error())
// 	}
// 	defer session.Close()

// 	var stdoutBuf bytes.Buffer
// 	session.Stdout = &stdoutBuf
// 	session.Run(cmd)

// 	fmt.Println(stdoutBuf.String())
// }

// func runPipe(user, pass, host, cmd string) {
// 	fmt.Println("runPipe...")
// 	sshConfig := &ssh.ClientConfig{
// 		User: user,
// 		Auth: []ssh.AuthMethod{ssh.Password(pass)},
// 	}
// 	sshConfig.SetDefaults()

// 	client, err := ssh.Dial("tcp", host, sshConfig)
// 	if err != nil {
// 		// errors.Wrap(err, err.Error())
// 		fmt.Println(err.Error())
// 	}

// 	var session *ssh.Session

// 	session, err = client.NewSession()
// 	if err != nil {
// 		// errors.Wrap(err, err.Error())
// 		fmt.Println(err.Error())
// 	}
// 	defer session.Close()

// 	// modes := ssh.TerminalModes{
// 	// 	ssh.ECHO:          0,     // disable echoing
// 	// 	ssh.TTY_OP_ISPEED: 14400, // input speed = 14.4kbaud
// 	// 	ssh.TTY_OP_OSPEED: 14400, // output speed = 14.4kbaud
// 	// }

// 	// if err := session.RequestPty("xterm", 80, 40, modes); err != nil {
// 	// 	log.Fatal(err)
// 	// }

// 	w, err := session.StdinPipe()
// 	if err != nil {
// 		panic(err)
// 	}
// 	//    r, err := session.StdoutPipe()
// 	//    if err != nil {
// 	//        panic(err)
// 	//    }

// 	var stdoutBuf bytes.Buffer
// 	session.Stdout = &stdoutBuf
// 	// session.Run(cmd)
// 	err = session.Shell()
// 	w.Write([]byte(fmt.Sprintf("%s\n", "configure")))
// 	w.Write([]byte(fmt.Sprintf("%s %s\n", "set interfaces ethernet eth4 description", cmd)))
// 	w.Write([]byte(fmt.Sprintf("%s\n", "commit")))
// 	w.Write([]byte(fmt.Sprintf("%s\n", "save")))
// 	w.Write([]byte(fmt.Sprintf("%s\n", "exit")))
// 	w.Write([]byte(fmt.Sprintf("%s\n", "exit")))

// 	fmt.Println(stdoutBuf.String())
// }
