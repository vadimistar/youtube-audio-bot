include .env

delete_webhook:
	curl --request POST --url "https://api.telegram.org/bot$(TELEGRAM_TOKEN)/deleteWebhook"

create_webhook: delete_webhook
	curl --request POST --url "https://api.telegram.org/bot$(TELEGRAM_TOKEN)/setWebhook" --header "content-type: application/json" --data "{\"url\": \"$(TELEGRAM_BOT_CONTAINER_URL)\"}"

build:
	docker-compose build

push: build
	docker-compose push

deploy: push
	yc serverless container revision deploy --container-id $(TELEGRAM_BOT_CONTAINER_ID) --image 'cr.yandex/$(YC_IMAGE_REGISTRY_ID)/telegram-bot:latest' --service-account-id $(SERVICE_ACCOUNT_ID) --environment='$(shell tr '\n' ',' < .env)' --core-fraction 5 --execution-timeout 60s
	yc serverless container revision deploy --container-id $(WORKER_CONTAINER_ID) --image 'cr.yandex/$(YC_IMAGE_REGISTRY_ID)/worker:latest' --service-account-id $(SERVICE_ACCOUNT_ID) --environment='$(shell tr '\n' ',' < .env)' --core-fraction 50 --execution-timeout 600s
