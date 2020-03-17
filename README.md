# shipsInventory
A simple inventory for use aboard ships.


```Inventory table looks like this
+---------------+--------------+------+-----+---------+----------------+
| Field         | Type         | Null | Key | Default | Extra          |
+---------------+--------------+------+-----+---------+----------------+
| id            | int(11)      | NO   | PRI | NULL    | auto_increment |
| itemName      | varchar(100) | NO   |     | NULL    |                |
| description   | text         | NO   |     | NULL    |                |
| created       | datetime     | NO   |     | NULL    |                |
| lastInventory | datetime     | NO   | MUL | NULL    |                |
| numOnHand     | int(11)      | YES  |     | NULL    |                |
| removed       | datetime     | YES  |     | NULL    |                |
| partNum       | varchar(100) | NO   |     | NULL    |                |
| site          | varchar(100) | YES  |     | NULL    |                |
| space         | varchar(100) | YES  |     | NULL    |                |
| drawer        | varchar(100) | YES  |     | NULL    |                |
+---------------+--------------+------+-----+---------+----------------+```
