<?php


namespace Api\Infrastructure\ArangoDb;



use Api\Domain\Entity\Product;
use Api\Domain\Repository\ProductRepository;
use ArangoDBClient\Cursor;
use ArangoDBClient\Document;
use ArangoDBClient\Exception;
use ArangoDBClient\Statement;
use Psr\Log\LoggerInterface;

class ArangoDbProductRepository implements ProductRepository
{

    private ProductCollection $productCollection;
    private LoggerInterface $logger;

    /**
     * ArangoDbProductRepository constructor.
     * @param ProductCollection $productCollection
     * @param LoggerInterface $logger
     */
    public function __construct(ProductCollection $productCollection, LoggerInterface $logger)
    {
        $this->productCollection = $productCollection;
        $this->logger = $logger;
    }

    /**
     * @return Product[]
     * @throws RepositoryError
     */
    public function getAllProducts(): array
    {
        try {
            $statement = new Statement($this->productCollection->connection, [
                'query' => 'FOR p IN @@collection RETURN p',
                'bindVars' => [
                    '@collection' => $this->productCollection->collectionName
                ]
            ]);
            $cursor = $statement->execute();
            if ($cursor->getCount() > 0) {
                return $this->deserialize($cursor);
            }
            return [];
        } catch (Exception $e) {
            $this->logger->error('impossible restore product list, reason: '.$e->getMessage());
            throw new RepositoryError('impossible restore product list');
        }
    }

    /**
     * @param Cursor $cursor
     * @return Product[]
     */
    private function deserialize(Cursor $cursor): array
    {
        return \Functional\map($cursor, function (Document $doc) {
           return new Product(
               $doc->getKey(), $doc->get('price_in_cents'), $doc->get('title'), $doc->get('description'));
        });
    }
}