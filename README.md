# go-query-generator

The function takes an interface as a parameter, allowing any type of data to be passed as an argument, although it only works with structures. From the structure tags, it retrieves the column names, and with the information stored in each field, it obtains the value and placeholders. Based on the structure name, it obtains the table and generates the data to build an SQL query.

