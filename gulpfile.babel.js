'use strict';

import gulp from 'gulp';
import babel from 'gulp-babel';

var plugins = require("gulp-load-plugins")();

function adminTasks() {
  var pathto = function (file) {
        return ('./public/' + file);
      };
  var scripts = {
        src: pathto('javascripts/*.js'),
        dest: pathto('dist')
      };
  var styles = {
        src: pathto('stylesheets/qor.scss'),
        scss: pathto('stylesheets/**/*.scss'),
        dest: pathto('dist')
      }

  gulp.task('js', function () {
    return gulp.src(scripts.src)
    .pipe(babel())
    .pipe(plugins.concat('app.js'))
    .pipe(plugins.uglify())
    .pipe(gulp.dest(scripts.dest));
  });

  gulp.task('sass', function () {
    return gulp.src(styles.src)
    .pipe(plugins.sourcemaps.init())
    .pipe(plugins.sass())
    .pipe(plugins.csscomb())
    .pipe(plugins.minifyCss())
    .pipe(plugins.sourcemaps.write('./'))
    .pipe(gulp.dest(styles.dest));
  });

  gulp.task('watch', function () {
    gulp.watch(scripts.src, ['js']);
    gulp.watch(styles.scss, ['sass']);
  });

  gulp.task('default', ['watch']);
}


// Init
// -----------------------------------------------------------------------------

console.log('Running "qor-example" module task...');
adminTasks();


// Task for compress js vendor assets
gulp.task('compressJavaScriptVendor', function () {
  return gulp.src(['!./public/vendors/jquery.js','./public/vendors/*.js'])
  .pipe(plugins.concat('vendors.js'))
  .pipe(gulp.dest('./public/dist'));
});
