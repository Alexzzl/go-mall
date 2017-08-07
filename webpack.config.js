const webpack = require('webpack');
const path = require('path');
const HtmlWebpackPlugin = require('html-webpack-plugin');
const OpenBrowserPlugin = require('open-browser-webpack-plugin');
// const casProxy = require('./proxy');
module.exports = {
    entry: {
        index: './app/index.jsx',
        good_detail: './app/good_detail.js',
        shopcart: "./app/shopcart.js",
        order_index: "./app/orderIndex.js",
        vendor: [
            'react', 'classnames', 'react-router', 'react-dom',
        ],
    },
    output: {
        // path: path.resolve(__dirname, 'dist'),
        // filename: './vendor.js',
        filename: '[name].js',
        path: path.join(__dirname, 'dist'),
        chunkFilename: '[name].js'
    },
    resolve: {
        extensions: [' ', '.js', '.json'],
        alias: {
            components: __dirname + '/app/components',
            actions: __dirname + '/app/actions',
            api: __dirname + '/app/api',
            reducers: __dirname + '/app/reducers',
            utils: __dirname + '/app/utils',
            constants: __dirname + '/app/constants',
            style: __dirname + '/app/style',
        },
    },
    module: {
        loaders: [
            {
                test: /\.js[x]?$/,
                exclude: /node_modules/,
                use: [{
                    loader: 'react-hot-loader'
                }, {
                    loader: 'babel-loader',
                    options: {
                        babelrc: false,
                        presets: [
                            'es2015',
                            'stage-0',
                            'react'
                        ],
                        plugins: [
                            ['import', {libraryName: 'antd', style: 'css'}],
                        ]
                    }
                }]
            },
            {
                test: /\.less$/,
                loader: 'style!css!postcss!less',
            },
            {
                test: /\.css/,
                loader: 'style-loader!css-loader',
            },
            {
                test: /\.(png|jpg)$/,
                loader: 'url-loader?limit=8192',
            },
        ],
    },

    devServer: {
        hot: true,
        inline: true
    },
    plugins: [
        new webpack.NoErrorsPlugin(),
        new webpack.HotModuleReplacementPlugin()
    ]
};
