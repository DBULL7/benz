package new 

// Vuex creates a multi page application using vue-router.
func Vuex(name string) {
	e2e 		:= E2e()
	backend := Backend()
	if backend == true {
		BeOnly(name, "html")
	}

	Mkdir(name + "/src")
	CreateFile("../files/frontend/vuex/index.js", name + "/src/index.js")		
	Mkdir(name + "/src/components")
	CreateFile("../files/frontend/vuex/App.vue", name + "/src/components/App.vue")
	CreateFile("../files/frontend/vuex/Home.vue", name + "/src/components/Home.vue")
	CreateFile("../files/frontend/vuex/router.js", name + "/src/router.js")	
	Mkdir(name + "/src/store")
	CreateFile("../files/frontend/vuex/storeIndex.js", name + "/src/store/index.js")		
	CreateFile("../files/frontend/vuex/actions.js", name + "/src/store/actions.js")	
	CreateFile("../files/frontend/vuex/mutations.js", name + "/src/store/mutations.js")	

	Mkdir(name + "/dist")
	CreateFile("../files/common/.gitignore", name + "/.gitignore")
	CreateFile("../files/common/README.md", name + "/README.md")
	CreateFile("../files/common/vueBabel", name + "/.babelrc")
	CreateFile("../files/common/packageJSON.tmpl", name + "/package.json") 
	WebpackDevServer(backend, name)	
	devdeps := []string{"add","vue", "vuex","vue-loader","vue-router", "vue-template-compiler", "webpack", "webpack-cli", "babel-preset-env","babel-loader","babel-core","css-loader","style-loader","sass-loader","postcss-loader","postcss-import","postcss-cssnext","cssnano", "node-sass","mini-css-extract-plugin", "--dev"}
	InstallJS(devdeps, name)

	// webpack and html file 
	CreateFile("../files/webpack/vuex.js", name + "/webpack.config.js")
	CreateFile("../files/frontend/react/postcss.config.js", name + "/postcss.config.js")
	CreateFile("../files/frontend/react/index.html", name + "/dist/index.html")
	//create e2e testing
	InstallE2E(e2e, name)
}
