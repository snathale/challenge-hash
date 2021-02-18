import { Request, Response } from "express";
import jsend from "jsend"
import Product from "../domain/entity/product";
import ProductRepository from "../domain/repository/productRepository";
import CalculatorDiscountService from "../infrastructure/grpc/calculatorDiscountService";

export default class ProductController {
    private repository: ProductRepository
    private service: CalculatorDiscountService

    constructor(repository: ProductRepository, service: CalculatorDiscountService) {
        this.repository = repository
        this.service = service
    }

    getDiscount = async (req: Request, res: Response) => {
        let products: Product[] = []
        try {
            const user_id = req.header('X-USER-ID')
            products = await this.repository.getAllProducts()
            if (user_id) {
                for (let product of products) {
                    await this.service.calculate(product.id, user_id)
                    .then((discount) => {
                        product.setDiscount(discount)
                    }).catch();
                }
            }
        } catch (err) {
            console.log(err)
        }
        res.status(200).json(jsend.success(products))
    }
}