package main

// make a grid of bricks
const (
	level1 = `--------------------------------------------------
--------------------------------------------------
--------------------------------------------------
----------------------xxxx------------------------
----------------------xxxx------------------------
--------------------------------------------------
`
	level2 = `--------------------------------------------------
--------------------------------------------------
-------------------xxxxxxxxxx---------------------
--------------------------------------------------
-------------------xxxxxxxxxx---------------------
--------------------------------------------------
`
	level3 = `--------------------------------------------------
--------------------------------------------------
-----------------xxxxxxxxxxxxxx-------------------
--------------------------------------------------
---------------xxxxxx--------xxxxxx---------------
--------------------------------------------------
`
	level4 = `--------------------------------------------------
--------------------------------------------------
-----------------xxxxxxxxxxxxxx-------------------
---------------xx--------------xx-----------------
---------------xx--------------xx-----------------
--------------------------------------------------
`
	level5 = `--------------------------------------------------
----------------xx--------------xx----------------
---------------xxxx------------xxxx---------------
--------------xxxxxx----------xxxxxx--------------
-------------xxxxxxxx--------xxxxxxxx-------------
--------------------------------------------------
`
	level6 = `--------------------------------------------------
---------------xxxxxxxxxxxxxxxxxx-----------------
----------------xx--------------xx----------------
----------------xx--------------xx----------------
---------------xxxxxxxxxxxxxxxxxx-----------------
--------------------------------------------------
`
	level7 = `--------------------------------------------------
---------------xxxxxxxxxxxxxxxxxx-----------------
--------------------------------------------------
---------------xxxxxx------xxxxxx-----------------
--------------------------------------------------
---------------xxxxxxxxxxxxxxxxxx-----------------
`
	level8 = `--------------------------------------------------
---------------xx----xx----xx----xx---------------
-------------xxxx--xxxx--xxxx--xxxx---------------
-----------xxxxxxxxxxxxxxxxxxxxxxxxxx-------------
--------------------------------------------------
--------------------------------------------------
`
	level9 = `--------------------------------------------------
------------xxxxxxxxxxxxxxxxxxxxxxxx--------------
------------xx----xx----xx----xx----xx------------
------------xx----xx----xx----xx----xx------------
------------xxxxxxxxxxxxxxxxxxxxxxxx--------------
--------------------------------------------------
`
	level10 = `--------------------------------------------------
-------------xxxxxxxxxxxxxxxxxxxxxxxx-------------
-------------xx----------------------xx-----------
-------------xx----------------------xx-----------
-------------xxxxxxxxxxxxxxxxxxxxxxxx-------------
--------------------------------------------------
`
	level11 = `--------------------------------------------------
-----------xxxxxxxxxxxxxx----xxxxxxxxxxxxxx-------
-------------xx----xx----------xx----xx-----------
-----------xxxxxxxxxxxxxx----xxxxxxxxxxxxxx-------
--------------------------------------------------
--------------------------------------------------
`
	level12 = `--------------------------------------------------
---------xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx---------
--------------------------------------------------
-----------xx----xx----xx----xx----xx-------------
--------------------------------------------------
---------xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx---------
`
	level13 = `--------------------------------------------------
---------xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx---------
--------------------------------------------------
---------xx--xx--xx--xx--xx--xx--xx--xx-----------
--------------------------------------------------
---------xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx---------
`
	level14 = `--------------------------------------------------
-------xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx-------
-------xx----xx----xx----xx----xx----xx----xx-----
--------------------------------------------------
--------------------------------------------------
-------xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx-------
`
	level15 = `--------------------------------------------------
---------xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx---------
--------xx------------------------------xx--------
-------xxxx----------------------------xxxx-------
------xxxxxx--------------------------xxxxxx------
--------------------------------------------------
`
	level16 = `--------------------------------------------------
---------xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx---------
---------xx----xx----xx----xx----xx----xx---------
---------xx----xx----xx----xx----xx----xx---------
---------xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx---------
--------------------------------------------------
`
	level17 = `--------------------------------------------------
-------xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx-------
-------xx----xx----xx----xx----xx----xx----xx-----
-------xx----xx----xx----xx----xx----xx----xx-----
-------xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx-------
--------------------------------------------------
`
	level18 = `--------------------------------------------------
-------xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx-------
-------xx----xx----xx----xx----xx----xx----xx-----
-----xxxx----xxxx----xxxx----xxxx----xxxx----xxxx-
--------------------------------------------------
-------xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx-------
`
	level19 = `--------------------------------------------------
-----xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx-----
------xx----xx----xx----xx----xx----xx----xx------
-----xxxx----xxxx----xxxx----xxxx----xxxx----xxxx-
------xx----xx----xx----xx----xx----xx----xx------
-----xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx-----
`
	level20 = `--------------------------------------------------
-----xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx-----
----xxxxxx----xxxxxx----xxxxxx----xxxxxx----xxxxxx
-----xx--xx----xx--xx----xx--xx----xx--xx----xx---
----xxxxxx----xxxxxx----xxxxxx----xxxxxx----xxxxxx
-----xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx-----
`
)

var Levels = [20]string{
	level1,
	level2,
	level3,
	level4,
	level5,
	level6,
	level7,
	level8,
	level9,
	level10,
	level11,
	level12,
	level13,
	level14,
	level15,
	level16,
	level17,
	level18,
	level19,
	level20,
}