<?php


namespace Api\Application\Handler;


use Slim\Psr7\Request;
use Slim\Psr7\Response;

class GetProductDiscount extends Handler
{

    /**
     * @param Request $request
     * @param Response $response
     * @param array $args
     * @return Response
     * @throws \Exception
     */
    public function __invoke(Request $request, Response $response, array $args): Response
    {
        $userId = $request->getAttribute('userId');
        $products = $this->controller->getProductDiscount($userId);
        $this->handlerPayload->successResponse($products);
        $encodedPayload = json_encode($this->handlerPayload, true);
        $response->getBody()->write($encodedPayload);
        $response = $response->withHeader('Content-Type', 'Application/json')
            ->withHeader('Access-Control-Allow-Credentials', 'true');
        return $response;
    }
}