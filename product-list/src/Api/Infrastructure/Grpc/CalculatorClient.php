<?php
// GENERATED CODE -- DO NOT EDIT!

namespace Api\Infrastructure\Grpc;

/**
 */
class CalculatorClient extends \Grpc\BaseStub {

    /**
     * @param string $hostname hostname
     * @param array $opts channel options
     * @param \Grpc\Channel $channel (optional) re-use channel object
     */
    public function __construct($hostname, $opts, $channel = null) {
        parent::__construct($hostname, $opts, $channel);
    }

    /**
     * @param \Api\Infrastructure\Grpc\Request $argument input argument
     * @param array $metadata metadata
     * @param array $options call options
     * @return \Grpc\UnaryCall
     */
    public function GetProductDiscount(\Api\Infrastructure\Grpc\Request $argument,
      $metadata = [], $options = []) {
        return $this->_simpleRequest('/proto.Calculator/GetProductDiscount',
        $argument,
        ['\Api\Infrastructure\Grpc\Discount', 'decode'],
        $metadata, $options);
    }

}
