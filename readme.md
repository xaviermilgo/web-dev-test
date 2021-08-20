# Web development challenge
In the following challenge, you will be required to build a backend and frontend service.
You are allowed to use any technologies / framework that you are comfortable with.
Additional support files can be found in the `/files` folder on this repository.

## Introduction
You are tasked with importing, serving and visualizing data from a food delivery center.
The backend should be able to import driver, trip and hotel data found in the `/files` folder.
It should serve this data on a rest api. The frontend should pull the data from this api and show it
in a table and a graph tracking number of deliveries over time. It should also allow the user to
filter based on hotel, driver and time range.

### Backend Api
Using a language and framework of your choice, import the data from the json files below into a SQL database.
You are required to model the data with relevant relationships. 

[Drivers](files/drivers.json)
[Hotels](files/hotels.json)
[Trips](files/trips.json)


Create an api that serves the data above.
Feel free to design the API in any way that may help you archive the requirements of the frontend.