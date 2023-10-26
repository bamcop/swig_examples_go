/* File : example.i */
%module(directors="1") pkg_director

%include "std_string.i"

%header %{
#include "director.h"
%}

%feature("director") FooBarAbstract;
%include "director.h"
