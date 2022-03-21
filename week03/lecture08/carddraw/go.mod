module carddraw

go 1.17

require (
	cardgame v1.2.3
	errors v1.2.3
)

replace (
	cardgame v1.2.3 => ../cardgame
)