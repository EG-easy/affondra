// const path = require('path');

module.exports = {
  devServer: {
    disableHostCheck: true,
    host:'0.0.0.0'
  },
  "transpileDependencies": [
    "vuetify"
  ],
  // configureWebpack: {
  //   module:{
  //     rules:[
  //       {
  //         test: /npm\.js$/,
  //         loader: 'string-replace-loader',
  //         include: path.resolve('node_modules/firebaseui/dist'),
  //         query: {
  //           search: 'require(\'firebase/app\');',
  //           replace: 'require(\'firebase/app\').default;',
  //         },
  //       },
  //     ]
  //   }
  // }
}