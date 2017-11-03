# Road Map
* Create some endpoints. No join tables yet except teams and divisions, and matches and rinks.
    1. Teams
        1. ~~Read~~
        2. Create
        3. Update
    2. Divisions
        1. ~~Read~~
        2. Create
        3. Update
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
* Create columns/tables to log creation and modified, and add a deletion flag
* Consider moving to gin-gonic/gin because it has had commits within the past few months until gocraft/web which last had a commit 2 years ago
* Restructure according to https://medium.com/@benbjohnson/standard-package-layout-7cdbc8391fc1
* Transactions for the database
