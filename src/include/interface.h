// -*- mode: C++; c-indent-level: 2; c-basic-offset: 2; tab-width: 8 -*-
///////////////////////////////////////////////////////////////////////////
// Copyright (C) 2014 Jason E. Aten                                      //
// License: Apache 2.0.                                                  //
// http://www.apache.org/licenses/                                       //
///////////////////////////////////////////////////////////////////////////


#ifndef INTERFACE_HPP
#define INTERFACE_HPP

#include <Rinternals.h>

#ifdef __cplusplus
extern "C" {
#endif


SEXP rmq(SEXP name_);

  int JasonsLinkeMe();

#ifdef __cplusplus
}
#endif


#endif // INTERFACE_HPP