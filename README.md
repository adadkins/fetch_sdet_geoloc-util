# FetchSDETWeatherGeocoder
Fetch.com SDET Take home weather geocoder assignment

# Instructions

Copy the .env-sample as .env and fill the API key with your value.

Build:

```
go build -o geoloc-util .
```


Run:
```
./geoloc-util "90292" "90034" "85014" "61240" "Chicago" "Ames, IA" "NYC"
```


Test Coverage:

```
go test ./... -cover
```


## Testing Strategy

The main app uses Depenedency Injection to inject a mock of an interface. This allows us to test expected results and errors. 

The open_weather_map_wrapper uses a recorder to record HTTP requests. The recorder allows us to replay the requests in our tests without actually calling the external API, but still get realistic results for testing.


# Assignment

## GeoLocation API

### Authorization
The API requires you to pass an appid url parameter with an API Key for authenticating the
request. For the purposes of this exercise, you can use the following API key:
f897a99d971b5eef57be6fafa0d83239

### Coordinates by location name
https://openweathermap.org/api/geocoding-api#direct_name
This API endpoint takes a single location city and state query and returns the possible results
for the location search. The results should include the place information needed to display to the
user.

### Coordinates by zip/post code
https://openweathermap.org/api/geocoding-api#direct_zip
This API endpoint takes a single zipcode query and returns the possible results for the location
search. The results should include the place information needed to display to the user.

### Test Code
Please write some test code that tests your utility. For the purpose of this exercise, we would
like you to write full integration tests and not unit tests. You should include whatever tests you
think would be necessary to properly test the utility as you have written it.

If you wish, you can submit two separate repositories, one for the utility and one for testing the
utility. If you do so, please provide detailed instructions on how the two repos should be
configured such that the test repo can run on the utility.

