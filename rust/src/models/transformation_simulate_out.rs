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
pub struct TransformationSimulateOut {
    #[serde(rename = "method", skip_serializing_if = "Option::is_none")]
    pub method: Option<models::TransformationHttpMethod>,
    #[serde(rename = "payload")]
    pub payload: String,
    #[serde(rename = "url")]
    pub url: String,
}

impl TransformationSimulateOut {
    pub fn new(payload: String, url: String) -> TransformationSimulateOut {
        TransformationSimulateOut {
            method: None,
            payload,
            url,
        }
    }
}
