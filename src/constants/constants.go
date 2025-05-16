package constants

import (
	"strconv"
	"time"
)

var hi = `
			 ██╗  ██╗██╗
			 ██║  ██║██║
			 ███████║██║
			 ██╔══██║██║
			 ██║  ██║██║
			 ╚═╝  ╚═╝╚═╝`

var nameAndEmail = "        \033[1;33m Pablo Triandafilide | " + "pablotrianda@gmail.com" + "\033[0m\n"

var contactMeLineTop1 = `┌─`
var contactText = "\033[1;32mContact\033[0m"
var contactMeLineTop2 = `───────────────────────────────────────────────────────────────┐
`
var github = "│ \033[1;34m Github   https://github.com/pablotrianda\033[0m                         	│\n"

var linkedIn = "│ \033[1;34m LinkedIn https://www.linkedin.com/in/pablo-triandafilide-641b24ba/\033[0m   │"

var contactMeLineBottom = `
└───────────────────────────────────────────────────────────────────────┘
`
var contactMe = contactMeLineTop1 +
	contactText +
	contactMeLineTop2 +
	github +
	linkedIn +
	contactMeLineBottom

var aboutMeLineTop1 = `┌─`
var aboutText = "\033[1;32mAbout me\033[0m"
var aboutMeLineTop2 = `────────────────────────────────────────────────────────────────────────────────────────────────────────────┐
`
var aboutMeDesc = "│ 🔭 I’m a backend developer specializing in Golang."+        
	"                                                                  │ \n" +
	"│ 🌱 I’m currently learning and building personal tools to improve my workflow and daily tasks using Golang and Rust. │ \n"+          
	"│ 👯 I’m looking to collaborate on cool and fun projects that help other developers work more efficiently.            │ \n" +
	"│ 📫 How to reach me: Email or LinkedIn. " +
	"                                                                             │" 

var aboutMeLineBottom = `
└─────────────────────────────────────────────────────────────────────────────────────────────────────────────────────┘
`
var aboutMe = aboutMeLineTop1 +
	aboutText +
	aboutMeLineTop2 +
	aboutMeDesc +
	aboutMeLineBottom

var space = "\n\n\n\n"

func footer() string {
	t := time.Now()
	year := t.Year()
	return " \033[1;32m" + strconv.Itoa(year) + " Pablo Triandafilide. Made with 🪄 and some 🧉 \033[0m\n"
}

var header = "\033[1;32m" + hi + "\033[0m\n"

var body = nameAndEmail +
	contactMe +
	aboutMe +
	space +
	footer()

func Page() string {
	return header + body
}
