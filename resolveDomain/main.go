package main

/*
	Goal:- Resolving domain names using Golang Channels
		1. A function that resolves the domain name.
		2. instead of going from the list of domains to the function that resolves them to the output feed the domains into a channel
		2. start multiple processing goroutines. They can each pull from the channel.
*/
