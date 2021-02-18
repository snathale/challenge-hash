// package: proto
// file: calculator.proto

/* tslint:disable */
/* eslint-disable */

import * as grpc from "grpc";
import * as calculator_pb from "./calculator_pb";

interface ICalculatorService extends grpc.ServiceDefinition<grpc.UntypedServiceImplementation> {
    getProductDiscount: ICalculatorService_IGetProductDiscount;
}

interface ICalculatorService_IGetProductDiscount extends grpc.MethodDefinition<calculator_pb.Request, calculator_pb.Discount> {
    path: "/proto.Calculator/GetProductDiscount";
    requestStream: false;
    responseStream: false;
    requestSerialize: grpc.serialize<calculator_pb.Request>;
    requestDeserialize: grpc.deserialize<calculator_pb.Request>;
    responseSerialize: grpc.serialize<calculator_pb.Discount>;
    responseDeserialize: grpc.deserialize<calculator_pb.Discount>;
}

export const CalculatorService: ICalculatorService;

export interface ICalculatorServer {
    getProductDiscount: grpc.handleUnaryCall<calculator_pb.Request, calculator_pb.Discount>;
}

export interface ICalculatorClient {
    getProductDiscount(request: calculator_pb.Request, callback: (error: grpc.ServiceError | null, response: calculator_pb.Discount) => void): grpc.ClientUnaryCall;
    getProductDiscount(request: calculator_pb.Request, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: calculator_pb.Discount) => void): grpc.ClientUnaryCall;
    getProductDiscount(request: calculator_pb.Request, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: calculator_pb.Discount) => void): grpc.ClientUnaryCall;
}

export class CalculatorClient extends grpc.Client implements ICalculatorClient {
    constructor(address: string, credentials: grpc.ChannelCredentials, options?: object);
    public getProductDiscount(request: calculator_pb.Request, callback: (error: grpc.ServiceError | null, response: calculator_pb.Discount) => void): grpc.ClientUnaryCall;
    public getProductDiscount(request: calculator_pb.Request, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: calculator_pb.Discount) => void): grpc.ClientUnaryCall;
    public getProductDiscount(request: calculator_pb.Request, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: calculator_pb.Discount) => void): grpc.ClientUnaryCall;
}
