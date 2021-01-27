<?php
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: schema.proto

namespace Hpaas;

use Google\Protobuf\Internal\GPBType;
use Google\Protobuf\Internal\RepeatedField;
use Google\Protobuf\Internal\GPBUtil;

/**
 * Charts is a list of helm charts
 *
 * Generated from protobuf message <code>hpaas.Charts</code>
 */
class Charts extends \Google\Protobuf\Internal\Message
{
    /**
     * Generated from protobuf field <code>repeated .hpaas.Chart charts = 1;</code>
     */
    private $charts;

    /**
     * Constructor.
     *
     * @param array $data {
     *     Optional. Data for populating the Message object.
     *
     *     @type \Hpaas\Chart[]|\Google\Protobuf\Internal\RepeatedField $charts
     * }
     */
    public function __construct($data = NULL) {
        \GPBMetadata\Schema::initOnce();
        parent::__construct($data);
    }

    /**
     * Generated from protobuf field <code>repeated .hpaas.Chart charts = 1;</code>
     * @return \Google\Protobuf\Internal\RepeatedField
     */
    public function getCharts()
    {
        return $this->charts;
    }

    /**
     * Generated from protobuf field <code>repeated .hpaas.Chart charts = 1;</code>
     * @param \Hpaas\Chart[]|\Google\Protobuf\Internal\RepeatedField $var
     * @return $this
     */
    public function setCharts($var)
    {
        $arr = GPBUtil::checkRepeatedField($var, \Google\Protobuf\Internal\GPBType::MESSAGE, \Hpaas\Chart::class);
        $this->charts = $arr;

        return $this;
    }

}

