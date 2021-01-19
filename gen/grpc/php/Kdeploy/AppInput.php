<?php
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: kdeploy.proto

namespace Kdeploy;

use Google\Protobuf\Internal\GPBType;
use Google\Protobuf\Internal\RepeatedField;
use Google\Protobuf\Internal\GPBUtil;

/**
 * AppInput creates/replaces an app
 *
 * Generated from protobuf message <code>kdeploy.AppInput</code>
 */
class AppInput extends \Google\Protobuf\Internal\Message
{
    /**
     * name of the application
     *
     * Generated from protobuf field <code>string name = 1 [(.validator.field) = {</code>
     */
    private $name = '';
    /**
     * application namespace
     *
     * Generated from protobuf field <code>string namespace = 2 [(.validator.field) = {</code>
     */
    private $namespace = '';
    /**
     * docker image of application
     *
     * Generated from protobuf field <code>string image = 3 [(.validator.field) = {</code>
     */
    private $image = '';
    /**
     * args are arguments given to the docker image at startup
     *
     * Generated from protobuf field <code>repeated string args = 4;</code>
     */
    private $args;
    /**
     * k/v map of environmental variables
     *
     * Generated from protobuf field <code>map<string, string> env = 5;</code>
     */
    private $env;
    /**
     * k/v map of ports to expose ex: http: 80 https: 443
     *
     * Generated from protobuf field <code>map<string, uint32> ports = 6 [(.validator.field) = {</code>
     */
    private $ports;
    /**
     * number of deployment replicas
     *
     * Generated from protobuf field <code>uint32 replicas = 7;</code>
     */
    private $replicas = 0;
    /**
     * Generated from protobuf field <code>.kdeploy.Networking networking = 9 [(.validator.field) = {</code>
     */
    private $networking = null;
    /**
     * Generated from protobuf field <code>map<string, string> labels = 10 [(.validator.field) = {</code>
     */
    private $labels;
    /**
     * Generated from protobuf field <code>map<string, string> selector = 11 [(.validator.field) = {</code>
     */
    private $selector;

    /**
     * Constructor.
     *
     * @param array $data {
     *     Optional. Data for populating the Message object.
     *
     *     @type string $name
     *           name of the application
     *     @type string $namespace
     *           application namespace
     *     @type string $image
     *           docker image of application
     *     @type string[]|\Google\Protobuf\Internal\RepeatedField $args
     *           args are arguments given to the docker image at startup
     *     @type array|\Google\Protobuf\Internal\MapField $env
     *           k/v map of environmental variables
     *     @type array|\Google\Protobuf\Internal\MapField $ports
     *           k/v map of ports to expose ex: http: 80 https: 443
     *     @type int $replicas
     *           number of deployment replicas
     *     @type \Kdeploy\Networking $networking
     *     @type array|\Google\Protobuf\Internal\MapField $labels
     *     @type array|\Google\Protobuf\Internal\MapField $selector
     * }
     */
    public function __construct($data = NULL) {
        \GPBMetadata\Kdeploy::initOnce();
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
     * application namespace
     *
     * Generated from protobuf field <code>string namespace = 2 [(.validator.field) = {</code>
     * @return string
     */
    public function getNamespace()
    {
        return $this->namespace;
    }

    /**
     * application namespace
     *
     * Generated from protobuf field <code>string namespace = 2 [(.validator.field) = {</code>
     * @param string $var
     * @return $this
     */
    public function setNamespace($var)
    {
        GPBUtil::checkString($var, True);
        $this->namespace = $var;

        return $this;
    }

    /**
     * docker image of application
     *
     * Generated from protobuf field <code>string image = 3 [(.validator.field) = {</code>
     * @return string
     */
    public function getImage()
    {
        return $this->image;
    }

    /**
     * docker image of application
     *
     * Generated from protobuf field <code>string image = 3 [(.validator.field) = {</code>
     * @param string $var
     * @return $this
     */
    public function setImage($var)
    {
        GPBUtil::checkString($var, True);
        $this->image = $var;

        return $this;
    }

    /**
     * args are arguments given to the docker image at startup
     *
     * Generated from protobuf field <code>repeated string args = 4;</code>
     * @return \Google\Protobuf\Internal\RepeatedField
     */
    public function getArgs()
    {
        return $this->args;
    }

    /**
     * args are arguments given to the docker image at startup
     *
     * Generated from protobuf field <code>repeated string args = 4;</code>
     * @param string[]|\Google\Protobuf\Internal\RepeatedField $var
     * @return $this
     */
    public function setArgs($var)
    {
        $arr = GPBUtil::checkRepeatedField($var, \Google\Protobuf\Internal\GPBType::STRING);
        $this->args = $arr;

        return $this;
    }

    /**
     * k/v map of environmental variables
     *
     * Generated from protobuf field <code>map<string, string> env = 5;</code>
     * @return \Google\Protobuf\Internal\MapField
     */
    public function getEnv()
    {
        return $this->env;
    }

    /**
     * k/v map of environmental variables
     *
     * Generated from protobuf field <code>map<string, string> env = 5;</code>
     * @param array|\Google\Protobuf\Internal\MapField $var
     * @return $this
     */
    public function setEnv($var)
    {
        $arr = GPBUtil::checkMapField($var, \Google\Protobuf\Internal\GPBType::STRING, \Google\Protobuf\Internal\GPBType::STRING);
        $this->env = $arr;

        return $this;
    }

    /**
     * k/v map of ports to expose ex: http: 80 https: 443
     *
     * Generated from protobuf field <code>map<string, uint32> ports = 6 [(.validator.field) = {</code>
     * @return \Google\Protobuf\Internal\MapField
     */
    public function getPorts()
    {
        return $this->ports;
    }

    /**
     * k/v map of ports to expose ex: http: 80 https: 443
     *
     * Generated from protobuf field <code>map<string, uint32> ports = 6 [(.validator.field) = {</code>
     * @param array|\Google\Protobuf\Internal\MapField $var
     * @return $this
     */
    public function setPorts($var)
    {
        $arr = GPBUtil::checkMapField($var, \Google\Protobuf\Internal\GPBType::STRING, \Google\Protobuf\Internal\GPBType::UINT32);
        $this->ports = $arr;

        return $this;
    }

    /**
     * number of deployment replicas
     *
     * Generated from protobuf field <code>uint32 replicas = 7;</code>
     * @return int
     */
    public function getReplicas()
    {
        return $this->replicas;
    }

    /**
     * number of deployment replicas
     *
     * Generated from protobuf field <code>uint32 replicas = 7;</code>
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
     * Generated from protobuf field <code>.kdeploy.Networking networking = 9 [(.validator.field) = {</code>
     * @return \Kdeploy\Networking
     */
    public function getNetworking()
    {
        return $this->networking;
    }

    /**
     * Generated from protobuf field <code>.kdeploy.Networking networking = 9 [(.validator.field) = {</code>
     * @param \Kdeploy\Networking $var
     * @return $this
     */
    public function setNetworking($var)
    {
        GPBUtil::checkMessage($var, \Kdeploy\Networking::class);
        $this->networking = $var;

        return $this;
    }

    /**
     * Generated from protobuf field <code>map<string, string> labels = 10 [(.validator.field) = {</code>
     * @return \Google\Protobuf\Internal\MapField
     */
    public function getLabels()
    {
        return $this->labels;
    }

    /**
     * Generated from protobuf field <code>map<string, string> labels = 10 [(.validator.field) = {</code>
     * @param array|\Google\Protobuf\Internal\MapField $var
     * @return $this
     */
    public function setLabels($var)
    {
        $arr = GPBUtil::checkMapField($var, \Google\Protobuf\Internal\GPBType::STRING, \Google\Protobuf\Internal\GPBType::STRING);
        $this->labels = $arr;

        return $this;
    }

    /**
     * Generated from protobuf field <code>map<string, string> selector = 11 [(.validator.field) = {</code>
     * @return \Google\Protobuf\Internal\MapField
     */
    public function getSelector()
    {
        return $this->selector;
    }

    /**
     * Generated from protobuf field <code>map<string, string> selector = 11 [(.validator.field) = {</code>
     * @param array|\Google\Protobuf\Internal\MapField $var
     * @return $this
     */
    public function setSelector($var)
    {
        $arr = GPBUtil::checkMapField($var, \Google\Protobuf\Internal\GPBType::STRING, \Google\Protobuf\Internal\GPBType::STRING);
        $this->selector = $arr;

        return $this;
    }

}

