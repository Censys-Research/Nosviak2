# nosviak2 change log


## version 1.4

- added
    1. AutoBuy support via sellix
        - custom fields required
            - *username*
            - *password*
        - sync products.json ids with the sellix product IDs
        - dynamic products for auto delivery

## version 1.3

- added 
    1. new gradient function allowing body ignorance with the escape code `\x1b[0m`
    2. forced table style option (ignored on when set to `-1`)
  
- fixed
    1. attack target validation on certain urls has been fixed
    2. [Maybe] SQL database connection being dropped over x amount of time
    3. sessions auto callback only displays other sessions and never your own session



- every version below *1.3* has no active change.log