var path = require("path");
var webpack = require('webpack');
var ExtractTextPlugin = require("extract-text-webpack-plugin");

var bower_dir = __dirname + '/bower_components';
var node_dir = __dirname + '/node_modules';

var config = {
    addVendor: function (name, path) {
        this.resolve.alias[name] = path;
        this.module.noParse.push(new RegExp('^' + name + '$'));
    },

    entry: {
        app: './public/javascripts/app.js',
    vendors: ['jquery','underscore']
    },

    output: {
        path: path.join(__dirname, process.env.NODE_ENV === 'production' ? 'public/dist' : 'public/dev'),
        filename: process.env.NODE_ENV === 'production' ? '[name].[hash].js' : '[name].js',
        publicPath: 'public/dev'
    },

    module:{
        noParse: [],
        loaders:[
            {test: /\.js$/,exclude: /node_modules/, loader:'babel?presets[]=es2015'},
            {test: /\.css$/,loader:ExtractTextPlugin.extract("style-loader", "css-loader")},
            {test: /\.(jpg|png)$/,loader: "url?limit=8192"},
            {test: /\.scss$/,loader:ExtractTextPlugin.extract("style-loader", "css-loader!sass-loader")}
        ]
    },


    resolve: { alias: {} },

    plugins: [
        new ExtractTextPlugin("charity.css"),
        new webpack.optimize.CommonsChunkPlugin('vendors', 'vendors.bundle.js'),
        new webpack.ProvidePlugin({
            $: "jquery",
            jQuery: "jquery",
            "window.jQuery": "jquery"
        }),
        new webpack.ProvidePlugin({
          _: "underscore",
        })
    ]
};

config.addVendor('jquery', node_dir + '/jquery/dist/jquery.js');
config.addVendor('underscore', node_dir + '/underscore/underscore.js');

module.exports = config;