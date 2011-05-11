# Web Gradients #
a tool for creating simple and smooth PNG linear gradients with plain HTTP requests.
This version is written in Go. Previous one was written in Python and is available [on Launchpad](http://bazaar.launchpad.net/~lol2fast4u/webgradients/trunk/files).

Fun to use with [Modernizr](http://modernizr.com):

    #something {
      height: 50px
      /* -webkit, -moz ... */
      background: linear-gradient(#123456, #abcdef)
    }
    .no-cssgradients #something {
      background: url(http://webgradients.appspot.com/make?start=123456&end=abcdef&height=50)
    }

