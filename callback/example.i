/* File : example.i */
%module(directors="1") callback
%{
#include "example.h"
%}

/* turn on director wrapping Callback */
%feature("director") Callback;

%include "example.h"

