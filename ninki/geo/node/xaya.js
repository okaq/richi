console.log("so funning");

/* 
 * default config setup
 *
requirejs.config({
    baseUrl: 'node',
    paths: {
        node: './'
    }
});

requirejs(['node/zaya']);

*/

define(function(require) {
    var regl = require('./regl');
    console.log(regl);
    var mat = require('./gl-matrix');
    console.log(mat);
    var cam = require('./canvas-orbit-camera');
    console.log(cam);
    var fit = require('./canvas-fit');
    console.log(fit);

    console.log("dependencies loaded, ready to app commence");
});


