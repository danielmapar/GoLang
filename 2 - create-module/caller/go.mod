module example.com/caller

go 1.16

replace example.com/greetings => ../module

require example.com/greetings v0.0.0-00010101000000-000000000000
