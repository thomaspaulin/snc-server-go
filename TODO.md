# Road Map
* Create some endpoints. No join tables yet except teams and divisions, and matches and rinks.
    1. Teams
        1. ~~Read~~
        2. Create
        3. Update
    2. Divisions
        1. ~~Read~~
        2. ~~Create~~
        3. ~~Update~~
    3. Matches
        1. Read
    4. Rinks
        1. ~~Read~~
        2. ~~Create~~
* Begin writing tests in go's built-in testing framework, then maybe, gingko (or cucumber/gherkin), and Frisby later
* Create DB queries for players
* Create players endpoint
    1. Alter Team struct to include players. This will require:
        1. Update the join table
        2. Logic to update the player struct to include who they play for
    2. Update queries
        1. Create
        2. Read
* Create endpoint goals (inc assists)
    a. Read
* Create endpoint penalties
    a. Read
* Add goals and penalties to Match struct
    a. Alter struct
    b. Update queries
* Add player match stats endpoint
* Add goalie match stats endpoint
* Create columns/tables to log creation and modified
* Return 404s instead of no rows in result from Scan()
* Think about how deleting an item in each table woul affect others in order to get the cascading right
* Consider moving to gin-gonic/gin because it has had commits within the past few months until gocraft/web which last had a commit 2 years ago
* Transactions for the database
