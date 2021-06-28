const path = require('path');
const resolve = (dir) => path.join(__dirname, dir); // 给public路径添加别名
 
module.exports = {
  publicPath: '/', // base目录，等同于router.js的base字段
  assetsDir: "static", // 打包后静态资源目录，注意public文件下目录（别名）配置，index.html的icon路径
  devServer: {
    open: true,
    host:"localhost",
    port: 8080,
    https: false,
    proxy: {
      "/apis": {
        target: "http://127.0.0.1:9000", //目标主机
        ws: true, //代理的WebSockets
        changeOrigin: true, //需要虚拟主机站点
        pathRewrite: {
          "^/apis": ""
        }
      }
    },
    before: app => {
    }
  },
  productionSourceMap: false, // 生产环境map文件
  chainWebpack: config => {
    // 添加别名（src默认为@，不用再次添加）
    config.resolve.alias
      .set('@pub', resolve('public')) // 设置public别名为@pub
  },
  configureWebpack: config => {
    if (process.env.NODE_ENV === 'production') {
      // 为生产环境修改配置...
    } else {
      // 为开发环境修改配置...
    }
  }
}