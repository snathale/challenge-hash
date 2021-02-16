<?php


namespace Api\Application\Controller;


use Api\Domain\Entity\Discount;
use Api\Domain\Entity\Product;
use Api\Domain\Repository\ProductRepository;
use Api\Infrastructure\ArangoDb\RepositoryError;
use Api\Infrastructure\Grpc\CalculatorClient;
use Api\Infrastructure\Grpc\Request;
use Psr\Log\LoggerInterface;

class GetProductDiscount
{

    const STATUS_OK = 0;

    /**
     * @var ProductRepository
     */
    private ProductRepository $productRepository;

    /**
     * @var LoggerInterface
     */
    private LoggerInterface $logger;

    /**
     * @var CalculatorClient
     */
    private CalculatorClient $client;

    /**
     * GetProductDiscount constructor.
     * @param ProductRepository $productRepository
     * @param LoggerInterface $logger
     * @param CalculatorClient $calculatorClient
     */
    public function __construct(
        ProductRepository $productRepository,
        LoggerInterface $logger,
        CalculatorClient $calculatorClient
    )
    {
        $this->productRepository = $productRepository;
        $this->logger = $logger;
        $this->client = $calculatorClient;
    }

    /**
     * @param string|null $userId
     * @return array
     */
    public function getProductDiscount(string $userId = null): array
    {
        try {
            $products = $this->productRepository->getAllProducts();
            if (isset($userId)) {
                foreach ($products as $product) {
                    $this->getProductDiscountOnGrpcServer($userId, $product);
                }
            }
            $this->logger->info("get product list");
            return $products;
        } catch (RepositoryError $e) {
            return [];
        }
    }

    /**
     * @param string $userId
     * @param Product $product
     */
    private function getProductDiscountOnGrpcServer(string $userId, Product $product): void
    {
        $request = new Request();
        $request->setUserId($userId);
        $request->setProductId($product->getId());
        $clientRequest = $this->client->GetProductDiscount($request);
        $response = $clientRequest->wait();
        if ($response[1]->code !== self::STATUS_OK) {
            $this->logger->error('impossible connect to grpc server, reason: ' . $response[1]->details);
            return;
        }
        $discount = new Discount($response[0]->getPercentage(), $response[0]->getValueInCents());
        $product->setDiscount($discount);
    }
}