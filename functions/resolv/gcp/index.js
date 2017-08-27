'use strict';

const config = require('./config.json');

var grpc = require('grpc');
var protoDescriptor = grpc.load(__dirname + '/service.proto');
var push = protoDescriptor.push;

exports.resolv = function resolv(event, callback) {
    const message = event.data;
    const data = Buffer.from(message.data, 'base64').toString();

    console.log(JSON.stringify(message));
    console.log(data);

    callback();
};
