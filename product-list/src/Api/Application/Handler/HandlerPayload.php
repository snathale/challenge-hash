<?php


namespace Api\Application\Handler;


class HandlerPayload implements \JsonSerializable
{
    /**
     * @var int
     */
    private int $statusCode;

    /**
     * @var array|object|null
     */
    private $data;

    /**
     * @var HandlerError|null
     */
    private ?HandlerError $error;

    /**
     * @var string|null
     */
    private string $status;

    /**
     * @param int $statusCode
     */
    public function __construct(int $statusCode = 200)
    {
        $this->statusCode = $statusCode;
    }

    public function successResponse($data) {
        $this->status = 'success';
        $this->data = $data;
    }

    public function errorResponse(HandlerError $error) {
        $this->statusCode = 'error';
        $this->error = $error;
    }

    /**
     * @return int
     */
    public function getStatusCode(): int
    {
        return $this->statusCode;
    }

    /**
     * @return array
     */
    public function jsonSerialize(): array
    {
        $payload = [
            'status' => $this->status
        ];
        if (isset($this->data)) {
            $payload['data'] = $this->data;
        }
        if (isset($this->error)) {
            $payload['message'] = $this->error->getMessage();
            $payload['code'] = $this->error->getCode();
        }
        return $payload;
    }
}