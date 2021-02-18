import { db } from "@arangodb";
import { Database } from "arangojs";
import express from "express";
import * as core from "express-serve-static-core";
import ProductController from "../controller/productController";
import ProductRepository from "../domain/repository/productRepository";
import CalculatorDiscountService from "../infrastructure/grpc/calculatorDiscountService";
import productRepository from "../infrastructure/productRepository";
import config from "./config";

export default class App {
    server: core.Express
    repository: ProductRepository

    
    constructor() {
        this.server = express()        
        const db = {
            url: config.dbHost,
            databaseName: config.dbName,
            auth: { username: config.user, password: config.password },
        }
        console.log(config)
        this.repository = new productRepository(new Database(db), config.collectionName)
    }


    public run() {
        this.registerRoutes()
        this.server.listen(config.port)
    }

    private registerRoutes() {
        let routes = express.Router();
        const controller = new ProductController(this.repository, new CalculatorDiscountService())
        routes.get("/product", controller.getDiscount)
        this.server.use(routes)
    }
}