# Civic QA Platform 
## Team RAVL: Rafi Bayer, Amit Galitzky, Lia Kitahata, and Vivian Hua


## Architecture
![arch](https://lucid.app/publicSegments/view/3b94b1a5-60ef-40cf-a499-41a6d1cbd504/image.png)

## ERD
![erd](https://lucid.app/publicSegments/view/6885381e-a569-47dc-9792-8b3df2ca0193/image.png)

## Documentation
Gateway API: `docs/services/gateway.yml`  
Services: `docs/services/{service-name}/api-v{major-version}.yml`

## Backend
### Run locally
- Requirements:
    - Go 1.15
    - Docker
    - GNU Make
- Steps:
    - Set `$DOCKER_USER` to your docker username
    - Navigate to `/services`
    - Run `$make up`
