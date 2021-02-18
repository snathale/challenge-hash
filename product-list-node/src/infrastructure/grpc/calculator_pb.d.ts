// package: proto
// file: calculator.proto

/* tslint:disable */
/* eslint-disable */

import * as jspb from "google-protobuf";

export class Request extends jspb.Message { 
    getUserId(): string;
    setUserId(value: string): Request;

    getProductId(): string;
    setProductId(value: string): Request;


    serializeBinary(): Uint8Array;
    toObject(includeInstance?: boolean): Request.AsObject;
    static toObject(includeInstance: boolean, msg: Request): Request.AsObject;
    static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
    static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
    static serializeBinaryToWriter(message: Request, writer: jspb.BinaryWriter): void;
    static deserializeBinary(bytes: Uint8Array): Request;
    static deserializeBinaryFromReader(message: Request, reader: jspb.BinaryReader): Request;
}

export namespace Request {
    export type AsObject = {
        userId: string,
        productId: string,
    }
}

export class Discount extends jspb.Message { 
    getPercentage(): number;
    setPercentage(value: number): Discount;

    getValueInCents(): number;
    setValueInCents(value: number): Discount;


    serializeBinary(): Uint8Array;
    toObject(includeInstance?: boolean): Discount.AsObject;
    static toObject(includeInstance: boolean, msg: Discount): Discount.AsObject;
    static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
    static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
    static serializeBinaryToWriter(message: Discount, writer: jspb.BinaryWriter): void;
    static deserializeBinary(bytes: Uint8Array): Discount;
    static deserializeBinaryFromReader(message: Discount, reader: jspb.BinaryReader): Discount;
}

export namespace Discount {
    export type AsObject = {
        percentage: number,
        valueInCents: number,
    }
}
