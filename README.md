# TheWall <small>(by lucius✨💞🌷)</small>

## Introduction

Simple website for writing whatever you want publicly on a website everyone can see. Original idea by [elcomix97edu](https://github.com/elcomix97edu/thewall).

Written mainly in good ol' JS/PHP.

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

v.0.1🐱‍
[ ✓ ] Make basic template
[ ✓ ] Make the website fetch data from the BD and show it on the wall
[ ✓ ] Let users add new posts to the wall
[   ] Insert date with every new post (possibly user too?)
[   ] Add fade animation for new posts
[   ] Put some nice font ✨
[   ] Add functionality to add photos/urls
[   ] Don't allow HTML code




