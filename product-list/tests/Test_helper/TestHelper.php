<?php


namespace Tests\Test_helper;


use ArangoDBClient\Collection;
use ArangoDBClient\CollectionHandler;
use ArangoDBClient\Connection;
use ArangoDBClient\ConnectionOptions;
use ArangoDBClient\Document;
use ArangoDBClient\DocumentHandler;
use ArangoDBClient\Exception;
use PHPUnit\Framework\TestCase;

class TestHelper
{
    const DB_HOST = "tcp://arangodb.svc";
    const DB_PORT = 8529;
    const DB_USER = "root";
    const DB_PASSWORD = "dummy_passowrd";
    const SYSTEM_DB = "dummy_discount_db";

    /**
     * @var TestCase
     */
    private TestCase $testCase;

    private Connection $connection;


    /**
     * TestHelper constructor.
     * @param TestCase $testCase
     */
    public function __construct(TestCase $testCase)
    {
        $this->testCase = $testCase;
        $this->connection = $this->createArangoNewConnection();
    }

    private function createArangoNewConnection(): Connection
    {
        $config = [
            ConnectionOptions::OPTION_DATABASE => self::SYSTEM_DB,
            ConnectionOptions::OPTION_ENDPOINT => self::DB_HOST . ':' . self::DB_PORT,
            ConnectionOptions::OPTION_AUTH_TYPE => 'Basic',
            ConnectionOptions::OPTION_AUTH_USER => self::DB_USER,
            ConnectionOptions::OPTION_AUTH_PASSWD => self::DB_PASSWORD,
        ];
        try {
            return new Connection($config);
        } catch (Exception $e) {
            $this->testCase->fail($e->getMessage());
        }
    }

    public function getArangoConnection(): Connection
    {
        if (isset($this->connection)) {
            return $this->connection;
        }
        return $this->createArangoNewConnection();
    }

    public function createCollection(string $name)
    {
        $handler = new CollectionHandler($this->connection);
        try {
            if ($handler->has($name)) {
                $handler->drop($name);
            }
            $collection = new Collection($name);
            $handler->create($collection);
        } catch (Exception $e) {
            $this->testCase->fail($e->getMessage());
        }
    }

    public function dropCollection(string $name) {
        $handler = new CollectionHandler($this->connection);
        try {
            if ($handler->has($name)) {
                $handler->drop($name);
            }
        } catch (Exception $e) {
            $this->testCase->fail($e->getMessage());
        }
    }

    public function insertDocument(string $collection, Document $document)
    {
        $handler = new DocumentHandler($this->connection);
        try {
            $handler->insert($collection, $document);
        } catch (Exception $e) {
            $this->testCase->fail($e->getMessage());
        }
    }

    public function getFakeProductDocument(int $priceInCents, string $title, string $description): Document {
        $doc = new Document();
        $doc->set('title', $title);
        $doc->set('price_in_cents', $priceInCents);
        $doc->set('description', $description);
        return $doc;
    }
}