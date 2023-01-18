module go_training/hello

go 1.19

require (
	go_training/greetings v1.0.0
	rsc.io/quote v1.5.2
)

require (
	golang.org/x/text v0.6.0 // indirect
	rsc.io/sampler v1.99.99 // indirect
)

replace go_training/greetings => ../greetings
