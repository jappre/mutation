var gulp = require('gulp');

var jshint = require('gulp-jshint');
var sass = require('gulp-sass');
var concat = require('gulp-concat');
var uglify = require('gulp-uglify');
var rename = require('gulp-rename');

gulp.task('lint', function(){
    gulp.src('./js/*.js')
        .pipe(jshint())
        .pipe(jshint.reporter('default'));
});

gulp.task('sass', function() {
    gulp.src('./scss/*.scss')
        .pipe(sass())
        .pipe(gulp.dest('./css'));//dest()写入文件
});

gulp.task('scripts', function() {
    gulp.src('./js/*.js')
        .pipe(concat('all.js'))
        .pipe(gulp.dest('./dist'))
        .pipe(rename('all.min.js'))
        .pipe(uglify())
        .pipe(gulp.dest('./dist'));
        console.log('gulp task is done');//自定义提醒信息
});

gulp.watch('./js/*.js', function(){
    gulp.run('lint', 'scripts');
});

gulp.task("default", ["clean"], () => {
        gulp.start("lint", "sass", "scripts");
});
