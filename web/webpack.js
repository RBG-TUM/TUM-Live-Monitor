const path = require("path");

module.exports = {
    mode: "production",
    optimization: {
        minimize: true,
        usedExports: true,
    },
    target: "web",
    entry: {
        main: "./src/index.ts",
    },
    module: {
        rules: [
            {
                test: /\.tsx?$/,
                use: "ts-loader",
                exclude: /node_modules/,
            },
        ],
    },
    resolve: {
        extensions: [".ts", ".js"],
    },
    output: {
        filename: "[name].bundle.js",
        path: path.resolve(__dirname, "./build/dist"),
        library: ["[name]"],
        libraryTarget: "umd",
    },
    plugins: [
    ],
};