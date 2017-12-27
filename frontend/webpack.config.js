const ExtractTextPlugin = require('extract-text-webpack-plugin');

module.exports = {
    entry: ['./static/stylesheets/main.scss'],
    output: {
        filename: 'dist/bundle.js'
    },
    module: {
        rules: [
            {
                test: /\.scss$/,
                use: ExtractTextPlugin.extract({
                    fallback: 'style-loader',
                    //resolve-url-loader may be chained before sass-loader if necessary
                    use: ['css-loader', 'sass-loader']
                })
            }
        ]
    },
    plugins: [
        new ExtractTextPlugin({
            filename: './static/stylesheets/main.css',
            allChunks: true
        })
        //if you want to pass in options, you can do so:
        //new ExtractTextPlugin({
        //  filename: 'style.css'
        //})
    ]
};
