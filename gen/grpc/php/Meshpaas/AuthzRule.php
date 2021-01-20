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
 * Generated from protobuf message <code>meshpaas.AuthzRule</code>
 */
class AuthzRule extends \Google\Protobuf\Internal\Message
{
    /**
     * Generated from protobuf field <code>.meshpaas.AuthzSource source = 1;</code>
     */
    private $source = null;
    /**
     * Generated from protobuf field <code>.meshpaas.AuthzDestination destination = 2;</code>
     */
    private $destination = null;
    /**
     * Generated from protobuf field <code>.meshpaas.AuthzSubject subject = 3;</code>
     */
    private $subject = null;

    /**
     * Constructor.
     *
     * @param array $data {
     *     Optional. Data for populating the Message object.
     *
     *     @type \Meshpaas\AuthzSource $source
     *     @type \Meshpaas\AuthzDestination $destination
     *     @type \Meshpaas\AuthzSubject $subject
     * }
     */
    public function __construct($data = NULL) {
        \GPBMetadata\Schema::initOnce();
        parent::__construct($data);
    }

    /**
     * Generated from protobuf field <code>.meshpaas.AuthzSource source = 1;</code>
     * @return \Meshpaas\AuthzSource
     */
    public function getSource()
    {
        return $this->source;
    }

    /**
     * Generated from protobuf field <code>.meshpaas.AuthzSource source = 1;</code>
     * @param \Meshpaas\AuthzSource $var
     * @return $this
     */
    public function setSource($var)
    {
        GPBUtil::checkMessage($var, \Meshpaas\AuthzSource::class);
        $this->source = $var;

        return $this;
    }

    /**
     * Generated from protobuf field <code>.meshpaas.AuthzDestination destination = 2;</code>
     * @return \Meshpaas\AuthzDestination
     */
    public function getDestination()
    {
        return $this->destination;
    }

    /**
     * Generated from protobuf field <code>.meshpaas.AuthzDestination destination = 2;</code>
     * @param \Meshpaas\AuthzDestination $var
     * @return $this
     */
    public function setDestination($var)
    {
        GPBUtil::checkMessage($var, \Meshpaas\AuthzDestination::class);
        $this->destination = $var;

        return $this;
    }

    /**
     * Generated from protobuf field <code>.meshpaas.AuthzSubject subject = 3;</code>
     * @return \Meshpaas\AuthzSubject
     */
    public function getSubject()
    {
        return $this->subject;
    }

    /**
     * Generated from protobuf field <code>.meshpaas.AuthzSubject subject = 3;</code>
     * @param \Meshpaas\AuthzSubject $var
     * @return $this
     */
    public function setSubject($var)
    {
        GPBUtil::checkMessage($var, \Meshpaas\AuthzSubject::class);
        $this->subject = $var;

        return $this;
    }

}
