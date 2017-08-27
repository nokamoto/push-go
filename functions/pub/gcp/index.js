'use strict';

const config = require('./config.json');

const PubSub = require(`@google-cloud/pubsub`);
const pubsub = PubSub();

var grpc = require('grpc');
var protoDescriptor = grpc.load(__dirname + '/service.proto');
var push = protoDescriptor.push;

exports.pub = function pub (req, res) {
    if (req.body.message === undefined) {
        res.status(400).send('No message defined!');
    } else if(req.body.topic === undefined) {
        res.status(400).send('No topic defined!');
    } else {
        console.log('topic=' + req.body.topic + ', message=' + req.body.message);

        var client = new push.SubscriptionService(
            config.SUBSCRIPTION_HOST + ':' + config.SUBSCRIPTION_PORT,
            grpc.credentials.createInsecure()
        );

        var meta = new grpc.Metadata();

        meta.add('x-api-key', config.API_KEY);

        var call = client.get({name: req.body.topic}, meta);

        console.log(call);

        const topic = pubsub.topic('resolve-endpoints');

        call.on('data', function(id) {
            console.log('data=' + JSON.stringify(id));

            const message = {
                data: {
                    message: req.body.message,
                    id: id
                }
            };

            topic.publish(message)
                .then(function(results){
                    const messageIds = results[0];
                    console.log('published=' + messageIds);
                }).catch(function(err){
                    console.error(err)
                });
        });

        call.on('end', function() {
            res.status(200).end();
        });
    }
};

exports.pub({body: {message: 'x', topic: 'all'}}, {});
