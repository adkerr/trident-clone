This directory contains the autogenerated Golang bindings for our ONTAP REST API.

This code uses go-swagger documented at https://goswagger.io

To regenerate the bindings, run the script named 'regen' within this folder.
* https://goswagger.io/install.html (to install the swagger command needed for regen)

As generation can take a very long time, we're checking these in instead of running
this regen script within CI on every build.
