import Product from "../entity/product"

export default interface ProductRepository {
    getAllProducts(): Promise<Product[]>
}