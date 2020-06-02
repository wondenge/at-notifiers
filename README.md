<p align="center">
<img src="src/img/AfricasTalkingNotifiers.png" alt="AfricasTalking Notifiers" title="AfricasTalking Notifiers" />
</p>

[![Go Report Card](https://goreportcard.com/badge/wondenge/at-notifiers)](https://goreportcard.com/report/wondenge/at-notifiers)
[![License](https://img.shields.io/badge/license-Apache-blue.svg)]

# AfricasTalking Notifiers

# 1. Overview

This is a Golang HTTP service hosting callbacks to enable rapid prototyping and production use for all AfricasTalking APIs.

This server has the ability to shutdown nicely to ensure that all requests have been completed. This is usually referred to as graceful shutdown.

AfricasTalking powers communications solutions across Africa. With simplified access to telco infrastructure, developers use their powerful SMS, USSD, Voice, Airtime and Payments APIs to bring their ideas to life, as they build and sustain scalable businesses.

# 2. Requirements

- Go 1.14 or newer
- Docker 19.03.9 or newer

# 3. Quickstart

## Local Development

When the above pre-requisites have been fullfilled, run the following commands:

```bash
$ git clone https://github.com/wondenge/at-notifiers.git

$ cd at-notifiers

$ go build ./cmd/at && go build ./cmd/at-cli
```

This creates two binaries, `at` and `at-cli`, which is a server and a client respectively. Start a local server by running `./at` which runs on `http://localhost:3000`.

On a separate terminal, run the CLI client for the notifier API by invoking the --help command using `./at-cli --help` and logs the following output.

## Production Use with Docker

Let's pull the docker image from docker hub

```bash
$ docker pull ondengew/at-notifiers
```

We can then run the docker image which starts the HTTP server listening on 0.0.0.0:8000

```bash
$ docker run -p 8000:8000 ondengew/at-notifiers
```

# 4. SMS Callback

For each request AfricasTalking receives, they respond with a notification back indicating whether the sms transaction was successful or failed.

Other events AfricasTalking sends notifications for are;

- Whenever they receive a message for an application.
- When a user opts out of receiving messages from the alphanumeric code, sender id, or if they subscribe/unsubscribe from the premium SMS service.

To receive these notifications you need to setup a callback URL depending on the type of notification you would like to receive.

SMS API notifications are sent for various SMS categories as shown below:

## Delivery Reports

These are sent whenever the mobile service provider confirms or rejects delivery of a message.

To receive delivery reports, you need to set a delivery report callback URL.

From the dashboard select SMS -> SMS Callback URLs -> Delivery Reports.

> Example Request

```bash
curl --request POST \
  --url http://0.0.0.0:8000/callbacks/africastalking/sms/deliveryreport \
  --header 'accept: application/json' \
  --header 'cache-control: no-cache' \
  --header 'content-type: application/json' \
  --data '{
          "failureReason": "UserAccountSuspended",
          "id": "Consequatur et dicta.",
          "networkCode": "64110",
          "phoneNumber": "Recusandae perspiciatis excepturi reiciendis est quasi est.",
          "retryCount": "Non quo omnis consequatur vero.",
          "status": "Buffered"
       }'
```

> Example Response

```bash
ts=2020-06-01T19:19:03.39104844Z caller=log.go:30 id=M_4eFRe_ req="POST /callbacks/africastalking/sms/deliveryreport" from=172.17.0.1
ts=2020-06-01T19:19:03.391291775Z caller=sms.go:24 info=sms.deliveryReport
ts=2020-06-01T19:19:03.391334458Z caller=log.go:37 id=M_4eFRe_ status=201 bytes=3 time=286.852µs
```

## Incoming Messages.

These are sent whenever a message is sent to any of your registered shortcodes.

To receive incoming messages, you need to set an incoming messages callback URL.

From the dashboard select SMS -> SMS Callback URLs -> Incoming Messages.

```bash
curl --request POST \
  --url http://0.0.0.0:8000/callbacks/africastalking/sms/incomingmessage \
  --header 'accept: application/json' \
  --header 'cache-control: no-cache' \
  --header 'content-type: application/json' \
  --data '{
      "date": "1977-01-31",
      "from": "Velit earum est.",
      "id": "Illo voluptate iste voluptate unde.",
      "linkId": "Vitae autem adipisci sed labore reprehenderit est.",
      "networkCode": "64007",
      "text": "Aut quas quia laboriosam vitae dolor hic.",
      "to": "Eligendi suscipit recusandae libero natus cupiditate libero."
   }'
```

```bash
ts=2020-06-01T19:20:32.06257147Z caller=log.go:30 id=3CrxESuC req="POST /callbacks/africastalking/sms/incomingmessage" from=172.17.0.1
ts=2020-06-01T19:20:32.0627394Z caller=sms.go:30 info=sms.incomingMessage
ts=2020-06-01T19:20:32.062780502Z caller=log.go:37 id=3CrxESuC status=201 bytes=3 time=209.2µs
```

## Bulk SMS Opt Out

These are sent whenever a user opts out of receiving messages from your alphanumeric sender ID.

To receive bulk sms opt out notifications, you need to set a bulk sms opt out callback URL.

From the dashboard select SMS -> SMS Callback URLs -> Bulk SMS Opt Out.

```bash
curl --request POST \
  --url http://0.0.0.0:8000/callbacks/africastalking/sms/bulksmsoptout \
  --header 'accept: application/json' \
  --header 'cache-control: no-cache' \
  --header 'content-type: application/json' \
  --data '{
      "phoneNumber": "Ut vitae sit.",
      "senderId": "Enim neque."
   }'
```

```bash
ts=2020-06-01T19:21:43.984298331Z caller=log.go:30 id=1p6NRt1J req="POST /callbacks/africastalking/sms/bulksmsoptout" from=172.17.0.1
ts=2020-06-01T19:21:43.984455752Z caller=sms.go:37 info=sms.bulkSMSOptOut
ts=2020-06-01T19:21:43.984502415Z caller=log.go:37 id=1p6NRt1J status=201 bytes=3 time=205.051µs
```

The instructions on how to opt out are automatically appended to the first message you send to the mobile subscriber. From then onwards, any other message will be sent ‘as is’ to the subscriber.

## Subscription Notification.

These are sent whenever someone subscribes or unsubscribes from any of your premium SMS products.

To receive premium sms subscription notifications, you need to set asubscription notification callback URL.

From the dashboard select SMS -> SMS Callback URLs -> Subscription Notifications.

```bash
curl --request POST \
  --url http://0.0.0.0:8000/callbacks/africastalking/sms/subscription \
  --header 'accept: application/json' \
  --header 'cache-control: no-cache' \
  --header 'content-type: application/json' \
  --data '{
      "keyword": "Magni molestiae molestiae esse.",
      "phoneNumber": "Ut perferendis sit ratione expedita.",
      "shortCode": "Pariatur aut omnis.",
      "updateType": "deletion"
   }'
```

```bash
ts=2020-06-01T19:22:30.674050184Z caller=log.go:30 id=0f6DpB1K req="POST /callbacks/africastalking/sms/subscription" from=172.17.0.1
ts=2020-06-01T19:22:30.67421404Z caller=sms.go:44 info=sms.subNotifier
ts=2020-06-01T19:22:30.674260295Z caller=log.go:37 id=0f6DpB1K status=201 bytes=3 time=212.748µs
```

# 5. Voice Callback

The Voice API sends a notification when a specific event happens.

To receive these notifications you need to setup a voice callback URL.

From the dashboard select Voice -> Phone Numbers -> Actions -> Callback.

Voice API notifications are sent for;

- Outbound calls: These are sent whenever you make a call from a registered SIP number.

- Inbound calls: These are sent when a call comes to your virtual or SIP number.

- After Input: These are sent whenever an action in your response requires user input (such as GetDigits and Record)

- When Call Ends: These are sent after a call ends. This is the final notification and contains some extra information about the call like the cost and duration.

```bash
curl --request POST \
  --url http://0.0.0.0:8000/callbacks/africastalking/voice/notifications \
  --header 'accept: application/json' \
  --header 'cache-control: no-cache' \
  --header 'content-type: application/json' \
  --data '{
      "amount": "Optio alias iure est.",
      "callSessionState": "Qui reiciendis in sed consequatur eos.",
      "callStartTime": "Saepe veniam.",
      "callerCountryCode": "Deserunt velit praesentium quos id repellat.",
      "callerNumber": "+254711XXXYYY",
      "currencyCode": "Suscipit deleniti beatae.",
      "destinationNumber": "+254711XXXYYY",
      "dialDestinationNumber": "Repellendus est suscipit aperiam in aut.",
      "dialDurationInSeconds": "Magnam dolores nihil eligendi perspiciatis dolore.",
      "dialStartTime": "Aspernatur est ut doloribus architecto est.",
      "direction": "Sunt vero totam sint qui.",
      "dtmfDigits": "Ut assumenda dignissimos cupiditate eius illo qui.",
      "durationInSeconds": "Quibusdam optio facilis accusantium assumenda eum nemo.",
      "hangupCause": "SERVICE_UNAVAILABLE",
      "isActive": "Voluptas fugiat sed numquam iure.",
      "recordingUrl": "Enim beatae ut velit porro.",
      "sessionId": "Nobis nostrum et quaerat quaerat accusantium earum."
   }'
```

```bash
ts=2020-06-01T19:23:19.674572551Z caller=log.go:30 id=jN1y8ZVd req="POST /callbacks/africastalking/voice/notifications" from=172.17.0.1
ts=2020-06-01T19:23:19.674887773Z caller=voice.go:24 info=voice.voiceNotifier
ts=2020-06-01T19:23:19.674912988Z caller=log.go:37 id=jN1y8ZVd status=201 bytes=3 time=342.365µs
```

```bash
curl --request POST \
  --url http://0.0.0.0:8000/callbacks/africastalking/voice/transferevents \
  --header 'accept: application/json' \
  --header 'cache-control: no-cache' \
  --header 'content-type: application/json' \
  --data '{
      "callSessionState": "Transferred",
      "callTransferHangupCause": "InvalidPhoneNumber",
      "callTransferParam": "Unde quos mollitia.",
      "callTransferState": " Active",
      "callTransferredToNumber": "Doloremque velit repellendus perferendis tempora veritatis quidem.",
      "isActive": "1",
      "status": "Success"
   }'
```

```bash
ts=2020-06-01T19:24:02.941335358Z caller=log.go:30 id=UHHNRDGm req="POST /callbacks/africastalking/voice/transferevents" from=172.17.0.1
ts=2020-06-01T19:24:02.941490098Z caller=voice.go:30 info=voice.transferEvents
ts=2020-06-01T19:24:02.941526761Z caller=log.go:37 id=UHHNRDGm status=201 bytes=3 time=193.731µs
```

# 6. USSD Callback

This service is hit when the user dials a USSD code and every time they respond to a menu. Processing USSD requests using the USSD API is very easy once your account is set up. In particular, you will need to:

- Register a service code with AfricasTalking.
- Register a URL that they can call whenever they get a request from a client coming into their system.

Once you register your callback URL, any requests that they receive belonging to you will trigger a callback that sends the request data to that URL using HTTP POST.

```bash
curl --request POST \
  --url http://0.0.0.0:8000/callbacks/africastalking/ussd/sessions \
  --header 'accept: text/plain' \
  --header 'cache-control: no-cache' \
  --header 'content-type: application/json' \
  --data '{
      "networkCode": "Provident ut.",
      "phoneNumber": "Et porro soluta.",
      "serviceCode": "Atque maiores asperiores.",
      "sessionId": "Itaque quam saepe ex.",
      "text": "Rerum ducimus at voluptas ipsa."
   }'
```

# 7 Airtime Callback

## Airtime Validation Notifications

The Airtime API provides optional functionality to validate airtime requests from your application.

To receive these notifications you need to setup an airtime validation callback URL.

From the dashboard select Airtime -> Airtime Callback URLs -> Validation Callback URL.

```bash
curl --request POST \
  --url http://0.0.0.0:8000/callbacks/africastalking/airtime/validation \
  --header 'accept: application/json' \
  --header 'cache-control: no-cache' \
  --header 'content-type: application/json' \
  --data '{
      "amount": 500,
      "currencyCode": "KES",
      "phoneNumber": "+254711XXXYYY",
      "transactionId": "SomeTransactionID"
   }'
```

```bash
ts=2020-06-01T19:27:30.439185432Z caller=log.go:30 id=1eyf-hiQ req="POST /callbacks/africastalking/airtime/validation" from=172.17.0.1
ts=2020-06-01T19:27:30.439358278Z caller=airtime.go:25 info=airtime.validation
ts=2020-06-01T19:27:30.439403337Z caller=log.go:37 id=1eyf-hiQ status=201 bytes=3 time=221.848µs
```

Once you receive a validation callback notification you’ll be expected to send back a JSON response that marks the transaction as Validated or Failed.

If validated we will proceed to send the airtime, if failed, we will block the airtime transaction.

## Airtime Status Notifications

The Airtime API sends delivery status notification from the mobile service provider to your application indicating success or failure of the request.

To receive these notifications you need to setup an airtime status callback URL.

From the dashboard select Airtime -> Airtime Callback URLs -> Status Callback URL.

```bash
curl --request POST \
  --url http://0.0.0.0:8000/callbacks/africastalking/airtime/status \
  --header 'accept: application/json' \
  --header 'cache-control: no-cache' \
  --header 'content-type: application/json' \
  --data '{
      "requestId": "ATQid_SampleTxnId123",
      "status": "Success"
   }'
```

```bash
ts=2020-06-01T19:28:43.286126869Z caller=log.go:30 id=lNKMqVbM req="POST /callbacks/africastalking/airtime/status" from=172.17.0.1
ts=2020-06-01T19:28:43.286228493Z caller=airtime.go:31 info=airtime.status
ts=2020-06-01T19:28:43.286257631Z caller=log.go:37 id=lNKMqVbM status=201 bytes=3 time=133.147µs
```

# 8. Payment Callback

The Payment API sends a notification when a specific event happens. To receive these notifications you need to setup a callback URL depending on the type of notification.

- BankCheckout: These are sent once the provided bank confirms or rejects the checkout request, or once the checkout request expires.

- CardCheckout: These are sent once the card provider confirms or rejects the checkout request, or once the checkout request expires.

- MobileCheckout: These are sent once the mobile subscriber confirms or rejects the checkout request, or once the checkout request expires.

- MobileC2B: These are sent once funds are moved from the mobile subscriber’s account to your payment wallet.

- MobileB2C: These are sent once funds are successfully moved from your payment wallet to the mobile subscriber’s account.
  If AT are not able to complete the transaction, they will refund your payment wallet with the value of the transaction, and also refund your Africa’s Talking Stash with any transaction Fees.

- MobileB2B: These are sent once funds are successfully moved from your payment wallet to the recipeints business account. If AT are not able to complete the transaction, they will refund your payment wallet with the value of the transaction, and also refund your Africa’s Talking Stash with any transaction Fees.

- BankTransfer: These are sent once funds are successfully moved from your payment wallet to the provided bank account account. If AT are not able to complete the transaction, they will refund your payment wallet with the value of the transaction, and also refund your Africa’s Talking Stash with any transaction Fees.

- WalletTransfer: These are sent once funds are successfully moved from your payment wallet to the target payment wallet.

- UserStashTopup: These are sent once funds are successfully moved from your payment wallet to your Africa’s Talking Stash.

```bash
curl --request POST \
  --url http://0.0.0.0:8000/callbacks/africastalking/payments/events \
  --header 'accept: application/json' \
  --header 'cache-control: no-cache' \
  --header 'content-type: application/json' \
  --data '{
      "category": "MobileB2B",
      "clientAccount": "Corrupti ducimus autem cum eum exercitationem.",
      "description": "Hic sequi quis nihil asperiores culpa.",
      "destination": "Eveniet doloribus.",
      "destinationType": "Card",
      "productName": "Molestiae animi officiis.",
      "provider": "Athena",
      "providerChannel": "Velit eveniet est veritatis.",
      "providerFee": "Minus corrupti voluptatem eos.",
      "providerMetadata": "Sunt nemo ex esse eveniet quas.",
      "providerRefId": "Vitae qui qui omnis.",
      "requestMetadata": "Corrupti corrupti nobis corporis qui.",
      "source": "Fugiat quod nemo natus.",
      "sourceType": "PhoneNumber",
      "status": "Success",
      "transactionDate": "Est quibusdam hic ut.",
      "transactionFee": "Quas est nostrum incidunt impedit vel.",
      "transactionId": "In eum.",
      "value": "Ut eveniet."
   }'
```

```bash
ts=2020-06-01T19:29:22.672857301Z caller=log.go:30 id=9fTUhiCO req="POST /callbacks/africastalking/payments/events" from=172.17.0.1
ts=2020-06-01T19:29:22.673067231Z caller=payments.go:24 info=payments.paymentNotifier
ts=2020-06-01T19:29:22.673115464Z caller=log.go:37 id=9fTUhiCO status=201 bytes=3 time=259.037µs
```

```bash
curl --request POST \
  --url http://0.0.0.0:8000/callbacks/africastalking/payments/c2b/validation \
  --header 'accept: application/json' \
  --header 'cache-control: no-cache' \
  --header 'content-type: application/json' \
  --data '{
      "clientAccount": "Culpa soluta quibusdam corporis rerum deleniti.",
      "phoneNumber": "Sequi non sit aliquid.",
      "productName": "Magni est aspernatur qui velit.",
      "provider": "TigoTanzania",
      "providerMetadata": {
         "Molestias laborum magnam numquam nostrum accusamus.": "Aliquam ut vel."
      },
      "value": "Provident nam non est quia et nobis."
   }'
```

```bash
ts=2020-06-01T19:30:05.299538595Z caller=log.go:30 id=jV-1vCdB req="POST /callbacks/africastalking/payments/c2b/validation" from=172.17.0.1
ts=2020-06-01T19:30:05.299679587Z caller=payments.go:30 info=payments.c2bNotifier
ts=2020-06-01T19:30:05.299715637Z caller=log.go:37 id=jV-1vCdB status=201 bytes=3 time=179.946µs
```

```bash
curl --request POST \
  --url http://0.0.0.0:8000/callbacks/africastalking/payments/b2c/validation \
  --header 'accept: application/json' \
  --header 'cache-control: no-cache' \
  --header 'content-type: application/json' \
  --data '{
      "amount": 500,
      "currencyCode": "KES",
      "metadata": {
         "Sit sit id voluptatibus occaecati nostrum.": "Fugiat sunt sed provident eos sunt."
      },
      "phoneNumber": "+254711XXXYYY",
      "sourceIpAddress": "12.34.56.78",
      "transactionId": "SomeTransactionID"
   }'
```

```bash
ts=2020-06-01T19:30:51.451874827Z caller=log.go:30 id=LgDHaWBX req="POST /callbacks/africastalking/payments/b2c/validation" from=172.17.0.1
ts=2020-06-01T19:30:51.452056778Z caller=payments.go:37 info=payments.b2cNotifier
ts=2020-06-01T19:30:51.452136713Z caller=log.go:37 id=LgDHaWBX status=201 bytes=3 time=264.64µs
```

# 9. IoT Callback

The IoT API sends a notification when a device publish event occurs.

To receive these notifications you need to setup a callback URL depending on the type of notification.

When creating your device group, you have the option to supply your application callback URL. This is the remote endpoint to which device messages will be sent.

Here is a curl example request testing the IoT callback.

```bash
curl --request POST \
  --url http://0.0.0.0:8000/callbacks/africastalking/iot/events \
  --header 'accept: application/json' \
  --header 'cache-control: no-cache' \
  --header 'content-type: application/json' \
  --data '{
      "payload": "42",
      "topic": "myusername/devicegroup/sensor/id/1/temperature"
   }'
```

```bash
ts=2020-06-01T19:31:27.110932634Z caller=log.go:30 id=0jnyFwax req="POST /callbacks/africastalking/iot/events" from=172.17.0.1
ts=2020-06-01T19:31:27.111037143Z caller=iot.go:24 info=iot.iotNotifier
ts=2020-06-01T19:31:27.111070365Z caller=log.go:37 id=0jnyFwax status=201 bytes=3 time=131.511µs
```
