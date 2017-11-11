# Road Map
* Trial using gorm. Will it make it faster to get everything down?
* Create DB queries for players
* Create players endpoint
    1. Alter Team struct to include players. This will require:
        1. Update the join table
        2. Logic to update the player struct to include who they play for
    2. Update queries
        1. Create
        2. Read
* Handle CORS
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
* Transactions for the database
* Make the server responses actually useful instead of "null" and similar
* Look into https://github.com/apiaryio/dredd for testing the API on the server