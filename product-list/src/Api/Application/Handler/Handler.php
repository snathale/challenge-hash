<?php


namespace Api\Application\Handler;


use Api\Application\Controller\GetProductDiscount as GetProductDiscountController;

abstract class Handler
{
    /**
     * @var GetProductDiscountController
     */
    protected GetProductDiscountController $controller;

    /**
     * @var HandlerPayload
     */
    protected HandlerPayload $handlerPayload;

    /**
     * GetProductDiscount constructor.
     * @param GetProductDiscountController $controller
     */
    public function __construct(GetProductDiscountController $controller)
    {
        $this->controller = $controller;
        $this->handlerPayload = new HandlerPayload();
    }
}