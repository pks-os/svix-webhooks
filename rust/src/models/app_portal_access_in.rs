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
                pub struct AppPortalAccessIn {
                        #[serde(rename = "application", skip_serializing_if = "Option::is_none")]
                        pub application: Option<Box<models::ApplicationIn>>,
                        /// How long the token will be valid for, in seconds.  Valid values are between 1 hour and 7 days. The default is 7 days.
                        #[serde(rename = "expiry", skip_serializing_if = "Option::is_none")]
                        pub expiry: Option<i32>,
                        /// The set of feature flags the created token will have access to.
                        #[serde(rename = "featureFlags", skip_serializing_if = "Option::is_none")]
                        pub feature_flags: Option<Vec<String>>,
                        /// Whether the app portal should be in read-only mode.
                        #[serde(rename = "readOnly", skip_serializing_if = "Option::is_none")]
                        pub read_only: Option<bool>,
                    }

                    impl AppPortalAccessIn {
                    pub fn new() -> AppPortalAccessIn {
                AppPortalAccessIn {
                    application: None,
                    expiry: None,
                    feature_flags: None,
                    read_only: None,
                    }
                    }
                    }

