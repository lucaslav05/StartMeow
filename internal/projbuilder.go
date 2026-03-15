package internal

func buildProject(proj *Project) {

	switch proj.projType {
	case WebApp:
		buildWebApp(proj)
	case ClientServer:

	case Terminal:

	case Mobile:

	default:

	}

}

func buildWebApp(proj *Project) {

}

func buildClientServer(proj *Project) {

}

func buildTerminal(proj *Project) {

}

func build(proj *Project) {

}
