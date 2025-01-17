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
pub struct EndpointUpdate {
    /// List of message channels this endpoint listens to (omit for all).
    #[serde(rename = "channels", skip_serializing_if = "Option::is_none")]
    pub channels: Option<Vec<String>>,
    #[serde(rename = "description", skip_serializing_if = "Option::is_none")]
    pub description: Option<String>,
    #[serde(rename = "disabled", skip_serializing_if = "Option::is_none")]
    pub disabled: Option<bool>,
    #[serde(rename = "filterTypes", skip_serializing_if = "Option::is_none")]
    pub filter_types: Option<Vec<String>>,
    #[serde(rename = "metadata", skip_serializing_if = "Option::is_none")]
    pub metadata: Option<std::collections::HashMap<String, String>>,
    #[serde(rename = "rateLimit", skip_serializing_if = "Option::is_none")]
    pub rate_limit: Option<i32>,
    /// Optional unique identifier for the endpoint.
    #[serde(rename = "uid", skip_serializing_if = "Option::is_none")]
    pub uid: Option<String>,
    #[serde(rename = "url")]
    pub url: String,
    #[serde(rename = "version", skip_serializing_if = "Option::is_none")]
    pub version: Option<i32>,
}

impl EndpointUpdate {
    pub fn new(url: String) -> EndpointUpdate {
        EndpointUpdate {
            channels: None,
            description: None,
            disabled: None,
            filter_types: None,
            metadata: None,
            rate_limit: None,
            uid: None,
            url,
            version: None,
        }
    }
}
