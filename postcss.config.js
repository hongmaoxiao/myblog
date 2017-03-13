var webpack = require('webpack');
module.exports = {
  plugins: [
    require('postcss-import')({
      addDependency: webpack,
      plugins: [
        require("stylelint")({ })
      ]
    }),
    require('precss')(),
    require('autoprefixer')({ browsers: ['last 3 versions']}),
  ]
}
