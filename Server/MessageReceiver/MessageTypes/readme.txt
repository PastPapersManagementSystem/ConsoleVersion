ACCEPTED REQUEST TYPES:

Server expects client that only can send the following requests:
1. Add: Add the requester to the lobby.
2. NextMove: Will change the state of the game.
3. Create: Will create a new game lobby and session and wait for all members to join.
4. End: Will dissolve the game session and all the players playing.
5. Quit: A player will quit from the game session, but the game session will exist.

GAME LOBBY:

This component will be responsible for lobby actions.
1. Create: Will create a new game lobby and session and wait for all members to join.
2. Start: This is the only point and function that can initiate the game.
3. Add: Add to lobby.
4. End: End the game.
5. Quit: The player will leave the lobby.

GAME ENGINE:

The core part of the game. Will contain all the logic.
1. Initialize: Will create the game session which will hold all the intermediate states of the game.

//Thinking more about game engine.