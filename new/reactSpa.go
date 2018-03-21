package new

// ReactSPA entry point for building create-react-app style applications
func ReactSPA(name string) {
	e2e 		:= E2e()
	backend := Backend()
	if backend == true {
		BeOnly(name, "html")
	}

	Mkdir(name + "/src")
	Mkdir(name + "/src/App")
	Mkdir(name + "/dist")
	CreateFile("../files/common/.gitignore", name + "/.gitignore")
	CreateFile("../files/common/README.md", name + "/README.md")
	CreateFile("../files/common/reactBabel", name + "/.babelrc")
	CreateFile("../files/common/jestPackageJSON.tmpl", name + "/package.json")
	// after package.json created add the start command for non-backend projects
	webpackDevServer(backend, name)
	devdeps := []string{"add","react", "react-dom", "webpack", "webpack-cli", "babel-preset-env","babel-loader","babel-core","babel-preset-react","css-loader","style-loader","sass-loader","postcss-loader","postcss-import","postcss-cssnext","cssnano", "jest","enzyme","enzyme-adapter-react-16", "identity-obj-proxy","node-sass", "--dev"}
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

func webpackDevServer(backend bool, name string) {
	if backend == false {
		AddScriptToPackageJSON("start", "webpack-dev-server --output-public-path=/dist/ --content-base dist/ --inline --hot --open --port 3000 --mode development", name) 
		devServer := []string{"add", "webpack-dev-server", "--dev"}
		InstallJS(devServer, name)
	} 
}

