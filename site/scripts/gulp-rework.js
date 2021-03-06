// Generated by CoffeeScript 1.7.1
(function() {
  var cssWhitespaceCompiler, mixins, rework, rework_calc, rework_clearfix, rework_imagesize, rework_import, rework_shade, rework_vars, sysPath, through;

  sysPath = require('path');

  through = require('through2');

  cssWhitespaceCompiler = require('css-whitespace');

  rework = require('rework');

  rework_vars = require('rework-vars');

  rework_calc = require('rework-calc');

  rework_shade = require('rework-shade');

  rework_import = require('rework-importer');

  rework_clearfix = require('rework-clearfix');

  rework_imagesize = require('rework-imagesize');

  mixins = {
    display: function(displayType) {
      var result;
      result = {
        display: displayType
      };
      if (displayType !== 'inline-block') {
        return result;
      }
      result['*display'] = 'inline';
      return result;
    }
  };

  module.exports = function() {
    return through.obj(function(file, enc, cb) {
      var css;
      if (file.isStream()) {
        this.emit('error', new gutil.PluginError('gulp-rework', 'Streaming not supported'));
        return cb();
      }
      css = file.contents.toString();
      if (sysPath.extname(file.path) === '.styl') {
        css = cssWhitespaceCompiler(css);
      }
      file.contents = new Buffer(rework(css).use(rework_import({
        path: 'src/',
        base: 'src/',
        whitespace: true
      })).use(rework_vars()).use(rework_shade()).use(rework.mixin(mixins)).use(rework.extend()).use(rework.references()).use(rework_imagesize('src/images')).use(rework_calc).use(rework_clearfix).toString());
      this.push(file);
      return cb();
    });
  };

}).call(this);
