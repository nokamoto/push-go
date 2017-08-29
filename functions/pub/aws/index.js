'use strict';

const config = require('./config.json');

var grpc = require('grpc');
var protoDescriptor = grpc.load(__dirname + '/service.proto');
var push = protoDescriptor.push;

const AWS = require('aws-sdk');
const SQS = new AWS.SQS({ apiVersion: '2012-11-05' });

exports.handler = (event, context, callback) => {
    console.log('Received event:', JSON.stringify(event, null, 2));

    const done = (err, res) => callback(null, {
        statusCode: err ? '400' : '200',
        body: err ? err.message : JSON.stringify(res),
        headers: {
            'Content-Type': 'application/json',
        },
    });

    const body = JSON.parse(event.body);

    if (body.message === undefined) {
        done(new Error('No message defined!'));
    } else if(body.topic === undefined) {
        done(new Error('No topic defined!'));
    } else {
        console.log('topic=' + body.topic + ', message=' + body.message);

        var client = new push.SubscriptionService(
            config.SUBSCRIPTION_HOST + ':' + config.SUBSCRIPTION_PORT,
            grpc.credentials.createInsecure()
        );

        var meta = new grpc.Metadata();

        meta.add('x-api-key', config.API_KEY);

        var call = client.get({name: body.topic}, meta);

        call.on('data', function(id) {
            console.log('data=' + JSON.stringify(id));

            var params = {
                MessageBody: JSON.stringify({
                    message: body.message,
                    id: id
                }),
                QueueUrl: config.QUEUE_URL,
                DelaySeconds: 0
            };

            SQS.sendMessage(params, function(err, data) {
                if (err) {
                    console.log(err);
                } else {
                    console.log(data)
                }
            });
        });

        call.on('end', function() {
            done(null, {message: "done"});
        });
    }
};
