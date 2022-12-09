module go_training/hello

go 1.19

require (
	go_training/greetings v1.0.0
	rsc.io/quote v1.5.2
	golang.org/x/text v0.0.0-20170915032832-14c0d48ead0c // indirect
	rsc.io/sampler v1.3.0 // indirect
)

replace go_training/greetings => ../greetings
