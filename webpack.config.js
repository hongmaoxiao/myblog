var path = require('path');
var projectRoot = path.resolve(__dirname, 'app')
var webpack = require('webpack');
var HtmlWebpackPlugin = require('html-webpack-plugin');
var ExtractTextPlugin = require('extract-text-webpack-plugin');

var cssloader = 'style-loader!css-loader?importLoaders=1!postcss-loader';

module.exports = {
  devtool: 'cheap-source-map',
  entry: {
    app: "./app/main.js"
  },
  output: {
    path: path.resolve(__dirname, './dist/static'),
    filename: "[name].js"
  },
  resolve: {
    modules: [
      "node_modules",
      "app",
    ],
    extensions: ['.vue', '.js', '.json', '.css'],
  },
  module: {
    noParse: /es6-promise\.js$/,
    loaders: [
      {
        enforce: 'pre',
        test: /\.(vue|js)$/,
        loader: 'eslint-loader',
        exclude: /node_modules/
      },
      {
        test: /\.vue$/,
        loader: 'vue-loader',
        options:  {
          loaders: {
            js: 'buble-loader?objectAssign=Object.assign',
            css: cssloader,
          }
        }
      },
      {
        test: /\.js$/,
        loader: 'buble-loader',
        exclude: /node_modules/,
        options: {
          objectAssign: 'Object.assign'
        }
      },
      {
        test: /\.css$/,
        loader: cssloader,
      },
      {
        test: /\.(png|jpe?g|gif|svg)(\?.*)?$/,
        use: [
          {
            loader: 'url-loader',
            options: {
              limit: 20000,
              name: '[name].[hash:7].[ext]'
            }
          },
        ]
      },
      {
        test: /\.(woff2?|eot|ttf|otf)(\?.*)?$/,
        use: [
          {
            loader: 'url-loader',
            options: {
              limit: 10000,
              name: '[name].[hash:7].[ext]'
            }
          },
        ]
      }
    ]
  },
  plugins: [
    new webpack.ProvidePlugin({
      //Promise: 'imports-loader?this=>global!exports-loader?global.Promise!es6-promise',
      $: "jquery",
      jQuery: "jquery",
    }),
    new HtmlWebpackPlugin({
        filename: 'index.html',
        template: path.resolve(__dirname, 'app/index.html'),
        inject: true
      }),
    new ExtractTextPlugin("[name]-[hash].css")
  ],
  devServer: {
    contentBase: path.join(__dirname, "dist"),
    proxy: {
      "/api": "http://localhost:3000",
    }
  },
};
