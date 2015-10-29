not.run <- function() {
  
require(rmq)
require(testthat)
require(data.table)

a= as.data.table(6:10)

m=to.msgpack(a)
m

attributes(a)

b=from.msgpack(m)
b

}
