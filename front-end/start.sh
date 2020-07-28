#!/bin/bash

if $IS_LIVE ; then
    echo 'Running production'
    rm -rf node_modules
    rm -rf __sapper__
    npm install
    npm run build
    npm start
else
   echo 'Running development'
   npm install
   export HOST=0.0.0.0
   npm run dev --host
fi