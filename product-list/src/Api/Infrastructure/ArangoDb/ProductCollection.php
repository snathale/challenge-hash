<?php


namespace Api\Infrastructure\ArangoDb;


use ArangoDBClient\Connection;

class ProductCollection
{
    public Connection $connection;

    public string $collectionName;

    /**
     * ProductArangoCollection constructor.
     * @param Connection $connection
     * @param $collectionName
     */
    public function __construct(Connection $connection, string $collectionName)
    {
        $this->connection = $connection;
        $this->collectionName = $collectionName;
    }

}