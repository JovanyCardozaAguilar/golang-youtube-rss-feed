HOW TO RUN (ONLY TESTED ON UBUNTU/POP! OS):
1. Run docker with
sudo docker compose up

2. Run go in terminal
go run .

3. Send request through postman or preferred method

Example:
http://localhost:8080/channel?channelId=1

NOTES:
CHANNEL_CATEGORY / VIDEO_CATEGORY - Do not have update/patch methods as they are not needed, as there construction makes updating them be effectively the same as deleting and inserting them
