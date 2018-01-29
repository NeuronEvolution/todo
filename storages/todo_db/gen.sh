#!/usr/bin/env bash

mysql-orm-gen -sql_file=./todo_db.sql -orm_file=./todo_db-gen.go -package_name="todo_db"