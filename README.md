# Civic QA Platform 
## Team RAVL: Rafi Bayer, Amit Galitzky, Lia Kitahata, and Vivian Hua

## Project Overview
Legislative Assistants and those in similar positions handle an overwhelming volume of constituent inquiries daily, making it challenging to respond to all constituents and still find time for their other daily duties. How might LAs better communicate with constituents and handle responses to improve transparency and increase engagement?

Some Legislative Assistants receive 100s of inquiries daily about various topics, ranging from House Bills to questions about citizenship. Due to this high volume of messages, LAs often do not get the chance to respond to all their constituents, creating a lack of engagement between the people and their representatives.

Existing platforms and technologies for LAs to filter and respond to their inquiries do not supply the necessary functionality to effectively and efficiently organize and display incoming messages. As such, our team wanted to create a new solution that caters directly to the specific needs of LAs and helps them to categorize messages while improving response times. With more efficiency on the LAs’ side, more constituents receive responses and more time can be allocated to helping constituents in other ways.

Beginning with extensive market and user research, we honed in on what Legislative Assistants need the most when it comes to a solution like ours, and how we can best create a platform that responds to those needs. Our process can be broken down into a few main steps:
1. Research
2. Concept Validation
3. Prototyping
4. User Testing
5. Feedback Implementation

We aim to make LAs’ jobs easier by removing some of the tediousness of their current system and making it so they can respond to constituents faster, while also gaining a better understanding of their engagement levels. By supplying engagement reports, daily trends, and filtered topics, LAs can see what constituents care about most and to what extent those in their jurisdiction are interacting with their representatives.

With improved responsiveness and greater overall efficiency, LAs will be able to communicate with constituents more effectively, likely encouraging more people to engage, as LAs will be more reliable and likely to respond.

## Project Presentation Video 
[Video Link](https://www.youtube.com/watch?v=COTimhg6Cs8)

## Architecture
![arch](https://lucid.app/publicSegments/view/02744f1b-bda1-41eb-8c9a-0b9c2d66a532/image.png)

## ERD
![erd](https://lucid.app/publicSegments/view/6885381e-a569-47dc-9792-8b3df2ca0193/image.png)

## Documentation
Gateway API: `docs/services/gateway.yml`  
Services: `docs/services/{service-name}/`: [`api-v{major-version}.yml` | `service.md`]

## License
MIT License: `LICENSE.txt`

## Backend
### Run locally
- Requirements:
    - Go 1.15
    - gcc 
    - Docker
    - GNU Make
- Steps:
    - Set `$DOCKER_USER` to your docker username
    - Navigate to `/services`
    - Run `$make up`

## Frontend
### Run locally
- Requirements: 
    - React 17.01 and (Recharts 2.0.0, React Redux 7.2.3, and React Router 5.2.0)
    - TypeScript 4.1.5
- Steps:
    - Go to my-app folder
    - npm install
    - Change Base constant in civic-qa/my-app/src/Constants/Endpoints.ts to http://localhost:3000/v0 to make requests         to local backend
    - npm start to run client on localhost:3000 

## Future Work
Some places for future maintainers to start:
- `logAggregator`'s Query functionality isn't really useful over raw SQL, consider creating a secure admin service to view database data from the log db and main db directly.
- `form`'s data access is kind of a mess due to limitations (mostly of my understanding) of GORM. Consider refactoring to reduce the number of JOINS produced by certain queries. 
- `unit testing`. I know it's not fun, but there are a lot of unit tests to right. Use them to find places where code is overly coupled too. Right now we rely pretty heavily on integration tests located in `test/e2e`.
- `templated responses`: Frontend + Storage/API for bulk responding via templates.
- `dynamic forms`: Frontend Editor/Management + backend storage/API/HTML Generation to create dynamic forms with custom fields, inputs, and validation. May require advanced knowledge of JSON data in MySQL or some other tricks to represent un-normalized data.
 