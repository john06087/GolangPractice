module example.com/hello

go 1.17

require (
	abc.com/GoProject2 v0.0.0-00010101000000-000000000000
	rsc.io/quote v1.5.2
)

require (
	go.uber.org/atomic v1.9.0 // indirect
	go.uber.org/multierr v1.8.0 // indirect
	go.uber.org/zap v1.21.0 // indirect
	golang.org/x/text v0.3.3 // indirect
	rsc.io/sampler v1.3.0 // indirect
)

replace abc.com/GoProject2 => ../GoProject2
