package internal

import (
	"bytes"
	"fmt"
	"os"
	"os/exec"
	"strings"
)

func BuildProject(proj *Project) {

	switch proj.ProjType {
	case WebApp:
		fmt.Println("Building Web App")
		BuildWebApp(proj)
	case ClientServer:
		BuildClientServer(proj)
	case Terminal:
		BuildTerminal(proj)
	case Mobile:
		BuildMobile(proj)
	default:

	}

}

func BuildWebApp(proj *Project) {
	pwd, err := os.Getwd()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Println(pwd)

	if proj.FrontFrame == React {
		fmt.Println("Building React")
		switch proj.BackFrame {

		//Use npx to make, then replace page.tsx
		case NextJS:
			selections := MakeTemplatePaths(proj)
			build := exec.Command("npx", "create-next-app@latest", proj.ProjName)

			buffer := bytes.Buffer{}
			buffer.Write([]byte("\n"))
			build.Stdin = &buffer

			build.Stdout = os.Stdout
			build.Stderr = os.Stderr

			err := build.Run()
			if err != nil {
				fmt.Fprintln(os.Stderr, err)
			}

			//FIX THIS ONCE FUNCTION IS UPDATED!
			GenerateManifest(selections, fmt.Sprintf("%s/manifest.json", pwd))

			var context = makeContext(proj)

			fmt.Println("Making Project!!")
			GenerateProject(context)

		//Use npx to make, then replace page.tsx
		case ReactRouter:
			selections := MakeTemplatePaths(proj)
			build := exec.Command("npx", "create-react-router@latest", "--template", "remix-run/react-router-templates/minimal", fmt.Sprint("./"+proj.ProjName))

			buffer := bytes.Buffer{}
			buffer.Write([]byte("\x1b[C\n"))
			build.Stdin = &buffer

			buffer.Reset()

			buffer.Write([]byte("\n"))
			build.Stdin = &buffer

			build.Stdout = os.Stdout
			build.Stderr = os.Stderr

			err := build.Run()
			if err != nil {
				fmt.Fprintln(os.Stderr, err)
			}

			GenerateManifest(selections, fmt.Sprintf("%s/manifest.json", pwd))

			var context = makeContext(proj)

			GenerateProject(context)

		//Use npx to make react client, then make server with manifest.json
		case ExpressJs:
			selections := MakeTemplatePaths(proj)
			build := exec.Command("npx", "create-react-app", "client")

			build.Stdout = os.Stdout
			build.Stderr = os.Stderr

			err := build.Run()
			if err != nil {
				fmt.Fprintln(os.Stderr, err)
			}

			//FIX THIS ONCE FUNCTION IS UPDATED!
			GenerateManifest(selections, fmt.Sprintf("%s/manifest.json", pwd))

			var context = makeContext(proj)

			GenerateProject(context)

		default:
		}

	} else {

	}

}

func BuildClientServer(proj *Project) {

}

func BuildTerminal(proj *Project) {

}

func BuildMobile(proj *Project) {

}

func MakeTemplatePaths(proj *Project) map[string]string {
	fmt.Println("Building Map")

	mapping := make(map[string]string)

	switch proj.ProjType {
	case WebApp:
		if proj.FrontFrame == React {

			switch proj.BackFrame {

			case NextJS:
				mapping[fmt.Sprintf("templates/ui/%s/page.tsx.tmpl", strings.ToLower(proj.Ui.String()))] = fmt.Sprintf("%s/app/page.tsx", proj.ProjName)

				switch proj.Database {
				case MongoDB:
					mapping["templates/backend/mongo/.env.tmpl"] = fmt.Sprintf("%s/.env", proj.ProjName)
					mapping["templates/backend/mongo/database.js.tmpl"] = fmt.Sprintf("%s/database.js", proj.ProjName)
				case SQLite:
					mapping["templates/backend/sqlite/database.db.tmpl"] = fmt.Sprintf("%s/database/database.db", proj.ProjName)
					mapping["templates/backend/sqlite/statements.js.tmpl"] = fmt.Sprintf("%s/database/statements.js", proj.ProjName)
					mapping["templates/backend/sqlite/statements.sql.tmpl"] = fmt.Sprintf("%s/database/statements.sql", proj.ProjName)
				}

			case ReactRouter:
				mapping[fmt.Sprintf("templates/ui/%s/home.tsx.tmpl", strings.ToLower(proj.Ui.String()))] = fmt.Sprintf("%s/app/routes/home.tsx", proj.ProjName)

			case ExpressJs:
				mapping[fmt.Sprintf("templates/ui/%s/App.js.tmpl", strings.ToLower(proj.Ui.String()))] = fmt.Sprintf("client/app/App.js")
				mapping["templates/backend/express/index.js.tmpl"] = "server/index.js"
				mapping["templates/backend/express/package.json.tmpl"] = "server/package.json"

			default:
			}

		} else {

			switch proj.BackFrame {

			case NodeJS:

			case ExpressJs:

			default:

			}

		}

	case ClientServer:
	case Terminal:
	case Mobile:
	default:

	}

	return mapping
}

func makeContext(proj *Project) (context Context) {
	fmt.Println("Building Context")

	context = Context{
		ProjectName: proj.ProjName,
		Force:       true,
		Template:    "",
	}

	return context
}
