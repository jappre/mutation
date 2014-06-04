gulp = require 'gulp'
gutil = require 'gulp-util'
gulp_jade = require 'gulp-jade'
gulp_coffee = require 'gulp-coffee'
gulp_rename = require 'gulp-rename'
gulp_replace = require 'gulp-replace'
gulp_rework = require './scripts/gulp-rework'

Q = require 'q'
_ = require 'lodash'
glob = require 'glob'
yaml = require 'js-yaml'
marked = require 'marked'

fs = require 'fs'
sysPath = require 'path'

markdownYamlType = new yaml.Type '!markdown',
  loadKind: 'scalar'
  loadResolver: (state) ->
    return true unless state.result
    state.result = marked state.result
    true

MARKDOWN_SCHEMA = yaml.Schema.create [markdownYamlType]

gulp.task 'asset', ->
  gulp.src([
    'src/**/*.!(jade|yaml|coffee|styl)'
    '!src/images/**/*'
  ])
    .pipe(gulp_replace /{%timestamp%}/g, Date.now())
    .pipe(gulp.dest 'www/')

gulp.task 'jade', ->
  Q.nfcall(glob, 'src/**/*.yaml').then (filelist) ->
    locals = _(filelist).map (filepath) ->
      [filepath, sysPath.basename filepath, '.yaml']
    .map ([filepath, dataname]) ->
      data = yaml.safeLoad(
        fs.readFileSync(filepath).toString()
        schema: MARKDOWN_SCHEMA
      )
      [dataname, data]
    .object()
    .value()

    gulp.src([
      'src/**/*.jade'
      '!src/**/_*.jade'
    ])
      .pipe(gulp_jade({pretty: true, locals}).on 'error', gutil.log)
      .pipe(gulp_replace /{%timestamp%}/g, Date.now())
      .pipe(gulp.dest 'www/')

gulp.task 'image', ->
  gulp.src([
    'src/images/**/*'
    '!src/images/**/_*'
  ])
    .pipe(gulp.dest 'www/images/')

gulp.task 'coffee', ['asset', 'jade'], ->
  gulp.src([
    'src/**/*.coffee'
    '!src/**/_*.coffee'
  ])
    .pipe(gulp_coffee().on 'error', gutil.log)
    .pipe(gulp.dest 'www/')

gulp.task 'rework', ['asset', 'jade'], ->
  gulp.src([
    'src/**/*.+(styl|css)'
    '!src/**/_*.styl'
  ])
    .pipe(gulp_rework().on 'error', gutil.log)
    .pipe(gulp_rename extname: '.css')
    .pipe(gulp.dest 'www/')

gulp.task 'watch', 'image coffee rework jade asset'.split(' '), ->
  gulp.watch 'src/images/**/*', ['image']
  gulp.watch 'src/**/*.coffee', ['coffee']
  gulp.watch 'src/**/*.+(styl|css)', ['rework']
  gulp.watch 'src/**/*.+(jade|yaml)', ['jade']
  gulp.watch ['src/**/*.!(jade|yaml|coffee|styl)'], ['asset']

  return

