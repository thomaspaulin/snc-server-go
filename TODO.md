# Road Map
1. Create some endpoints. No join tables yet except teams and divisions, and matches and rinks.
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
2. Begin writing tests in go's built-in testing framework, then maybe, gingko (or cucumber/gherkin), and Frisby later
3. Create DB queries for players
4. Create players endpoint
    1. Alter Team struct to include players. This will require:
        1. Update the join table
        2. Logic to update the player struct to include who they play for
    2. Update queries
        1. Create
        2. Read
5. Create endpoint goals (inc assists)
    a. Read
6. Create endpoint penalties
    a. Read
7. Add goals and penalties to Match struct
    a. Alter struct
    b. Update queries
8. Add player match stats endpoint
9. Add goalie match stats endpoint
11. Create columns/tables to log creation and modified, and add a deletion flag
12. Consider moving to gin-gonic/gin because it has had commits within the past few months until gocraft/web which last had a commit 2 years ago
13. Consider using views

.
.
.

???. Create materialised views
???+1. Transactions for the database