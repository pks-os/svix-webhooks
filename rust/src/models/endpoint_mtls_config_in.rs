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
pub struct EndpointMtlsConfigIn {
    /// A PEM encoded private key and X509 certificate to identify the webhook
    /// sender.
    #[serde(rename = "identity")]
    pub identity: String,
    /// A PEM encoded X509 certificate used to verify the webhook receiver's
    /// certificate.
    #[serde(rename = "serverCaCert", skip_serializing_if = "Option::is_none")]
    pub server_ca_cert: Option<String>,
}

impl EndpointMtlsConfigIn {
    pub fn new(identity: String) -> EndpointMtlsConfigIn {
        EndpointMtlsConfigIn {
            identity,
            server_ca_cert: None,
        }
    }
}
