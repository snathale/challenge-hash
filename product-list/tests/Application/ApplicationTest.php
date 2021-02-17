<?php


namespace Tests\Application;


use ArangoDBClient\Document;
use GuzzleHttp\Client;
use PHPUnit\Framework\TestCase;
use Tests\Test_helper\TestHelper;

class ApplicationTest extends TestCase
{
    private TestHelper $test_helper;
    private Client $client;

    public function setUp(): void
    {
        $this->test_helper = new TestHelper($this);
        $this->client = new Client(['base_uri' => 'http://product.svc:80']);
    }

    public function testGetProductWithBirthdayDiscount()
    {
        $this->test_helper->createCollection('product');
        $this->test_helper->createCollection('user');
        $product1 = $this->test_helper->getFakeProductDocument(3000, 'iphone 12', 'apple smartphone');
        $this->test_helper->insertDocument('product', $product1);
        $product2 = $this->test_helper->getFakeProductDocument(6000, 'mackbook pro', 'apple notebook');
        $this->test_helper->insertDocument('product', $product2);
        $date = new \DateTime('now', new \DateTimeZone('UTC'));
        $format = $date->format('Y-m-d\TH:i:s.u\Z');
        $user = $this->getFakeUser('Ada', 'Lovelace', $format);
        $this->test_helper->insertDocument('user', $user);

        $response = $this->client->get('/product', ['headers' => ['X-USER-ID' => $user->getKey()]]);

        $assert =  [
            'status' => 'success',
            'data' => [
                [
                    'id' => $product1->getKey(),
                    'price_in_cents' => 3000,
                    'title' => 'iphone 12',
                    'description' => 'apple smartphone',
                    'discount' => [
                        'percentage' => 0.05,
                        'value_in_cents' => 2850
                    ]
                ],
                [
                    'id' => $product2->getKey(),
                    'price_in_cents' => 6000,
                    'title' => 'mackbook pro',
                    'description' => 'apple notebook',
                    'discount' => [
                        'percentage' => 0.05,
                        'value_in_cents' => 5700
                    ]
                ]
            ]
        ];
        $this->assertEquals($response->getStatusCode(), 200);
        $this->assertEquals(json_decode($response->getBody()->getContents(), true), $assert);
    }

    public function testGetProductWithoutDiscount()
    {
        $this->test_helper->createCollection('product');
        $this->test_helper->createCollection('user');
        $product1 = $this->test_helper->getFakeProductDocument(3000, 'iphone 12', 'apple smartphone');
        $this->test_helper->insertDocument('product', $product1);
        $product2 = $this->test_helper->getFakeProductDocument(6000, 'mackbook pro', 'apple notebook');
        $this->test_helper->insertDocument('product', $product2);
        $date = new \DateTime('now', new \DateTimeZone('UTC'));
        $interval = new \DateInterval('P1M');
        $date->sub($interval);
        $format = $date->format('Y-m-d\TH:i:s.u\Z');
        $user = $this->getFakeUser('Ada', 'Lovelace', $format);
        $this->test_helper->insertDocument('user', $user);

        $response = $this->client->get('/product', ['headers' => ['X-USER-ID' => $user->getKey()]]);

        $assert =  [
            'status' => 'success',
            'data' => [
                [
                    'id' => $product1->getKey(),
                    'price_in_cents' => 3000,
                    'title' => 'iphone 12',
                    'description' => 'apple smartphone',
                    'discount' => [
                        'percentage' => 0,
                        'value_in_cents' => 0
                    ]
                ],
                [
                    'id' => $product2->getKey(),
                    'price_in_cents' => 6000,
                    'title' => 'mackbook pro',
                    'description' => 'apple notebook',
                    'discount' => [
                        'percentage' => 0,
                        'value_in_cents' => 0
                    ]
                ]
            ]
        ];
        $this->assertEquals($response->getStatusCode(), 200);
        $this->assertEquals(json_decode($response->getBody()->getContents(), true), $assert);
    }

    public function testAllGetProductWithoutDiscountWithUnknownUser()
    {
        $this->test_helper->createCollection('product');
        $this->test_helper->createCollection('user');
        $product1 = $this->test_helper->getFakeProductDocument(3000, 'iphone 12', 'apple smartphone');
        $this->test_helper->insertDocument('product', $product1);
        $product2 = $this->test_helper->getFakeProductDocument(6000, 'mackbook pro', 'apple notebook');
        $this->test_helper->insertDocument('product', $product2);

        $response = $this->client->get('/product', ['headers' => ['X-USER-ID' => 1]]);

        $assert =  [
            'status' => 'success',
            'data' => [
                [
                    'id' => $product1->getKey(),
                    'price_in_cents' => 3000,
                    'title' => 'iphone 12',
                    'description' => 'apple smartphone',
                    'discount' => [
                        'percentage' => 0,
                        'value_in_cents' => 0
                    ]
                ],
                [
                    'id' => $product2->getKey(),
                    'price_in_cents' => 6000,
                    'title' => 'mackbook pro',
                    'description' => 'apple notebook',
                    'discount' => [
                        'percentage' => 0,
                        'value_in_cents' => 0
                    ]
                ]
            ]
        ];
        $this->assertEquals($response->getStatusCode(), 200);
        $this->assertEquals(json_decode($response->getBody()->getContents(), true), $assert);
    }

    public function testGetAllProductWithoutConsultDiscountService()
    {
        $this->test_helper->createCollection('product');
        $product1 = $this->test_helper->getFakeProductDocument(3000, 'iphone 12', 'apple smartphone');
        $this->test_helper->insertDocument('product', $product1);
        $product2 = $this->test_helper->getFakeProductDocument(6000, 'mackbook pro', 'apple notebook');
        $this->test_helper->insertDocument('product', $product2);

        $response = $this->client->get('/product');

        $assert =  [
            'status' => 'success',
            'data' => [
                [
                    'id' => $product1->getKey(),
                    'price_in_cents' => 3000,
                    'title' => 'iphone 12',
                    'description' => 'apple smartphone',
                    'discount' => [
                        'percentage' => 0,
                        'value_in_cents' => 0
                    ]
                ],
                [
                    'id' => $product2->getKey(),
                    'price_in_cents' => 6000,
                    'title' => 'mackbook pro',
                    'description' => 'apple notebook',
                    'discount' => [
                        'percentage' => 0,
                        'value_in_cents' => 0
                    ]
                ]
            ]
        ];
        $this->assertEquals($response->getStatusCode(), 200);
        $this->assertEquals(json_decode($response->getBody()->getContents(), true), $assert);
    }

    private function getFakeUser(string $firstName, string $lastName, string $dateOfBirth): Document
    {
        $doc = new Document();
        $doc->set('first_name', $firstName);
        $doc->set('last_name', $lastName);
        $doc->set('date_of_birth', $dateOfBirth);
        return $doc;
    }
}