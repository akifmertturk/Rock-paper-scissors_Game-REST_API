# Rock-paper-scissors_Game-REST_API
A game, REST API by using Go language
Author: Miraç Akif Merttürk

Rock–paper–scissors one of the most well-known
hand game which usually played between two people,
in which each player simultaneously forms one of
three shapes with an outstretched hand.

I created the Rock-Paper-Scissors (RPS) game 
which can be played with an HTTP API by using
Go language. This project is the first usage 
of Go language for me. I created this REST API 
because of applying an internship for an company.

In the game, I decided to use random number 
generator to  decide the steps of the opponent one.
When you play, in deed you play with a computer
( random number generator ).

* Starts a new game and responses with session ID which will be used later to play.
-> http://localhost:8080/newGame?round=3

* The user sends his/her decision to the server with this request.
-> http://localhost:8080/play?choose=rock&id=23

The player needs to choose what to play by using "choose=rock" comment.
to play Rock: http://localhost:8080/play?choose=rock&id=23
to play Paper: http://localhost:8080/play?choose=paper&id=23
to play Scissors: http://localhost:8080/play?choose=scissors&id=23

You can find an example usage of the game on:
https://github.com/akifmertturk/Rock-paper-scissors_Game-REST_API.wiki.git.
