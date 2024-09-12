<h3 id="ex01">Exercise 01: Assessing Damage</h3>

Okay, so now the owner decided to compare the databases. You've seen that the stolen database has modified versions of the same recipes, meaning there are several possible cases:

1. New cake is added or old one removed
2. Cooking time is different for the same cake
3. New ingredient is added or removed for the same cake. _Important:_ the order of ingredients doesn't matter. Only the names are.
4. The count of units for the same ingredient has changed.
5. The unit itself for measuring the ingredient has changed.
6. Ingredient unit is missing or added

Quickly looking through the database, the owner also noticed that nobody bothered renaming the cakes or the ingredients (possibly, only added some new ones), so you may assume names are the same across both databases.

Your application should be runnable like this:

`~$ ./compareDB --old original_database.xml --new stolen_database.json`

It should work with both formats (JSON and XML) for original AND new database, reusing the code from Exercise 00.

The output should look like this (the same cases explained above):

```
ADDED cake "Moonshine Muffin"
REMOVED cake "Blueberry Muffin Cake"
CHANGED cooking time for cake "Red Velvet Strawberry Cake" - "45 min" instead of "40 min"
ADDED ingredient "Coffee beans" for cake  "Red Velvet Strawberry Cake"
REMOVED ingredient "Vanilla extract" for cake  "Red Velvet Strawberry Cake"
CHANGED unit for ingredient "Flour" for cake  "Red Velvet Strawberry Cake" - "mugs" instead of "cups"
CHANGED unit count for ingredient "Strawberries" for cake  "Red Velvet Strawberry Cake" - "8" instead of "7"
REMOVED unit "pieces" for ingredient "Cinnamon" for cake  "Red Velvet Strawberry Cake"
```

<h2 id="chapter-vi" >Chapter VI</h2>
<h3 id="ex02">Exercise 02: Afterparty</h3>

Digging through the database, the Villariba bakery owner suddenly realized - this guy is brilliant! Some recipes were improved a lot comparing to the old version, and these new ideas were really creative! He rushed into Villabajo and found the guy who, as he first thought, stole his most precious legacy.

...The same evening in the tavern two old bakers were hugging, drinking and laughing so hard that it was heard in both villages. They became best friends during the last couple of hours, and each of them was really happy to finally find the person who could blabber all night about cakes! Also turns out, the guy did't steal the database, he was just trying to guess by the taste and tried to improvise around a bit.

After this whole mess they both agreed to use your code, but asked you to do one last task for them. They were quite impressed by how you've managed to do the comparison between the databases, so they've also asked you to do the same thing with their server filesystem backups, so neither bakery would run into any technical issues in the future.

So, your program should take two filesystem dumps.

`~$ ./compareFS --old snapshot1.txt --new snapshot2.txt`

They are both plain text files, unsorted, and each of them includes a filepath on every like, like this:

```
/etc/stove/config.xml
/Users/baker/recipes/database.xml
/Users/baker/recipes/database_version3.yaml
/var/log/orders.log
/Users/baker/pokemon.avi
```

Your tool should output the very similar thing to a previous code (without CHANGED case though):

```
ADDED /etc/systemd/system/very_important/stash_location.jpg
REMOVED /var/log/browser_history.txt
```

There is one issue though - the files can be really big, so you can assume both of them won't fit into RAM on the same time. There are two possible ways to overcome this - either to compress the file in memory somehow, or just read one of them and then avoid reading the other. **Side note:** this is actually a very popular interview question.
