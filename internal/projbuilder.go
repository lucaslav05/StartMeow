package internal

func buildProject(proj *Project) {

	switch proj.projType {
	case WebApp:
		buildWebApp(proj)
	case ClientServer:
		buildClientServer(proj)
	case Terminal:
		buildTerminal(proj)
	case Mobile:
		buildMobile(proj)
	default:

	}

}

func buildWebApp(proj *Project) {

	if proj.frontFrame == React {

		switch proj.backFrame {
		case NextJS:
		case ReactRouter:
		case ExpressJs:
		default:
		}

	} else {

	}

}

func buildClientServer(proj *Project) {

}

func buildTerminal(proj *Project) {

}

func buildMobile(proj *Project) {

}
