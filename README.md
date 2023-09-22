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
5. Create a bucket for storing objects, set object read access and object listing access to public, place the name of the bucket to 'YC_BUCKET_NAME' in .env and the URL to 'YC_BUCKET_URL' in .env (you can get URL by getting the URL of the file in the bucket and trimming the filename). It's crucial to set the lifecycle of objects, so objects are automatically deleted after they were used.
6. Configure credentials for Object Storage using 'AWS_ACCESS_KEY_ID' and 'AWS_SECRET_ACCESS_KEY' in .env. You can do that by creating a new static access key for your service account.
7. Create a message queue, fill the 'MESSAGE_QUEUE_URL' field in .env and add the newly created message queue as a trigger for your serverless containers.

### Telegram

Create a new bot in Telegram using BotFather. Get it's API token and store it into 'TELEGRAM_TOKEN' field in .env. 
