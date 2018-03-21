package new 

import (
	// "fmt"
)

// VueSPA single page app entry point.
func VueSPA(name string) {
	e2e 		:= E2e()
	backend := Backend()
	if backend == true {
		BeOnly(name, "html")
	}

	Mkdir(name + "/src")
	Mkdir(name + "/dist")
	CreateFile("../files/common/.gitignore", name + "/.gitignore")
	CreateFile("../files/common/README.md", name + "/README.md")
	CreateFile("../files/common/vueBabel", name + "/.babelrc")
	// package json
	CreateFile("../files/common/packageJSON.tmpl", name + "/package.json") 
	WebpackDevServer(backend, name)
	devdeps := []string{"add","vue", "vue-loader", "vue-template-compiler", "webpack", "webpack-cli", "babel-preset-env","babel-loader","babel-core","css-loader","style-loader","sass-loader","postcss-loader","postcss-import","postcss-cssnext","cssnano", "node-sass", "--dev"}
	InstallJS(devdeps, name)
	CreateFile("../files/frontend/vue/index.js", name + "/src/index.js")
	CreateFile("../files/frontend/vue/App.vue", name + "/src/App.vue")
	CreateFile("../files/webpack/vue.js", name + "/webpack.config.js")
	CreateFile("../files/frontend/react/postcss.config.js", name + "/postcss.config.js")
	CreateFile("../files/frontend/react/index.html", name + "/dist/index.html")
	//create e2e testing
	InstallE2E(e2e, name)
}