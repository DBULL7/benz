package new 

// Redux creates multi page React projects using React, Redux, and React Router
func Redux(name string) {
	e2e 		:= E2e()
	backend := Backend()
	if backend == true {
		BeOnly(name, "html")
	}

	Mkdir(name + "/src")
	Mkdir(name + "/src/components")
	Mkdir(name + "/src/components/App")
	Mkdir(name + "/src/components/Home")
	Mkdir(name + "/src/components/PageNotFound")
	Mkdir(name + "/src/actions")
	Mkdir(name + "/src/reducers")
	Mkdir(name + "/dist")
	CreateFile("../files/common/.gitignore", name + "/.gitignore")
	CreateFile("../files/common/README.md", name + "/README.md")
	CreateFile("../files/common/reactBabel", name + "/.babelrc")
	CreateFile("../files/common/jestPackageJSON.tmpl", name + "/package.json")
	WebpackDevServer(backend, name)
	devdeps := []string{"add","react", "react-dom", "react-redux","react-router-dom", "react-router", "redux","webpack", "webpack-cli", "babel-preset-env","babel-loader","babel-core","babel-preset-react","css-loader","style-loader","sass-loader","postcss-loader","postcss-import","postcss-cssnext","cssnano", "jest","enzyme","enzyme-adapter-react-16", "identity-obj-proxy","node-sass", "--dev"}
	InstallJS(devdeps, name)
  // frontend files 
	CreateFile("../files/frontend/redux/index.js", name + "/src/index.js")
	CreateFile("../files/frontend/redux/App.js", name + "/src/components/App/App.js")
	CreateFile("../files/frontend/redux/AppContainer.js", name + "/src/components/App/AppContainer.js")
	CreateFile("../files/frontend/redux/Home.js", name + "/src/components/Home/Home.js")
	CreateFile("../files/frontend/redux/HomeContainer.js", name + "/src/components/Home/HomeContainer.js")
	CreateFile("../files/frontend/redux/Home.css", name + "/src/components/Home/Home.css")
	CreateFile("../files/frontend/redux/Home.spec.js", name + "/src/components/Home/Home.spec.js")
	CreateFile("../files/frontend/redux/PageNotFound.js", name + "/src/components/PageNotFound/PageNotFound.js")
	// actions, reducers, and redux store configuration
	CreateFile("../files/frontend/redux/configStore.js", name + "/src/configStore.js")
	CreateFile("../files/frontend/redux/rootReducer.js", name + "/src/reducers/rootReducer.js")
	CreateFile("../files/frontend/redux/actions.js", name + "/src/actions/index.js")

	CreateFile("../files/webpack/react.js", name + "/webpack.config.js")
	CreateFile("../files/frontend/react/postcss.config.js", name + "/postcss.config.js")
	CreateFile("../files/frontend/react/index.html", name + "/dist/index.html")

	// create e2e testing
	InstallE2E(e2e, name)
}