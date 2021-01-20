<?php
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: schema.proto

namespace Meshpaas;

use Google\Protobuf\Internal\GPBType;
use Google\Protobuf\Internal\RepeatedField;
use Google\Protobuf\Internal\GPBUtil;

/**
 * App is a stateless application
 *
 * Generated from protobuf message <code>meshpaas.App</code>
 */
class App extends \Google\Protobuf\Internal\Message
{
    /**
     * name of the application
     *
     * Generated from protobuf field <code>string name = 1 [(.validator.field) = {</code>
     */
    private $name = '';
    /**
     * application project
     *
     * Generated from protobuf field <code>string project = 2 [(.validator.field) = {</code>
     */
    private $project = '';
    /**
     * Generated from protobuf field <code>repeated .meshpaas.Container containers = 3 [(.validator.field) = {</code>
     */
    private $containers;
    /**
     * number of deployment replicas
     *
     * Generated from protobuf field <code>uint32 replicas = 8;</code>
     */
    private $replicas = 0;
    /**
     * gateway/service-mesh networking
     *
     * Generated from protobuf field <code>.meshpaas.Networking networking = 11 [(.validator.field) = {</code>
     */
    private $networking = null;
    /**
     * application authentication options
     *
     * Generated from protobuf field <code>.meshpaas.Authn authentication = 12 [(.validator.field) = {</code>
     */
    private $authentication = null;
    /**
     * status tracks the state of the application during it's lifecycle
     *
     * Generated from protobuf field <code>.meshpaas.AppStatus status = 20 [(.validator.field) = {</code>
     */
    private $status = null;

    /**
     * Constructor.
     *
     * @param array $data {
     *     Optional. Data for populating the Message object.
     *
     *     @type string $name
     *           name of the application
     *     @type string $project
     *           application project
     *     @type \Meshpaas\Container[]|\Google\Protobuf\Internal\RepeatedField $containers
     *     @type int $replicas
     *           number of deployment replicas
     *     @type \Meshpaas\Networking $networking
     *           gateway/service-mesh networking
     *     @type \Meshpaas\Authn $authentication
     *           application authentication options
     *     @type \Meshpaas\AppStatus $status
     *           status tracks the state of the application during it's lifecycle
     * }
     */
    public function __construct($data = NULL) {
        \GPBMetadata\Schema::initOnce();
        parent::__construct($data);
    }

    /**
     * name of the application
     *
     * Generated from protobuf field <code>string name = 1 [(.validator.field) = {</code>
     * @return string
     */
    public function getName()
    {
        return $this->name;
    }

    /**
     * name of the application
     *
     * Generated from protobuf field <code>string name = 1 [(.validator.field) = {</code>
     * @param string $var
     * @return $this
     */
    public function setName($var)
    {
        GPBUtil::checkString($var, True);
        $this->name = $var;

        return $this;
    }

    /**
     * application project
     *
     * Generated from protobuf field <code>string project = 2 [(.validator.field) = {</code>
     * @return string
     */
    public function getProject()
    {
        return $this->project;
    }

    /**
     * application project
     *
     * Generated from protobuf field <code>string project = 2 [(.validator.field) = {</code>
     * @param string $var
     * @return $this
     */
    public function setProject($var)
    {
        GPBUtil::checkString($var, True);
        $this->project = $var;

        return $this;
    }

    /**
     * Generated from protobuf field <code>repeated .meshpaas.Container containers = 3 [(.validator.field) = {</code>
     * @return \Google\Protobuf\Internal\RepeatedField
     */
    public function getContainers()
    {
        return $this->containers;
    }

    /**
     * Generated from protobuf field <code>repeated .meshpaas.Container containers = 3 [(.validator.field) = {</code>
     * @param \Meshpaas\Container[]|\Google\Protobuf\Internal\RepeatedField $var
     * @return $this
     */
    public function setContainers($var)
    {
        $arr = GPBUtil::checkRepeatedField($var, \Google\Protobuf\Internal\GPBType::MESSAGE, \Meshpaas\Container::class);
        $this->containers = $arr;

        return $this;
    }

    /**
     * number of deployment replicas
     *
     * Generated from protobuf field <code>uint32 replicas = 8;</code>
     * @return int
     */
    public function getReplicas()
    {
        return $this->replicas;
    }

    /**
     * number of deployment replicas
     *
     * Generated from protobuf field <code>uint32 replicas = 8;</code>
     * @param int $var
     * @return $this
     */
    public function setReplicas($var)
    {
        GPBUtil::checkUint32($var);
        $this->replicas = $var;

        return $this;
    }

    /**
     * gateway/service-mesh networking
     *
     * Generated from protobuf field <code>.meshpaas.Networking networking = 11 [(.validator.field) = {</code>
     * @return \Meshpaas\Networking
     */
    public function getNetworking()
    {
        return $this->networking;
    }

    /**
     * gateway/service-mesh networking
     *
     * Generated from protobuf field <code>.meshpaas.Networking networking = 11 [(.validator.field) = {</code>
     * @param \Meshpaas\Networking $var
     * @return $this
     */
    public function setNetworking($var)
    {
        GPBUtil::checkMessage($var, \Meshpaas\Networking::class);
        $this->networking = $var;

        return $this;
    }

    /**
     * application authentication options
     *
     * Generated from protobuf field <code>.meshpaas.Authn authentication = 12 [(.validator.field) = {</code>
     * @return \Meshpaas\Authn
     */
    public function getAuthentication()
    {
        return $this->authentication;
    }

    /**
     * application authentication options
     *
     * Generated from protobuf field <code>.meshpaas.Authn authentication = 12 [(.validator.field) = {</code>
     * @param \Meshpaas\Authn $var
     * @return $this
     */
    public function setAuthentication($var)
    {
        GPBUtil::checkMessage($var, \Meshpaas\Authn::class);
        $this->authentication = $var;

        return $this;
    }

    /**
     * status tracks the state of the application during it's lifecycle
     *
     * Generated from protobuf field <code>.meshpaas.AppStatus status = 20 [(.validator.field) = {</code>
     * @return \Meshpaas\AppStatus
     */
    public function getStatus()
    {
        return $this->status;
    }

    /**
     * status tracks the state of the application during it's lifecycle
     *
     * Generated from protobuf field <code>.meshpaas.AppStatus status = 20 [(.validator.field) = {</code>
     * @param \Meshpaas\AppStatus $var
     * @return $this
     */
    public function setStatus($var)
    {
        GPBUtil::checkMessage($var, \Meshpaas\AppStatus::class);
        $this->status = $var;

        return $this;
    }

}

