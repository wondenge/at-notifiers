# AfricasTalking Notifiers

# Overview

This is a Golang HTTP service hosting callbacks for all AfricasTalking APIs. This server has the ability to shutdown nicely to ensure that all requests have been completed. This is usually referred to as graceful shutdown.

```bash
^Cts=2020-05-31T14:35:49.780992137Z caller=main.go:155 info="exiting (interrupt)"
ts=2020-05-31T14:35:49.781055553Z caller=http.go:236 info="shutting down HTTP server at \"0.0.0.0:8000\""
ts=2020-05-31T14:35:49.781140662Z caller=main.go:161 info=exited
```

# Build from Source

Requirements:

- Go 1.14 or newer
- Docker 19.03.9 or newer

## For development

```bash
$ git clone
$ cd at-notifiers
$ go build ./cmd/atsvr && go build ./cmd/atsvr-cli
```

This creates two binaries, atsvr and atsvr-cli, which is a server and a client respectively. Start a local server by running `./atsvr`
which runs on `http://localhost:3000` and outputs the following logs.

```bash
ts=2020-05-31T14:39:05.520825711Z caller=http.go:201 info="HTTP \"DeliveryReport\" mounted on POST /callbacks/africastalking/sms/deliveryreport"
ts=2020-05-31T14:39:05.52087906Z caller=http.go:201 info="HTTP \"IncomingMessage\" mounted on POST /callbacks/africastalking/sms/incomingmessage"
ts=2020-05-31T14:39:05.520909653Z caller=http.go:201 info="HTTP \"BulkSMSOptOut\" mounted on POST /callbacks/africastalking/sms/bulksmsoptout"
ts=2020-05-31T14:39:05.520925898Z caller=http.go:201 info="HTTP \"SubNotifier\" mounted on POST /callbacks/africastalking/sms/subscription"
ts=2020-05-31T14:39:05.520947528Z caller=http.go:204 info="HTTP \"VoiceNotifier\" mounted on POST /callbacks/africastalking/voice/notifications"
ts=2020-05-31T14:39:05.520967045Z caller=http.go:204 info="HTTP \"TransferEvents\" mounted on POST /callbacks/africastalking/voice/transferevents"
ts=2020-05-31T14:39:05.520983162Z caller=http.go:207 info="HTTP \"UssdNotifier\" mounted on POST /callbacks/africastalking/ussd/sessions"
ts=2020-05-31T14:39:05.520999475Z caller=http.go:210 info="HTTP \"Validation\" mounted on POST /callbacks/africastalking/airtime/validation"
ts=2020-05-31T14:39:05.521021341Z caller=http.go:210 info="HTTP \"Status\" mounted on POST /callbacks/africastalking/airtime/status"
ts=2020-05-31T14:39:05.521039289Z caller=http.go:213 info="HTTP \"PaymentNotifier\" mounted on POST /callbacks/africastalking/payments/events"
ts=2020-05-31T14:39:05.521055436Z caller=http.go:213 info="HTTP \"C2bNotifier\" mounted on POST /callbacks/africastalking/payments/c2b/validation"
ts=2020-05-31T14:39:05.521071599Z caller=http.go:213 info="HTTP \"B2cNotifier\" mounted on POST /callbacks/africastalking/payments/b2c/validation"
ts=2020-05-31T14:39:05.521089257Z caller=http.go:216 info="HTTP \"IotNotifier\" mounted on POST /callbacks/africastalking/iot/events"
ts=2020-05-31T14:39:05.521114861Z caller=http.go:219 info="HTTP \"Show\" mounted on GET /health"
ts=2020-05-31T14:39:05.521133021Z caller=http.go:222 info="HTTP \"../../gen/http/openapi.json\" mounted on GET /swagger/swagger.json"
ts=2020-05-31T14:39:05.52115686Z caller=http.go:231 info="HTTP server listening on \"localhost:3000\""
```

On a separate terminal, run the CLI client for the notifier API by invoking the --help command using `./atsvr-cli --help` and logs the following output.

```bash
./atsvr-cli is a command line client for the notifier API.

Usage:
    ./atsvr-cli [-host HOST][-url URL][-timeout SECONDS][-verbose|-v] SERVICE ENDPOINT [flags]

    -host HOST:  server host (local). valid values: local, docker
    -url URL:    specify service URL overriding host URL (http://localhost:8080)
    -timeout:    maximum number of seconds to wait for response (30)
    -verbose|-v: print request and response details (false)

Commands:
    airtime (validation|status)
    health show
    iot iot-notifier
    payments (payment-notifier|c2b-notifier|b2c-notifier)
    sms (delivery-report|incoming-message|bulk-sms-opt-out|sub-notifier)
    ussd ussd-notifier
    voice (voice-notifier|transfer-events)

Additional help:
    ./atsvr-cli SERVICE [ENDPOINT] --help
```

## Docker

We can build from source and use a Docker Image locally for our Golang HTTP Listener using our [Dockerfile](). This is a multi-stage build Dockerfile built with two images:

- From [golang:1.14.3-alpine3.11](https://hub.docker.com/layers/golang/library/golang/1.14.3-alpine3.11/images/sha256-911ebd34ce76d69beac233711215ece09b81f9000bb0fc4615ef3ee732a1c495?context=explore) as Builder
- and then the resulting binaries built from [scratch](https://hub.docker.com/_/scratch/).

Let's build the image by running;

```bash
$ git clone
$ cd at-notifiers
$ docker build . -t at-notifier:latest
```

We can then run the docker image locally;

```bash
$ docker run -p 8000:8000 at-notifier
```

```bash
(base) wondenge@hpelitebook:~/projects/sdks/at-notifiers$ docker run -p 8000:8000 at-notifier
ts=2020-05-31T14:06:24.31800445Z caller=http.go:201 info="HTTP \"DeliveryReport\" mounted on POST /callbacks/africastalking/sms/deliveryreport"
ts=2020-05-31T14:06:24.318065289Z caller=http.go:201 info="HTTP \"IncomingMessage\" mounted on POST /callbacks/africastalking/sms/incomingmessage"
ts=2020-05-31T14:06:24.318093753Z caller=http.go:201 info="HTTP \"BulkSMSOptOut\" mounted on POST /callbacks/africastalking/sms/bulksmsoptout"
ts=2020-05-31T14:06:24.318114975Z caller=http.go:201 info="HTTP \"SubNotifier\" mounted on POST /callbacks/africastalking/sms/subscription"
ts=2020-05-31T14:06:24.318127887Z caller=http.go:204 info="HTTP \"VoiceNotifier\" mounted on POST /callbacks/africastalking/voice/notifications"
ts=2020-05-31T14:06:24.318140974Z caller=http.go:204 info="HTTP \"TransferEvents\" mounted on POST /callbacks/africastalking/voice/transferevents"
ts=2020-05-31T14:06:24.318155288Z caller=http.go:207 info="HTTP \"UssdNotifier\" mounted on POST /callbacks/africastalking/ussd/sessions"
ts=2020-05-31T14:06:24.31817146Z caller=http.go:210 info="HTTP \"Validation\" mounted on POST /callbacks/africastalking/airtime/validation"
ts=2020-05-31T14:06:24.318186354Z caller=http.go:210 info="HTTP \"Status\" mounted on POST /callbacks/africastalking/airtime/status"
ts=2020-05-31T14:06:24.31820336Z caller=http.go:213 info="HTTP \"PaymentNotifier\" mounted on POST /callbacks/africastalking/payments/events"
ts=2020-05-31T14:06:24.31822186Z caller=http.go:213 info="HTTP \"C2bNotifier\" mounted on POST /callbacks/africastalking/payments/c2b/validation"
ts=2020-05-31T14:06:24.318231384Z caller=http.go:213 info="HTTP \"B2cNotifier\" mounted on POST /callbacks/africastalking/payments/b2c/validation"
ts=2020-05-31T14:06:24.318247624Z caller=http.go:216 info="HTTP \"IotNotifier\" mounted on POST /callbacks/africastalking/iot/events"
ts=2020-05-31T14:06:24.318261437Z caller=http.go:219 info="HTTP \"Show\" mounted on GET /health"
ts=2020-05-31T14:06:24.318276134Z caller=http.go:222 info="HTTP \"../../gen/http/openapi.json\" mounted on GET /swagger/swagger.json"
ts=2020-05-31T14:06:24.318302371Z caller=http.go:231 info="HTTP server listening on \"0.0.0.0:8000\""
```

We also have an existing maintained official docker image from docker hub.

## 1. SMS Callback Service

SMSs are sent from an application by making a HTTP POST request to the AfricasTalking SMS API. For each request AfricasTalking receives, they respond with a notification back indicating whether the sms transaction was successful or failed.

Other events AfricasTalking sends notifications for are;

- Whenever they receive a message for an application.
- When a user opts out of receiving messages from the alphanumeric code, sender id, or if they subscribe/unsubscribe from the premium SMS service.

To receive these notifications you need to setup a callback URL depending on the type of notification you would like to receive. SMS API notifications are sent for various SMS categories as shown below:

#### 1.1 Delivery Reports

These are sent whenever the mobile service provider confirms or rejects delivery of a message. To receive delivery reports, you need to set a delivery report callback URL. From the dashboard select SMS -> SMS Callback URLs -> Delivery Reports.

#### Delivery Reports Callback Example

```bash
    ./atsvr-cli sms delivery-report --body '{
          "failureReason": "UserAccountSuspended",
          "id": "Consequatur et dicta.",
          "networkCode": "64110",
          "phoneNumber": "Recusandae perspiciatis excepturi reiciendis est quasi est.",
          "retryCount": "Non quo omnis consequatur vero.",
          "status": "Buffered"
       }'
```

#### Delivery Reports Response Example

```bash
ts=2020-05-31T15:12:19.311626187Z caller=log.go:30 id=Ad3v0i-K req="POST /callbacks/africastalking/sms/deliveryreport" from=127.0.0.1
ts=2020-05-31T15:12:19.311917897Z caller=sms.go:24 info=sms.deliveryReport
ts=2020-05-31T15:12:19.311981209Z caller=log.go:37 id=Ad3v0i-K status=201 bytes=3 time=358.417µs
```

#### 1.2 Incoming Messages.

These are sent whenever a message is sent to any of your registered shortcodes. To receive incoming messages, you need to set an incoming messages callback URL. From the dashboard select SMS -> SMS Callback URLs -> Incoming Messages.

#### 1.3 Bulk SMS Opt Out

These are sent whenever a user opts out of receiving messages from your alphanumeric sender ID. To receive bulk sms opt out notifications, you need to set a bulk sms opt out callback URL. From the dashboard select SMS -> SMS Callback URLs -> Bulk SMS Opt Out. The instructions on how to opt out are automatically appended to the first message you send to the mobile subscriber. From then onwards, any other message will be sent ‘as is’ to the subscriber.

#### 1.4 Subscription Notification.

These are sent whenever someone subscribes or unsubscribes from any of your premium SMS products. To receive premium sms subscription notifications, you need to set asubscription notification callback URL. From the dashboard select SMS -> SMS Callback URLs -> Subscription Notifications.

## 2. Voice Callback Service

The Voice API sends a notification when a specific event happens. To receive these notifications you need to setup a voice callback URL. From the dashboard select Voice -> Phone Numbers -> Actions -> Callback.

Voice API notifications are sent for;

#### 2.1 Outbound calls

These are sent whenever you make a call from a registered SIP number.

#### 2.2 Inbound calls

These are sent when a call comes to your virtual or SIP number.

#### 2.3 After Input

These are sent whenever an action in your response requires user input (such as GetDigits and Record)

#### 2.4 When Call Ends

These are sent after a call ends. This is the final notification and contains some extra information about the call like the cost and duration.

## 3. USSD Callback Service

This service is hit when the user dials a USSD code and every time they respond to a menu. Processing USSD requests using the USSD API is very easy once your account is set up. In particular, you will need to:

- Register a service code with AfricasTalking.
- Register a URL that they can call whenever they get a request from a client coming into their system.

Once you register your callback URL, any requests that they receive belonging to you will trigger a callback that sends the request data to that URL using HTTP POST.

## 4. Airtime Callback Service

#### 4.1 Airtime Validation Notifications

The Airtime API provides optional functionality to validate airtime requests from your application. To receive these notifications you need to setup an airtime validation callback URL. From the dashboard select Airtime -> Airtime Callback URLs -> Validation Callback URL. Airtime validation notifications are sent as HTTP POST requests to the validation callback URL provided.

Once you receive a validation callback notification you’ll be expected to send back a JSON response that marks the transaction as Validated or Failed. If validated we will proceed to send the airtime, if failed, we will block the airtime transaction

#### 4.2 Airtime Status Notifications

The Airtime API sends delivery status notification from the mobile service provider to your application indicating success or failure of the request. To receive these notifications you need to setup an airtime status callback URL. From the dashboard select Airtime -> Airtime Callback URLs -> Status Callback URL. Status notification content Airtime status notifications are sent as HTTP POST requests to the status callback URL provided.

## 5. Payment Callback Service

The Payment API sends a notification when a specific event happens. To receive these notifications you need to setup a callback URL depending on the type of notification.

#### 5.1 BankCheckout

These are sent once the provided bank confirms or rejects the checkout request, or once the checkout request expires.

#### 5.2 CardCheckout

These are sent once the card provider confirms or rejects the checkout request, or once the checkout request expires.

#### 5.3 MobileCheckout

These are sent once the mobile subscriber confirms or rejects the checkout request, or once the checkout request expires.

#### 5.4 MobileC2B

These are sent once funds are moved from the mobile subscriber’s account to your payment wallet.

#### 5.5 MobileB2C

These are sent once funds are successfully moved from your payment wallet to the mobile subscriber’s account.
If AT are not able to complete the transaction, they will refund your payment wallet with the value of the transaction, and also refund your Africa’s Talking Stash with any transaction Fees.

#### 5.6 MobileB2B

These are sent once funds are successfully moved from your payment wallet to the recipeints business account. If AT are not able to complete the transaction, they will refund your payment wallet with the value of the transaction, and also refund your Africa’s Talking Stash with any transaction Fees.

#### 5.7 BankTransfer

These are sent once funds are successfully moved from your payment wallet to the provided bank account account. If AT are not able to complete the transaction, they will refund your payment wallet with the value of the transaction, and also refund your Africa’s Talking Stash with any transaction Fees.

#### 5.8 WalletTransfer

These are sent once funds are successfully moved from your payment wallet to the target payment wallet.

#### 5.9 UserStashTopup

These are sent once funds are successfully moved from your payment wallet to your Africa’s Talking Stash.

## 6. IoT Callback Service

The IoT API sends a notification when a device publish event occurs. To receive these notifications you need to setup a callback URL depending on the type of notification. When creating your device group, you have the option to supply your application callback URL. This is the remote endpoint to which device messages will be sent.
