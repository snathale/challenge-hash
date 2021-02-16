<?php


namespace Api\Application\Error;


use Api\Application\Handler\HandlerError;
use Api\Application\Handler\HandlerPayload;
use Psr\Http\Message\ResponseInterface;
use Slim\Exception\HttpBadRequestException;
use Slim\Exception\HttpException;
use Slim\Exception\HttpNotFoundException;
use Slim\Handlers\ErrorHandler as SlimErrorHandler;
use Throwable;

class HttpError extends SlimErrorHandler
{

    /**
     * @inheritdoc
     */
    protected function respond(): ResponseInterface
    {
        $exception = $this->exception;
        $statusCode = 500;
        $error = new HandlerError(
            HandlerError::SERVER_ERROR,
            'An internal error has occurred while processing your request.'
        );
        if ($exception instanceof HttpException) {
            $statusCode = $exception->getCode();
            $error->setMessage($exception->getMessage());

            if ($exception instanceof HttpNotFoundException) {
                $error->setCode(HandlerError::RESOURCE_NOT_FOUND);
            } elseif ($exception instanceof HttpBadRequestException) {
                $error->setCode(HandlerError::BAD_REQUEST);
            }
        }
        if (!($exception instanceof HttpException) && $exception instanceof Throwable && $this->displayErrorDetails) {
            $error->setMessage($exception->getMessage());
        }
        $payload = new HandlerPayload($statusCode);
        $payload->errorResponse($error);
        $encodedPayload = json_encode($payload, JSON_PRETTY_PRINT);
        $response = $this->responseFactory->createResponse($statusCode);
        $response->getBody()->write($encodedPayload);
        $response = $response->withHeader('Content-Type', 'Application/json')
            ->withHeader('Access-Control-Allow-Credentials', 'true');
        return $response;
    }
}