docker rm -f textanalytics_1
# build docker container locally
echo "building python server.."
docker build -t textanalytics .
echo "build done"
# run docker container
docker run -d \
-p 9090:9090 \
--name textanalytics_1 \
-e AZURE_TEXT_ANALYTICS_ENDPOINT=${AZURE_TEXT_ANALYTICS_ENDPOINT} \
-e COGNITIVE_SERVICE_KEY=${COGNITIVE_SERVICE_KEY} \
-e ADDR=9090 \
textanalytics
