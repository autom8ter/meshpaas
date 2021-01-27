<?php
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: schema.proto

namespace Meshpaas;

use Google\Protobuf\Internal\GPBType;
use Google\Protobuf\Internal\RepeatedField;
use Google\Protobuf\Internal\GPBUtil;

/**
 * At least on AuthzRule rule must pass for a request to reach its final destination(an application) in the service mesh
 *
 * Generated from protobuf message <code>hpaas.AuthzRule</code>
 */
class AuthzRule extends \Google\Protobuf\Internal\Message
{
    /**
     * subject restricts access based on the subject of a request in the service mesh
     *
     * Generated from protobuf field <code>.hpaas.AuthzSubject subject = 1;</code>
     */
    private $subject = null;
    /**
     * source restricts access based on the destination of a request in the service mesh
     *
     * Generated from protobuf field <code>.hpaas.AuthzDestination destination = 2;</code>
     */
    private $destination = null;

    /**
     * Constructor.
     *
     * @param array $data {
     *     Optional. Data for populating the Message object.
     *
     *     @type \Meshpaas\AuthzSubject $subject
     *           subject restricts access based on the subject of a request in the service mesh
     *     @type \Meshpaas\AuthzDestination $destination
     *           source restricts access based on the destination of a request in the service mesh
     * }
     */
    public function __construct($data = NULL) {
        \GPBMetadata\Schema::initOnce();
        parent::__construct($data);
    }

    /**
     * subject restricts access based on the subject of a request in the service mesh
     *
     * Generated from protobuf field <code>.hpaas.AuthzSubject subject = 1;</code>
     * @return \Meshpaas\AuthzSubject
     */
    public function getSubject()
    {
        return $this->subject;
    }

    /**
     * subject restricts access based on the subject of a request in the service mesh
     *
     * Generated from protobuf field <code>.hpaas.AuthzSubject subject = 1;</code>
     * @param \Meshpaas\AuthzSubject $var
     * @return $this
     */
    public function setSubject($var)
    {
        GPBUtil::checkMessage($var, \Meshpaas\AuthzSubject::class);
        $this->subject = $var;

        return $this;
    }

    /**
     * source restricts access based on the destination of a request in the service mesh
     *
     * Generated from protobuf field <code>.hpaas.AuthzDestination destination = 2;</code>
     * @return \Meshpaas\AuthzDestination
     */
    public function getDestination()
    {
        return $this->destination;
    }

    /**
     * source restricts access based on the destination of a request in the service mesh
     *
     * Generated from protobuf field <code>.hpaas.AuthzDestination destination = 2;</code>
     * @param \Meshpaas\AuthzDestination $var
     * @return $this
     */
    public function setDestination($var)
    {
        GPBUtil::checkMessage($var, \Meshpaas\AuthzDestination::class);
        $this->destination = $var;

        return $this;
    }

}

