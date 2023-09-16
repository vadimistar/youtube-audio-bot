# Youtube Audio Bot

Based on Yandex Cloud serverless containers.

## Architecture

There are two containers: 'telegram-bot' and 'worker'. 'telegram-bot' gets messages from Telegram using webhook and sends responses to users. 'worker' downloads YouTube video and converts it into audio format.

## How to set up

### Prepare .env file

Copy the .env.example into .env. You will fill the empty fields later.

### Yandex Cloud

1. Create a Yandex Cloud container registry. Get it's ID and fill 'YC_IMAGE_REGISTRY_ID' field in .env. 
2. Create a service account, give it an admin role. Get it's ID and fill 'SERVICE_ACCOUNT_ID' field in .env.
3. [Authorize using your service account](https://cloud.yandex.ru/docs/container-registry/operations/authentication)
4. Create two serverless containers (first is for Telegram bot, second is for worker), make them public and store their IDs and URLs into 'TELEGRAM_BOT_CONTAINER_ID', 'TELEGRAM_BOT_CONTAINER_URL', 'WORKER_CONTAINER_ID', 'WORKER_CONTAINER_URL' in .env.
5. Create a bucket for storing objects, set object read access and object listing access to public, place the name of the bucket to 'YC_BUCKET_NAME' in .env'.

### Telegram

Create a new bot in Telegram using BotFather. Get it's API token and store it into 'TELEGRAM_TOKEN' field in .env. 
