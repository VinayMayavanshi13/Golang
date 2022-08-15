package main

func main() {

}

/*
Memory allocation and de-allocation happens automatically
new()
		allocate memory but no INIT (doesn't initialize)
		you will get a memory address
		zeroed storage

make()
		allocate memory and INIT (initialize)
		you will get memory address
		non-zeroed storage

Garbage Collections happens automatically,anything which becomes Out Of Scope or nil comes under garbage collection.
when initially Go lang was wrote the Garbage Collector was not performing well then the Go developers realised that
re-wrote the entire garbage collector then the efficiency sky-rocketed.

The GOGC variable sets the initial garbage collection target percentage.A collection is triggered when the ratio of freshly
allocated data to live data remaining  after the previous collection reaches this percentage.
The default is GOGC is 100. Setting this to off disables the garbage collector entirely. The runtime/debug package's SetGCPercent
function allows changing this percentage at run time.

check out https://pkg.go.dev/runtime

*/
