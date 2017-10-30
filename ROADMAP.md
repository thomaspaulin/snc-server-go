# Road Map
0. Create views in the database of the joins to avoid having to do them
1. Create some endpoints. No join tables yet except teams and divisions, and matches and rinks.
    1. Teams
        1. Read
        2. Create
        3. Update
    2. Divisions
        1. Read
        2. Create
        3. Update
    3. Matches
        1. Read
    4. Rinks
        1. Read
        2. Create
2. Export members of the model structs instead of keeping them all private
3. Create DB queries for players
4. Create players endpoint
    1. Alter Team struct to include players
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
10. Create columns/tables to log creation and modified, and add a deletion flag

.
.
.

???. Create materialised views
???+1. Transactions for the database
