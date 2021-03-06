// http://webpack.github.io/docs/configuration.html
// http://webpack.github.io/docs/webpack-dev-server.html
var app_root = 'src'; // the app root folder: src, src_users, etc
var path = require('path');
var webpack = require('webpack');
module.exports = {
  app_root: app_root, // the app root folder, needed by the other webpack configs
  entry: [
    // http://gaearon.github.io/react-hot-loader/getstarted/
    'webpack-dev-server/client?http://localhost:8080',
    'webpack/hot/only-dev-server',
    __dirname + '/' + app_root + '/main.js',
  ],
  output: {
    path: __dirname + '/public/js',
    publicPath: 'js/',
    filename: 'bundle.js',
  },
  module: {
    preLoaders: [
      {
        test: /\.js$|\.jsx$/,
        exclude: [/node_modules/, /src\/pbuf/],
        loaders: ["eslint"]
      }
    ],

    loaders: [
     {
        test: /\.js$/,
        loader: 'babel',
        exclude: /node_modules/,
     },
     {
        // https://github.com/jtangelder/sass-loader
        test: /\.scss$/,
        loaders: ['style', 'css', 'sass'],
     },
     {
        test: /.*\.(gif|png|jpe?g|svg)$/i,
        loaders: ['file?hash=sha512&digest=hex&name=[hash].[ext]',
                  'image-webpack'],
     },
     {
        test: /\.proto$/,
        use:  'raw-loader'
     }
    ],
  },
  devServer: {
    contentBase: __dirname + '/public',
    proxy: { '/api': 'http://localhost:8081' }
  },
  plugins: [
  ],
  resolve: {
    root: path.resolve('src'),
  },
  devtool: "source-map"
};
