// GENERATED CODE -- DO NOT EDIT!

'use strict';
var grpc = require('grpc');
var calculator_pb = require('./calculator_pb.js');

function serialize_proto_Discount(arg) {
  if (!(arg instanceof calculator_pb.Discount)) {
    throw new Error('Expected argument of type proto.Discount');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_proto_Discount(buffer_arg) {
  return calculator_pb.Discount.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_proto_Request(arg) {
  if (!(arg instanceof calculator_pb.Request)) {
    throw new Error('Expected argument of type proto.Request');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_proto_Request(buffer_arg) {
  return calculator_pb.Request.deserializeBinary(new Uint8Array(buffer_arg));
}


var CalculatorService = exports.CalculatorService = {
  getProductDiscount: {
    path: '/proto.Calculator/GetProductDiscount',
    requestStream: false,
    responseStream: false,
    requestType: calculator_pb.Request,
    responseType: calculator_pb.Discount,
    requestSerialize: serialize_proto_Request,
    requestDeserialize: deserialize_proto_Request,
    responseSerialize: serialize_proto_Discount,
    responseDeserialize: deserialize_proto_Discount,
  },
};

exports.CalculatorClient = grpc.makeGenericClientConstructor(CalculatorService);
