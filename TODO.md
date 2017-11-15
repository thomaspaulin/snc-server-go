# Road Map
* Create endpoint goals (inc assists)
    a. Read
* Create endpoint penalties
    a. Read
* Add goals and penalties to Match struct
    a. Alter struct
    b. Update queries
* Fix the error from parsing the path param being ignored
* Add player match stats endpoint
* Add goalie match stats endpoint
* Transactions for the database
* Look into https://github.com/apiaryio/dredd for testing the API on the server
* Handle CORS
* HAL links. Initial request gets some of the metadata then the next one follows the _links property to get things for example the examples at https://api.football-data.org/documentation


# GORM todos
* Relations (foreign keys)
* Indexes