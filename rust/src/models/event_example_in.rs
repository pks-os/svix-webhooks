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

#[derive(Clone, Default, Debug, PartialEq, Serialize, Deserialize)]
pub struct EventExampleIn {
    /// The event type's name
    #[serde(rename = "eventType")]
    pub event_type: String,
    /// If the event type schema contains an array of examples, chooses which
    /// one to send.  Defaults to the first example. Ignored if the schema
    /// doesn't contain an array of examples.
    #[serde(rename = "exampleIndex", skip_serializing_if = "Option::is_none")]
    pub example_index: Option<i32>,
}

impl EventExampleIn {
    pub fn new(event_type: String) -> EventExampleIn {
        EventExampleIn {
            event_type,
            example_index: None,
        }
    }
}
