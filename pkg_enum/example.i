/* File : example.i */
%module pkg_enum

%{
#include "example.h"
%}

/* Let's just grab the original header file here */

%include "example.h"

