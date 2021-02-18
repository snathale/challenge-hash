const config  = {
    port: process.env.SERVER_PORT || 3333,
    dbHost: process.env.DB_HOST || 'http://arangodb.svc:8529',
    dbName: process.env.DB_NAME || 'dummy_discount_db',
    user: process.env.DB_USER || 'root',
    password: process.env.DB_PASSWORD || 'dummy_passowrd',
    collectionName: process.env.DB_COLLECTION || 'product',
    calculatorServiceHost: process.env.CALCULATOR_HOST|| 'calculator.svc:3000',
}

export default config