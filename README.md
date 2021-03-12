# Civic QA Platform 
## Team RAVL: Rafi Bayer, Amit Galitzky, Lia Kitahata, and Vivian Hua


## Architecture
![arch](https://lucid.app/publicSegments/view/02744f1b-bda1-41eb-8c9a-0b9c2d66a532/image.png)

## ERD
![erd](https://lucid.app/publicSegments/view/6885381e-a569-47dc-9792-8b3df2ca0193/image.png)

## Documentation
Gateway API: `docs/services/gateway.yml`  
Services: `docs/services/{service-name}/api-v{major-version}.yml`

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
