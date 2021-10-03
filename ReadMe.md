
# Covid Data REST API

  

This is the Covid Data Api. It Fetches data from https://data.covid19india.org/v4/min/data.min.json and stores the data in mongodb. Based on the state it gives specific data. The Api autorefreshes teh data after every 30 min. It return the last updated data in IST. For geocoding mapmyindia's api are used. It returns the state data if the location are of India. If the location are not of india it returns an error. Currently the api has data of 37 States and UTs. The list of state and UTs are:

    Delhi,Nagaland,Kerala,Mizoram,Odisha,
    Andaman and Nicobar,Gujarat,Jharkhand,
    Meghalaya,Rajasthan,Uttarakhand,
    Total,Bihar,Ladakh,Tripura,Andhra Pradesh,
    Assam,Telangana,Haryana,Sikkim,West Bengal,
    Chandigarh,Chhattisgarh,Karnataka,Jammu and Kashmir,
    Lakshadweep,Madhya Pradesh,Maharashtra,Manipur,Arunachal Pradesh,
    Goa,Himachal Pradesh,Puducherry,Punjab,Tamil Nadu,Uttar Pradesh

For UI  Swagger and OpenAPI has been implemented in this project. For dependency management it has the vendor. The structure of project is standard. 
The api is hosted on heruko. 
[Heruko Link ](https://covid-data-go-api.herokuapp.com/swagger/index.html#)

  

# Specifications

 1. Go Language
 2. Echo Framework
 3. Swagger Documentation(echoSwagger)







  


## Install

    install: go 1.17
    
    clone from github https://github.com/Pankaj-Kadian/CovidApi
    
    run: go mod tidy

  

## Run the app

  

    go run main.go
    
      

  
  

# REST API

The Covid Api has 3 End points.

 1. Get Data of a particular State
 2. Get Data of all states
 3. Get Data based on latitude and longitude

  

## Get Data of a state
  

### Request

**`GET /GetStateData?state=statename`**

  

curl -i -H 'Accept: application/json' http://localhost:8080/swagger/index.html/GetStateData

  

### Response

  

#### Response Header

    connection: keep-alive
    
    content-length: 210
    
    content-type: application/json; charset=UTF-8
    
    date: Sun, 03 Oct 2021 16:39:45 GMT
    
    server: Cowboy
    
    via: 1.1 vegur

  

#### Response Body

    {
    
    "state_code": "PB",
    
    "total_cases": 601698,
    
    "total_recovered": 584898,
    
    "total_death": 16520,
    
    "total_vaccinated1": 14555719,
    
    "total_vaccinated2": 5060276,
    
    "total_tested": 14606979,
    
    "last_updated": "2021-10-03T10:47:07+05:30"
    
    }

  

#### Error Body

    {
    
    "code": 400,
    
    "message": "status bad request"
    
    }

## Get Data of all states
  

### Request

**`GET /GetAllData`**

  

curl -i -H 'Accept: application/json' http://localhost:8080/swagger/index.html/GetAllData

  

### Response

  

#### Response Header

    connection: keep-alive
    
    content-length: 210
    
    content-type: application/json; charset=UTF-8
    
    date: Sun, 03 Oct 2021 16:39:45 GMT
    
    server: Cowboy
    
    via: 1.1 vegur

  

#### Response Body
[
    {
    
    "state_code": "PB",
    
    "total_cases": 601698,
    
    "total_recovered": 584898,
    
    "total_death": 16520,
    
    "total_vaccinated1": 14555719,
    
    "total_vaccinated2": 5060276,
    
    "total_tested": 14606979,
    
    "last_updated": "2021-10-03T10:47:07+05:30"
    
    }
    ]

  

#### Error Body

    {
    
    "code": 400,
    
    "message": "status bad request"
    
    }
 
## Get data of a state using geocoding
  

### Request

**`GET /GetByGeoLocation?latitude=lat&longitude=lng`**

  

curl -i -H 'Accept: application/json' http://localhost:8080/swagger/index.html//GetByGeoLocation

  

### Response

  

#### Response Header

    connection: keep-alive
    
    content-length: 210
    
    content-type: application/json; charset=UTF-8
    
    date: Sun, 03 Oct 2021 16:39:45 GMT
    
    server: Cowboy
    
    via: 1.1 vegur

  

#### Response Body

    {
    
    "state_code": "PB",
    
    "total_cases": 601698,
    
    "total_recovered": 584898,
    
    "total_death": 16520,
    
    "total_vaccinated1": 14555719,
    
    "total_vaccinated2": 5060276,
    
    "total_tested": 14606979,
    
    "last_updated": "2021-10-03T10:47:07+05:30"
    
    }

  

#### Error Body

    {
    
    "code": 400,
    
    "message": "status bad request"
    
    }
  
## Examples 
### Home Page
    ![](UI.jpg)
### Get State Data
    ![](GetStateData.jpg)
### Get All Data
    ![](Get%20All%20Data.jpg)
### Get By GeoLoaction Data
    ![](Get%20By%20location.jpg)

 
