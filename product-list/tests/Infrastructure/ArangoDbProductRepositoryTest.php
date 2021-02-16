<?php


namespace Tests\Infrastructure;


use Api\Domain\Entity\Product;
use Api\Infrastructure\ArangoDb\ArangoDbProductRepository;
use Api\Infrastructure\ArangoDb\ProductCollection;
use ArangoDBClient\Document;
use Monolog\Logger;
use PHPUnit\Framework\TestCase;
use Tests\Test_helper\TestHelper;

class ArangoDbProductRepositoryTest extends TestCase
{
    private TestHelper $test_helper;

    public function setUp(): void
    {
        $this->test_helper = new TestHelper($this);
    }

    public function testGetAllProductsEmpty() {
        $this->test_helper->createCollection('product');
        $productCollection = new ProductCollection($this->test_helper->getArangoConnection(), 'product');
        $repository = new ArangoDbProductRepository($productCollection, new Logger('log-test'));
        $this->assertEquals([], $repository->getAllProducts());
    }

    public function testGetAllProducts(){
        $this->test_helper->createCollection('product');
        $product1 = $this->test_helper->getFakeProductDocument(3000, 'iphone 12', 'apple smartphone');
        $this->test_helper->insertDocument('product', $product1);
        $product2 = $this->test_helper->getFakeProductDocument(6000, 'mackbook pro', 'apple notebook');
        $this->test_helper->insertDocument('product', $product2);
        $productCollection = new ProductCollection($this->test_helper->getArangoConnection(), 'product');
        $repository = new ArangoDbProductRepository($productCollection, new Logger('log-test'));
        $product1Assert = $this->getFakeProduct(
            $product1->getKey(), 3000, 'iphone 12', 'apple smartphone'
        );
        $product2Assert = $this->getFakeProduct(
            $product2->getKey(), 6000, 'mackbook pro', 'apple notebook'
        );
        $this->assertEquals([$product1Assert,$product2Assert], $repository->getAllProducts());
    }

    private function getFakeProduct(string $id, int $priceInCents, string $title, string $description) {
        return new Product($id, $priceInCents, $title, $description);
    }
}