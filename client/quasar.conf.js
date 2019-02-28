module.exports = function (ctx) {
  return {
    plugins: [
      'axios'
    ],
    css: [
      'app.styl'
    ],
    extras: [
      ctx.theme.mat ? 'roboto-font' : null,
      'material-icons',
      'mdi',
      'fontawesome'
    ],
    supportIE: false,
    build: {
      scopeHoisting: true,
      vueRouterMode: 'history',
      publicPath: '/',
      showProgress: true,
      extendWebpack (cfg) {
        cfg.module.rules.push({
          enforce: 'pre',
          test: /\.(js|vue)$/,
          loader: 'eslint-loader',
          exclude: /node_modules/
        })
      },
      env: ctx.dev
        ? {
          API: JSON.stringify('http://localhost:8000')
        }
        : {
          API: JSON.stringify('http://localhost:8000')
        }
    },
    devServer: {
      port: 7000,
      open: true
    },
    framework: {
      components: [
        'QLayout',
        'QLayoutHeader',
        'QLayoutDrawer',
        'QPageContainer',
        'QPage',
        'QToolbar',
        'QToolbarTitle',
        'QBtn',
        'QIcon',
        'QList',
        'QListHeader',
        'QItem',
        'QItemMain',
        'QItemSide',
        'QTable',
        'QCard',
        'QCardTitle',
        'QCardMain',
        'QCardMedia',
        'QCardSeparator',
        'QCardActions',
        'QInput'
      ],
      directives: [
        'Ripple',
        'CloseOverlay'
      ],
      plugins: [
        'Notify',
        'Loading',
        'LocalStorage',
        'Dialog',
        'Screen'
      ],
      iconSet: ctx.theme.mat ? 'material-icons' : 'ionicons',
      i18n: 'pt-br'
    },
    animations: 'all',
    ssr: {
      pwa: false
    },
    pwa: {},
    cordova: {},
    electron: {}
  }
}
