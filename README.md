# TheWall <small>(by lucius✨💞🌷)</small>

## Introduction

Simple website for writing whatever you want publicly on a website everyone can see. Original idea by [elcomix97edu](https://github.com/elcomix97edu/thewall).

It was written mainly in good ol' PHP but I'm currently porting it over to Go, so I made a new branch for that you can check here.

## Notes

* Connecting to the DB

connect.php is assumed to be outside of the html_public folder, so it's not included here.

A basic connect.php file could be as follows:

```
    <?php
    $link = mysqli_connect("localhost", "root", "password", "database");
    mb_language('uni');
    mb_internal_encoding('UTF-8');
    mysqli_set_charset($link, "utf8");
   ```
   
  * Importing the database
  
  TheWall currently runs on MySQL. Inside the "db" folder there's everything you need to set it up or modify it.
 

## To-Do List

[ ✓ ] Alert the user when SQL connection is not working

[ ] Port from PHP to Go

[ ] Have a nicer CSS with animations and shit

[ ] Make a JavaScript for drawing small black and white pictures a la _Splatoon_

[ ] Auto detect 

[ ] Send data through JSON

[ ] Deny JS tags
