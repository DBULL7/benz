package new

import (
	// "fmt"
)

// ReactSPA entry point for building create-react-app style applications
func ReactSPA(name string) {
	e2e := E2e()
	if backend := Backend(); backend == true {
		BeOnly(name, "html")
	} else {
		// create html file in root of project. 
		// package json needs to have command for webpack dev server with hot reload 
	}

	Mkdir(name + "/src")
	Mkdir(name + "/src/App")
	Mkdir(name + "/dist")
	CreateFile("../files/common/.gitignore", name + "/.gitignore")
	CreateFile("../files/common/README.md", name + "/README.md")
	CreateFile("../files/common/reactBabel", name + "/.babelrc")
	CreateFile("../files/common/jestPackageJSON.tmpl", name + "/package.json")
	deps := []string{"add", "react", "react-dom"}
	InstallJS(deps, name)
	devdeps := []string{"add","webpack@3.10.0", "webpack-cli", "babel-preset-env","babel-loader","babel-core","babel-preset-react","css-loader","style-loader","sass-loader","postcss-loader","postcss-import","postcss-cssnext","cssnano", "jest","enzyme","enzyme-adapter-react-16", "identity-obj-proxy", "--dev"}
	InstallJS(devdeps, name)
	CreateFile("../files/frontend/react/reactIndex.js", name + "/src/index.js")
	CreateFile("../files/frontend/react/reactApp.js", name + "/src/App/App.js")
	CreateFile("../files/frontend/react/reactCSS.css", name + "/src/App/App.css")
	CreateFile("../files/webpack/react.js", name + "/webpack.config.js")
	// webpack.prod.js
	CreateFile("../files/frontend/react/postcss.config.js", name + "/postcss.config.js")
	CreateFile("../files/frontend/react/App.spec.js", name + "/src/App/App.spec.js")
	CreateFile("../files/frontend/react/index.html", name + "/dist/index.html")

	// create e2e testing
	InstallE2E(e2e, name)
}

