const path = require('path')

module.exports = {
  entry: './src/index.js',
  output: {
    filename: 'bundle.js',
    path: path.resolve(__dirname, 'dist')
  },
  module: {
    rules: [
      {
        test: /\.js$/,
        use: "babel-loader",
        exclude: /node_modules/
      },
      {
        test: /\.css$/,
        use: ['style-loader', 'postcss-loader']
      },
      {
        test: /\.scss$/,
        use: ['style-loader', 'css-loader', 'sass-loader', 'postcss-loader']
      },
      {
        test: /\.vue$/,
        loaders: "vue-loader",
        exclude: /node_modules/
      }
    ]
  },
  resolve: {
    extensions: ['.js', '.css', '.scss'],
    alias: {
      vue: 'vue/dist/vue.js'
    }
  }
}