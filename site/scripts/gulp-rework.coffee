
sysPath = require 'path'
through = require 'through2'
cssWhitespaceCompiler = require 'css-whitespace'

rework = require 'rework'
rework_vars = require 'rework-vars'
rework_calc = require 'rework-calc'
rework_shade = require 'rework-shade'
rework_import = require 'rework-importer'
rework_clearfix = require 'rework-clearfix'
rework_imagesize = require 'rework-imagesize'

mixins =
  display: (displayType) ->
    result = display: displayType
    return result if displayType isnt 'inline-block'
    result['*display'] = 'inline'
    result

module.exports = ->
  through.obj (file, enc, cb) ->
    if file.isStream()
      @emit 'error', new gutil.PluginError 'gulp-rework', 'Streaming not supported'
      return cb()

    css = file.contents.toString()
    if sysPath.extname(file.path) is '.styl'
      css = cssWhitespaceCompiler css

    file.contents = new Buffer(rework css
      .use rework_import
        path: 'src/'
        base: 'src/'
        whitespace: true
      .use rework_vars()
      .use rework_shade()
      .use rework.mixin mixins
      .use rework.extend()
      .use rework.references()
      .use rework_imagesize 'src/images'
      .use rework_calc
      .use rework_clearfix
      .toString()
    )

    @push file
    cb()

