package internal

import (
	"fmt"
	"os/exec"
)

func BuildProject(proj *Project) {

	switch proj.ProjType {
	case WebApp:
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

	if proj.FrontFrame == React {

		switch proj.BackFrame {

		//Use npx to make, then replace page.tsx
		case NextJS:
			selections := MakeTemplatePaths(proj)
			exec.Command("npx", "create-next-app@latest")

			//FIX THIS ONCE FUNCTION IS UPDATED!
			GenerateManifest(selections, "./")

			var context = makeContext(proj)

			GenerateProject(context)

		//Use npx to make, then replace page.tsx
		case ReactRouter:
			selections := MakeTemplatePaths(proj)
			exec.Command("npx", "create-react-router@latest", "--template", "remix-run/react-router-templates/minimal")

			//FIX THIS ONCE FUNCTION IS UPDATED!
			GenerateManifest(selections, "./")

			var context = makeContext(proj)

			GenerateProject(context)

		//Use npx to make react client, then make server with manifest.json
		case ExpressJs:
			selections := MakeTemplatePaths(proj)
			exec.Command("npx", "create-react-app", "client")

			//FIX THIS ONCE FUNCTION IS UPDATED!
			GenerateManifest(selections, "./")

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

func MakeTemplatePaths(proj *Project) (mapping map[string]string) {

	switch proj.ProjType {
	case WebApp:
		if proj.FrontFrame == React {

			switch proj.BackFrame {

			case NextJS:
				mapping[fmt.Sprintf("templates/ui/%s/page.tsx.tpml", proj.Ui)] = fmt.Sprintf("%s/app/page.tsx.tmpl", proj.ProjName)

			case ReactRouter:
				mapping[fmt.Sprintf("templates/ui/%s/home.tsx.tpml", proj.Ui)] = fmt.Sprintf("%s/app/home.tsx.tmpl", proj.ProjName)

			case ExpressJs:
				mapping[fmt.Sprintf("templates/ui/%s/App.js.tpml", proj.Ui)] = fmt.Sprintf("%s/app/App.js.tmpl", proj.ProjName)
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

}
