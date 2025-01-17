/*
 * Svix API
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * The version of the OpenAPI document: 1.1.1
 *
 * Generated by: https://openapi-generator.tech
 */

#[allow(unused_imports)]
use crate::models;
#[allow(unused_imports)]
use serde::{Deserialize, Serialize};

#[derive(Clone, Debug, PartialEq, Serialize, Deserialize)]
#[serde(untagged)]
pub enum SinkOut {
    SinkInOneOf(Box<models::SinkInOneOf>),
    SinkInOneOf1(Box<models::SinkInOneOf1>),
    SinkInOneOf2(Box<models::SinkInOneOf2>),
    SinkInOneOf3(Box<models::SinkInOneOf3>),
    SinkInOneOf4(Box<models::SinkInOneOf4>),
}

impl Default for SinkOut {
    fn default() -> Self {
        Self::SinkInOneOf(Default::default())
    }
}
///
#[derive(Clone, Copy, Debug, Eq, PartialEq, Ord, PartialOrd, Hash, Serialize, Deserialize)]
pub enum Type {
    #[serde(rename = "rabbitMQ")]
    RabbitMq,
    #[serde(rename = "sqs")]
    Sqs,
    #[serde(rename = "kafka")]
    Kafka,
    #[serde(rename = "http")]
    Http,
    #[serde(rename = "eventStream")]
    EventStream,
}

impl Default for Type {
    fn default() -> Type {
        Self::RabbitMq
    }
}
