ACCEPTED REQUEST TYPES:
-----------------------

Server expects client that only can send the following requests:
1. Add: Add the requester to the lobby.
2. NextMove: Will change the state of the game.
3. Create: Will create a new game lobby and session and wait for all members to join.
4. End: Will dissolve the game session and all the players playing.
5. Quit: A player will quit from the game session, but the game session will exist.

GAME LOBBY:

This component will be responsible for lobby actions.
1. Create: Will create a new game lobby and session and wait for all members to join. (completed)
2. Start: This is the only point and function that can initiate the game. (completed)
3. Add: Add to lobby. (completed)
4. End: End the game. (completed)
5. Quit: The player will leave the lobby. (completed)

GAME ENGINE:

The core part of the game. Will contain all the logic.
1. Initialize: Will create the game session which will hold all the intermediate states of the
game.
2. NextMove: Will update the state of the game.

SERVER INITIATED REQUESTS:
--------------------------

1. Update: Will update the state of all players accordingly.

----------------------------------------------------------------------------------------------

IMPLEMENTATION LOGIC:
---------------------

GAME LOBBY:

- CREATE:
So my first instinct to have some sort of global variable that contains all the state of the
games that are happening. So I will first create a global variable called `sessions`.

    > Sessions will be of `typeof session[]`.

    > State type -> {
        CurrentCall string
        PreviousMove string
        NextMove string //player who has to call the move now
        Cards dictionary<string, string[]>
    }

    > Session type -> {
         Creator int
         NoOfMembers int
         Members string[]
         GameNo int
         PreviousWinners string[]
         CurrentState state
         PreviousStates state[]
    }
