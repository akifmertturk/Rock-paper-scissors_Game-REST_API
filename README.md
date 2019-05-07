# Rock-paper-scissors_Game-REST_API
A game, REST API by using Go language


Rock–paper–scissors one of the most well-known
hand game which usually played between two people,
in which each player simultaneously forms one of
three shapes with an outstretched hand.

I created the Rock-Paper-Scissors (RPS) game 
which can be played with an HTTP API by using
Go language. This project is the first usage 
of Go language for me. I created this REST API 
because of applying an internship for an company.


* Starts a new game and responses with session ID which will be used later to play.
-> http://localhost:8080/newGame?round=3

* The user sends his/her decision to the server with this request.
-> http://localhost:8080/play?choose=rock&id=23

***** EXAMPLE USAGE *****

$ curl "http://localhost:8080/newGame?round=2"

New Rock-Paper-Scissors game started
Session ID = 23

to play Rock: http://localhost:8080/play?choose=rock&id=23
to play Paper: http://localhost:8080/play?choose=paper&id=23
to play Scissors: http://localhost:8080/play?choose=scissors&id=23

$ curl "http://localhost:8080/play?id=23&choose=paper"

-> ROUND 1

me: ROCK
you: PAPER

YOU WON THIS ROUND!!

There is 1 more round.
$ curl "http://localhost:8080/play?id=23&choose=rock"

-> ROUND 2

me: SCISSORS
you: ROCK

YOU WON THIS ROUND!!

-> GAME COMPLETED
Round 1: Rock vs Paper 
Round 2: Scissors vs Rock 

0 vs 2
YOU WON !!
