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
});


