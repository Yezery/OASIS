const { defineConfig } = require('@vue/cli-service')
module.exports = defineConfig({
  transpileDependencies: true,

  devServer: {
    open: false, // 自动打开浏览器
    port: 8085,
}
})
