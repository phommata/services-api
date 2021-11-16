Services API
Includes the data model and API. 
A Services API that can be used to implement this dashboard widget. It supports
- Returning a list of services
- support filtering, sorting, pagination
- Fetching a particular service
- including a method for retrieving its versions

Additional considerations
- Include tests (unit, integration) or a test plan
- Provide authentication/authorization on the API
- Add CRUD operations to the API

## Getting started 
Build our custom image(i.e dockerfile provided on build key)

    docker-compose build 

Starts all our container(service) configured on yaml file

    docker-compose up

Removes all our container(service)

    docker-compose down

## Quick start

    docker-compose down; docker-compose build; docker-compose up

## Example HTTP request:

    GET localhost:8080/services?limit=1&offset=0&search=us
    
    {
       "data":[
          {
             "id":7,
             "name":"FX Rates International...",
             "description":"Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor",
             "versions":[
                {
                   "id":19,
                   "name":"1.0.0",
                   "service_id":"7",
                   "created_at":"2021-11-09T03:55:51Z",
                   "updated_at":"2021-11-09T03:55:51Z"
                },
                {
                   "id":20,
                   "name":"1.0.1",
                   "service_id":"7",
                   "created_at":"2021-11-09T03:55:51Z",
                   "updated_at":"2021-11-09T03:55:51Z"
                },
                {
                   "id":21,
                   "name":"1.0.2",
                   "service_id":"7",
                   "created_at":"2021-11-09T03:55:51Z",
                   "updated_at":"2021-11-09T03:55:51Z"
                }
             ],
             "version_count":3,
             "created_at":"2021-11-09T03:55:51Z"
          }
       ],
       "limit":1,
       "offset":1,
       "total":5
    }

## Design considerations
1. Extensibility
2. Modularity
3. Maintainability
4. Reusability
5. Robustness
6. Scalability

## Assumptions
1. Service has id, name, description, version, version_count, created_at
2. Version has id, name, created_at
3. API can sort, filter, paginate, and list versions
4. UI will handle individual service cards
5. Search substring in name or description 
6. Response with pagination will contain offset and limit
7. Services will be sorted by created_at

## Trade-offs
1. Go is fast and scalable
2. MySQL is good for structured, normalized data
3. Docker containerizes application to easily build, test, and deploy. 