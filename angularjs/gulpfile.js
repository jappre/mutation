var gulp = require('gulp'),
 sass = require('gulp-ruby-sass'),
 autoprefixer = require('gulp-autoprefixer');
gulp.task('styles', function() {
  return gulp.src('src/styles/main.scss')
	.pipe(sass({style: 'expanded'}))
	.pipe(autoprefixer('last 2 version', 'safari 5', 'ie 9'))
	.pipe(gulp.dest('dist/assets/css'))
	.pipe(rename({suffix: '.min'}))
	.pipe(minifycss())
	.pipe(gulp.dest('dist/assets/css'))
	.pipe(notify({ message: 'Styles task complete' }));
});
