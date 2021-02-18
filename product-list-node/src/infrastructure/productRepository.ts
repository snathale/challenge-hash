import { Database } from "arangojs";
import _ from "lodash";
import Product from "../domain/entity/product";
import ProductRepository from "../domain/repository/productRepository";

export default class productRepository implements ProductRepository {

    public db: Database
    public collectionName: string

    constructor(db: Database, collectionName: string) {
        this.db = db
        this.collectionName = collectionName
    }

    async getAllProducts(): Promise<Array<Product>> {
        const cursor = await this.db.query({
            query: `FOR p IN @@collection RETURN p`,
            bindVars: {
                "@collection": this.collectionName
            }
        });
        const docs = await cursor.all();
        return  _.map(docs, (doc) => {
            return new Product(doc._key, doc.price_in_cents, doc.title, doc.description)
        })
    }
}