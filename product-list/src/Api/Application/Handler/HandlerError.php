<?php


namespace Api\Application\Handler;


class HandlerError implements \JsonSerializable
{
    public const BAD_REQUEST = 'BAD_REQUEST';
    public const SERVER_ERROR = 'SERVER_ERROR';
    const RESOURCE_NOT_FOUND = 'RESOURCE_NOT_FOUND';

    /**
     * @var string|null
     */
    private ?string $code;

    /**
     * @var string
     */
    private string $message;

    /**
     * @param string|null $code
     * @param string $message
     */
    public function __construct(?string $code, string $message)
    {
        $this->code = $code;
        $this->message = $message;
    }

    /**
     * @return string
     */
    public function getCode(): string
    {
        return $this->code;
    }

    /**
     * @param string $code
     * @return self
     */
    public function setCode(string $code): self
    {
        $this->code = $code;
        return $this;
    }

    /**
     * @return string
     */
    public function getMessage(): string
    {
        return $this->message;
    }

    /**
     * @param string $message
     * @return self
     */
    public function setMessage(string $message): self
    {
        $this->message = $message;
        return $this;
    }

    /**
     * @return array
     */
    public function jsonSerialize()
    {
        return [
            'code' => $this->code,
            'message' => $this->message,
        ];
    }
}