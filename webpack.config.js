const path = require("path");
const webpack = require('webpack');
const ExtractTextPlugin = require("extract-text-webpack-plugin");

const node_dir = __dirname + '/node_modules';

const sassLoaders = [
  'css-loader',
  'sass-loader?indentedSyntax=sass&includePaths[]=' + path.resolve(__dirname, './public/stylesheets/')
]

const config = {
    addVendor: function (name, path) {
        this.resolve.alias[name] = path;
        this.module.noParse.push(new RegExp('^' + name + '$'));
    },

    entry: {
        app: './public/javascripts/app.js',
    vendors: ['jquery','underscore','swiper']
    },

    output: {
        path: path.join(__dirname, process.env.NODE_ENV === 'production' ? './public/dist' : './public/dev'),
        filename: process.env.NODE_ENV === 'production' ? '[name].[hash].js' : '[name].js',
        publicPath: 'public/dev'
    },

    module:{
        noParse: [],
        loaders:[
            {
                test: /\.css$/,
                loader: 'style-loader!css-loader'
            },
            {
                test: /\.sass/,
                loader: 'style-loader!css-loader!sass-loader?outputStyle=expanded&indentedSyntax'
            },
            {
                test: /\.scss/,
                loader: 'style-loader!css-loader!sass-loader?outputStyle=expanded'
            },
            {
                test: /\.(png|jpg|gif|woff|woff2)$/,
                loader: 'url-loader?limit=8192'
            },
            {
                test: /\.js$/,
                exclude: /node_modules/,
                loader:'babel?presets[]=es2015'
            }
        ]
    },

    resolve: {
        alias: {},
        extensions: ['', '.js', '.scss'],
    },

    plugins: [
        new webpack.NoErrorsPlugin(),
        new webpack.optimize.CommonsChunkPlugin('vendors', 'vendors.bundle.js'),
        new webpack.ProvidePlugin({
            $: "jquery",
            jQuery: "jquery",
            "window.jQuery": "jquery"
        }),
        new webpack.ProvidePlugin({
          _: "underscore"
        })
    ]
};

config.addVendor('jquery', node_dir + '/jquery/dist/jquery.js');
config.addVendor('underscore', node_dir + '/underscore/underscore.js');
config.addVendor('swiper', node_dir + '/swiper/dist/js/swiper.jquery.min.js');

module.exports = config;