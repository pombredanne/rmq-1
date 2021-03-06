% Generated by roxygen2: do not edit by hand
% Please edit documentation in R/rmq.R
\name{read.tmframe}
\alias{read.tmframe}
\title{read TMFRAME formatted timeseries}
\usage{
read.tmframe(path)
}
\arguments{
\item{path}{A string holding the filesystem path to the TMFRAME file.}
}
\description{
Given a filesystem path to a TMFRAME file, \code{read.tmframe} will return an R dataframe with the series.
}
\seealso{
Other rmq.functions: \code{\link{from.msgpack}},
  \code{\link{r2r.call}}, \code{\link{r2r.server}},
  \code{\link{read.msgpack.frame}},
  \code{\link{read.ndjson}}, \code{\link{rmq.call}},
  \code{\link{rmq.default.addr}}, \code{\link{rmq.server}},
  \code{\link{to.msgpack}}
}

