import { Metadata } from "grpc";
import Discount from "../../domain/entity/discount";
import { Request } from "./calculator_pb";

import client from './client';

export default class CalculatorDiscountService {

    async calculate(product_id: string, user_id: string): Promise<Discount> {
        return new Promise<Discount>((resolve, reject) => {
            const request = new Request
            request.setUserId(user_id)
            request.setProductId(product_id)
            client.getProductDiscount(request, new Metadata(), {deadline: new Date(Date.now() + 100)}, (err, discount) => {
                if (err) {
                    return reject(err);
                }
                const obj = new Discount(parseFloat(discount.getPercentage().toFixed(2)), discount.getValueInCents())
                return resolve(obj);
            });
        });
    }    
}