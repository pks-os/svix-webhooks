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
                pub struct ApplicationOut {
                        #[serde(rename = "createdAt")]
                        pub created_at: String,
                        /// The app's ID
                        #[serde(rename = "id")]
                        pub id: String,
                        #[serde(rename = "metadata")]
                        pub metadata: std::collections::HashMap<String, String>,
                        #[serde(rename = "name")]
                        pub name: String,
                        #[serde(rename = "rateLimit", skip_serializing_if = "Option::is_none")]
                        pub rate_limit: Option<i32>,
                        /// The app's UID
                        #[serde(rename = "uid", skip_serializing_if = "Option::is_none")]
                        pub uid: Option<String>,
                        #[serde(rename = "updatedAt")]
                        pub updated_at: String,
                    }

                    impl ApplicationOut {
                    pub fn new(created_at: String, id: String, metadata: std::collections::HashMap<String, String>, name: String, updated_at: String) -> ApplicationOut {
                ApplicationOut {
                    created_at,
                    id,
                    metadata,
                    name,
                    rate_limit: None,
                    uid: None,
                    updated_at,
                    }
                    }
                    }

