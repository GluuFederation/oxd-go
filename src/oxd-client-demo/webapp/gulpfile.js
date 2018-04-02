var gulp = require('gulp');
var concat = require('gulp-concat');
var watch = require('gulp-watch');

//script paths
var jsFiles = 'js/**/*.js',
    jsDest = '';

gulp.task('build', function() {
    return gulp.src(jsFiles)
        .pipe(concat('demo.js'))
        .pipe(gulp.dest(jsDest));
});

gulp.task('watch', function () {
    gulp.watch("js/**/*.js", ['build']);
});