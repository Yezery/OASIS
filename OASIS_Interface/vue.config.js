const { defineConfig } = require('@vue/cli-service')
const NodePolyfillPlugin = require('node-polyfill-webpack-plugin')
module.exports = defineConfig({
  transpileDependencies: true,
  devServer: {
    open: false, // 自动打开浏览器
    port: 8085,
  },
  configureWebpack: {
    plugins: [new NodePolyfillPlugin()],
  },
  chainWebpack: (config) => {
    const oneOfsMap = config.module.rule('scss').oneOfs.store
    oneOfsMap.forEach((item) => {
      item
        .use('sass-resources-loader')
        .loader('sass-resources-loader')
        .options({
          resources: './src/style/index.scss',
        })
        .end()
    })
  },
})
